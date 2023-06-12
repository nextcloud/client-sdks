# \WeatherStatusWeatherStatusApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WeatherStatusWeatherStatusGetFavorites**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusGetFavorites) | **Get** /ocs/v2.php/apps/weather_status/api/v1/favorites | Get favorites list
[**WeatherStatusWeatherStatusGetForecast**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusGetForecast) | **Get** /ocs/v2.php/apps/weather_status/api/v1/forecast | Get forecast for current location
[**WeatherStatusWeatherStatusGetLocation**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusGetLocation) | **Get** /ocs/v2.php/apps/weather_status/api/v1/location | Get stored user location
[**WeatherStatusWeatherStatusSetFavorites**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusSetFavorites) | **Put** /ocs/v2.php/apps/weather_status/api/v1/favorites | Set favorites list
[**WeatherStatusWeatherStatusSetLocation**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusSetLocation) | **Put** /ocs/v2.php/apps/weather_status/api/v1/location | Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding
[**WeatherStatusWeatherStatusSetMode**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusSetMode) | **Put** /ocs/v2.php/apps/weather_status/api/v1/mode | Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address
[**WeatherStatusWeatherStatusUsePersonalAddress**](WeatherStatusWeatherStatusApi.md#WeatherStatusWeatherStatusUsePersonalAddress) | **Put** /ocs/v2.php/apps/weather_status/api/v1/use-personal | Try to use the address set in user personal settings as weather location



## WeatherStatusWeatherStatusGetFavorites

> ProvisioningApiGroupsGetSubAdminsOfGroup200Response WeatherStatusWeatherStatusGetFavorites(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get favorites list

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
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetFavorites(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetFavorites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusGetFavorites`: ProvisioningApiGroupsGetSubAdminsOfGroup200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusGetFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**ProvisioningApiGroupsGetSubAdminsOfGroup200Response**](ProvisioningApiGroupsGetSubAdminsOfGroup200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusGetForecast

> WeatherStatusWeatherStatusGetForecast200Response WeatherStatusWeatherStatusGetForecast(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get forecast for current location

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
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetForecast(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetForecast``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusGetForecast`: WeatherStatusWeatherStatusGetForecast200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetForecast`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusGetForecastRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**WeatherStatusWeatherStatusGetForecast200Response**](WeatherStatusWeatherStatusGetForecast200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusGetLocation

> WeatherStatusWeatherStatusGetLocation200Response WeatherStatusWeatherStatusGetLocation(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Get stored user location

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
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetLocation(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusGetLocation`: WeatherStatusWeatherStatusGetLocation200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusGetLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusGetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**WeatherStatusWeatherStatusGetLocation200Response**](WeatherStatusWeatherStatusGetLocation200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusSetFavorites

> CoreReferenceApiTouchProvider200Response WeatherStatusWeatherStatusSetFavorites(ctx).Favorites(favorites).OCSAPIRequest(oCSAPIRequest).Execute()

Set favorites list

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
    favorites := "favorites_example" // string | Favorite addresses
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetFavorites(context.Background()).Favorites(favorites).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetFavorites``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusSetFavorites`: CoreReferenceApiTouchProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetFavorites`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusSetFavoritesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **favorites** | **string** | Favorite addresses | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreReferenceApiTouchProvider200Response**](CoreReferenceApiTouchProvider200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusSetLocation

> WeatherStatusWeatherStatusUsePersonalAddress200Response WeatherStatusWeatherStatusSetLocation(ctx).OCSAPIRequest(oCSAPIRequest).Address(address).Lat(lat).Lon(lon).Execute()

Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding

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
    address := "address_example" // string | Any approximative or exact address (optional)
    lat := float32(3.4) // float32 | Latitude in decimal degree format (optional)
    lon := float32(3.4) // float32 | Longitude in decimal degree format (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetLocation(context.Background()).OCSAPIRequest(oCSAPIRequest).Address(address).Lat(lat).Lon(lon).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusSetLocation`: WeatherStatusWeatherStatusUsePersonalAddress200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetLocation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusSetLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]
 **address** | **string** | Any approximative or exact address | 
 **lat** | **float32** | Latitude in decimal degree format | 
 **lon** | **float32** | Longitude in decimal degree format | 

### Return type

[**WeatherStatusWeatherStatusUsePersonalAddress200Response**](WeatherStatusWeatherStatusUsePersonalAddress200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusSetMode

> CoreReferenceApiTouchProvider200Response WeatherStatusWeatherStatusSetMode(ctx).Mode(mode).OCSAPIRequest(oCSAPIRequest).Execute()

Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address

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
    mode := int64(789) // int64 | New mode
    oCSAPIRequest := "oCSAPIRequest_example" // string |  (default to "true")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetMode(context.Background()).Mode(mode).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetMode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusSetMode`: CoreReferenceApiTouchProvider200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusSetMode`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusSetModeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mode** | **int64** | New mode | 
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**CoreReferenceApiTouchProvider200Response**](CoreReferenceApiTouchProvider200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WeatherStatusWeatherStatusUsePersonalAddress

> WeatherStatusWeatherStatusUsePersonalAddress200Response WeatherStatusWeatherStatusUsePersonalAddress(ctx).OCSAPIRequest(oCSAPIRequest).Execute()

Try to use the address set in user personal settings as weather location

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
    resp, r, err := apiClient.WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusUsePersonalAddress(context.Background()).OCSAPIRequest(oCSAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusUsePersonalAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WeatherStatusWeatherStatusUsePersonalAddress`: WeatherStatusWeatherStatusUsePersonalAddress200Response
    fmt.Fprintf(os.Stdout, "Response from `WeatherStatusWeatherStatusApi.WeatherStatusWeatherStatusUsePersonalAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWeatherStatusWeatherStatusUsePersonalAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **oCSAPIRequest** | **string** |  | [default to &quot;true&quot;]

### Return type

[**WeatherStatusWeatherStatusUsePersonalAddress200Response**](WeatherStatusWeatherStatusUsePersonalAddress200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

