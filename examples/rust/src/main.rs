use reqwest;
use reqwest::header;

use openapi::apis::configuration::{Configuration};
use openapi::apis::provisioning_api_users_api::{provisioning_api_users_get_current_user, ProvisioningApiUsersGetCurrentUserSuccess};

#[tokio::main]
async fn main() {
    let mut headers = header::HeaderMap::new();
    headers.insert(header::ACCEPT, header::HeaderValue::from_str("application/json").unwrap());
    let client = reqwest::ClientBuilder::new().default_headers(headers).build().unwrap();

    let configuration = Configuration {
        base_path: "http://localhost".to_string(),
        client,
        basic_auth: Option::from(("admin".to_string(), Option::from("admin".to_string()))),
        ..Default::default()
    };

    let response = provisioning_api_users_get_current_user(&configuration, "true").await;
    match response.unwrap().entity.unwrap() {
        ProvisioningApiUsersGetCurrentUserSuccess::Status200(response) => {
            println!("{}", response.ocs.data.displayname);
        }
        _ => {}
    }
}
