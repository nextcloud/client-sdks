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
pub struct ProvisioningApiAppConfigGetApps200ResponseOcsData {
    #[serde(rename = "data")]
    pub data: Vec<String>,
}

impl ProvisioningApiAppConfigGetApps200ResponseOcsData {
    pub fn new(data: Vec<String>) -> ProvisioningApiAppConfigGetApps200ResponseOcsData {
        ProvisioningApiAppConfigGetApps200ResponseOcsData {
            data,
        }
    }
}


