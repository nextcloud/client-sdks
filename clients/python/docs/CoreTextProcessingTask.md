# CoreTextProcessingTask


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**id** | **int** |  | 
**type** | **str** |  | 
**status** | **int** |  | 
**user_id** | **str** |  | 
**app_id** | **str** |  | 
**input** | **str** |  | 
**output** | **str** |  | 
**identifier** | **str** |  | 

## Example

```python
from nextcloud_client.models.core_text_processing_task import CoreTextProcessingTask

# TODO update the JSON string below
json = "{}"
# create an instance of CoreTextProcessingTask from a JSON string
core_text_processing_task_instance = CoreTextProcessingTask.from_json(json)
# print the JSON string representation of the object
print CoreTextProcessingTask.to_json()

# convert the object into a dict
core_text_processing_task_dict = core_text_processing_task_instance.to_dict()
# create an instance of CoreTextProcessingTask from a dict
core_text_processing_task_form_dict = core_text_processing_task.from_dict(core_text_processing_task_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


