# FilesTemplateFile


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**basename** | **str** |  | 
**etag** | **str** |  | 
**fileid** | **int** |  | 
**filename** | **str** |  | 
**lastmod** | **int** |  | 
**mime** | **str** |  | 
**size** | **int** |  | 
**type** | **str** |  | 
**has_preview** | **bool** |  | 

## Example

```python
from nextcloud_client.models.files_template_file import FilesTemplateFile

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplateFile from a JSON string
files_template_file_instance = FilesTemplateFile.from_json(json)
# print the JSON string representation of the object
print FilesTemplateFile.to_json()

# convert the object into a dict
files_template_file_dict = files_template_file_instance.to_dict()
# create an instance of FilesTemplateFile from a dict
files_template_file_form_dict = files_template_file.from_dict(files_template_file_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


