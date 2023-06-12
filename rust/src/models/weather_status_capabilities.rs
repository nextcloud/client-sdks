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
pub struct WeatherStatusCapabilities {
    #[serde(rename = "weather_status")]
    pub weather_status: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate>,
}

impl WeatherStatusCapabilities {
    pub fn new(weather_status: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate) -> WeatherStatusCapabilities {
        WeatherStatusCapabilities {
            weather_status: Box::new(weather_status),
        }
    }
}


