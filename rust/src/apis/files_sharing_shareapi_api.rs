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


/// struct for typed successes of method [`files_sharing_shareapi_accept_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiAcceptShareSuccess {
    Status200(crate::models::CoreWhatsNewDismiss200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_create_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiCreateShareSuccess {
    Status200(crate::models::FilesSharingShareapiCreateShare200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_delete_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiDeleteShareSuccess {
    Status200(crate::models::CoreWhatsNewDismiss200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_get_inherited_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetInheritedSharesSuccess {
    Status200(crate::models::FilesSharingShareapiGetShares200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_get_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetShareSuccess {
    Status200(crate::models::FilesSharingShareapiCreateShare200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_get_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetSharesSuccess {
    Status200(crate::models::FilesSharingShareapiGetShares200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_pending_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiPendingSharesSuccess {
    Status200(crate::models::FilesSharingShareapiGetShares200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed successes of method [`files_sharing_shareapi_update_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiUpdateShareSuccess {
    Status200(crate::models::FilesSharingShareapiCreateShare200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_accept_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiAcceptShareError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_create_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiCreateShareError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_delete_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiDeleteShareError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_get_inherited_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetInheritedSharesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_get_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetShareError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_get_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiGetSharesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_pending_shares`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiPendingSharesError {
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_shareapi_update_share`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareapiUpdateShareError {
    UnknownValue(serde_json::Value),
}


pub async fn files_sharing_shareapi_accept_share(configuration: &configuration::Configuration, id: &str, ocs_api_request: &str) -> Result<ResponseContent<FilesSharingShareapiAcceptShareSuccess>, Error<FilesSharingShareapiAcceptShareError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/pending/{id}", local_var_configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

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
        let local_var_entity: Option<FilesSharingShareapiAcceptShareSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiAcceptShareError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_create_share(configuration: &configuration::Configuration, ocs_api_request: &str, path: Option<&str>, permissions: Option<i64>, share_type: Option<i64>, share_with: Option<&str>, public_upload: Option<&str>, password: Option<&str>, send_password_by_talk: Option<&str>, expire_date: Option<&str>, note: Option<&str>, label: Option<&str>, attributes: Option<&str>) -> Result<ResponseContent<FilesSharingShareapiCreateShareSuccess>, Error<FilesSharingShareapiCreateShareError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = path {
        local_var_req_builder = local_var_req_builder.query(&[("path", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = permissions {
        local_var_req_builder = local_var_req_builder.query(&[("permissions", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = share_type {
        local_var_req_builder = local_var_req_builder.query(&[("shareType", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = share_with {
        local_var_req_builder = local_var_req_builder.query(&[("shareWith", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = public_upload {
        local_var_req_builder = local_var_req_builder.query(&[("publicUpload", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = password {
        local_var_req_builder = local_var_req_builder.query(&[("password", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = send_password_by_talk {
        local_var_req_builder = local_var_req_builder.query(&[("sendPasswordByTalk", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = expire_date {
        local_var_req_builder = local_var_req_builder.query(&[("expireDate", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = note {
        local_var_req_builder = local_var_req_builder.query(&[("note", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = label {
        local_var_req_builder = local_var_req_builder.query(&[("label", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = attributes {
        local_var_req_builder = local_var_req_builder.query(&[("attributes", &local_var_str.to_string())]);
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
        let local_var_entity: Option<FilesSharingShareapiCreateShareSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiCreateShareError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_delete_share(configuration: &configuration::Configuration, id: &str, ocs_api_request: &str) -> Result<ResponseContent<FilesSharingShareapiDeleteShareSuccess>, Error<FilesSharingShareapiDeleteShareError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/{id}", local_var_configuration.base_path, id=crate::apis::urlencode(id));
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
        let local_var_entity: Option<FilesSharingShareapiDeleteShareSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiDeleteShareError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_get_inherited_shares(configuration: &configuration::Configuration, path: &str, ocs_api_request: &str) -> Result<ResponseContent<FilesSharingShareapiGetInheritedSharesSuccess>, Error<FilesSharingShareapiGetInheritedSharesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/inherited", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    local_var_req_builder = local_var_req_builder.query(&[("path", &path.to_string())]);
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
        let local_var_entity: Option<FilesSharingShareapiGetInheritedSharesSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiGetInheritedSharesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_get_share(configuration: &configuration::Configuration, id: &str, ocs_api_request: &str, include_tags: Option<i32>) -> Result<ResponseContent<FilesSharingShareapiGetShareSuccess>, Error<FilesSharingShareapiGetShareError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/{id}", local_var_configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = include_tags {
        local_var_req_builder = local_var_req_builder.query(&[("include_tags", &local_var_str.to_string())]);
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
        let local_var_entity: Option<FilesSharingShareapiGetShareSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiGetShareError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_get_shares(configuration: &configuration::Configuration, ocs_api_request: &str, shared_with_me: Option<&str>, reshares: Option<&str>, subfiles: Option<&str>, path: Option<&str>, include_tags: Option<&str>) -> Result<ResponseContent<FilesSharingShareapiGetSharesSuccess>, Error<FilesSharingShareapiGetSharesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = shared_with_me {
        local_var_req_builder = local_var_req_builder.query(&[("shared_with_me", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = reshares {
        local_var_req_builder = local_var_req_builder.query(&[("reshares", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = subfiles {
        local_var_req_builder = local_var_req_builder.query(&[("subfiles", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = path {
        local_var_req_builder = local_var_req_builder.query(&[("path", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = include_tags {
        local_var_req_builder = local_var_req_builder.query(&[("include_tags", &local_var_str.to_string())]);
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
        let local_var_entity: Option<FilesSharingShareapiGetSharesSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiGetSharesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_pending_shares(configuration: &configuration::Configuration, ocs_api_request: &str) -> Result<ResponseContent<FilesSharingShareapiPendingSharesSuccess>, Error<FilesSharingShareapiPendingSharesError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/pending", local_var_configuration.base_path);
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
        let local_var_entity: Option<FilesSharingShareapiPendingSharesSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiPendingSharesError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

pub async fn files_sharing_shareapi_update_share(configuration: &configuration::Configuration, id: &str, ocs_api_request: &str, permissions: Option<i64>, password: Option<&str>, send_password_by_talk: Option<&str>, public_upload: Option<&str>, expire_date: Option<&str>, note: Option<&str>, label: Option<&str>, hide_download: Option<&str>, attributes: Option<&str>) -> Result<ResponseContent<FilesSharingShareapiUpdateShareSuccess>, Error<FilesSharingShareapiUpdateShareError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/apps/files_sharing/api/v1/shares/{id}", local_var_configuration.base_path, id=crate::apis::urlencode(id));
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::PUT, local_var_uri_str.as_str());

    if let Some(ref local_var_str) = permissions {
        local_var_req_builder = local_var_req_builder.query(&[("permissions", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = password {
        local_var_req_builder = local_var_req_builder.query(&[("password", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = send_password_by_talk {
        local_var_req_builder = local_var_req_builder.query(&[("sendPasswordByTalk", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = public_upload {
        local_var_req_builder = local_var_req_builder.query(&[("publicUpload", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = expire_date {
        local_var_req_builder = local_var_req_builder.query(&[("expireDate", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = note {
        local_var_req_builder = local_var_req_builder.query(&[("note", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = label {
        local_var_req_builder = local_var_req_builder.query(&[("label", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = hide_download {
        local_var_req_builder = local_var_req_builder.query(&[("hideDownload", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = attributes {
        local_var_req_builder = local_var_req_builder.query(&[("attributes", &local_var_str.to_string())]);
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
        let local_var_entity: Option<FilesSharingShareapiUpdateShareSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareapiUpdateShareError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

