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
pub struct ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf {
    #[serde(rename = "id")]
    pub id: String,
}

impl ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf {
    pub fn new(id: String) -> ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf {
        ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf {
            id,
        }
    }
}

