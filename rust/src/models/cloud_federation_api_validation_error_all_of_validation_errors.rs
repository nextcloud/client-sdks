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
pub struct CloudFederationApiValidationErrorAllOfValidationErrors {
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "message", deserialize_with = "Option::deserialize")]
    pub message: Option<String>,
}

impl CloudFederationApiValidationErrorAllOfValidationErrors {
    pub fn new(name: String, message: Option<String>) -> CloudFederationApiValidationErrorAllOfValidationErrors {
        CloudFederationApiValidationErrorAllOfValidationErrors {
            name,
            message,
        }
    }
}


