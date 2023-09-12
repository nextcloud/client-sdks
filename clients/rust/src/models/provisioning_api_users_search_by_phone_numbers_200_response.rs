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
pub struct ProvisioningApiUsersSearchByPhoneNumbers200Response {
    #[serde(rename = "ocs")]
    pub ocs: Box<crate::models::ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs>,
}

impl ProvisioningApiUsersSearchByPhoneNumbers200Response {
    pub fn new(ocs: crate::models::ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs) -> ProvisioningApiUsersSearchByPhoneNumbers200Response {
        ProvisioningApiUsersSearchByPhoneNumbers200Response {
            ocs: Box::new(ocs),
        }
    }
}


