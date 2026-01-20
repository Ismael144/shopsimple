use async_trait::async_trait; 
use domain::UserId; 
use domain::entities::user::User;

#[derive(Debug, thiserror::Error)]
pub enum AppError {
    #[error("User already exists")]
    UserAlreadyExists, 
    #[error("User not found")]
    NotFound, 
    #[error("Authentication error: ")]
    AuthenticationError(#[from] domain::AuthError), 
    #[error("Internal error: {0}")]
    Internal(String),
}

#[async_trait]
pub trait UserRepository: Send + Sync {
    async fn save(&self, user: User) -> Result<(), AppError>; 
    async fn find_by_id(&self, email: &str) -> Result<Option<User>, AppError>; 
    async fn find_by_email(&self, email: &str) -> Result<Option<User>, AppError>; 
    async fn list(&elf, limit: u32, offset: u32) -> Result<Vec<User>, AppError>; 
}