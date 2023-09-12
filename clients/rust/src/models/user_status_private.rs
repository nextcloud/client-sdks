/*
 * nextcloud
 *
 * Nextcloud APIs
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct UserStatusPrivate {
    #[serde(rename = "userId")]
    pub user_id: String,
    #[serde(rename = "message", deserialize_with = "Option::deserialize")]
    pub message: Option<String>,
    #[serde(rename = "icon", deserialize_with = "Option::deserialize")]
    pub icon: Option<String>,
    #[serde(rename = "clearAt", deserialize_with = "Option::deserialize")]
    pub clear_at: Option<i64>,
    #[serde(rename = "status")]
    pub status: String,
    #[serde(rename = "messageId", deserialize_with = "Option::deserialize")]
    pub message_id: Option<String>,
    #[serde(rename = "messageIsPredefined")]
    pub message_is_predefined: bool,
    #[serde(rename = "statusIsUserDefined")]
    pub status_is_user_defined: bool,
}

impl UserStatusPrivate {
    pub fn new(user_id: String, message: Option<String>, icon: Option<String>, clear_at: Option<i64>, status: String, message_id: Option<String>, message_is_predefined: bool, status_is_user_defined: bool) -> UserStatusPrivate {
        UserStatusPrivate {
            user_id,
            message,
            icon,
            clear_at,
            status,
            message_id,
            message_is_predefined,
            status_is_user_defined,
        }
    }
}


