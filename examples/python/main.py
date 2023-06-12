import sys

sys.path.append("../../python")

import openapi_client

configuration = openapi_client.Configuration(
    host="http://localhost",
    username="admin",
    password="admin",
)

with openapi_client.ApiClient(configuration) as api_client:
    instance = openapi_client.ProvisioningApiUsersApi(api_client)

    response = instance.provisioning_api_users_get_current_user(ocs_api_request="true")
    print(response.ocs.data.display_name)  # admin
