# FilesDirectEditingTemplates200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**FilesDirectEditingTemplates200ResponseOcs**](FilesDirectEditingTemplates200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.files_direct_editing_templates200_response import FilesDirectEditingTemplates200Response

# TODO update the JSON string below
json = "{}"
# create an instance of FilesDirectEditingTemplates200Response from a JSON string
files_direct_editing_templates200_response_instance = FilesDirectEditingTemplates200Response.from_json(json)
# print the JSON string representation of the object
print FilesDirectEditingTemplates200Response.to_json()

# convert the object into a dict
files_direct_editing_templates200_response_dict = files_direct_editing_templates200_response_instance.to_dict()
# create an instance of FilesDirectEditingTemplates200Response from a dict
files_direct_editing_templates200_response_form_dict = files_direct_editing_templates200_response.from_dict(files_direct_editing_templates200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


