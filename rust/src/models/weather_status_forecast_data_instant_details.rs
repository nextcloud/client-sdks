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
pub struct WeatherStatusForecastDataInstantDetails {
    #[serde(rename = "air_pressure_at_sea_level")]
    pub air_pressure_at_sea_level: f32,
    #[serde(rename = "air_temperature")]
    pub air_temperature: f32,
    #[serde(rename = "cloud_area_fraction")]
    pub cloud_area_fraction: f32,
    #[serde(rename = "cloud_area_fraction_high")]
    pub cloud_area_fraction_high: f32,
    #[serde(rename = "cloud_area_fraction_low")]
    pub cloud_area_fraction_low: f32,
    #[serde(rename = "cloud_area_fraction_medium")]
    pub cloud_area_fraction_medium: f32,
    #[serde(rename = "dew_point_temperature")]
    pub dew_point_temperature: f32,
    #[serde(rename = "fog_area_fraction")]
    pub fog_area_fraction: f32,
    #[serde(rename = "relative_humidity")]
    pub relative_humidity: f32,
    #[serde(rename = "ultraviolet_index_clear_sky")]
    pub ultraviolet_index_clear_sky: f32,
    #[serde(rename = "wind_from_direction")]
    pub wind_from_direction: f32,
    #[serde(rename = "wind_speed")]
    pub wind_speed: f32,
    #[serde(rename = "wind_speed_of_gust")]
    pub wind_speed_of_gust: f32,
}

impl WeatherStatusForecastDataInstantDetails {
    pub fn new(air_pressure_at_sea_level: f32, air_temperature: f32, cloud_area_fraction: f32, cloud_area_fraction_high: f32, cloud_area_fraction_low: f32, cloud_area_fraction_medium: f32, dew_point_temperature: f32, fog_area_fraction: f32, relative_humidity: f32, ultraviolet_index_clear_sky: f32, wind_from_direction: f32, wind_speed: f32, wind_speed_of_gust: f32) -> WeatherStatusForecastDataInstantDetails {
        WeatherStatusForecastDataInstantDetails {
            air_pressure_at_sea_level,
            air_temperature,
            cloud_area_fraction,
            cloud_area_fraction_high,
            cloud_area_fraction_low,
            cloud_area_fraction_medium,
            dew_point_temperature,
            fog_area_fraction,
            relative_humidity,
            ultraviolet_index_clear_sky,
            wind_from_direction,
            wind_speed,
            wind_speed_of_gust,
        }
    }
}


