import * as nextcloud from "nextcloud-api-sdk";

const configurationParameters = {
    username: "admin",
    password: "admin",
}
const config = new nextcloud.Configuration(configurationParameters);
const api = new nextcloud.ProvisioningApiUsersApi(config);

const response = await api.provisioningApiUsersGetCurrentUser("true");
console.log(response.data.ocs.data.id); // admin
