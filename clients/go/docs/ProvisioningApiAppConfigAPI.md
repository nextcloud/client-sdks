# \ProvisioningApiAppConfigAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningApiAppConfigDeleteKey**](ProvisioningApiAppConfigAPI.md#ProvisioningApiAppConfigDeleteKey) | **Delete** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Delete a config key of an app
[**ProvisioningApiAppConfigGetApps**](ProvisioningApiAppConfigAPI.md#ProvisioningApiAppConfigGetApps) | **Get** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps | Get a list of apps
[**ProvisioningApiAppConfigGetKeys**](ProvisioningApiAppConfigAPI.md#ProvisioningApiAppConfigGetKeys) | **Get** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app} | Get the config keys of an app
[**ProvisioningApiAppConfigGetValue**](ProvisioningApiAppConfigAPI.md#ProvisioningApiAppConfigGetValue) | **Get** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Get a the config value of an app
[**ProvisioningApiAppConfigSetValue**](ProvisioningApiAppConfigAPI.md#ProvisioningApiAppConfigSetValue) | **Post** /ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key} | Update the config value of an app



## ProvisioningApiAppConfigDeleteKey

> CoreWhatsNewDismiss200Response ProvisioningApiAppConfigDeleteKey(ctx, app, key).OCSAPIRequest(oCSAPIRequest).Execute()

Delete a config key of an app



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
    key := "key_example" // string | Key to delete
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigDeleteKey(context.Background(), app, key).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigDeleteKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppConfigDeleteKey`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigDeleteKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 
**key** | **string** | Key to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppConfigDeleteKeyRequest struct via the builder pattern


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


## ProvisioningApiAppConfigGetApps

> ProvisioningApiAppConfigGetApps200Response ProvisioningApiAppConfigGetApps(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get a list of apps



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
    resp, r, err := apiClient.ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetApps(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppConfigGetApps`: ProvisioningApiAppConfigGetApps200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppConfigGetAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiAppConfigGetApps200Response**](ProvisioningApiAppConfigGetApps200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiAppConfigGetKeys

> ProvisioningApiAppConfigGetApps200Response ProvisioningApiAppConfigGetKeys(ctx, app).OCSAPIRequest(oCSAPIRequest).Execute()

Get the config keys of an app



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
    resp, r, err := apiClient.ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetKeys(context.Background(), app).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppConfigGetKeys`: ProvisioningApiAppConfigGetApps200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppConfigGetKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiAppConfigGetApps200Response**](ProvisioningApiAppConfigGetApps200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiAppConfigGetValue

> ProvisioningApiAppConfigGetValue200Response ProvisioningApiAppConfigGetValue(ctx, app, key).OCSAPIRequest(oCSAPIRequest).DefaultValue(defaultValue).Execute()

Get a the config value of an app



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
    key := "key_example" // string | Key
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    defaultValue := "defaultValue_example" // string | Default returned value if the value is empty (optional) (default to "")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetValue(context.Background(), app, key).OCSAPIRequest(oCSAPIRequest).DefaultValue(defaultValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppConfigGetValue`: ProvisioningApiAppConfigGetValue200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigGetValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 
**key** | **string** | Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppConfigGetValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **defaultValue** | **string** | Default returned value if the value is empty | [default to &quot;&quot;]

### Return type

[**ProvisioningApiAppConfigGetValue200Response**](ProvisioningApiAppConfigGetValue200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProvisioningApiAppConfigSetValue

> CoreWhatsNewDismiss200Response ProvisioningApiAppConfigSetValue(ctx, app, key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()

Update the config value of an app

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
    value := "value_example" // string | New value for the key
    app := "app_example" // string | ID of the app
    key := "key_example" // string | Key to update
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigSetValue(context.Background(), app, key).Value(value).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigSetValue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiAppConfigSetValue`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiAppConfigAPI.ProvisioningApiAppConfigSetValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 
**key** | **string** | Key to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiAppConfigSetValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **value** | **string** | New value for the key | 


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

