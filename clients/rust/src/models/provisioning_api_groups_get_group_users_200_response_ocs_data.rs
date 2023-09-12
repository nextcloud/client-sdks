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
pub struct ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
    #[serde(rename = "users")]
    pub users: Vec<String>,
}

impl ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
    pub fn new(users: Vec<String>) -> ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
        ProvisioningApiGroupsGetGroupUsers200ResponseOcsData {
            users,
        }
    }
}

