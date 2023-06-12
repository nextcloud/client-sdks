# \ThemingThemingApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ThemingThemingGetImage**](ThemingThemingApi.md#ThemingThemingGetImage) | **Get** /index.php/apps/theming/image/{key} | Get an image
[**ThemingThemingGetManifest**](ThemingThemingApi.md#ThemingThemingGetManifest) | **Get** /index.php/apps/theming/manifest/{app} | Get the manifest for an app
[**ThemingThemingGetThemeStylesheet**](ThemingThemingApi.md#ThemingThemingGetThemeStylesheet) | **Get** /index.php/apps/theming/theme/{themeId}.css | Get the CSS stylesheet for a theme



## ThemingThemingGetImage

> *os.File ThemingThemingGetImage(ctx, key).UseSvg(useSvg).Execute()

Get an image

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
    key := "key_example" // string | Key of the image
    useSvg := int32(56) // int32 | Return image as SVG (optional) (default to 1)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingThemingApi.ThemingThemingGetImage(context.Background(), key).UseSvg(useSvg).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingThemingApi.ThemingThemingGetImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingThemingGetImage`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingThemingApi.ThemingThemingGetImage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**key** | **string** | Key of the image | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingThemingGetImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useSvg** | **int32** | Return image as SVG | [default to 1]

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


## ThemingThemingGetManifest

> ThemingThemingGetManifest200Response ThemingThemingGetManifest(ctx, app).Execute()

Get the manifest for an app

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
    app := "app_example" // string | ID of the app

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingThemingApi.ThemingThemingGetManifest(context.Background(), app).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingThemingApi.ThemingThemingGetManifest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingThemingGetManifest`: ThemingThemingGetManifest200Response
    fmt.Fprintf(os.Stdout, "Response from `ThemingThemingApi.ThemingThemingGetManifest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string** | ID of the app | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingThemingGetManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ThemingThemingGetManifest200Response**](ThemingThemingGetManifest200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ThemingThemingGetThemeStylesheet

> *os.File ThemingThemingGetThemeStylesheet(ctx, themeId).Plain(plain).WithCustomCss(withCustomCss).Execute()

Get the CSS stylesheet for a theme

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
    themeId := "themeId_example" // string | ID of the theme
    plain := int32(56) // int32 | Let the browser decide the CSS priority (optional) (default to 0)
    withCustomCss := int32(56) // int32 | Include custom CSS (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ThemingThemingApi.ThemingThemingGetThemeStylesheet(context.Background(), themeId).Plain(plain).WithCustomCss(withCustomCss).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ThemingThemingApi.ThemingThemingGetThemeStylesheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ThemingThemingGetThemeStylesheet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `ThemingThemingApi.ThemingThemingGetThemeStylesheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**themeId** | **string** | ID of the theme | 

### Other Parameters

Other parameters are passed through a pointer to a apiThemingThemingGetThemeStylesheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **plain** | **int32** | Let the browser decide the CSS priority | [default to 0]
 **withCustomCss** | **int32** | Include custom CSS | [default to 0]

### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/css

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

