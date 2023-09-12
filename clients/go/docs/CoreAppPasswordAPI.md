# \CoreAppPasswordAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreAppPasswordDeleteAppPassword**](CoreAppPasswordAPI.md#CoreAppPasswordDeleteAppPassword) | **Delete** /ocs/v2.php/core/apppassword | Delete app password
[**CoreAppPasswordGetAppPassword**](CoreAppPasswordAPI.md#CoreAppPasswordGetAppPassword) | **Get** /ocs/v2.php/core/getapppassword | Create app password
[**CoreAppPasswordRotateAppPassword**](CoreAppPasswordAPI.md#CoreAppPasswordRotateAppPassword) | **Post** /ocs/v2.php/core/apppassword/rotate | Rotate app password



## CoreAppPasswordDeleteAppPassword

> CoreWhatsNewDismiss200Response CoreAppPasswordDeleteAppPassword(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Delete app password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAppPasswordAPI.CoreAppPasswordDeleteAppPassword(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAppPasswordAPI.CoreAppPasswordDeleteAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAppPasswordDeleteAppPassword`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreAppPasswordAPI.CoreAppPasswordDeleteAppPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreAppPasswordDeleteAppPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
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


## CoreAppPasswordGetAppPassword

> CoreAppPasswordGetAppPassword200Response CoreAppPasswordGetAppPassword(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Create app password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAppPasswordAPI.CoreAppPasswordGetAppPassword(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAppPasswordAPI.CoreAppPasswordGetAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAppPasswordGetAppPassword`: CoreAppPasswordGetAppPassword200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreAppPasswordAPI.CoreAppPasswordGetAppPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreAppPasswordGetAppPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreAppPasswordGetAppPassword200Response**](CoreAppPasswordGetAppPassword200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreAppPasswordRotateAppPassword

> CoreAppPasswordGetAppPassword200Response CoreAppPasswordRotateAppPassword(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Rotate app password

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/nextcloud/client-sdks"
)

func main() {
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAppPasswordAPI.CoreAppPasswordRotateAppPassword(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAppPasswordAPI.CoreAppPasswordRotateAppPassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAppPasswordRotateAppPassword`: CoreAppPasswordGetAppPassword200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreAppPasswordAPI.CoreAppPasswordRotateAppPassword`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreAppPasswordRotateAppPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreAppPasswordGetAppPassword200Response**](CoreAppPasswordGetAppPassword200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

