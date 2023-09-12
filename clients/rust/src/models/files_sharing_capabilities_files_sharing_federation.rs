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
pub struct FilesSharingCapabilitiesFilesSharingFederation {
    #[serde(rename = "outgoing")]
    pub outgoing: bool,
    #[serde(rename = "incoming")]
    pub incoming: bool,
    #[serde(rename = "expire_date")]
    pub expire_date: Box<crate::models::FilesSharingCapabilitiesFilesSharingUserExpireDate>,
    #[serde(rename = "expire_date_supported")]
    pub expire_date_supported: Box<crate::models::FilesSharingCapabilitiesFilesSharingUserExpireDate>,
}

impl FilesSharingCapabilitiesFilesSharingFederation {
    pub fn new(outgoing: bool, incoming: bool, expire_date: crate::models::FilesSharingCapabilitiesFilesSharingUserExpireDate, expire_date_supported: crate::models::FilesSharingCapabilitiesFilesSharingUserExpireDate) -> FilesSharingCapabilitiesFilesSharingFederation {
        FilesSharingCapabilitiesFilesSharingFederation {
            outgoing,
            incoming,
            expire_date: Box::new(expire_date),
            expire_date_supported: Box::new(expire_date_supported),
        }
    }
}


