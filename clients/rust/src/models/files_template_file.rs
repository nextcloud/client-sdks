/*
 * nextcloud
 *
 * Nextcloud APIs
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct FilesTemplateFile {
    #[serde(rename = "basename")]
    pub basename: String,
    #[serde(rename = "etag")]
    pub etag: String,
    #[serde(rename = "fileid")]
    pub fileid: i64,
    #[serde(rename = "filename", deserialize_with = "Option::deserialize")]
    pub filename: Option<String>,
    #[serde(rename = "lastmod")]
    pub lastmod: i64,
    #[serde(rename = "mime")]
    pub mime: String,
    #[serde(rename = "size")]
    pub size: i64,
    #[serde(rename = "type")]
    pub r#type: String,
    #[serde(rename = "hasPreview")]
    pub has_preview: bool,
}

impl FilesTemplateFile {
    pub fn new(basename: String, etag: String, fileid: i64, filename: Option<String>, lastmod: i64, mime: String, size: i64, r#type: String, has_preview: bool) -> FilesTemplateFile {
        FilesTemplateFile {
            basename,
            etag,
            fileid,
            filename,
            lastmod,
            mime,
            size,
            r#type,
            has_preview,
        }
    }
}

