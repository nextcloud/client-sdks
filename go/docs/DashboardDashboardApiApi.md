# \DashboardDashboardApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DashboardDashboardApiGetWidgetItems**](DashboardDashboardApiApi.md#DashboardDashboardApiGetWidgetItems) | **Get** /ocs/v2.php/apps/dashboard/api/v1/widget-items | Get the items for the widgets
[**DashboardDashboardApiGetWidgets**](DashboardDashboardApiApi.md#DashboardDashboardApiGetWidgets) | **Get** /ocs/v2.php/apps/dashboard/api/v1/widgets | Get the widgets



## DashboardDashboardApiGetWidgetItems

> DashboardDashboardApiGetWidgetItems200Response DashboardDashboardApiGetWidgetItems(ctx).OCSAPIRequest(oCSAPIRequest).SinceIds(sinceIds).Limit(limit).Widgets(widgets).Execute()

Get the items for the widgets

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
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")
    sinceIds := "sinceIds_example" // string | Array indexed by widget Ids, contains date/id from which we want the new items (optional)
    limit := int64(789) // int64 | Limit number of result items per widget (optional) (default to 7)
    widgets := "widgets_example" // string | Limit results to specific widgets (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardDashboardApiApi.DashboardDashboardApiGetWidgetItems(context.Background()).OCSAPIRequest(oCSAPIRequest).SinceIds(sinceIds).Limit(limit).Widgets(widgets).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardDashboardApiApi.DashboardDashboardApiGetWidgetItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardDashboardApiGetWidgetItems`: DashboardDashboardApiGetWidgetItems200Response
    fmt.Fprintf(os.Stdout, "Response from `DashboardDashboardApiApi.DashboardDashboardApiGetWidgetItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardDashboardApiGetWidgetItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **sinceIds** | **string** | Array indexed by widget Ids, contains date/id from which we want the new items | 
 **limit** | **int64** | Limit number of result items per widget | [default to 7]
 **widgets** | **string** | Limit results to specific widgets | 

### Return type

[**DashboardDashboardApiGetWidgetItems200Response**](DashboardDashboardApiGetWidgetItems200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DashboardDashboardApiGetWidgets

> DashboardDashboardApiGetWidgets200Response DashboardDashboardApiGetWidgets(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get the widgets

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
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DashboardDashboardApiApi.DashboardDashboardApiGetWidgets(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DashboardDashboardApiApi.DashboardDashboardApiGetWidgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DashboardDashboardApiGetWidgets`: DashboardDashboardApiGetWidgets200Response
    fmt.Fprintf(os.Stdout, "Response from `DashboardDashboardApiApi.DashboardDashboardApiGetWidgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDashboardDashboardApiGetWidgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**DashboardDashboardApiGetWidgets200Response**](DashboardDashboardApiGetWidgets200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

