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
pub struct WeatherStatusForecastDataNext12HoursSummary {
    #[serde(rename = "symbol_code")]
    pub symbol_code: String,
}

impl WeatherStatusForecastDataNext12HoursSummary {
    pub fn new(symbol_code: String) -> WeatherStatusForecastDataNext12HoursSummary {
        WeatherStatusForecastDataNext12HoursSummary {
            symbol_code,
        }
    }
}


