# \CoreCollaborationResourcesApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_collaboration_resources_add_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_add_resource) | **POST** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Add a resource to a collection
[**core_collaboration_resources_create_collection_on_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_create_collection_on_resource) | **POST** /ocs/v2.php/collaboration/resources/{baseResourceType}/{baseResourceId} | Create a collection for a resource
[**core_collaboration_resources_get_collections_by_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_get_collections_by_resource) | **GET** /ocs/v2.php/collaboration/resources/{resourceType}/{resourceId} | Get collections by resource
[**core_collaboration_resources_list_collection**](CoreCollaborationResourcesApi.md#core_collaboration_resources_list_collection) | **GET** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Get a collection
[**core_collaboration_resources_remove_resource**](CoreCollaborationResourcesApi.md#core_collaboration_resources_remove_resource) | **DELETE** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Remove a resource from a collection
[**core_collaboration_resources_rename_collection**](CoreCollaborationResourcesApi.md#core_collaboration_resources_rename_collection) | **PUT** /ocs/v2.php/collaboration/resources/collections/{collectionId} | Rename a collection
[**core_collaboration_resources_search_collections**](CoreCollaborationResourcesApi.md#core_collaboration_resources_search_collections) | **GET** /ocs/v2.php/collaboration/resources/collections/search/{filter} | Search for collections



## core_collaboration_resources_add_resource

> crate::models::CoreCollaborationResourcesListCollection200Response core_collaboration_resources_add_resource(resource_type, resource_id, collection_id, ocs_api_request)
Add a resource to a collection

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**resource_type** | **String** | Name of the resource | [required] |
**resource_id** | **String** | ID of the resource | [required] |
**collection_id** | **i64** | ID of the collection | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesListCollection200Response**](core_collaboration_resources_list_collection_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_create_collection_on_resource

> crate::models::CoreCollaborationResourcesListCollection200Response core_collaboration_resources_create_collection_on_resource(name, base_resource_type, base_resource_id, ocs_api_request)
Create a collection for a resource

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**name** | **String** | Name of the collection | [required] |
**base_resource_type** | **String** | Type of the base resource | [required] |
**base_resource_id** | **String** | ID of the base resource | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesListCollection200Response**](core_collaboration_resources_list_collection_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_get_collections_by_resource

> crate::models::CoreCollaborationResourcesSearchCollections200Response core_collaboration_resources_get_collections_by_resource(resource_type, resource_id, ocs_api_request)
Get collections by resource

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**resource_type** | **String** | Type of the resource | [required] |
**resource_id** | **String** | ID of the resource | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesSearchCollections200Response**](core_collaboration_resources_search_collections_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_list_collection

> crate::models::CoreCollaborationResourcesListCollection200Response core_collaboration_resources_list_collection(collection_id, ocs_api_request)
Get a collection

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**collection_id** | **i64** | ID of the collection | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesListCollection200Response**](core_collaboration_resources_list_collection_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_remove_resource

> crate::models::CoreCollaborationResourcesListCollection200Response core_collaboration_resources_remove_resource(resource_type, resource_id, collection_id, ocs_api_request)
Remove a resource from a collection

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**resource_type** | **String** | Name of the resource | [required] |
**resource_id** | **String** | ID of the resource | [required] |
**collection_id** | **i64** | ID of the collection | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesListCollection200Response**](core_collaboration_resources_list_collection_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_rename_collection

> crate::models::CoreCollaborationResourcesListCollection200Response core_collaboration_resources_rename_collection(collection_name, collection_id, ocs_api_request)
Rename a collection

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**collection_name** | **String** | New name | [required] |
**collection_id** | **i64** | ID of the collection | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesListCollection200Response**](core_collaboration_resources_list_collection_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


## core_collaboration_resources_search_collections

> crate::models::CoreCollaborationResourcesSearchCollections200Response core_collaboration_resources_search_collections(filter, ocs_api_request)
Search for collections

### Parameters


Name | Type | Description  | Required | Notes
------------- | ------------- | ------------- | ------------- | -------------
**filter** | **String** | Filter collections | [required] |
**ocs_api_request** | **String** |  | [required] |[default to true]

### Return type

[**crate::models::CoreCollaborationResourcesSearchCollections200Response**](core_collaboration_resources_search_collections_200_response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

