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
pub struct CoreWhatsNewGet200ResponseOcsData {
    #[serde(rename = "changelogURL")]
    pub changelog_url: String,
    #[serde(rename = "product")]
    pub product: String,
    #[serde(rename = "version")]
    pub version: String,
    #[serde(rename = "whatsNew", skip_serializing_if = "Option::is_none")]
    pub whats_new: Option<Box<crate::models::CoreWhatsNewGet200ResponseOcsDataWhatsNew>>,
}

impl CoreWhatsNewGet200ResponseOcsData {
    pub fn new(changelog_url: String, product: String, version: String) -> CoreWhatsNewGet200ResponseOcsData {
        CoreWhatsNewGet200ResponseOcsData {
            changelog_url,
            product,
            version,
            whats_new: None,
        }
    }
}


