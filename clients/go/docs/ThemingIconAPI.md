# \ThemingIconAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThemingIconGetFavicon**](ThemingIconAPI.md#ThemingIconGetFavicon) | **Get** /index.php/apps/theming/favicon/{app} | Return a 32x32 favicon as png
[**ThemingIconGetThemedIcon**](ThemingIconAPI.md#ThemingIconGetThemedIcon) | **Get** /index.php/apps/theming/img/{app}/{image} | Get a themed icon
[**ThemingIconGetTouchIcon**](ThemingIconAPI.md#ThemingIconGetTouchIcon) | **Get** /index.php/apps/theming/icon/{app} | Return a 512x512 icon for touch devices



## ThemingIconGetFavicon

> *os.File ThemingIconGetFavicon(ctx, app).Execute()

Return a 32x32 favicon as png

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
    app := "app_example" // string | ID of the app (default to "core")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingIconAPI.ThemingIconGetFavicon(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingIconAPI.ThemingIconGetFavicon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingIconGetFavicon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingIconAPI.ThemingIconGetFavicon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | [default to &quot;core&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiThemingIconGetFaviconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/x-icon

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemingIconGetThemedIcon

> *os.File ThemingIconGetThemedIcon(ctx, app, image).Execute()

Get a themed icon

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
    image := "image_example" // string | image file name (svg required)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingIconAPI.ThemingIconGetThemedIcon(context.Background(), app, image).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingIconAPI.ThemingIconGetThemedIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingIconGetThemedIcon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingIconAPI.ThemingIconGetThemedIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 
**image** | **string** | image file name (svg required) | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingIconGetThemedIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemingIconGetTouchIcon

> *os.File ThemingIconGetTouchIcon(ctx, app).Execute()

Return a 512x512 icon for touch devices

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
    app := "app_example" // string | ID of the app (default to "core")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingIconAPI.ThemingIconGetTouchIcon(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingIconAPI.ThemingIconGetTouchIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingIconGetTouchIcon`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingIconAPI.ThemingIconGetTouchIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | [default to &quot;core&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiThemingIconGetTouchIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/png, image/x-icon

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

