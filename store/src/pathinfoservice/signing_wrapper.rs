//! This module provides a [PathInfoService] implementation that signs narinfos

use super::PathInfoService;
use crate::proto::PathInfo;
use futures::stream::BoxStream;
use std::path::PathBuf;
use std::sync::Arc;
use tonic::async_trait;

use tvix_castore::composition::{CompositionContext, ServiceBuilder};

use tvix_castore::Error;

use nix_compat::narinfo::{parse_keypair, SigningKey};
use nix_compat::nixbase32;
use tracing::{instrument, warn};

#[cfg(test)]
use super::MemoryPathInfoService;

/// PathInfoService that wraps around an inner [PathInfoService] and when put is called it extracts
/// the underlying narinfo and signs it using a [SigningKey]. For the moment only the
/// [ed25519::signature::Signer<ed25519::Signature>] is available using a keyfile (see
/// [KeyFileSigningPathInfoServiceConfig] for more informations). However the implementation is
/// generic (see [nix_compat::narinfo::SigningKey] documentation).
///
/// The [PathInfo] with the added signature is then put into the inner [PathInfoService].
///
/// The service signs the [PathInfo] **only if it has a narinfo attribute**
pub struct SigningPathInfoService<T, S> {
    /// The inner [PathInfoService]
    inner: T,
    /// The key to sign narinfos
    signing_key: Arc<SigningKey<S>>,
}

impl<T, S> SigningPathInfoService<T, S> {
    pub fn new(inner: T, signing_key: Arc<SigningKey<S>>) -> Self {
        Self { inner, signing_key }
    }
}

#[async_trait]
impl<T, S> PathInfoService for SigningPathInfoService<T, S>
where
    T: PathInfoService,
    S: ed25519::signature::Signer<ed25519::Signature> + Sync + Send,
{
    #[instrument(level = "trace", skip_all, fields(path_info.digest = nixbase32::encode(&digest)))]
    async fn get(&self, digest: [u8; 20]) -> Result<Option<PathInfo>, Error> {
        self.inner.get(digest).await
    }

    async fn put(&self, path_info: PathInfo) -> Result<PathInfo, Error> {
        let store_path = path_info.validate().map_err(|e| {
            warn!(err=%e, "invalid PathInfo");
            Error::StorageError(e.to_string())
        })?;
        let root_node = path_info.node.clone();
        // If we have narinfo then sign it, else passthrough to the upper pathinfoservice
        let path_info_to_put = match path_info.to_narinfo(store_path.as_ref()) {
            Some(mut nar_info) => {
                nar_info.add_signature(self.signing_key.as_ref());
                let mut signed_path_info = PathInfo::from(&nar_info);
                signed_path_info.node = root_node;
                signed_path_info
            }
            None => path_info,
        };
        self.inner.put(path_info_to_put).await
    }

    fn list(&self) -> BoxStream<'static, Result<PathInfo, Error>> {
        self.inner.list()
    }
}

/// [ServiceBuilder] implementation that builds a [SigningPathInfoService] that signs narinfos using
/// a keyfile. The keyfile is parsed using [parse_keypair], the expected format is the nix one
/// (`nix-store --generate-binary-cache-key` for more informations).
#[derive(serde::Deserialize)]
pub struct KeyFileSigningPathInfoServiceConfig {
    /// Inner [PathInfoService], will be resolved using a [CompositionContext].
    pub inner: String,
    /// Path to the keyfile in the nix format. It will be accessed once when building the service
    pub keyfile: PathBuf,
}

impl TryFrom<url::Url> for KeyFileSigningPathInfoServiceConfig {
    type Error = Box<dyn std::error::Error + Send + Sync>;
    fn try_from(_url: url::Url) -> Result<Self, Self::Error> {
        Err(Error::StorageError(
            "Instantiating a SigningPathInfoService from a url is not supported".into(),
        )
        .into())
    }
}

