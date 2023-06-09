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
pub struct ProvisioningApiGroupDetails {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "displayname")]
    pub displayname: String,
    #[serde(rename = "usercount")]
    pub usercount: Box<crate::models::ProvisioningApiGroupDetailsUsercount>,
    #[serde(rename = "disabled")]
    pub disabled: Box<crate::models::ProvisioningApiGroupDetailsUsercount>,
    #[serde(rename = "canAdd")]
    pub can_add: bool,
    #[serde(rename = "canRemove")]
    pub can_remove: bool,
}

impl ProvisioningApiGroupDetails {
    pub fn new(id: String, displayname: String, usercount: crate::models::ProvisioningApiGroupDetailsUsercount, disabled: crate::models::ProvisioningApiGroupDetailsUsercount, can_add: bool, can_remove: bool) -> ProvisioningApiGroupDetails {
        ProvisioningApiGroupDetails {
            id,
            displayname,
            usercount: Box::new(usercount),
            disabled: Box::new(disabled),
            can_add,
            can_remove,
        }
    }
}


