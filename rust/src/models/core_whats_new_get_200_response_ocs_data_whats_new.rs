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
pub struct CoreWhatsNewGet200ResponseOcsDataWhatsNew {
    #[serde(rename = "regular")]
    pub regular: Vec<String>,
    #[serde(rename = "admin")]
    pub admin: Vec<String>,
}

impl CoreWhatsNewGet200ResponseOcsDataWhatsNew {
    pub fn new(regular: Vec<String>, admin: Vec<String>) -> CoreWhatsNewGet200ResponseOcsDataWhatsNew {
        CoreWhatsNewGet200ResponseOcsDataWhatsNew {
            regular,
            admin,
        }
    }
}


