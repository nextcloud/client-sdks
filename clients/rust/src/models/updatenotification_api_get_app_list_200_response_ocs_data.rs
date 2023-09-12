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
pub struct UpdatenotificationApiGetAppList200ResponseOcsData {
    #[serde(rename = "missing")]
    pub missing: Vec<crate::models::UpdatenotificationApp>,
    #[serde(rename = "available")]
    pub available: Vec<crate::models::UpdatenotificationApp>,
}

impl UpdatenotificationApiGetAppList200ResponseOcsData {
    pub fn new(missing: Vec<crate::models::UpdatenotificationApp>, available: Vec<crate::models::UpdatenotificationApp>) -> UpdatenotificationApiGetAppList200ResponseOcsData {
        UpdatenotificationApiGetAppList200ResponseOcsData {
            missing,
            available,
        }
    }
}


