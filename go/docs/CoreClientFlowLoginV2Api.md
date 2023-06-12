# \CoreClientFlowLoginV2Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreClientFlowLoginV2Init**](CoreClientFlowLoginV2Api.md#CoreClientFlowLoginV2Init) | **Post** /index.php/login/v2 | Init a login flow
[**CoreClientFlowLoginV2Poll**](CoreClientFlowLoginV2Api.md#CoreClientFlowLoginV2Poll) | **Post** /index.php/login/v2/poll | Poll the login flow credentials



## CoreClientFlowLoginV2Init

> CoreLoginFlowV2 CoreClientFlowLoginV2Init(ctx).Execute()

Init a login flow

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Init(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Init``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreClientFlowLoginV2Init`: CoreLoginFlowV2
    fmt.Fprintf(os.Stdout, "Response from `CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Init`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCoreClientFlowLoginV2InitRequest struct via the builder pattern


### Return type

[**CoreLoginFlowV2**](CoreLoginFlowV2.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreClientFlowLoginV2Poll

> CoreLoginFlowV2Credentials CoreClientFlowLoginV2Poll(ctx).Token(token).Execute()

Poll the login flow credentials

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
    token := "token_example" // string | Token of the flow

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Poll(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Poll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreClientFlowLoginV2Poll`: CoreLoginFlowV2Credentials
    fmt.Fprintf(os.Stdout, "Response from `CoreClientFlowLoginV2Api.CoreClientFlowLoginV2Poll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreClientFlowLoginV2PollRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token of the flow | 

### Return type

[**CoreLoginFlowV2Credentials**](CoreLoginFlowV2Credentials.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

