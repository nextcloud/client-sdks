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
pub struct FilesSharingShareeCircle {
    #[serde(rename = "count", deserialize_with = "Option::deserialize")]
    pub count: Option<i64>,
    #[serde(rename = "label")]
    pub label: String,
    #[serde(rename = "shareWithDescription")]
    pub share_with_description: String,
    #[serde(rename = "value")]
    pub value: Box<crate::models::FilesSharingShareeCircleAllOfValue>,
}

impl FilesSharingShareeCircle {
    pub fn new(count: Option<i64>, label: String, share_with_description: String, value: crate::models::FilesSharingShareeCircleAllOfValue) -> FilesSharingShareeCircle {
        FilesSharingShareeCircle {
            count,
            label,
            share_with_description,
            value: Box::new(value),
        }
    }
}


