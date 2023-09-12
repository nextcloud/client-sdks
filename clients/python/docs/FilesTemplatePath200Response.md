# FilesTemplatePath200Response


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ocs** | [**FilesTemplatePath200ResponseOcs**](FilesTemplatePath200ResponseOcs.md) |  | 

## Example

```python
from nextcloud_client.models.files_template_path200_response import FilesTemplatePath200Response

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplatePath200Response from a JSON string
files_template_path200_response_instance = FilesTemplatePath200Response.from_json(json)
# print the JSON string representation of the object
print FilesTemplatePath200Response.to_json()

# convert the object into a dict
files_template_path200_response_dict = files_template_path200_response_instance.to_dict()
# create an instance of FilesTemplatePath200Response from a dict
files_template_path200_response_form_dict = files_template_path200_response.from_dict(files_template_path200_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


