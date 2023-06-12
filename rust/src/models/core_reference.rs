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
pub struct CoreReference {
    #[serde(rename = "richObjectType")]
    pub rich_object_type: String,
    #[serde(rename = "richObject")]
    pub rich_object: ::std::collections::HashMap<String, serde_json::Value>,
    #[serde(rename = "openGraphObject")]
    pub open_graph_object: Box<crate::models::CoreReferenceOpenGraphObject>,
    #[serde(rename = "accessible")]
    pub accessible: bool,
}

impl CoreReference {
    pub fn new(rich_object_type: String, rich_object: ::std::collections::HashMap<String, serde_json::Value>, open_graph_object: crate::models::CoreReferenceOpenGraphObject, accessible: bool) -> CoreReference {
        CoreReference {
            rich_object_type,
            rich_object,
            open_graph_object: Box::new(open_graph_object),
            accessible,
        }
    }
}


