# \CoreHoverCardApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_hover_card_get_user**](CoreHoverCardApi.md#core_hover_card_get_user) | **GET** /ocs/v2.php/hovercard/v1/{userId} | Get the user details for a hovercard



## core_hover_card_get_user

> crate::models::CoreHoverCardGetUser200Response core_hover_card_get_user(user_id, ocs_api_request)
Get the user details for a hovercard

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**user_id** | **String** | ID of the user | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreHoverCardGetUser200Response**](core_hover_card_get_user_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

