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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
    #[serde(rename = "chunking")]
    pub chunking: String,
    #[serde(rename = "bulkupload", skip_serializing_if = "Option::is_none")]
    pub bulkupload: Option<String>,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
    pub fn new(chunking: String) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav {
            chunking,
            bulkupload: None,
        }
    }
}


