import * as nextcloud from "@nextcloud/client-sdk";

const configurationParameters = {
    username: "admin",
    password: "admin",
}
const config = new nextcloud.Configuration(configurationParameters);
const api = new nextcloud.ProvisioningApiUsersApi(config);

const response = await api.provisioningApiUsersGetCurrentUser({oCSAPIRequest: "true"});
console.log(response.data.ocs.data.id); // admin
