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
pub struct FilesTemplatePath200ResponseOcsData {
    #[serde(rename = "template_path")]
    pub template_path: String,
    #[serde(rename = "templates")]
    pub templates: Vec<crate::models::FilesTemplateFileCreator>,
}

impl FilesTemplatePath200ResponseOcsData {
    pub fn new(template_path: String, templates: Vec<crate::models::FilesTemplateFileCreator>) -> FilesTemplatePath200ResponseOcsData {
        FilesTemplatePath200ResponseOcsData {
            template_path,
            templates,
        }
    }
}

