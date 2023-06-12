# \ProvisioningApiUsersApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**provisioning_api_users_add_sub_admin**](ProvisioningApiUsersApi.md#provisioning_api_users_add_sub_admin) | **POST** /ocs/v2.php/cloud/users/{userId}/subadmins | Make a user a subadmin of a group
[**provisioning_api_users_add_to_group**](ProvisioningApiUsersApi.md#provisioning_api_users_add_to_group) | **POST** /ocs/v2.php/cloud/users/{userId}/groups | Add a user to a group
[**provisioning_api_users_add_user**](ProvisioningApiUsersApi.md#provisioning_api_users_add_user) | **POST** /ocs/v2.php/cloud/users | Create a new user
[**provisioning_api_users_delete_user**](ProvisioningApiUsersApi.md#provisioning_api_users_delete_user) | **DELETE** /ocs/v2.php/cloud/users/{userId} | Delete a user
[**provisioning_api_users_disable_user**](ProvisioningApiUsersApi.md#provisioning_api_users_disable_user) | **PUT** /ocs/v2.php/cloud/users/{userId}/disable | Disable a user
[**provisioning_api_users_edit_user**](ProvisioningApiUsersApi.md#provisioning_api_users_edit_user) | **PUT** /ocs/v2.php/cloud/users/{userId} | Update a value of the user's details
[**provisioning_api_users_edit_user_multi_value**](ProvisioningApiUsersApi.md#provisioning_api_users_edit_user_multi_value) | **PUT** /ocs/v2.php/cloud/users/{userId}/{collectionName} | Update multiple values of the user's details
[**provisioning_api_users_enable_user**](ProvisioningApiUsersApi.md#provisioning_api_users_enable_user) | **PUT** /ocs/v2.php/cloud/users/{userId}/enable | Enable a user
[**provisioning_api_users_get_current_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_current_user) | **GET** /ocs/v2.php/cloud/user | Get the details of the current user
[**provisioning_api_users_get_editable_fields**](ProvisioningApiUsersApi.md#provisioning_api_users_get_editable_fields) | **GET** /ocs/v2.php/cloud/user/fields | Get a list of fields that are editable for the current user
[**provisioning_api_users_get_editable_fields_for_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_editable_fields_for_user) | **GET** /ocs/v2.php/cloud/user/fields/{userId} | Get a list of fields that are editable for a user
[**provisioning_api_users_get_user**](ProvisioningApiUsersApi.md#provisioning_api_users_get_user) | **GET** /ocs/v2.php/cloud/users/{userId} | Get the details of a user
[**provisioning_api_users_get_user_sub_admin_groups**](ProvisioningApiUsersApi.md#provisioning_api_users_get_user_sub_admin_groups) | **GET** /ocs/v2.php/cloud/users/{userId}/subadmins | Get the groups a user is a subadmin of
[**provisioning_api_users_get_users**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users) | **GET** /ocs/v2.php/cloud/users | Get a list of users
[**provisioning_api_users_get_users_details**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users_details) | **GET** /ocs/v2.php/cloud/users/details | Get a list of users and their details
[**provisioning_api_users_get_users_groups**](ProvisioningApiUsersApi.md#provisioning_api_users_get_users_groups) | **GET** /ocs/v2.php/cloud/users/{userId}/groups | Get a list of groups the user belongs to
[**provisioning_api_users_remove_from_group**](ProvisioningApiUsersApi.md#provisioning_api_users_remove_from_group) | **DELETE** /ocs/v2.php/cloud/users/{userId}/groups | Remove a user from a group
[**provisioning_api_users_remove_sub_admin**](ProvisioningApiUsersApi.md#provisioning_api_users_remove_sub_admin) | **DELETE** /ocs/v2.php/cloud/users/{userId}/subadmins | Remove a user from the subadmins of a group
[**provisioning_api_users_resend_welcome_message**](ProvisioningApiUsersApi.md#provisioning_api_users_resend_welcome_message) | **POST** /ocs/v2.php/cloud/users/{userId}/welcome | Resend the welcome message
[**provisioning_api_users_search_by_phone_numbers**](ProvisioningApiUsersApi.md#provisioning_api_users_search_by_phone_numbers) | **POST** /ocs/v2.php/cloud/users/search/by-phone | Search users by their phone numbers
[**provisioning_api_users_wipe_user_devices**](ProvisioningApiUsersApi.md#provisioning_api_users_wipe_user_devices) | **POST** /ocs/v2.php/cloud/users/{userId}/wipe | Wipe all devices of a user



## provisioning_api_users_add_sub_admin

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_add_sub_admin(groupid, user_id, ocs_api_request)
Make a user a subadmin of a group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**groupid** | **String** | ID of the group | [required] |
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_add_to_group

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_add_to_group(user_id, ocs_api_request, groupid)
Add a user to a group

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**groupid** | Option<**String**> | ID of the group |  |[default to ]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_add_user

> crate::models::ProvisioningApiUsersAddUser200Response provisioning_api_users_add_user(userid, ocs_api_request, password, display_name, email, groups, subadmin, quota, language, manager)
Create a new user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**userid** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]
**password** | Option<**String**> | Password of the user |  |[default to ]
**display_name** | Option<**String**> | Display name of the user |  |[default to ]
**email** | Option<**String**> | Email of the user |  |[default to ]
**groups** | Option<**String**> | Groups of the user |  |
**subadmin** | Option<**String**> | Groups where the user is subadmin |  |
**quota** | Option<**String**> | Quota of the user |  |[default to ]
**language** | Option<**String**> | Language of the user |  |[default to ]
**manager** | Option<**String**> | Manager of the user |  |

### Return type

[**crate::models::ProvisioningApiUsersAddUser200Response**](provisioning_api_users_add_user_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_delete_user

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_delete_user(user_id, ocs_api_request)
Delete a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_disable_user

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_disable_user(user_id, ocs_api_request)
Disable a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_edit_user

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_edit_user(key, value, user_id, ocs_api_request)
Update a value of the user's details

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**key** | **String** | Key that will be updated | [required] |
**value** | **String** | New value for the key | [required] |
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_edit_user_multi_value

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_edit_user_multi_value(key, value, user_id, collection_name, ocs_api_request)
Update multiple values of the user's details

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**key** | **String** | Key that will be updated | [required] |
**value** | **String** | New value for the key | [required] |
**user_id** | **String** | ID of the user | [required] |
**collection_name** | **String** | Collection to update | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_enable_user

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_enable_user(user_id, ocs_api_request)
Enable a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_current_user

> crate::models::ProvisioningApiUsersGetUser200Response provisioning_api_users_get_current_user(ocs_api_request)
Get the details of the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiUsersGetUser200Response**](provisioning_api_users_get_user_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_editable_fields

> crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_editable_fields(ocs_api_request)
Get a list of fields that are editable for the current user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](provisioning_api_groups_get_sub_admins_of_group_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_editable_fields_for_user

> crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_editable_fields_for_user(user_id, ocs_api_request)
Get a list of fields that are editable for a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](provisioning_api_groups_get_sub_admins_of_group_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_user

> crate::models::ProvisioningApiUsersGetUser200Response provisioning_api_users_get_user(user_id, ocs_api_request)
Get the details of a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiUsersGetUser200Response**](provisioning_api_users_get_user_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_user_sub_admin_groups

> crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response provisioning_api_users_get_user_sub_admin_groups(user_id, ocs_api_request)
Get the groups a user is a subadmin of

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID if the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](provisioning_api_groups_get_sub_admins_of_group_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_users

> crate::models::ProvisioningApiGroupsGetGroupUsers200Response provisioning_api_users_get_users(ocs_api_request, search, limit, offset)
Get a list of users

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**search** | Option<**String**> | Text to search for |  |[default to ]
**limit** | Option<**i64**> | Limit the amount of groups returned |  |
**offset** | Option<**i64**> | Offset for searching for groups |  |[default to 0]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroupUsers200Response**](provisioning_api_groups_get_group_users_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_get_users_details

> crate::models::ProvisioningApiGroupsGetGroupUsersDetails200Response provisioning_api_users_get_users_details(ocs_api_request, search, limit, offset)
Get a list of users and their details

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
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


## provisioning_api_users_get_users_groups

> crate::models::ProvisioningApiGroupsGetGroups200Response provisioning_api_users_get_users_groups(user_id, ocs_api_request)
Get a list of groups the user belongs to

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiGroupsGetGroups200Response**](provisioning_api_groups_get_groups_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_remove_from_group

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_remove_from_group(groupid, user_id, ocs_api_request)
Remove a user from a group

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**groupid** | **String** | ID of the group | [required] |
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_remove_sub_admin

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_remove_sub_admin(groupid, user_id, ocs_api_request)
Remove a user from the subadmins of a group

This endpoint requires admin access

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**groupid** | **String** | ID of the group | [required] |
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_resend_welcome_message

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_resend_welcome_message(user_id, ocs_api_request)
Resend the welcome message

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID if the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_search_by_phone_numbers

> crate::models::ProvisioningApiUsersSearchByPhoneNumbers200Response provisioning_api_users_search_by_phone_numbers(location, search, ocs_api_request)
Search users by their phone numbers

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**location** | **String** | Location of the phone number (for country code) | [required] |
**search** | **String** | Phone numbers to search for | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::ProvisioningApiUsersSearchByPhoneNumbers200Response**](provisioning_api_users_search_by_phone_numbers_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## provisioning_api_users_wipe_user_devices

> crate::models::CoreWhatsNewDismiss200Response provisioning_api_users_wipe_user_devices(user_id, ocs_api_request)
Wipe all devices of a user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

