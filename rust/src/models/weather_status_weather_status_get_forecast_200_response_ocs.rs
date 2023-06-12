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
pub struct WeatherStatusWeatherStatusGetForecast200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Vec<crate::models::WeatherStatusForecast>,
}

impl WeatherStatusWeatherStatusGetForecast200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: Vec<crate::models::WeatherStatusForecast>) -> WeatherStatusWeatherStatusGetForecast200ResponseOcs {
        WeatherStatusWeatherStatusGetForecast200ResponseOcs {
            meta: Box::new(meta),
            data,
        }
    }
}


