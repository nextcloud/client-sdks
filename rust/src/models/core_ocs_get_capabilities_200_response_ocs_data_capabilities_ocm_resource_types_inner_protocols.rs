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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols {
    #[serde(rename = "webdav")]
    pub webdav: String,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols {
    pub fn new(webdav: String) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols {
            webdav,
        }
    }
}


