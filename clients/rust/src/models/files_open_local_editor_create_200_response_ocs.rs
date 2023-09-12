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
pub struct FilesOpenLocalEditorCreate200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Box<crate::models::FilesOpenLocalEditorCreate200ResponseOcsData>,
}

impl FilesOpenLocalEditorCreate200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: crate::models::FilesOpenLocalEditorCreate200ResponseOcsData) -> FilesOpenLocalEditorCreate200ResponseOcs {
        FilesOpenLocalEditorCreate200ResponseOcs {
            meta: Box::new(meta),
            data: Box::new(data),
        }
    }
}


