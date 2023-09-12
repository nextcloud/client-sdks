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
pub struct CoreHoverCardGetUser200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreHoverCardGetUser200ResponseOcs>,
}

impl CoreHoverCardGetUser200Response {
    pub fn new(ocs: crate::models::CoreHoverCardGetUser200ResponseOcs) -> CoreHoverCardGetUser200Response {
        CoreHoverCardGetUser200Response {
            ocs: Box::new(ocs),
        }
    }
}


