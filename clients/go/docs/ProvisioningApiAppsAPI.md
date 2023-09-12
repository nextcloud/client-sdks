# \ProvisioningApiAppsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningApiAppsDisable**](ProvisioningApiAppsAPI.md#ProvisioningApiAppsDisable) | **Delete** /ocs/v2.php/cloud/apps/{app} | Disable an app
[**ProvisioningApiAppsEnable**](ProvisioningApiAppsAPI.md#ProvisioningApiAppsEnable) | **Post** /ocs/v2.php/cloud/apps/{app} | Enable an app
[**ProvisioningApiAppsGetAppInfo**](ProvisioningApiAppsAPI.md#ProvisioningApiAppsGetAppInfo) | **Get** /ocs/v2.php/cloud/apps/{app} | Get the app info for an app
[**ProvisioningApiAppsGetApps**](ProvisioningApiAppsAPI.md#ProvisioningApiAppsGetApps) | **Get** /ocs/v2.php/cloud/apps | Get a list of installed apps



## ProvisioningApiAppsDisable

> CoreWhatsNewDismiss200Response ProvisioningApiAppsDisable(ctx, app).OCSAPIRequest(oCSAPIRequest).Execute()

Disable an app



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
    app := "app_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppsAPI.ProvisioningApiAppsDisable(context.Background(), app).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppsAPI.ProvisioningApiAppsDisable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppsDisable`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppsAPI.ProvisioningApiAppsDisable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppsDisableRequest struct via the builder pattern


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


## ProvisioningApiAppsEnable

> CoreWhatsNewDismiss200Response ProvisioningApiAppsEnable(ctx, app).OCSAPIRequest(oCSAPIRequest).Execute()

Enable an app



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
    app := "app_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppsAPI.ProvisioningApiAppsEnable(context.Background(), app).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppsAPI.ProvisioningApiAppsEnable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppsEnable`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppsAPI.ProvisioningApiAppsEnable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppsEnableRequest struct via the builder pattern


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


## ProvisioningApiAppsGetAppInfo

> ProvisioningApiAppsGetAppInfo200Response ProvisioningApiAppsGetAppInfo(ctx, app).OCSAPIRequest(oCSAPIRequest).Execute()

Get the app info for an app



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
    app := "app_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppsAPI.ProvisioningApiAppsGetAppInfo(context.Background(), app).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppsAPI.ProvisioningApiAppsGetAppInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppsGetAppInfo`: ProvisioningApiAppsGetAppInfo200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppsAPI.ProvisioningApiAppsGetAppInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppsGetAppInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiAppsGetAppInfo200Response**](ProvisioningApiAppsGetAppInfo200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiAppsGetApps

> ProvisioningApiAppsGetApps200Response ProvisioningApiAppsGetApps(ctx).OCSAPIRequest(oCSAPIRequest).Filter(filter).Execute()

Get a list of installed apps



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
    filter := "filter_example" // string | Filter for enabled or disabled apps (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppsAPI.ProvisioningApiAppsGetApps(context.Background()).OCSAPIRequest(oCSAPIRequest).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppsAPI.ProvisioningApiAppsGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppsGetApps`: ProvisioningApiAppsGetApps200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppsAPI.ProvisioningApiAppsGetApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppsGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **filter** | **string** | Filter for enabled or disabled apps | 

### Return type

[**ProvisioningApiAppsGetApps200Response**](ProvisioningApiAppsGetApps200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

