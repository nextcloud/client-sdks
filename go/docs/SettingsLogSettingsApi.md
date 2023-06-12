# \SettingsLogSettingsApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SettingsLogSettingsDownload**](SettingsLogSettingsApi.md#SettingsLogSettingsDownload) | **Get** /index.php/settings/admin/log/download | download logfile



## SettingsLogSettingsDownload

> *os.File SettingsLogSettingsDownload(ctx).Execute()

download logfile



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SettingsLogSettingsApi.SettingsLogSettingsDownload(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SettingsLogSettingsApi.SettingsLogSettingsDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SettingsLogSettingsDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `SettingsLogSettingsApi.SettingsLogSettingsDownload`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSettingsLogSettingsDownloadRequest struct via the builder pattern


### Return type

[***os.File**](*os.File.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

