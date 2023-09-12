# \CoreAvatarAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CoreAvatarGetAvatar**](CoreAvatarAPI.md#CoreAvatarGetAvatar) | **Get** /index.php/avatar/{userId}/{size} | Get the avatar
[**CoreAvatarGetAvatarDark**](CoreAvatarAPI.md#CoreAvatarGetAvatarDark) | **Get** /index.php/avatar/{userId}/{size}/dark | Get the dark avatar



## CoreAvatarGetAvatar

> *os.File CoreAvatarGetAvatar(ctx, userId, size).Execute()

Get the avatar

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
    userId := "userId_example" // string | ID of the user
    size := int64(789) // int64 | Size of the avatar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAvatarAPI.CoreAvatarGetAvatar(context.Background(), userId, size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAvatarAPI.CoreAvatarGetAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAvatarGetAvatar`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CoreAvatarAPI.CoreAvatarGetAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 
**size** | **int64** | Size of the avatar | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreAvatarGetAvatarRequest struct via the builder pattern


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


## CoreAvatarGetAvatarDark

> *os.File CoreAvatarGetAvatarDark(ctx, userId, size).Execute()

Get the dark avatar

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
    userId := "userId_example" // string | ID of the user
    size := int64(789) // int64 | Size of the avatar

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CoreAvatarAPI.CoreAvatarGetAvatarDark(context.Background(), userId, size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CoreAvatarAPI.CoreAvatarGetAvatarDark``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CoreAvatarGetAvatarDark`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CoreAvatarAPI.CoreAvatarGetAvatarDark`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | ID of the user | 
**size** | **int64** | Size of the avatar | 

### Other Parameters

Other parameters are passed through a pointer to a apiCoreAvatarGetAvatarDarkRequest struct via the builder pattern


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

