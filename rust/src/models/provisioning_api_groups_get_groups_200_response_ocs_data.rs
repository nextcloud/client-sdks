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
pub struct ProvisioningApiGroupsGetGroups200ResponseOcsData {
    #[serde(rename = "groups")]
    pub groups: Vec<String>,
}

impl ProvisioningApiGroupsGetGroups200ResponseOcsData {
    pub fn new(groups: Vec<String>) -> ProvisioningApiGroupsGetGroups200ResponseOcsData {
        ProvisioningApiGroupsGetGroups200ResponseOcsData {
            groups,
        }
    }
}


