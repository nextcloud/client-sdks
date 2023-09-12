# FilesTemplatePath200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**template_path** | **str** |  | 
**templates** | [**List[FilesTemplateFileCreator]**](FilesTemplateFileCreator.md) |  | 

## Example

```python
from nextcloud_client.models.files_template_path200_response_ocs_data import FilesTemplatePath200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplatePath200ResponseOcsData from a JSON string
files_template_path200_response_ocs_data_instance = FilesTemplatePath200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print FilesTemplatePath200ResponseOcsData.to_json()

# convert the object into a dict
files_template_path200_response_ocs_data_dict = files_template_path200_response_ocs_data_instance.to_dict()
# create an instance of FilesTemplatePath200ResponseOcsData from a dict
files_template_path200_response_ocs_data_form_dict = files_template_path200_response_ocs_data.from_dict(files_template_path200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


