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
pub struct CoreWhatsNewGet200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::CoreWhatsNewGet200ResponseOcs>,
}

impl CoreWhatsNewGet200Response {
    pub fn new(ocs: crate::models::CoreWhatsNewGet200ResponseOcs) -> CoreWhatsNewGet200Response {
        CoreWhatsNewGet200Response {
            ocs: Box::new(ocs),
        }
    }
}


