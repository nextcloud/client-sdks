# nextcloud_client.CoreTranslationApiApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**core_translation_api_languages**](CoreTranslationApiApi.md#core_translation_api_languages) | **GET** /ocs/v2.php/translation/languages | Get the list of supported languages
[**core_translation_api_translate**](CoreTranslationApiApi.md#core_translation_api_translate) | **POST** /ocs/v2.php/translation/translate | Translate a text


# **core_translation_api_languages**
> CoreTranslationApiLanguages200Response core_translation_api_languages(ocs_api_request)

Get the list of supported languages

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_translation_api_languages200_response import CoreTranslationApiLanguages200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.CoreTranslationApiApi(api_client)
    ocs_api_request = 'true' # str |  (default to 'true')

    try:
        # Get the list of supported languages
        api_response = api_instance.core_translation_api_languages(ocs_api_request)
        print("The response of CoreTranslationApiApi->core_translation_api_languages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTranslationApiApi->core_translation_api_languages: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]

### Return type

[**CoreTranslationApiLanguages200Response**](CoreTranslationApiLanguages200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** |  |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **core_translation_api_translate**
> CoreTranslationApiTranslate200Response core_translation_api_translate(text, to_language, ocs_api_request, from_language=from_language)

Translate a text

### Example

* Basic Authentication (basic_auth):
* Bearer Authentication (bearer_auth):
```python
import time
import os
import nextcloud_client
from nextcloud_client.models.core_translation_api_translate200_response import CoreTranslationApiTranslate200Response
from nextcloud_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to http://localhost
# See configuration.py for a list of all supported configuration parameters.
configuration = nextcloud_client.Configuration(
    host = "http://localhost"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure HTTP basic authorization: basic_auth
configuration = nextcloud_client.Configuration(
    username = os.environ["USERNAME"],
    password = os.environ["PASSWORD"]
)

# Configure Bearer authorization: bearer_auth
configuration = nextcloud_client.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with nextcloud_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = nextcloud_client.CoreTranslationApiApi(api_client)
    text = 'text_example' # str | Text to be translated
    to_language = 'to_language_example' # str | Language to translate to
    ocs_api_request = 'true' # str |  (default to 'true')
    from_language = 'from_language_example' # str | Language to translate from (optional)

    try:
        # Translate a text
        api_response = api_instance.core_translation_api_translate(text, to_language, ocs_api_request, from_language=from_language)
        print("The response of CoreTranslationApiApi->core_translation_api_translate:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling CoreTranslationApiApi->core_translation_api_translate: %s\n" % e)
```



### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **text** | **str**| Text to be translated | 
 **to_language** | **str**| Language to translate to | 
 **ocs_api_request** | **str**|  | [default to &#39;true&#39;]
 **from_language** | **str**| Language to translate from | [optional] 

### Return type

[**CoreTranslationApiTranslate200Response**](CoreTranslationApiTranslate200Response.md)

### Authorization

[basic_auth](../README.md#basic_auth), [bearer_auth](../README.md#bearer_auth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details
| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Translated text returned |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

