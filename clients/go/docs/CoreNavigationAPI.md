# \CoreNavigationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreNavigationGetAppsNavigation**](CoreNavigationAPI.md#CoreNavigationGetAppsNavigation) | **Get** /ocs/v2.php/core/navigation/apps | Get the apps navigation
[**CoreNavigationGetSettingsNavigation**](CoreNavigationAPI.md#CoreNavigationGetSettingsNavigation) | **Get** /ocs/v2.php/core/navigation/settings | Get the settings navigation



## CoreNavigationGetAppsNavigation

> CoreNavigationGetAppsNavigation200Response CoreNavigationGetAppsNavigation(ctx).OCSAPIRequest(oCSAPIRequest).Absolute(absolute).Execute()

Get the apps navigation

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
    absolute := int32(56) // int32 | Rewrite URLs to absolute ones (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreNavigationAPI.CoreNavigationGetAppsNavigation(context.Background()).OCSAPIRequest(oCSAPIRequest).Absolute(absolute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNavigationAPI.CoreNavigationGetAppsNavigation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreNavigationGetAppsNavigation`: CoreNavigationGetAppsNavigation200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreNavigationAPI.CoreNavigationGetAppsNavigation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreNavigationGetAppsNavigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **absolute** | **int32** | Rewrite URLs to absolute ones | [default to 0]

### Return type

[**CoreNavigationGetAppsNavigation200Response**](CoreNavigationGetAppsNavigation200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CoreNavigationGetSettingsNavigation

> CoreNavigationGetAppsNavigation200Response CoreNavigationGetSettingsNavigation(ctx).OCSAPIRequest(oCSAPIRequest).Absolute(absolute).Execute()

Get the settings navigation

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
    absolute := int32(56) // int32 | Rewrite URLs to absolute ones (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreNavigationAPI.CoreNavigationGetSettingsNavigation(context.Background()).OCSAPIRequest(oCSAPIRequest).Absolute(absolute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreNavigationAPI.CoreNavigationGetSettingsNavigation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreNavigationGetSettingsNavigation`: CoreNavigationGetAppsNavigation200Response
    fmt.Fprintf(os.Stdout, "Response from `CoreNavigationAPI.CoreNavigationGetSettingsNavigation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCoreNavigationGetSettingsNavigationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **absolute** | **int32** | Rewrite URLs to absolute ones | [default to 0]

### Return type

[**CoreNavigationGetAppsNavigation200Response**](CoreNavigationGetAppsNavigation200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

