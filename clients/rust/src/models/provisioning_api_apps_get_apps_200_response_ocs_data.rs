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
pub struct ProvisioningApiAppsGetApps200ResponseOcsData {
    #[serde(rename = "apps")]
    pub apps: Vec<String>,
}

impl ProvisioningApiAppsGetApps200ResponseOcsData {
    pub fn new(apps: Vec<String>) -> ProvisioningApiAppsGetApps200ResponseOcsData {
        ProvisioningApiAppsGetApps200ResponseOcsData {
            apps,
        }
    }
}


