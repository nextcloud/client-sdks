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
pub struct FilesSharingShareeValue {
    #[serde(rename = "shareType")]
    pub share_type: i64,
    #[serde(rename = "shareWith")]
    pub share_with: String,
}

impl FilesSharingShareeValue {
    pub fn new(share_type: i64, share_with: String) -> FilesSharingShareeValue {
        FilesSharingShareeValue {
            share_type,
            share_with,
        }
    }
}


