# openapi_client.Oauth2OauthApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**oauth2_oauth_api_get_token**](Oauth2OauthApiApi.md#oauth2_oauth_api_get_token) | **POST** /index.php/apps/oauth2/api/v1/token | Get a token


# **oauth2_oauth_api_get_token**
> Oauth2OauthApiGetToken200Response oauth2_oauth_api_get_token(grant_type, code, refresh_token, client_id, client_secret)

Get a token

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import openapi_client
from openapi_client.models.oauth2_oauth_api_get_token200_response import Oauth2OauthApiGetToken200Response
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = openapi_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = openapi_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.Oauth2OauthApiApi(api_client)
    grant_type = 'grant_type_example' # str | Token type that should be granted
    code = 'code_example' # str | Code of the flow
    refresh_token = 'refresh_token_example' # str | Refresh token
    client_id = 'client_id_example' # str | Client ID
    client_secret = 'client_secret_example' # str | Client secret

    try:
        # Get a token
        api_response = api_instance.oauth2_oauth_api_get_token(grant_type, code, refresh_token, client_id, client_secret)
        print("The response of Oauth2OauthApiApi->oauth2_oauth_api_get_token:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling Oauth2OauthApiApi->oauth2_oauth_api_get_token: %s\n" % e)
```


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grant_type** | **str**| Token type that should be granted | 
 **code** | **str**| Code of the flow | 
 **refresh_token** | **str**| Refresh token | 
 **client_id** | **str**| Client ID | 
 **client_secret** | **str**| Client secret | 

### Return type

[**Oauth2OauthApiGetToken200Response**](Oauth2OauthApiGetToken200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Token returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

