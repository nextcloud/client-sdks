# Oauth2OauthApiGetToken200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**access_token** | **str** |  | 
**token_type** | **str** |  | 
**expires_in** | **int** |  | 
**refresh_token** | **str** |  | 
**user_id** | **str** |  | 

## Example

```python
from openapi_client.models.oauth2_oauth_api_get_token200_response import Oauth2OauthApiGetToken200Response

# TODO update the JSON string below
json = "{}"
# create an instance of Oauth2OauthApiGetToken200Response from a JSON string
oauth2_oauth_api_get_token200_response_instance = Oauth2OauthApiGetToken200Response.from_json(json)
# print the JSON string representation of the object
print Oauth2OauthApiGetToken200Response.to_json()

# convert the object into a dict
oauth2_oauth_api_get_token200_response_dict = oauth2_oauth_api_get_token200_response_instance.to_dict()
# create an instance of Oauth2OauthApiGetToken200Response from a dict
oauth2_oauth_api_get_token200_response_form_dict = oauth2_oauth_api_get_token200_response.from_dict(oauth2_oauth_api_get_token200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


