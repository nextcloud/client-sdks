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
pub struct FilesSharingShareesSearchResult {
    #[serde(rename = "exact")]
    pub exact: Box<crate::models::FilesSharingShareesSearchResultExact>,
    #[serde(rename = "circles")]
    pub circles: Vec<crate::models::FilesSharingShareeCircle>,
    #[serde(rename = "emails")]
    pub emails: Vec<crate::models::FilesSharingShareeEmail>,
    #[serde(rename = "groups")]
    pub groups: Vec<crate::models::FilesSharingSharee>,
    #[serde(rename = "lookup")]
    pub lookup: Vec<crate::models::FilesSharingShareeLookup>,
    #[serde(rename = "remote_groups")]
    pub remote_groups: Vec<crate::models::FilesSharingShareeRemoteGroup>,
    #[serde(rename = "remotes")]
    pub remotes: Vec<crate::models::FilesSharingShareeRemote>,
    #[serde(rename = "rooms")]
    pub rooms: Vec<crate::models::FilesSharingSharee>,
    #[serde(rename = "users")]
    pub users: Vec<crate::models::FilesSharingShareeUser>,
    #[serde(rename = "lookupEnabled")]
    pub lookup_enabled: bool,
}

impl FilesSharingShareesSearchResult {
    pub fn new(exact: crate::models::FilesSharingShareesSearchResultExact, circles: Vec<crate::models::FilesSharingShareeCircle>, emails: Vec<crate::models::FilesSharingShareeEmail>, groups: Vec<crate::models::FilesSharingSharee>, lookup: Vec<crate::models::FilesSharingShareeLookup>, remote_groups: Vec<crate::models::FilesSharingShareeRemoteGroup>, remotes: Vec<crate::models::FilesSharingShareeRemote>, rooms: Vec<crate::models::FilesSharingSharee>, users: Vec<crate::models::FilesSharingShareeUser>, lookup_enabled: bool) -> FilesSharingShareesSearchResult {
        FilesSharingShareesSearchResult {
            exact: Box::new(exact),
            circles,
            emails,
            groups,
            lookup,
            remote_groups,
            remotes,
            rooms,
            users,
            lookup_enabled,
        }
    }
}


