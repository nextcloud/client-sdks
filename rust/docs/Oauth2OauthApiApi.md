# \Oauth2OauthApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**oauth2_oauth_api_get_token**](Oauth2OauthApiApi.md#oauth2_oauth_api_get_token) | **POST** /index.php/apps/oauth2/api/v1/token | Get a token



## oauth2_oauth_api_get_token

> crate::models::Oauth2OauthApiGetToken200Response oauth2_oauth_api_get_token(grant_type, code, refresh_token, client_id, client_secret)
Get a token

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**grant_type** | **String** | Token type that should be granted | [required] |
**code** | **String** | Code of the flow | [required] |
**refresh_token** | **String** | Refresh token | [required] |
**client_id** | **String** | Client ID | [required] |
**client_secret** | **String** | Client secret | [required] |

### Return type

[**crate::models::Oauth2OauthApiGetToken200Response**](oauth2_oauth_api_get_token_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

