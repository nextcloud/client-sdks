import sys

sys.path.append("../../clients/python")

import nextcloud_client

configuration = nextcloud_client.Configuration(
    host="http://localhost",
    username="admin",
    password="admin",
)

with nextcloud_client.ApiClient(configuration) as api_client:
    instance = nextcloud_client.ProvisioningApiUsersApi(api_client)

    response = instance.provisioning_api_users_get_current_user(ocs_api_request="true")
    print(response.ocs.data.display_name)  # admin
