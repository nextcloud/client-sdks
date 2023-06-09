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
pub struct CoreUnifiedSearchProvider {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "order")]
    pub order: i64,
}

impl CoreUnifiedSearchProvider {
    pub fn new(id: String, name: String, order: i64) -> CoreUnifiedSearchProvider {
        CoreUnifiedSearchProvider {
            id,
            name,
            order,
        }
    }
}


