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
pub struct FilesSharingShareesRecommendedResult {
    #[serde(rename = "exact")]
    pub exact: Box<crate::models::FilesSharingShareesRecommendedResultExact>,
    #[serde(rename = "emails")]
    pub emails: Vec<crate::models::FilesSharingShareeEmail>,
    #[serde(rename = "groups")]
    pub groups: Vec<crate::models::FilesSharingSharee>,
    #[serde(rename = "remote_groups")]
    pub remote_groups: Vec<crate::models::FilesSharingShareeRemoteGroup>,
    #[serde(rename = "remotes")]
    pub remotes: Vec<crate::models::FilesSharingShareeRemote>,
    #[serde(rename = "users")]
    pub users: Vec<crate::models::FilesSharingShareeUser>,
}

impl FilesSharingShareesRecommendedResult {
    pub fn new(exact: crate::models::FilesSharingShareesRecommendedResultExact, emails: Vec<crate::models::FilesSharingShareeEmail>, groups: Vec<crate::models::FilesSharingSharee>, remote_groups: Vec<crate::models::FilesSharingShareeRemoteGroup>, remotes: Vec<crate::models::FilesSharingShareeRemote>, users: Vec<crate::models::FilesSharingShareeUser>) -> FilesSharingShareesRecommendedResult {
        FilesSharingShareesRecommendedResult {
            exact: Box::new(exact),
            emails,
            groups,
            remote_groups,
            remotes,
            users,
        }
    }
}

