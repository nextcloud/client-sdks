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
pub struct FilesSharingShareeUser {
    #[serde(rename = "count", deserialize_with = "Option::deserialize")]
    pub count: Option<i64>,
    #[serde(rename = "label")]
    pub label: String,
    #[serde(rename = "subline")]
    pub subline: String,
    #[serde(rename = "icon")]
    pub icon: String,
    #[serde(rename = "shareWithDisplayNameUnique")]
    pub share_with_display_name_unique: String,
    #[serde(rename = "status")]
    pub status: Box<crate::models::FilesSharingShareeUserAllOfStatus>,
    #[serde(rename = "value")]
    pub value: Box<crate::models::FilesSharingShareeValue>,
}

impl FilesSharingShareeUser {
    pub fn new(count: Option<i64>, label: String, subline: String, icon: String, share_with_display_name_unique: String, status: crate::models::FilesSharingShareeUserAllOfStatus, value: crate::models::FilesSharingShareeValue) -> FilesSharingShareeUser {
        FilesSharingShareeUser {
            count,
            label,
            subline,
            icon,
            share_with_display_name_unique,
            status: Box::new(status),
            value: Box::new(value),
        }
    }
}


