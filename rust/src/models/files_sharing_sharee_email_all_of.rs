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
pub struct FilesSharingShareeEmailAllOf {
    #[serde(rename = "uuid")]
    pub uuid: String,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "type")]
    pub r#type: String,
    #[serde(rename = "shareWithDisplayNameUnique")]
    pub share_with_display_name_unique: String,
    #[serde(rename = "value")]
    pub value: Box<crate::models::FilesSharingShareeValue>,
}

impl FilesSharingShareeEmailAllOf {
    pub fn new(uuid: String, name: String, r#type: String, share_with_display_name_unique: String, value: crate::models::FilesSharingShareeValue) -> FilesSharingShareeEmailAllOf {
        FilesSharingShareeEmailAllOf {
            uuid,
            name,
            r#type,
            share_with_display_name_unique,
            value: Box::new(value),
        }
    }
}


