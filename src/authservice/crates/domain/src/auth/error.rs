use thiserror::Error;

#[derive(Clone, Debug, Error)]
pub enum AuthError {
    #[error("Password Hashing failed: {0}")]
    HashError(String),

    #[error("Invalid Password")]
    InvalidToken,

    #[error("Invalid token")]
    InvalidToken,

    #[error("Token error")]
    TokenError(String),

    #[error("Token Expired")]
    TokenExpired,

    #[error("Internal auth error: {0}")]
    Internal(String),
}
