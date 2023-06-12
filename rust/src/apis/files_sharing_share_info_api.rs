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


/// struct for typed successes of method [`files_sharing_share_info_info`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareInfoInfoSuccess {
    Status200(crate::models::FilesSharingShareInfo),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`files_sharing_share_info_info`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum FilesSharingShareInfoInfoError {
    UnknownValue(serde_json::Value),
}


pub async fn files_sharing_share_info_info(configuration: &configuration::Configuration, t: &str, password: Option<&str>, dir: Option<&str>, depth: Option<i64>) -> Result<ResponseContent<FilesSharingShareInfoInfoSuccess>, Error<FilesSharingShareInfoInfoError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/index.php/apps/files_sharing/shareinfo", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    local_var_req_builder = local_var_req_builder.query(&[("t", &t.to_string())]);
    if let Some(ref local_var_str) = password {
        local_var_req_builder = local_var_req_builder.query(&[("password", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = dir {
        local_var_req_builder = local_var_req_builder.query(&[("dir", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = depth {
        local_var_req_builder = local_var_req_builder.query(&[("depth", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_user_agent) = local_var_configuration.user_agent {
        local_var_req_builder = local_var_req_builder.header(reqwest::header::USER_AGENT, local_var_user_agent.clone());
    }
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
        let local_var_entity: Option<FilesSharingShareInfoInfoSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<FilesSharingShareInfoInfoError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

