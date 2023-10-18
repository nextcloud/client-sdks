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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
    #[serde(rename = "enabled")]
    pub enabled: bool,
    #[serde(rename = "apiVersion")]
    pub api_version: String,
    #[serde(rename = "endPoint")]
    pub end_point: String,
    #[serde(rename = "resourceTypes")]
    pub resource_types: Vec<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner>,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
    pub fn new(enabled: bool, api_version: String, end_point: String, resource_types: Vec<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner>) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm {
            enabled,
            api_version,
            end_point,
            resource_types,
        }
    }
}

