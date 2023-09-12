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
pub struct FilesTemplatePath200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::FilesTemplatePath200ResponseOcs>,
}

impl FilesTemplatePath200Response {
    pub fn new(ocs: crate::models::FilesTemplatePath200ResponseOcs) -> FilesTemplatePath200Response {
        FilesTemplatePath200Response {
            ocs: Box::new(ocs),
        }
    }
}


