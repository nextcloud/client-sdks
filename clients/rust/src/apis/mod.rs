use std::error;
use std::fmt;

#[derive(Debug, Clone)]
pub struct ResponseContent<T> {
    pub status: reqwest::StatusCode,
    pub content: String,
    pub entity: Option<T>,
}

#[derive(Debug)]
pub enum Error<T> {
    Reqwest(reqwest::Error),
    Serde(serde_json::Error),
    Io(std::io::Error),
    ResponseError(ResponseContent<T>),
}

impl <T> fmt::Display for Error<T> {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        let (module, e) = match self {
            Error::Reqwest(e) => ("reqwest", e.to_string()),
            Error::Serde(e) => ("serde", e.to_string()),
            Error::Io(e) => ("IO", e.to_string()),
            Error::ResponseError(e) => ("response", format!("status code {}", e.status)),
        };
        write!(f, "error in {}: {}", module, e)
    }
}

impl <T: fmt::Debug> error::Error for Error<T> {
    fn source(&self) -> Option<&(dyn error::Error + 'static)> {
        Some(match self {
            Error::Reqwest(e) => e,
            Error::Serde(e) => e,
            Error::Io(e) => e,
            Error::ResponseError(_) => return None,
        })
    }
}

impl <T> From<reqwest::Error> for Error<T> {
    fn from(e: reqwest::Error) -> Self {
        Error::Reqwest(e)
    }
}

impl <T> From<serde_json::Error> for Error<T> {
    fn from(e: serde_json::Error) -> Self {
        Error::Serde(e)
    }
}

impl <T> From<std::io::Error> for Error<T> {
    fn from(e: std::io::Error) -> Self {
        Error::Io(e)
    }
}

pub fn urlencode<T: AsRef<str>>(s: T) -> String {
    ::url::form_urlencoded::byte_serialize(s.as_ref().as_bytes()).collect()
}

pub fn parse_deep_object(prefix: &str, value: &serde_json::Value) -> Vec<(String, String)> {
    if let serde_json::Value::Object(object) = value {
        let mut params = vec![];

        for (key, value) in object {
            match value {
                serde_json::Value::Object(_) => params.append(&mut parse_deep_object(
                    &format!("{}[{}]", prefix, key),
                    value,
                )),
                serde_json::Value::Array(array) => {
                    for (i, value) in array.iter().enumerate() {
                        params.append(&mut parse_deep_object(
                            &format!("{}[{}][{}]", prefix, key, i),
                            value,
                        ));
                    }
                },
                serde_json::Value::String(s) => params.push((format!("{}[{}]", prefix, key), s.clone())),
                _ => params.push((format!("{}[{}]", prefix, key), value.to_string())),
            }
        }

        return params;
    }

    unimplemented!("Only objects are supported with style=deepObject")
}

pub mod core_api;
pub mod core_app_password_api;
pub mod core_auto_complete_api;
pub mod core_avatar_api;
pub mod core_client_flow_login_v2_api;
pub mod core_collaboration_resources_api;
pub mod core_guest_avatar_api;
pub mod core_hover_card_api;
pub mod core_navigation_api;
pub mod core_ocs_api;
pub mod core_preview_api;
pub mod core_profile_api_api;
pub mod core_reference_api;
pub mod core_reference_api_api;
pub mod core_text_processing_api_api;
pub mod core_translation_api_api;
pub mod core_unified_search_api;
pub mod core_whats_new_api;
pub mod core_wipe_api;
pub mod dashboard_dashboard_api_api;
pub mod dav_direct_api;
pub mod files_api_api;
pub mod files_direct_editing_api;
pub mod files_external_api_api;
pub mod files_open_local_editor_api;
pub mod files_reminders_api_api;
pub mod files_sharing_deleted_shareapi_api;
pub mod files_sharing_public_preview_api;
pub mod files_sharing_remote_api;
pub mod files_sharing_share_info_api;
pub mod files_sharing_shareapi_api;
pub mod files_sharing_shareesapi_api;
pub mod files_template_api;
pub mod files_transfer_ownership_api;
pub mod files_trashbin_preview_api;
pub mod files_versions_preview_api;
pub mod provisioning_api_app_config_api;
pub mod provisioning_api_apps_api;
pub mod provisioning_api_groups_api;
pub mod provisioning_api_preferences_api;
pub mod provisioning_api_users_api;
pub mod settings_log_settings_api;
pub mod theming_icon_api;
pub mod theming_theming_api;
pub mod theming_user_theme_api;
pub mod updatenotification_api_api;
pub mod user_status_heartbeat_api;
pub mod user_status_predefined_status_api;
pub mod user_status_statuses_api;
pub mod user_status_user_status_api;
pub mod weather_status_weather_status_api;

pub mod configuration;
