# FilesTemplateFileCreator


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**app** | **str** |  | 
**label** | **str** |  | 
**extension** | **str** |  | 
**icon_class** | **str** |  | 
**mimetypes** | **List[str]** |  | 
**ratio** | **float** |  | 
**action_label** | **str** |  | 

## Example

```python
from nextcloud_client.models.files_template_file_creator import FilesTemplateFileCreator

# TODO update the JSON string below
json = "{}"
# create an instance of FilesTemplateFileCreator from a JSON string
files_template_file_creator_instance = FilesTemplateFileCreator.from_json(json)
# print the JSON string representation of the object
print FilesTemplateFileCreator.to_json()

# convert the object into a dict
files_template_file_creator_dict = files_template_file_creator_instance.to_dict()
# create an instance of FilesTemplateFileCreator from a dict
files_template_file_creator_form_dict = files_template_file_creator.from_dict(files_template_file_creator_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


