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
pub struct CoreCollaborationResourcesSearchCollections200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreCollaborationResourcesSearchCollections200ResponseOcs>,
}

impl CoreCollaborationResourcesSearchCollections200Response {
    pub fn new(ocs: crate::models::CoreCollaborationResourcesSearchCollections200ResponseOcs) -> CoreCollaborationResourcesSearchCollections200Response {
        CoreCollaborationResourcesSearchCollections200Response {
            ocs: Box::new(ocs),
        }
    }
}

