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
pub struct CoreCollaborationResourcesListCollection200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreCollaborationResourcesListCollection200ResponseOcs>,
}

impl CoreCollaborationResourcesListCollection200Response {
    pub fn new(ocs: crate::models::CoreCollaborationResourcesListCollection200ResponseOcs) -> CoreCollaborationResourcesListCollection200Response {
        CoreCollaborationResourcesListCollection200Response {
            ocs: Box::new(ocs),
        }
    }
}


