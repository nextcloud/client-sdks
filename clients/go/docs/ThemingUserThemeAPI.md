# \ThemingUserThemeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThemingUserThemeDeleteBackground**](ThemingUserThemeAPI.md#ThemingUserThemeDeleteBackground) | **Delete** /index.php/apps/theming/background/custom | Delete the background
[**ThemingUserThemeDisableTheme**](ThemingUserThemeAPI.md#ThemingUserThemeDisableTheme) | **Delete** /ocs/v2.php/apps/theming/api/v1/theme/{themeId} | Disable theme
[**ThemingUserThemeEnableTheme**](ThemingUserThemeAPI.md#ThemingUserThemeEnableTheme) | **Put** /ocs/v2.php/apps/theming/api/v1/theme/{themeId}/enable | Enable theme
[**ThemingUserThemeGetBackground**](ThemingUserThemeAPI.md#ThemingUserThemeGetBackground) | **Get** /index.php/apps/theming/background | Get the background image
[**ThemingUserThemeSetBackground**](ThemingUserThemeAPI.md#ThemingUserThemeSetBackground) | **Post** /index.php/apps/theming/background/{type} | Set the background



## ThemingUserThemeDeleteBackground

> ThemingBackground ThemingUserThemeDeleteBackground(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Delete the background

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
    resp, r, err := apiClient.ThemingUserThemeAPI.ThemingUserThemeDeleteBackground(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingUserThemeAPI.ThemingUserThemeDeleteBackground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingUserThemeDeleteBackground`: ThemingBackground
    fmt.Fprintf(os.Stdout, "Response from `ThemingUserThemeAPI.ThemingUserThemeDeleteBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThemingUserThemeDeleteBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ThemingBackground**](ThemingBackground.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemingUserThemeDisableTheme

> CoreWhatsNewDismiss200Response ThemingUserThemeDisableTheme(ctx, themeId).OCSAPIRequest(oCSAPIRequest).Execute()

Disable theme

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
    themeId := "themeId_example" // string | the theme ID
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingUserThemeAPI.ThemingUserThemeDisableTheme(context.Background(), themeId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingUserThemeAPI.ThemingUserThemeDisableTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingUserThemeDisableTheme`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ThemingUserThemeAPI.ThemingUserThemeDisableTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string** | the theme ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingUserThemeDisableThemeRequest struct via the builder pattern


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


## ThemingUserThemeEnableTheme

> CoreWhatsNewDismiss200Response ThemingUserThemeEnableTheme(ctx, themeId).OCSAPIRequest(oCSAPIRequest).Execute()

Enable theme

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
    themeId := "themeId_example" // string | the theme ID
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingUserThemeAPI.ThemingUserThemeEnableTheme(context.Background(), themeId).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingUserThemeAPI.ThemingUserThemeEnableTheme``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingUserThemeEnableTheme`: CoreWhatsNewDismiss200Response
    fmt.Fprintf(os.Stdout, "Response from `ThemingUserThemeAPI.ThemingUserThemeEnableTheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string** | the theme ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingUserThemeEnableThemeRequest struct via the builder pattern


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


## ThemingUserThemeGetBackground

> *os.File ThemingUserThemeGetBackground(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the background image

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
    resp, r, err := apiClient.ThemingUserThemeAPI.ThemingUserThemeGetBackground(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingUserThemeAPI.ThemingUserThemeGetBackground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingUserThemeGetBackground`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingUserThemeAPI.ThemingUserThemeGetBackground`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiThemingUserThemeGetBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemingUserThemeSetBackground

> ThemingBackground ThemingUserThemeSetBackground(ctx, type_).OCSAPIRequest(oCSAPIRequest).Value(value).Color(color).Execute()

Set the background

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
    type_ := "type__example" // string | Type of background
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    value := "value_example" // string | Path of the background image (optional) (default to "")
    color := "color_example" // string | Color for the background (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingUserThemeAPI.ThemingUserThemeSetBackground(context.Background(), type_).OCSAPIRequest(oCSAPIRequest).Value(value).Color(color).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingUserThemeAPI.ThemingUserThemeSetBackground``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingUserThemeSetBackground`: ThemingBackground
    fmt.Fprintf(os.Stdout, "Response from `ThemingUserThemeAPI.ThemingUserThemeSetBackground`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | Type of background | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingUserThemeSetBackgroundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **value** | **string** | Path of the background image | [default to &quot;&quot;]
 **color** | **string** | Color for the background | 

### Return type

[**ThemingBackground**](ThemingBackground.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

