# \CoreWipeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreWipeCheckWipe**](CoreWipeApi.md#CoreWipeCheckWipe) | **Post** /index.php/core/wipe/check | Check if the device should be wiped
[**CoreWipeWipeDone**](CoreWipeApi.md#CoreWipeWipeDone) | **Post** /index.php/core/wipe/success | Finish the wipe



## CoreWipeCheckWipe

> CoreWipeCheckWipe200Response CoreWipeCheckWipe(ctx).Token(token).Execute()

Check if the device should be wiped

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
    token := "token_example" // string | App password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreWipeApi.CoreWipeCheckWipe(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWipeApi.CoreWipeCheckWipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreWipeCheckWipe`: CoreWipeCheckWipe200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreWipeApi.CoreWipeCheckWipe`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreWipeCheckWipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | App password | 

### Return type

[**CoreWipeCheckWipe200Response**](CoreWipeCheckWipe200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreWipeWipeDone

> map[string]interface{} CoreWipeWipeDone(ctx).Token(token).Execute()

Finish the wipe

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
    token := "token_example" // string | App password

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreWipeApi.CoreWipeWipeDone(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreWipeApi.CoreWipeWipeDone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreWipeWipeDone`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `CoreWipeApi.CoreWipeWipeDone`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreWipeWipeDoneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | App password | 

### Return type

**map[string]interface{}**

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

