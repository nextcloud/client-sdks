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
pub struct WeatherStatusForecastDataNext12HoursDetails {
    #[serde(rename = "probability_of_precipitation")]
    pub probability_of_precipitation: f32,
}

impl WeatherStatusForecastDataNext12HoursDetails {
    pub fn new(probability_of_precipitation: f32) -> WeatherStatusForecastDataNext12HoursDetails {
        WeatherStatusForecastDataNext12HoursDetails {
            probability_of_precipitation,
        }
    }
}


