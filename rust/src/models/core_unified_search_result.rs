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
pub struct CoreUnifiedSearchResult {
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "isPaginated")]
    pub is_paginated: bool,
    #[serde(rename = "entries")]
    pub entries: Vec<crate::models::CoreUnifiedSearchResultEntry>,
    #[serde(rename = "cursor", deserialize_with = "Option::deserialize")]
    pub cursor: Option<Box<crate::models::CoreUnifiedSearchResultCursor>>,
}

impl CoreUnifiedSearchResult {
    pub fn new(name: String, is_paginated: bool, entries: Vec<crate::models::CoreUnifiedSearchResultEntry>, cursor: Option<crate::models::CoreUnifiedSearchResultCursor>) -> CoreUnifiedSearchResult {
        CoreUnifiedSearchResult {
            name,
            is_paginated,
            entries,
            cursor: if let Some(x) = cursor {Some(Box::new(x))} else {None},
        }
    }
}

