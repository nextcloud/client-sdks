/*
 * nextcloud
 *
 * Nextcloud APIs
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Default, Serialize, Deserialize)]
pub struct UserStatusClearAt {
    #[serde(rename = "type")]
    pub r#type: RHashType,
    #[serde(rename = "time")]
    pub time: Box<crate::models::UserStatusClearAtTime>,
}

impl UserStatusClearAt {
    pub fn new(r#type: RHashType, time: crate::models::UserStatusClearAtTime) -> UserStatusClearAt {
        UserStatusClearAt {
            r#type,
            time: Box::new(time),
        }
    }
}

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum RHashType {
    #[serde(rename = "period")]
    Period,
    #[serde(rename = "end-of")]
    EndOf,
}

impl Default for RHashType {
    fn default() -> RHashType {
        Self::Period
    }
}

