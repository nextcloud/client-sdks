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
pub struct CoreCollaborationResourcesListCollection200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Box<crate::models::CoreCollection>,
}

impl CoreCollaborationResourcesListCollection200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: crate::models::CoreCollection) -> CoreCollaborationResourcesListCollection200ResponseOcs {
        CoreCollaborationResourcesListCollection200ResponseOcs {
            meta: Box::new(meta),
            data: Box::new(data),
        }
    }
}

