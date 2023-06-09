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
pub struct WeatherStatusForecastDataNext12Hours {
    #[serde(rename = "summary")]
    pub summary: Box<crate::models::WeatherStatusForecastDataNext12HoursSummary>,
    #[serde(rename = "details")]
    pub details: Box<crate::models::WeatherStatusForecastDataNext12HoursDetails>,
}

impl WeatherStatusForecastDataNext12Hours {
    pub fn new(summary: crate::models::WeatherStatusForecastDataNext12HoursSummary, details: crate::models::WeatherStatusForecastDataNext12HoursDetails) -> WeatherStatusForecastDataNext12Hours {
        WeatherStatusForecastDataNext12Hours {
            summary: Box::new(summary),
            details: Box::new(details),
        }
    }
}


