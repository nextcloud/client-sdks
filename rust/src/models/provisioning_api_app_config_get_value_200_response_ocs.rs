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
pub struct ProvisioningApiAppConfigGetValue200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Box<crate::models::ProvisioningApiAppConfigGetValue200ResponseOcsData>,
}

impl ProvisioningApiAppConfigGetValue200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: crate::models::ProvisioningApiAppConfigGetValue200ResponseOcsData) -> ProvisioningApiAppConfigGetValue200ResponseOcs {
        ProvisioningApiAppConfigGetValue200ResponseOcs {
            meta: Box::new(meta),
            data: Box::new(data),
        }
    }
}


