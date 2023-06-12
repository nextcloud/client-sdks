# \Oauth2LoginRedirectorApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Oauth2LoginRedirectorAuthorize**](Oauth2LoginRedirectorApi.md#Oauth2LoginRedirectorAuthorize) | **Get** /index.php/apps/oauth2/authorize | Authorize the user



## Oauth2LoginRedirectorAuthorize

> string Oauth2LoginRedirectorAuthorize(ctx).ClientId(clientId).State(state).ResponseType(responseType).Execute()

Authorize the user

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
    clientId := "clientId_example" // string | Client ID
    state := "state_example" // string | State of the flow
    responseType := "responseType_example" // string | Response type for the flow

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.Oauth2LoginRedirectorApi.Oauth2LoginRedirectorAuthorize(context.Background()).ClientId(clientId).State(state).ResponseType(responseType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `Oauth2LoginRedirectorApi.Oauth2LoginRedirectorAuthorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `Oauth2LoginRedirectorAuthorize`: string
    fmt.Fprintf(os.Stdout, "Response from `Oauth2LoginRedirectorApi.Oauth2LoginRedirectorAuthorize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauth2LoginRedirectorAuthorizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clientId** | **string** | Client ID | 
 **state** | **string** | State of the flow | 
 **responseType** | **string** | Response type for the flow | 

### Return type

**string**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

