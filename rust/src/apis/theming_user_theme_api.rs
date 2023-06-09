/*
 * nextcloud
 *
 * Nextcloud APIs
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */


use reqwest;

use crate::apis::ResponseContent;
use super::{Error, configuration};


/// struct for typed successes of method [`theming_user_theme_delete_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeDeleteBackgroundSuccess {
    Status200(crate::models::ThemingBackground),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`theming_user_theme_disable_theme`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeDisableThemeSuccess {
    Status200(crate::models::CoreWhatsNewDismiss200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`theming_user_theme_enable_theme`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeEnableThemeSuccess {
    Status200(crate::models::CoreWhatsNewDismiss200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`theming_user_theme_get_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeGetBackgroundSuccess {
    Status200(std::path::PathBuf),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`theming_user_theme_set_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeSetBackgroundSuccess {
    Status200(crate::models::ThemingBackground),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`theming_user_theme_delete_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeDeleteBackgroundError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`theming_user_theme_disable_theme`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeDisableThemeError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`theming_user_theme_enable_theme`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeEnableThemeError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`theming_user_theme_get_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeGetBackgroundError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`theming_user_theme_set_background`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum ThemingUserThemeSetBackgroundError {
    UnknownValue(serde_json::Value),
}


pub async fn theming_user_theme_delete_background(configuration: &configuration::Configuration, ocs_api_request: &str) -> Result<ResponseContent<ThemingUserThemeDeleteBackgroundSuccess>, Error<ThemingUserThemeDeleteBackgroundError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/index.php/apps/theming/background/custom", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.header("OCS-APIRequest", ocs_api_request.to_string());
    if let Some(ref local_var_auth_conf) = local_var_configuration.basic_auth {
        local_var_req_builder = local_var_req_builder.basic_auth(local_var_auth_conf.0.to_owned(), local_var_auth_conf.1.to_owned());
    };
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        let local_var_entity: Option<ThemingUserThemeDeleteBackgroundSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<ThemingUserThemeDeleteBackgroundError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn theming_user_theme_disable_theme(configuration: &configuration::Configuration, theme_id: &str, ocs_api_request: &str) -> Result<ResponseContent<ThemingUserThemeDisableThemeSuccess>, Error<ThemingUserThemeDisableThemeError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/theming/api/v1/theme/{themeId}", local_var_configuration.base_path, themeId=crate::apis::urlencode(theme_id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::DELETE, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.header("OCS-APIRequest", ocs_api_request.to_string());
    if let Some(ref local_var_auth_conf) = local_var_configuration.basic_auth {
        local_var_req_builder = local_var_req_builder.basic_auth(local_var_auth_conf.0.to_owned(), local_var_auth_conf.1.to_owned());
    };
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        let local_var_entity: Option<ThemingUserThemeDisableThemeSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<ThemingUserThemeDisableThemeError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn theming_user_theme_enable_theme(configuration: &configuration::Configuration, theme_id: &str, ocs_api_request: &str) -> Result<ResponseContent<ThemingUserThemeEnableThemeSuccess>, Error<ThemingUserThemeEnableThemeError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/theming/api/v1/theme/{themeId}/enable", local_var_configuration.base_path, themeId=crate::apis::urlencode(theme_id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.header("OCS-APIRequest", ocs_api_request.to_string());
    if let Some(ref local_var_auth_conf) = local_var_configuration.basic_auth {
        local_var_req_builder = local_var_req_builder.basic_auth(local_var_auth_conf.0.to_owned(), local_var_auth_conf.1.to_owned());
    };
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        let local_var_entity: Option<ThemingUserThemeEnableThemeSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<ThemingUserThemeEnableThemeError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn theming_user_theme_get_background(configuration: &configuration::Configuration, ocs_api_request: &str) -> Result<ResponseContent<ThemingUserThemeGetBackgroundSuccess>, Error<ThemingUserThemeGetBackgroundError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/index.php/apps/theming/background", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.header("OCS-APIRequest", ocs_api_request.to_string());
    if let Some(ref local_var_auth_conf) = local_var_configuration.basic_auth {
        local_var_req_builder = local_var_req_builder.basic_auth(local_var_auth_conf.0.to_owned(), local_var_auth_conf.1.to_owned());
    };
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        let local_var_entity: Option<ThemingUserThemeGetBackgroundSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<ThemingUserThemeGetBackgroundError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn theming_user_theme_set_background(configuration: &configuration::Configuration, r#type: &str, ocs_api_request: &str, value: Option<&str>, color: Option<&str>) -> Result<ResponseContent<ThemingUserThemeSetBackgroundSuccess>, Error<ThemingUserThemeSetBackgroundError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/index.php/apps/theming/background/{type}", local_var_configuration.base_path, type=crate::apis::urlencode(r#type));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = value {
        local_var_req_builder = local_var_req_builder.query(&[("value", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = color {
        local_var_req_builder = local_var_req_builder.query(&[("color", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
    local_var_req_builder = local_var_req_builder.header("OCS-APIRequest", ocs_api_request.to_string());
    if let Some(ref local_var_auth_conf) = local_var_configuration.basic_auth {
        local_var_req_builder = local_var_req_builder.basic_auth(local_var_auth_conf.0.to_owned(), local_var_auth_conf.1.to_owned());
    };
    if let Some(ref local_var_token) = local_var_configuration.bearer_access_token {
        local_var_req_builder = local_var_req_builder.bearer_auth(local_var_token.to_owned());
    };

    let local_var_req = local_var_req_builder.build()?;
    let local_var_resp = local_var_client.execute(local_var_req).await?;

    let local_var_status = local_var_resp.status();
    let local_var_content = local_var_resp.text().await?;

    if !local_var_status.is_client_error() && !local_var_status.is_server_error() {
        let local_var_entity: Option<ThemingUserThemeSetBackgroundSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<ThemingUserThemeSetBackgroundError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

