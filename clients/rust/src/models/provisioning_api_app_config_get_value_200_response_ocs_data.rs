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
pub struct ProvisioningApiAppConfigGetValue200ResponseOcsData {
    #[serde(rename = "data")]
    pub data: String,
}

impl ProvisioningApiAppConfigGetValue200ResponseOcsData {
    pub fn new(data: String) -> ProvisioningApiAppConfigGetValue200ResponseOcsData {
        ProvisioningApiAppConfigGetValue200ResponseOcsData {
            data,
        }
    }
}


