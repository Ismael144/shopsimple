use crate::auth::error::AuthError;
use crate::entities::user::User;
use async_trait::async_trait;

#[async_trait]
pub trait AuthProvider: Send + Sync {
    fn hash_password(&self, password: &str) -> Result<String, AuthError>;
    fn verify_password(&self, password: &str, hash: &str) -> Result<bool, AuthError>;
    fn verify_token(&self, token_str: &str) -> Result<String, AuthError>;
    fn generate_token(&self, user: &User) -> Result<String, AuthError>;
}
