# \CoreClientFlowLoginV2Api

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_client_flow_login_v2_init**](CoreClientFlowLoginV2Api.md#core_client_flow_login_v2_init) | **POST** /index.php/login/v2 | Init a login flow
[**core_client_flow_login_v2_poll**](CoreClientFlowLoginV2Api.md#core_client_flow_login_v2_poll) | **POST** /index.php/login/v2/poll | Poll the login flow credentials



## core_client_flow_login_v2_init

> crate::models::CoreLoginFlowV2 core_client_flow_login_v2_init()
Init a login flow

### Parameters

This endpoint does not need any parameter.

### Return type

[**crate::models::CoreLoginFlowV2**](CoreLoginFlowV2.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_client_flow_login_v2_poll

> crate::models::CoreLoginFlowV2Credentials core_client_flow_login_v2_poll(token)
Poll the login flow credentials

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**token** | **String** | Token of the flow | [required] |

### Return type

[**crate::models::CoreLoginFlowV2Credentials**](CoreLoginFlowV2Credentials.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

