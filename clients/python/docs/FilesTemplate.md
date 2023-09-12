# FilesTemplate


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**template_type** | **str** |  | 
**template_id** | **str** |  | 
**basename** | **str** |  | 
**etag** | **str** |  | 
**fileid** | **int** |  | 
**filename** | **str** |  | 
**lastmod** | **int** |  | 
**mime** | **str** |  | 
**size** | **int** |  | 
**type** | **str** |  | 
**has_preview** | **bool** |  | 
**preview_url** | **str** |  | 

## Example

```python
from nextcloud_client.models.files_template import FilesTemplate

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplate from a JSON string
files_template_instance = FilesTemplate.from_json(json)
# print the JSON string representation of the object
print FilesTemplate.to_json()

# convert the object into a dict
files_template_dict = files_template_instance.to_dict()
# create an instance of FilesTemplate from a dict
files_template_form_dict = files_template.from_dict(files_template_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


