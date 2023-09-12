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
pub struct UserStatusStatusesFindAll200ResponseOcs {
    #[serde(rename = "meta")]
    pub meta: Box<crate::models::OcsMeta>,
    #[serde(rename = "data")]
    pub data: Vec<crate::models::UserStatusPublic>,
}

impl UserStatusStatusesFindAll200ResponseOcs {
    pub fn new(meta: crate::models::OcsMeta, data: Vec<crate::models::UserStatusPublic>) -> UserStatusStatusesFindAll200ResponseOcs {
        UserStatusStatusesFindAll200ResponseOcs {
            meta: Box::new(meta),
            data,
        }
    }
}


