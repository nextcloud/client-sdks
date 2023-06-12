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
pub struct CoreOcsGetCapabilities200ResponseOcsData {
    #[serde(rename = "version")]
    pub version: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataVersion>,
    #[serde(rename = "capabilities")]
    pub capabilities: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilities,
}

impl CoreOcsGetCapabilities200ResponseOcsData {
    pub fn new(version: crate::models::CoreOcsGetCapabilities200ResponseOcsDataVersion, capabilities: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilities) -> CoreOcsGetCapabilities200ResponseOcsData {
        CoreOcsGetCapabilities200ResponseOcsData {
            version: Box::new(version),
            capabilities,
        }
    }
}


