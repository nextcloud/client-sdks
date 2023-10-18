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
pub struct CloudFederationApiValidationError {
    #[serde(rename = "message")]
    pub message: String,
    #[serde(rename = "validationErrors")]
    pub validation_errors: Vec<crate::models::CloudFederationApiValidationErrorAllOfValidationErrors>,
}

impl CloudFederationApiValidationError {
    pub fn new(message: String, validation_errors: Vec<crate::models::CloudFederationApiValidationErrorAllOfValidationErrors>) -> CloudFederationApiValidationError {
        CloudFederationApiValidationError {
            message,
            validation_errors,
        }
    }
}

