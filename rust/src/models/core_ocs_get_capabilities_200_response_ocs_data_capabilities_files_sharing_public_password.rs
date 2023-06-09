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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword {
    #[serde(rename = "enforced")]
    pub enforced: bool,
    #[serde(rename = "askForOptionalPassword")]
    pub ask_for_optional_password: bool,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword {
    pub fn new(enforced: bool, ask_for_optional_password: bool) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword {
            enforced,
            ask_for_optional_password,
        }
    }
}


