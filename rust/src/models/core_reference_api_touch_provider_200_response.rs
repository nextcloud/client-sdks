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
pub struct CoreReferenceApiTouchProvider200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreReferenceApiTouchProvider200ResponseOcs>,
}

impl CoreReferenceApiTouchProvider200Response {
    pub fn new(ocs: crate::models::CoreReferenceApiTouchProvider200ResponseOcs) -> CoreReferenceApiTouchProvider200Response {
        CoreReferenceApiTouchProvider200Response {
            ocs: Box::new(ocs),
        }
    }
}


