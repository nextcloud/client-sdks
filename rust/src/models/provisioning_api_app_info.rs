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
pub struct ProvisioningApiAppInfo {
    #[serde(rename = "active", deserialize_with = "Option::deserialize")]
    pub active: Option<bool>,
    #[serde(rename = "activity", deserialize_with = "Option::deserialize")]
    pub activity: Option<serde_json::Value>,
    #[serde(rename = "author", deserialize_with = "Option::deserialize")]
    pub author: Option<serde_json::Value>,
    #[serde(rename = "background-jobs", deserialize_with = "Option::deserialize")]
    pub background_jobs: Option<serde_json::Value>,
    #[serde(rename = "bugs", deserialize_with = "Option::deserialize")]
    pub bugs: Option<serde_json::Value>,
    #[serde(rename = "category", deserialize_with = "Option::deserialize")]
    pub category: Option<serde_json::Value>,
    #[serde(rename = "collaboration", deserialize_with = "Option::deserialize")]
    pub collaboration: Option<serde_json::Value>,
    #[serde(rename = "commands", deserialize_with = "Option::deserialize")]
    pub commands: Option<serde_json::Value>,
    #[serde(rename = "default_enable", deserialize_with = "Option::deserialize")]
    pub default_enable: Option<serde_json::Value>,
    #[serde(rename = "dependencies", deserialize_with = "Option::deserialize")]
    pub dependencies: Option<serde_json::Value>,
    #[serde(rename = "description")]
    pub description: String,
    #[serde(rename = "discussion", deserialize_with = "Option::deserialize")]
    pub discussion: Option<serde_json::Value>,
    #[serde(rename = "documentation", deserialize_with = "Option::deserialize")]
    pub documentation: Option<serde_json::Value>,
    #[serde(rename = "groups", deserialize_with = "Option::deserialize")]
    pub groups: Option<serde_json::Value>,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "info", deserialize_with = "Option::deserialize")]
    pub info: Option<serde_json::Value>,
    #[serde(rename = "internal", deserialize_with = "Option::deserialize")]
    pub internal: Option<bool>,
    #[serde(rename = "level", deserialize_with = "Option::deserialize")]
    pub level: Option<i64>,
    #[serde(rename = "licence", deserialize_with = "Option::deserialize")]
    pub licence: Option<serde_json::Value>,
    #[serde(rename = "name")]
    pub name: String,
    #[serde(rename = "namespace", deserialize_with = "Option::deserialize")]
    pub namespace: Option<serde_json::Value>,
    #[serde(rename = "navigations", deserialize_with = "Option::deserialize")]
    pub navigations: Option<serde_json::Value>,
    #[serde(rename = "preview", deserialize_with = "Option::deserialize")]
    pub preview: Option<serde_json::Value>,
    #[serde(rename = "previewAsIcon", deserialize_with = "Option::deserialize")]
    pub preview_as_icon: Option<bool>,
    #[serde(rename = "public", deserialize_with = "Option::deserialize")]
    pub public: Option<serde_json::Value>,
    #[serde(rename = "remote", deserialize_with = "Option::deserialize")]
    pub remote: Option<serde_json::Value>,
    #[serde(rename = "removable", deserialize_with = "Option::deserialize")]
    pub removable: Option<bool>,
    #[serde(rename = "repair-steps", deserialize_with = "Option::deserialize")]
    pub repair_steps: Option<serde_json::Value>,
    #[serde(rename = "repository", deserialize_with = "Option::deserialize")]
    pub repository: Option<serde_json::Value>,
    #[serde(rename = "sabre", deserialize_with = "Option::deserialize")]
    pub sabre: Option<serde_json::Value>,
    #[serde(rename = "screenshot", deserialize_with = "Option::deserialize")]
    pub screenshot: Option<serde_json::Value>,
    #[serde(rename = "settings", deserialize_with = "Option::deserialize")]
    pub settings: Option<serde_json::Value>,
    #[serde(rename = "summary")]
    pub summary: String,
    #[serde(rename = "trash", deserialize_with = "Option::deserialize")]
    pub trash: Option<serde_json::Value>,
    #[serde(rename = "two-factor-providers", deserialize_with = "Option::deserialize")]
    pub two_factor_providers: Option<serde_json::Value>,
    #[serde(rename = "types", deserialize_with = "Option::deserialize")]
    pub types: Option<serde_json::Value>,
    #[serde(rename = "version")]
    pub version: String,
    #[serde(rename = "versions", deserialize_with = "Option::deserialize")]
    pub versions: Option<serde_json::Value>,
    #[serde(rename = "website", deserialize_with = "Option::deserialize")]
    pub website: Option<serde_json::Value>,
}

impl ProvisioningApiAppInfo {
    pub fn new(active: Option<bool>, activity: Option<serde_json::Value>, author: Option<serde_json::Value>, background_jobs: Option<serde_json::Value>, bugs: Option<serde_json::Value>, category: Option<serde_json::Value>, collaboration: Option<serde_json::Value>, commands: Option<serde_json::Value>, default_enable: Option<serde_json::Value>, dependencies: Option<serde_json::Value>, description: String, discussion: Option<serde_json::Value>, documentation: Option<serde_json::Value>, groups: Option<serde_json::Value>, id: String, info: Option<serde_json::Value>, internal: Option<bool>, level: Option<i64>, licence: Option<serde_json::Value>, name: String, namespace: Option<serde_json::Value>, navigations: Option<serde_json::Value>, preview: Option<serde_json::Value>, preview_as_icon: Option<bool>, public: Option<serde_json::Value>, remote: Option<serde_json::Value>, removable: Option<bool>, repair_steps: Option<serde_json::Value>, repository: Option<serde_json::Value>, sabre: Option<serde_json::Value>, screenshot: Option<serde_json::Value>, settings: Option<serde_json::Value>, summary: String, trash: Option<serde_json::Value>, two_factor_providers: Option<serde_json::Value>, types: Option<serde_json::Value>, version: String, versions: Option<serde_json::Value>, website: Option<serde_json::Value>) -> ProvisioningApiAppInfo {
        ProvisioningApiAppInfo {
            active,
            activity,
            author,
            background_jobs,
            bugs,
            category,
            collaboration,
            commands,
            default_enable,
            dependencies,
            description,
            discussion,
            documentation,
            groups,
            id,
            info,
            internal,
            level,
            licence,
            name,
            namespace,
            navigations,
            preview,
            preview_as_icon,
            public,
            remote,
            removable,
            repair_steps,
            repository,
            sabre,
            screenshot,
            settings,
            summary,
            trash,
            two_factor_providers,
            types,
            version,
            versions,
            website,
        }
    }
}


