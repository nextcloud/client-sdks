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


/// struct for typed errors of method [`core_auto_complete_get`]
#[derive(Debug, Clone, Serialize, Deserialize)]
#[serde(untagged)]
pub enum CoreAutoCompleteGetError {
    UnknownValue(serde_json::Value),
}


pub async fn core_auto_complete_get(configuration: &configuration::Configuration, search: &str, ocs_api_request: &str, item_type: Option<&str>, item_id: Option<&str>, sorter: Option<&str>, share_types_left_square_bracket_right_square_bracket: Option<Vec<i64>>, limit: Option<i64>) -> Result<crate::models::CoreAutoCompleteGet200Response, Error<CoreAutoCompleteGetError>> {
    let local_var_configuration = configuration;

    let local_var_client = &local_var_configuration.client;

    let local_var_uri_str = format!("{}/ocs/v2.php/core/autocomplete/get", local_var_configuration.base_path);
    let mut local_var_req_builder = local_var_client.request(reqwest::Method::GET, local_var_uri_str.as_str());

    local_var_req_builder = local_var_req_builder.query(&[("search", &search.to_string())]);
    if let Some(ref local_var_str) = item_type {
        local_var_req_builder = local_var_req_builder.query(&[("itemType", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = item_id {
        local_var_req_builder = local_var_req_builder.query(&[("itemId", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = sorter {
        local_var_req_builder = local_var_req_builder.query(&[("sorter", &local_var_str.to_string())]);
    }
    if let Some(ref local_var_str) = share_types_left_square_bracket_right_square_bracket {
        local_var_req_builder = match "multi" {
            "multi" => local_var_req_builder.query(&local_var_str.into_iter().map(|p| ("shareTypes[]".to_owned(), p.to_string())).collect::<Vec<(std::string::String, std::string::String)>>()),
            _ => local_var_req_builder.query(&[("shareTypes[]", &local_var_str.into_iter().map(|p| p.to_string()).collect::<Vec<String>>().join(",").to_string())]),
        };
    }
    if let Some(ref local_var_str) = limit {
        local_var_req_builder = local_var_req_builder.query(&[("limit", &local_var_str.to_string())]);
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
        serde_json::from_str(&local_var_content).map_err(Error::from)
    } else {
        let local_var_entity: Option<CoreAutoCompleteGetError> = serde_json::from_str(&local_var_content).ok();
        let local_var_error = ResponseContent { status: local_var_status, content: local_var_content, entity: local_var_entity };
        Err(Error::ResponseError(local_var_error))
    }
}
