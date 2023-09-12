# CoreLoginFlowV2Credentials


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**server** | **str** |  | 
**login_name** | **str** |  | 
**app_password** | **str** |  | 

## Example

```python
from nextcloud_client.models.core_login_flow_v2_credentials import CoreLoginFlowV2Credentials

# TODO update the JSON string below
json = "{}"
# create an instance of CoreLoginFlowV2Credentials from a JSON string
core_login_flow_v2_credentials_instance = CoreLoginFlowV2Credentials.from_json(json)
# print the JSON string representation of the object
print CoreLoginFlowV2Credentials.to_json()

# convert the object into a dict
core_login_flow_v2_credentials_dict = core_login_flow_v2_credentials_instance.to_dict()
# create an instance of CoreLoginFlowV2Credentials from a dict
core_login_flow_v2_credentials_form_dict = core_login_flow_v2_credentials.from_dict(core_login_flow_v2_credentials_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


