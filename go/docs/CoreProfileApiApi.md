# \CoreProfileApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreProfileApiSetVisibility**](CoreProfileApiApi.md#CoreProfileApiSetVisibility) | **Put** /ocs/v2.php/profile/{targetUserId} | Update the visiblity of a parameter



## CoreProfileApiSetVisibility

> CoreWhatsNewDismiss200Response CoreProfileApiSetVisibility(ctx, targetUserId).ParamId(paramId).Visibility(visibility).OCSAPIRequest(oCSAPIRequest).Execute()

Update the visiblity of a parameter

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
    paramId := "paramId_example" // string | ID of the parameter
    visibility := "visibility_example" // string | New visibility
    targetUserId := "targetUserId_example" // string | ID of the user
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreProfileApiApi.CoreProfileApiSetVisibility(context.Background(), targetUserId).ParamId(paramId).Visibility(visibility).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreProfileApiApi.CoreProfileApiSetVisibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreProfileApiSetVisibility`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreProfileApiApi.CoreProfileApiSetVisibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetUserId** | **string** | ID of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreProfileApiSetVisibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paramId** | **string** | ID of the parameter | 
 **visibility** | **string** | New visibility | 

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreWhatsNewDismiss200Response**](CoreWhatsNewDismiss200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

