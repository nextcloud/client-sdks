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
pub struct CoreTextProcessingApiListTasksByApp200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreTextProcessingApiListTasksByApp200ResponseOcs>,
}

impl CoreTextProcessingApiListTasksByApp200Response {
    pub fn new(ocs: crate::models::CoreTextProcessingApiListTasksByApp200ResponseOcs) -> CoreTextProcessingApiListTasksByApp200Response {
        CoreTextProcessingApiListTasksByApp200Response {
            ocs: Box::new(ocs),
        }
    }
}

