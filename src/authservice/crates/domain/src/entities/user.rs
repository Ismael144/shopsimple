use crate::{value_objects::role::Role, UserId};
use chrono::{DateTime, Utc};

#[derive(Clone, Debug)]
pub struct User {
    pub id: UserId,
    pub email: String,
    pub password_hash: String,
    pub role: Role,
    pub is_active: bool,
    pub created_at: DateTime<Utc>,
    pub updated_at: DateTime<Utc>,
}

impl User {
    // Initialize new user
    pub fn new(email: String, password_hash: String, role: Role) -> Self {
        let now = Utc::now();
        Self {
            id: UserId::new(),
            email,
            password_hash,
            role,
            is_active: true,
            created_at: now,
            updated_at: now,
        }
    }
}
