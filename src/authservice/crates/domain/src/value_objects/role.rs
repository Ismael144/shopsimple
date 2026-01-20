use std::str::FromStr;

#[derive(Debug, Default, Clone, Eq, PartialEq, serde::Serialize, serde::Deserialize)]
pub enum Role {
    #[default]
    User,
    Admin,
}

impl FromStr for Role {
    type Err = ();

    fn from_str(row_str: &str) -> Result<Role, Self::Err> {
        match row_str.to_lowercase().as_str() {
            "user" => Ok(Role::User),
            "admin" => Ok(Role::Admin),
            _ => Err(()),
        }
    }
}
