use std::sync::Arc;

use domain::AuthProvider;

use crate::{
    ports::user_repository::UserRepository,
    use_cases::{
        deactivate_user::DeactivateUserUseCase, get_user::GetUserUseCase,
        list_users::ListUsersUseCase, login::LoginUseCase, signup::SignUpUseCase,
    },
};

/// User Service
/// With all user service use cases 
pub struct UserService {
    pub get_user: GetUserUseCase,
    pub list_users: ListUsersUseCase,
    pub login: LoginUseCase,
    pub signup: SignUpUseCase,
    pub deactivate_user: DeactivateUserUseCase,
}

impl UserService {
    pub fn new(repo: Arc<dyn UserRepository>, auth: Arc<dyn AuthProvider>) -> Self {
        Self {
            signup: SignUpUseCase::new(repo.clone().auth.clone()),
            list_users: ListUsersUseCase::new(repo.clone()),
            login: LoginUseCase::new(repo.clone()),
            deactivate_user: DeactivateUserUseCase::new(repo.clone()),
            get_user: GetUserUseCase::new(repo.clone()),
        }
    }
}
