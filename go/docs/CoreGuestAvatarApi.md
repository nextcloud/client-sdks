# \CoreGuestAvatarApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreGuestAvatarGetAvatar**](CoreGuestAvatarApi.md#CoreGuestAvatarGetAvatar) | **Get** /index.php/avatar/guest/{guestName}/{size} | Returns a guest avatar image response
[**CoreGuestAvatarGetAvatarDark**](CoreGuestAvatarApi.md#CoreGuestAvatarGetAvatarDark) | **Get** /index.php/avatar/guest/{guestName}/{size}/dark | Returns a dark guest avatar image response



## CoreGuestAvatarGetAvatar

> *os.File CoreGuestAvatarGetAvatar(ctx, guestName, size).DarkTheme(darkTheme).Execute()

Returns a guest avatar image response

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
    guestName := "guestName_example" // string | The guest name, e.g. \"Albert\"
    size := "size_example" // string | The desired avatar size, e.g. 64 for 64x64px
    darkTheme := int32(56) // int32 | Return dark avatar (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreGuestAvatarApi.CoreGuestAvatarGetAvatar(context.Background(), guestName, size).DarkTheme(darkTheme).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreGuestAvatarApi.CoreGuestAvatarGetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreGuestAvatarGetAvatar`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CoreGuestAvatarApi.CoreGuestAvatarGetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guestName** | **string** | The guest name, e.g. \&quot;Albert\&quot; | 
**size** | **string** | The desired avatar size, e.g. 64 for 64x64px | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGuestAvatarGetAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **darkTheme** | **int32** | Return dark avatar | [default to 0]

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


## CoreGuestAvatarGetAvatarDark

> *os.File CoreGuestAvatarGetAvatarDark(ctx, guestName, size).Execute()

Returns a dark guest avatar image response

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
    guestName := "guestName_example" // string | The guest name, e.g. \"Albert\"
    size := "size_example" // string | The desired avatar size, e.g. 64 for 64x64px

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreGuestAvatarApi.CoreGuestAvatarGetAvatarDark(context.Background(), guestName, size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreGuestAvatarApi.CoreGuestAvatarGetAvatarDark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreGuestAvatarGetAvatarDark`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CoreGuestAvatarApi.CoreGuestAvatarGetAvatarDark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**guestName** | **string** | The guest name, e.g. \&quot;Albert\&quot; | 
**size** | **string** | The desired avatar size, e.g. 64 for 64x64px | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreGuestAvatarGetAvatarDarkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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

