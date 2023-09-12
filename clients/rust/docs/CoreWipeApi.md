# \CoreWipeApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_wipe_check_wipe**](CoreWipeApi.md#core_wipe_check_wipe) | **POST** /index.php/core/wipe/check | Check if the device should be wiped
[**core_wipe_wipe_done**](CoreWipeApi.md#core_wipe_wipe_done) | **POST** /index.php/core/wipe/success | Finish the wipe



## core_wipe_check_wipe

> crate::models::CoreWipeCheckWipe200Response core_wipe_check_wipe(token)
Check if the device should be wiped

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**token** | **String** | App password | [required] |

### Return type

[**crate::models::CoreWipeCheckWipe200Response**](core_wipe_check_wipe_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_wipe_wipe_done

> core_wipe_wipe_done(token)
Finish the wipe

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**token** | **String** | App password | [required] |

### Return type

 (empty response body)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

