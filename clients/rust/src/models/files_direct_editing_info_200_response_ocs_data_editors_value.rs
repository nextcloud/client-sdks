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
pub struct FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "mimetypes")]
    pub mimetypes: Vec<String>,
    #[serde(rename = "optionalMimetypes")]
    pub optional_mimetypes: Vec<String>,
    #[serde(rename = "secure")]
    pub secure: bool,
}

impl FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
    pub fn new(id: String, name: String, mimetypes: Vec<String>, optional_mimetypes: Vec<String>, secure: bool) -> FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
        FilesDirectEditingInfo200ResponseOcsDataEditorsValue {
            id,
            name,
            mimetypes,
            optional_mimetypes,
            secure,
        }
    }
}

