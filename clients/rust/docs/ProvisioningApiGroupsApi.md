# \ProvisioningApiGroupsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_groups_add_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_add_group) | **POST** /ocs/v2.php/cloud/groups | Create a new group
[**provisioning_api_groups_delete_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_delete_group) | **DELETE** /ocs/v2.php/cloud/groups/{groupId} | Delete a group
[**provisioning_api_groups_get_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group) | **GET** /ocs/v2.php/cloud/groups/{groupId} | Get a list of users in the specified group
[**provisioning_api_groups_get_group_users**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group_users) | **GET** /ocs/v2.php/cloud/groups/{groupId}/users | Get a list of users in the specified group
[**provisioning_api_groups_get_group_users_details**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_group_users_details) | **GET** /ocs/v2.php/cloud/groups/{groupId}/users/details | Get a list of users details in the specified group
[**provisioning_api_groups_get_groups**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_groups) | **GET** /ocs/v2.php/cloud/groups | Get a list of groups
[**provisioning_api_groups_get_groups_details**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_groups_details) | **GET** /ocs/v2.php/cloud/groups/details | Get a list of groups details
[**provisioning_api_groups_get_sub_admins_of_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_get_sub_admins_of_group) | **GET** /ocs/v2.php/cloud/groups/{groupId}/subadmins | Get the list of user IDs that are a subadmin of the group
[**provisioning_api_groups_update_group**](ProvisioningApiGroupsApi.md#provisioning_api_groups_update_group) | **PUT** /ocs/v2.php/cloud/groups/{groupId} | Update a group



## provisioning_api_groups_add_group

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_groups_add_group(groupid, ocs_api_request, displayname)
Create a new group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**groupid** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**displayname** | Option<**String**> | Display name of the group |  |[default to ]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_delete_group

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_groups_delete_group(group_id, ocs_api_request)
Delete a group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_group

> crate::models::ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_groups_get_group(group_id, ocs_api_request)
Get a list of users in the specified group

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroupUsers200Response**](provisioning_api_groups_get_group_users_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_group_users

> crate::models::ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_groups_get_group_users(group_id, ocs_api_request)
Get a list of users in the specified group

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroupUsers200Response**](provisioning_api_groups_get_group_users_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_group_users_details

> crate::models::ProvisioningApiGroupsGetGroupUsersDetails200Response provisioning_api_groups_get_group_users_details(group_id, ocs_api_request, search, limit, offset)
Get a list of users details in the specified group

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**search** | Option<**String**> | Text to search for |  |[default to ]
**limit** | Option<**i64**> | Limit the amount of groups returned |  |
**offset** | Option<**i64**> | Offset for searching for groups |  |[default to 0]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroupUsersDetails200Response**](provisioning_api_groups_get_group_users_details_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_groups

> crate::models::ProvisioningApiGroupsGetGroups200Response provisioning_api_groups_get_groups(ocs_api_request, search, limit, offset)
Get a list of groups

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**search** | Option<**String**> | Text to search for |  |[default to ]
**limit** | Option<**i64**> | Limit the amount of groups returned |  |
**offset** | Option<**i64**> | Offset for searching for groups |  |[default to 0]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroups200Response**](provisioning_api_groups_get_groups_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_groups_details

> crate::models::ProvisioningApiGroupsGetGroupsDetails200Response provisioning_api_groups_get_groups_details(ocs_api_request, search, limit, offset)
Get a list of groups details

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**search** | Option<**String**> | Text to search for |  |[default to ]
**limit** | Option<**i64**> | Limit the amount of groups returned |  |
**offset** | Option<**i64**> | Offset for searching for groups |  |[default to 0]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroupsDetails200Response**](provisioning_api_groups_get_groups_details_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_get_sub_admins_of_group

> crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_groups_get_sub_admins_of_group(group_id, ocs_api_request)
Get the list of user IDs that are a subadmin of the group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](provisioning_api_groups_get_sub_admins_of_group_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_groups_update_group

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_groups_update_group(key, value, group_id, ocs_api_request)
Update a group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**key** | **String** | Key to update, only 'displayname' | [required] |
**value** | **String** | New value for the key | [required] |
**group_id** | **String** | ID of the group | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

