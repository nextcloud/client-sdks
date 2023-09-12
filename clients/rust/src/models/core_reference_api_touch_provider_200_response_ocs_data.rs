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
pub struct CoreReferenceApiTouchProvider200ResponseOcsData {
    #[serde(rename = "success")]
    pub success: bool,
}

impl CoreReferenceApiTouchProvider200ResponseOcsData {
    pub fn new(success: bool) -> CoreReferenceApiTouchProvider200ResponseOcsData {
        CoreReferenceApiTouchProvider200ResponseOcsData {
            success,
        }
    }
}


