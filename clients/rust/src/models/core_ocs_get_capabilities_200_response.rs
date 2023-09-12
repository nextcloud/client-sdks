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
pub struct CoreOcsGetCapabilities200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreOcsGetCapabilities200ResponseOcs>,
}

impl CoreOcsGetCapabilities200Response {
    pub fn new(ocs: crate::models::CoreOcsGetCapabilities200ResponseOcs) -> CoreOcsGetCapabilities200Response {
        CoreOcsGetCapabilities200Response {
            ocs: Box::new(ocs),
        }
    }
}


