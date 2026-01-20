use domain::UserId; 

use crate::ports::user_repository::{AppError, UserRepository}; 

pub struct DeactivateUserUseCase {
    repo: Arc<dyn UserRepository>,
}

impl DeactivateUserUseCase {
    pub fn new(repo: Arc<dyn UserRepository>) -> Self {
        Self { repo }
    }

    pub async fn execute(&self, user_id: UserId) -> Result<(), AppError> {
        let mut user = self
            .repo
            .find_by_id(user_id)
            .await?
            .ok_or(AppError::NotFound)?;
        user.is_active = false; 
        self.repo.save(user).await
    }
}