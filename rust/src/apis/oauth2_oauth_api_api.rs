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


/// struct for typed successes of method [`oauth2_oauth_api_get_token`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum Oauth2OauthApiGetTokenSuccess {
    Status200(crate::models::Oauth2OauthApiGetToken200Response),
    UnknownValue(serde_json::Value),
}

/// struct for typed errors of method [`oauth2_oauth_api_get_token`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum Oauth2OauthApiGetTokenError {
    UnknownValue(serde_json::Value),
}


pub async fn oauth2_oauth_api_get_token(configuration: &configuration::Configuration, grant_type: &str, code: &str, refresh_token: &str, client_id: &str, client_secret: &str) -> Result<ResponseContent<Oauth2OauthApiGetTokenSuccess>, Error<Oauth2OauthApiGetTokenError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/index.php/apps/oauth2/api/v1/token", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::POST, local_var_uri_str.as_str());

    local_var_req_builder = local_var_req_builder.query(&[("grant_type", &grant_type.to_string())]);
    local_var_req_builder = local_var_req_builder.query(&[("code", &code.to_string())]);
    local_var_req_builder = local_var_req_builder.query(&[("refresh_token", &refresh_token.to_string())]);
    local_var_req_builder = local_var_req_builder.query(&[("client_id", &client_id.to_string())]);
    local_var_req_builder = local_var_req_builder.query(&[("client_secret", &client_secret.to_string())]);
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
        let local_var_entity: Option<Oauth2OauthApiGetTokenSuccess> = serde_json::from_str(&local_var_content).ok();
        let local_var_result = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Ok(local_var_result)
    } else {
        let local_var_entity: Option<Oauth2OauthApiGetTokenError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}

