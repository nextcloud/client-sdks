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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
    #[serde(rename = "enabled")]
    pub enabled: bool,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
    pub fn new(enabled: bool) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate {
            enabled,
        }
    }
}