#[async_trait]
impl ServiceBuilder for KeyFileSigningPathInfoServiceConfig {
    type Output = dyn PathInfoService;
    async fn build<'a>(
        &'a self,
        _instance_name: &str,
        context: &CompositionContext,
    ) -> Result<Arc<dyn PathInfoService>, Box<dyn std::error::Error + Send + Sync + 'static>> {
        let inner = context.resolve(self.inner.clone()).await?;
        let signing_key = Arc::new(
            parse_keypair(tokio::fs::read_to_string(&self.keyfile).await?.trim())
                .map_err(|e| Error::StorageError(e.to_string()))?
                .0,
        );
        Ok(Arc::new(SigningPathInfoService { inner, signing_key }))
    }
}

#[cfg(test)]
pub(crate) fn test_signing_service() -> Arc<dyn PathInfoService> {
    let memory_svc: Arc<dyn PathInfoService> = Arc::new(MemoryPathInfoService::default());
    Arc::new(SigningPathInfoService {
        inner: memory_svc,
        signing_key: Arc::new(
            parse_keypair(DUMMY_KEYPAIR)
                .expect("DUMMY_KEYPAIR to be valid")
                .0,
        ),
    })
}

#[cfg(test)]
pub const DUMMY_KEYPAIR: &str = "do.not.use:sGPzxuK5WvWPraytx+6sjtaff866sYlfvErE6x0hFEhy5eqe7OVZ8ZMqZ/ME/HaRdKGNGvJkyGKXYTaeA6lR3A==";
#[cfg(test)]
pub const DUMMY_VERIFYING_KEY: &str = "do.not.use:cuXqnuzlWfGTKmfzBPx2kXShjRryZMhil2E2ngOpUdw=";

#[cfg(test)]
mod test {
    use crate::{
        pathinfoservice::PathInfoService,
        proto::PathInfo,
        tests::fixtures::{DUMMY_PATH, PATH_INFO_WITH_NARINFO},
    };
    use nix_compat::narinfo::VerifyingKey;

    use lazy_static::lazy_static;
    use nix_compat::store_path::StorePath;

    lazy_static! {
        static ref PATHINFO_1: PathInfo = PATH_INFO_WITH_NARINFO.clone();
        static ref PATHINFO_1_DIGEST: [u8; 20] = [0; 20];
    }

    #[tokio::test]
    async fn put_and_verify_signature() {
        let svc = super::test_signing_service();

        // pathinfo_1 should not be there ...
        assert!(svc
            .get(*PATHINFO_1_DIGEST)
            .await
            .expect("no error")
            .is_none());

        // ... and not be signed
        assert!(PATHINFO_1.narinfo.clone().unwrap().signatures.is_empty());

        // insert it
        svc.put(PATHINFO_1.clone()).await.expect("no error");

        // now it should be there ...
        let signed = svc
            .get(*PATHINFO_1_DIGEST)
            .await
            .expect("no error")
            .unwrap();

        // and signed
        let narinfo = signed
            .to_narinfo(
                StorePath::from_bytes(DUMMY_PATH.as_bytes()).expect("DUMMY_PATH to be parsed"),
            )
            .expect("no error");
        let fp = narinfo.fingerprint();

        // load our keypair from the fixtures
        let (signing_key, _verifying_key) =
            super::parse_keypair(super::DUMMY_KEYPAIR).expect("must succeed");

        // ensure the signature is added
        let new_sig = narinfo
            .signatures
            .last()
            .expect("The retrieved narinfo to be signed");
        assert_eq!(signing_key.name(), *new_sig.name());

        // verify the new signature against the verifying key
        let verifying_key =
            VerifyingKey::parse(super::DUMMY_VERIFYING_KEY).expect("parsing dummy verifying key");

        assert!(
            verifying_key.verify(&fp, new_sig),
            "expect signature to be valid"
        );
    }
}
