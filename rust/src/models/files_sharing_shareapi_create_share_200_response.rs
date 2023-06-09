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
pub struct FilesSharingShareapiCreateShare200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::FilesSharingShareapiCreateShare200ResponseOcs>,
}

impl FilesSharingShareapiCreateShare200Response {
    pub fn new(ocs: crate::models::FilesSharingShareapiCreateShare200ResponseOcs) -> FilesSharingShareapiCreateShare200Response {
        FilesSharingShareapiCreateShare200Response {
            ocs: Box::new(ocs),
        }
    }
}


