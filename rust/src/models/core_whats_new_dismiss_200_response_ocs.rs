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
pub struct CoreWhatsNewDismiss200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: ::std::collections::HashMap<String, serde_json::Value>,
}

impl CoreWhatsNewDismiss200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: ::std::collections::HashMap<String, serde_json::Value>) -> CoreWhatsNewDismiss200ResponseOcs {
        CoreWhatsNewDismiss200ResponseOcs {
            meta: Box::new(meta),
            data,
        }
    }
}


