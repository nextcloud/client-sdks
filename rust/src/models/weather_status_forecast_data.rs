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
pub struct WeatherStatusForecastData {
    #[serde(rename = "instant")]
    pub instant: Box<crate::models::WeatherStatusForecastDataInstant>,
    #[serde(rename = "next_12_hours")]
    pub next_12_hours: Box<crate::models::WeatherStatusForecastDataNext12Hours>,
    #[serde(rename = "next_1_hours")]
    pub next_1_hours: Box<crate::models::WeatherStatusForecastDataNext1Hours>,
    #[serde(rename = "next_6_hours")]
    pub next_6_hours: Box<crate::models::WeatherStatusForecastDataNext6Hours>,
}

impl WeatherStatusForecastData {
    pub fn new(instant: crate::models::WeatherStatusForecastDataInstant, next_12_hours: crate::models::WeatherStatusForecastDataNext12Hours, next_1_hours: crate::models::WeatherStatusForecastDataNext1Hours, next_6_hours: crate::models::WeatherStatusForecastDataNext6Hours) -> WeatherStatusForecastData {
        WeatherStatusForecastData {
            instant: Box::new(instant),
            next_12_hours: Box::new(next_12_hours),
            next_1_hours: Box::new(next_1_hours),
            next_6_hours: Box::new(next_6_hours),
        }
    }
}


