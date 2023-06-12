# \Oauth2OauthApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2OauthApiGetToken**](Oauth2OauthApiApi.md#Oauth2OauthApiGetToken) | **Post** /index.php/apps/oauth2/api/v1/token | Get a token



## Oauth2OauthApiGetToken

> Oauth2OauthApiGetToken200Response Oauth2OauthApiGetToken(ctx).GrantType(grantType).Code(code).RefreshToken(refreshToken).ClientId(clientId).ClientSecret(clientSecret).Execute()

Get a token

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/api-sdk"
)

func main() {
    grantType := "grantType_example" // string | Token type that should be granted
    code := "code_example" // string | Code of the flow
    refreshToken := "refreshToken_example" // string | Refresh token
    clientId := "clientId_example" // string | Client ID
    clientSecret := "clientSecret_example" // string | Client secret

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2OauthApiApi.Oauth2OauthApiGetToken(context.Background()).GrantType(grantType).Code(code).RefreshToken(refreshToken).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2OauthApiApi.Oauth2OauthApiGetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2OauthApiGetToken`: Oauth2OauthApiGetToken200Response
    fmt.Fprintf(os.Stdout, "Response from `Oauth2OauthApiApi.Oauth2OauthApiGetToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2OauthApiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** | Token type that should be granted | 
 **code** | **string** | Code of the flow | 
 **refreshToken** | **string** | Refresh token | 
 **clientId** | **string** | Client ID | 
 **clientSecret** | **string** | Client secret | 

### Return type

[**Oauth2OauthApiGetToken200Response**](Oauth2OauthApiGetToken200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

