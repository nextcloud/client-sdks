# \Oauth2LoginRedirectorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**oauth2_login_redirector_authorize**](Oauth2LoginRedirectorApi.md#oauth2_login_redirector_authorize) | **GET** /index.php/apps/oauth2/authorize | Authorize the user



## oauth2_login_redirector_authorize

> String oauth2_login_redirector_authorize(client_id, state, response_type)
Authorize the user

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**client_id** | **String** | Client ID | [required] |
**state** | **String** | State of the flow | [required] |
**response_type** | **String** | Response type for the flow | [required] |

### Return type

**String**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

