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
pub struct DashboardDashboardApiGetWidgets200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Vec<crate::models::DashboardWidget>,
}

impl DashboardDashboardApiGetWidgets200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: Vec<crate::models::DashboardWidget>) -> DashboardDashboardApiGetWidgets200ResponseOcs {
        DashboardDashboardApiGetWidgets200ResponseOcs {
            meta: Box::new(meta),
            data,
        }
    }
}


