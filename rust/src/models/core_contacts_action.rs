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
pub struct CoreContactsAction {
    #[serde(rename = "title")]
    pub title: String,
    #[serde(rename = "icon")]
    pub icon: String,
    #[serde(rename = "hyperlink")]
    pub hyperlink: String,
    #[serde(rename = "appId")]
    pub app_id: String,
}

impl CoreContactsAction {
    pub fn new(title: String, icon: String, hyperlink: String, app_id: String) -> CoreContactsAction {
        CoreContactsAction {
            title,
            icon,
            hyperlink,
            app_id,
        }
    }
}


