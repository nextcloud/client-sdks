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
pub struct CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing {
    #[serde(rename = "api_enabled")]
    pub api_enabled: bool,
    #[serde(rename = "public")]
    pub public: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic>,
    #[serde(rename = "user")]
    pub user: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser>,
    #[serde(rename = "resharing")]
    pub resharing: bool,
    #[serde(rename = "group_sharing", skip_serializing_if = "Option::is_none")]
    pub group_sharing: Option<bool>,
    #[serde(rename = "group", skip_serializing_if = "Option::is_none")]
    pub group: Option<Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup>>,
    #[serde(rename = "default_permissions", skip_serializing_if = "Option::is_none")]
    pub default_permissions: Option<i64>,
    #[serde(rename = "federation")]
    pub federation: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation>,
    #[serde(rename = "sharee")]
    pub sharee: Box<crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee>,
}

impl CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing {
    pub fn new(api_enabled: bool, public: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic, user: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser, resharing: bool, federation: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation, sharee: crate::models::CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee) -> CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing {
        CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing {
            api_enabled,
            public: Box::new(public),
            user: Box::new(user),
            resharing,
            group_sharing: None,
            group: None,
            default_permissions: None,
            federation: Box::new(federation),
            sharee: Box::new(sharee),
        }
    }
}


