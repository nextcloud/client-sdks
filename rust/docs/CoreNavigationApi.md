# \CoreNavigationApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_navigation_get_apps_navigation**](CoreNavigationApi.md#core_navigation_get_apps_navigation) | **GET** /ocs/v2.php/core/navigation/apps | Get the apps navigation
[**core_navigation_get_settings_navigation**](CoreNavigationApi.md#core_navigation_get_settings_navigation) | **GET** /ocs/v2.php/core/navigation/settings | Get the settings navigation



## core_navigation_get_apps_navigation

> crate::models::CoreNavigationGetAppsNavigation200Response core_navigation_get_apps_navigation(ocs_api_request, absolute)
Get the apps navigation

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**absolute** | Option<**i32**> | Rewrite URLs to absolute ones |  |[default to 0]

### Return type

[**crate::models::CoreNavigationGetAppsNavigation200Response**](core_navigation_get_apps_navigation_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_navigation_get_settings_navigation

> crate::models::CoreNavigationGetAppsNavigation200Response core_navigation_get_settings_navigation(ocs_api_request, absolute)
Get the settings navigation

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**ocs_api_request** | **String** |  | [required] |[default to true]
**absolute** | Option<**i32**> | Rewrite URLs to absolute ones |  |[default to 0]

### Return type

[**crate::models::CoreNavigationGetAppsNavigation200Response**](core_navigation_get_apps_navigation_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

