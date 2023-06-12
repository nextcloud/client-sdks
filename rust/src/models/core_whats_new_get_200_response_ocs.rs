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
pub struct CoreWhatsNewGet200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Box<crate::models::CoreWhatsNewGet200ResponseOcsData>,
}

impl CoreWhatsNewGet200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: crate::models::CoreWhatsNewGet200ResponseOcsData) -> CoreWhatsNewGet200ResponseOcs {
        CoreWhatsNewGet200ResponseOcs {
            meta: Box::new(meta),
            data: Box::new(data),
        }
    }
}


