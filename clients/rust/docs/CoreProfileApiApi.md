# \CoreProfileApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_profile_api_set_visibility**](CoreProfileApiApi.md#core_profile_api_set_visibility) | **PUT** /ocs/v2.php/profile/{targetUserId} | Update the visibility of a parameter



## core_profile_api_set_visibility

> crate::models::CoreWhatsNewDismiss200Response core_profile_api_set_visibility(param_id, visibility, target_user_id, ocs_api_request)
Update the visibility of a parameter

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**param_id** | **String** | ID of the parameter | [required] |
**visibility** | **String** | New visibility | [required] |
**target_user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreWhatsNewDismiss200Response**](core_whats_new_dismiss_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

