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
pub struct FilesSharingShareeRemoteGroupAllOf {
    #[serde(rename = "guid")]
    pub guid: String,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "value")]
    pub value: Box<crate::models::FilesSharingShareeRemoteAllOfValue>,
}

impl FilesSharingShareeRemoteGroupAllOf {
    pub fn new(guid: String, name: String, value: crate::models::FilesSharingShareeRemoteAllOfValue) -> FilesSharingShareeRemoteGroupAllOf {
        FilesSharingShareeRemoteGroupAllOf {
            guid,
            name,
            value: Box::new(value),
        }
    }
}


