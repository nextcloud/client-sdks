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
pub struct UpdatenotificationApiGetAppList200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::UpdatenotificationApiGetAppList200ResponseOcs>,
}

impl UpdatenotificationApiGetAppList200Response {
    pub fn new(ocs: crate::models::UpdatenotificationApiGetAppList200ResponseOcs) -> UpdatenotificationApiGetAppList200Response {
        UpdatenotificationApiGetAppList200Response {
            ocs: Box::new(ocs),
        }
    }
}

