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
pub struct FilesSharingShareeLookupAllOfExtra {
    #[serde(rename = "federationId")]
    pub federation_id: String,
    #[serde(rename = "name")]
    pub name: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "email")]
    pub email: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "address")]
    pub address: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "website")]
    pub website: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "twitter")]
    pub twitter: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "phone")]
    pub phone: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "twitter_signature")]
    pub twitter_signature: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "website_signature")]
    pub website_signature: Box<crate::models::FilesSharingLookup>,
    #[serde(rename = "userid")]
    pub userid: Box<crate::models::FilesSharingLookup>,
}

impl FilesSharingShareeLookupAllOfExtra {
    pub fn new(federation_id: String, name: crate::models::FilesSharingLookup, email: crate::models::FilesSharingLookup, address: crate::models::FilesSharingLookup, website: crate::models::FilesSharingLookup, twitter: crate::models::FilesSharingLookup, phone: crate::models::FilesSharingLookup, twitter_signature: crate::models::FilesSharingLookup, website_signature: crate::models::FilesSharingLookup, userid: crate::models::FilesSharingLookup) -> FilesSharingShareeLookupAllOfExtra {
        FilesSharingShareeLookupAllOfExtra {
            federation_id,
            name: Box::new(name),
            email: Box::new(email),
            address: Box::new(address),
            website: Box::new(website),
            twitter: Box::new(twitter),
            phone: Box::new(phone),
            twitter_signature: Box::new(twitter_signature),
            website_signature: Box::new(website_signature),
            userid: Box::new(userid),
        }
    }
}


