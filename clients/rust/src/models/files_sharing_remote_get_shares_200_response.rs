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
pub struct FilesSharingRemoteGetShares200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::FilesSharingRemoteGetShares200ResponseOcs>,
}

impl FilesSharingRemoteGetShares200Response {
    pub fn new(ocs: crate::models::FilesSharingRemoteGetShares200ResponseOcs) -> FilesSharingRemoteGetShares200Response {
        FilesSharingRemoteGetShares200Response {
            ocs: Box::new(ocs),
        }
    }
}


