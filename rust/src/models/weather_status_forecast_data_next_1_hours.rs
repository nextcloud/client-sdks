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
pub struct WeatherStatusForecastDataNext1Hours {
    #[serde(rename = "summary")]
    pub summary: Box<crate::models::WeatherStatusForecastDataNext12HoursSummary>,
    #[serde(rename = "details")]
    pub details: Box<crate::models::WeatherStatusForecastDataNext1HoursDetails>,
}

impl WeatherStatusForecastDataNext1Hours {
    pub fn new(summary: crate::models::WeatherStatusForecastDataNext12HoursSummary, details: crate::models::WeatherStatusForecastDataNext1HoursDetails) -> WeatherStatusForecastDataNext1Hours {
        WeatherStatusForecastDataNext1Hours {
            summary: Box::new(summary),
            details: Box::new(details),
        }
    }
}


