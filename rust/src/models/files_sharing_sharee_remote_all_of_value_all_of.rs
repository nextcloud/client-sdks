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
pub struct FilesSharingShareeRemoteAllOfValueAllOf {
    #[serde(rename = "server")]
    pub server: String,
}

impl FilesSharingShareeRemoteAllOfValueAllOf {
    pub fn new(server: String) -> FilesSharingShareeRemoteAllOfValueAllOf {
        FilesSharingShareeRemoteAllOfValueAllOf {
            server,
        }
    }
}


