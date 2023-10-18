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
pub struct ThemingThemingGetManifest200ResponseIconsInner {
    #[serde(rename = "src")]
    pub src: String,
    #[serde(rename = "type")]
    pub r#type: String,
    #[serde(rename = "sizes")]
    pub sizes: String,
}

impl ThemingThemingGetManifest200ResponseIconsInner {
    pub fn new(src: String, r#type: String, sizes: String) -> ThemingThemingGetManifest200ResponseIconsInner {
        ThemingThemingGetManifest200ResponseIconsInner {
            src,
            r#type,
            sizes,
        }
    }
}

