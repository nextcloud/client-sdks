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
pub struct FilesSharingRemoteGetShare200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Box<crate::models::FilesSharingRemoteShare>,
}

impl FilesSharingRemoteGetShare200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: crate::models::FilesSharingRemoteShare) -> FilesSharingRemoteGetShare200ResponseOcs {
        FilesSharingRemoteGetShare200ResponseOcs {
            meta: Box::new(meta),
            data: Box::new(data),
        }
    }
}


