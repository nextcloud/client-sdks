# FilesTemplateList200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**FilesTemplateList200ResponseOcs**](FilesTemplateList200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.files_template_list200_response import FilesTemplateList200Response

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplateList200Response from a JSON string
files_template_list200_response_instance = FilesTemplateList200Response.from_json(json)
# print the JSON string representation of the object
print FilesTemplateList200Response.to_json()

# convert the object into a dict
files_template_list200_response_dict = files_template_list200_response_instance.to_dict()
# create an instance of FilesTemplateList200Response from a dict
files_template_list200_response_form_dict = files_template_list200_response.from_dict(files_template_list200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


