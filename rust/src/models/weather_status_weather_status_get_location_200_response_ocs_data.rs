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
pub struct WeatherStatusWeatherStatusGetLocation200ResponseOcsData {
    #[serde(rename = "lat")]
    pub lat: f32,
    #[serde(rename = "lon")]
    pub lon: f32,
    #[serde(rename = "address")]
    pub address: String,
    #[serde(rename = "mode")]
    pub mode: i64,
}

impl WeatherStatusWeatherStatusGetLocation200ResponseOcsData {
    pub fn new(lat: f32, lon: f32, address: String, mode: i64) -> WeatherStatusWeatherStatusGetLocation200ResponseOcsData {
        WeatherStatusWeatherStatusGetLocation200ResponseOcsData {
            lat,
            lon,
            address,
            mode,
        }
    }
}


