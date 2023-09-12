# \ProvisioningApiPreferencesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProvisioningApiPreferencesDeleteMultiplePreference**](ProvisioningApiPreferencesAPI.md#ProvisioningApiPreferencesDeleteMultiplePreference) | **Delete** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Delete multiple preferences for an app
[**ProvisioningApiPreferencesDeletePreference**](ProvisioningApiPreferencesAPI.md#ProvisioningApiPreferencesDeletePreference) | **Delete** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Delete a preference for an app
[**ProvisioningApiPreferencesSetMultiplePreferences**](ProvisioningApiPreferencesAPI.md#ProvisioningApiPreferencesSetMultiplePreferences) | **Post** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId} | Update multiple preference values of an app
[**ProvisioningApiPreferencesSetPreference**](ProvisioningApiPreferencesAPI.md#ProvisioningApiPreferencesSetPreference) | **Post** /ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey} | Update a preference value of an app



## ProvisioningApiPreferencesDeleteMultiplePreference

> CoreWhatsNewDismiss200Response ProvisioningApiPreferencesDeleteMultiplePreference(ctx, appId).ConfigKeys(configKeys).OCSAPIRequest(oCSAPIRequest).Execute()

Delete multiple preferences for an app

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
    configKeys := []string{"Inner_example"} // []string | Keys to delete
    appId := "appId_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeleteMultiplePreference(context.Background(), appId).ConfigKeys(configKeys).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeleteMultiplePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiPreferencesDeleteMultiplePreference`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeleteMultiplePreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiPreferencesDeleteMultiplePreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configKeys** | **[]string** | Keys to delete | 

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


## ProvisioningApiPreferencesDeletePreference

> CoreWhatsNewDismiss200Response ProvisioningApiPreferencesDeletePreference(ctx, appId, configKey).OCSAPIRequest(oCSAPIRequest).Execute()

Delete a preference for an app

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
    appId := "appId_example" // string | ID of the app
    configKey := "configKey_example" // string | Key to delete
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeletePreference(context.Background(), appId, configKey).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeletePreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiPreferencesDeletePreference`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesDeletePreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the app | 
**configKey** | **string** | Key to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiPreferencesDeletePreferenceRequest struct via the builder pattern


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


## ProvisioningApiPreferencesSetMultiplePreferences

> CoreWhatsNewDismiss200Response ProvisioningApiPreferencesSetMultiplePreferences(ctx, appId).Configs(configs).OCSAPIRequest(oCSAPIRequest).Execute()

Update multiple preference values of an app

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
    configs := "configs_example" // string | Key-value pairs of the preferences
    appId := "appId_example" // string | ID of the app
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetMultiplePreferences(context.Background(), appId).Configs(configs).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetMultiplePreferences``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiPreferencesSetMultiplePreferences`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetMultiplePreferences`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiPreferencesSetMultiplePreferencesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configs** | **string** | Key-value pairs of the preferences | 

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


## ProvisioningApiPreferencesSetPreference

> CoreWhatsNewDismiss200Response ProvisioningApiPreferencesSetPreference(ctx, appId, configKey).ConfigValue(configValue).OCSAPIRequest(oCSAPIRequest).Execute()

Update a preference value of an app

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
    configValue := "configValue_example" // string | New value
    appId := "appId_example" // string | ID of the app
    configKey := "configKey_example" // string | Key of the preference
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetPreference(context.Background(), appId, configKey).ConfigValue(configValue).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetPreference``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProvisioningApiPreferencesSetPreference`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ProvisioningApiPreferencesAPI.ProvisioningApiPreferencesSetPreference`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** | ID of the app | 
**configKey** | **string** | Key of the preference | 

### Other Parameters

Other parameters are passed through a pointer to a apiProvisioningApiPreferencesSetPreferenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configValue** | **string** | New value | 


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

