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
pub struct DavDirectGetUrl200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::DavDirectGetUrl200ResponseOcs>,
}

impl DavDirectGetUrl200Response {
    pub fn new(ocs: crate::models::DavDirectGetUrl200ResponseOcs) -> DavDirectGetUrl200Response {
        DavDirectGetUrl200Response {
            ocs: Box::new(ocs),
        }
    }
}

