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
pub struct FilesSharingShareInfo {
    #[serde(rename = "id")]
    pub id: i64,
    #[serde(rename = "parentId")]
    pub parent_id: i64,
    #[serde(rename = "mtime")]
    pub mtime: i64,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "permissions")]
    pub permissions: i64,
    #[serde(rename = "mimetype")]
    pub mimetype: String,
    #[serde(rename = "size")]
    pub size: i64,
    #[serde(rename = "type")]
    pub r#type: String,
    #[serde(rename = "etag")]
    pub etag: String,
    #[serde(rename = "children", deserialize_with = "Option::deserialize")]
    pub children: Option<Vec<::std::collections::HashMap<String, serde_json::Value>>>,
}

impl FilesSharingShareInfo {
    pub fn new(id: i64, parent_id: i64, mtime: i64, name: String, permissions: i64, mimetype: String, size: i64, r#type: String, etag: String, children: Option<Vec<::std::collections::HashMap<String, serde_json::Value>>>) -> FilesSharingShareInfo {
        FilesSharingShareInfo {
            id,
            parent_id,
            mtime,
            name,
            permissions,
            mimetype,
            size,
            r#type,
            etag,
            children,
        }
    }
}


