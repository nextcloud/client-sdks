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
pub struct CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner {
    #[serde(rename = "from")]
    pub from: String,
    #[serde(rename = "fromLabel")]
    pub from_label: String,
    #[serde(rename = "to")]
    pub to: String,
    #[serde(rename = "toLabel")]
    pub to_label: String,
}

impl CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner {
    pub fn new(from: String, from_label: String, to: String, to_label: String) -> CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner {
        CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner {
            from,
            from_label,
            to,
            to_label,
        }
    }
}


