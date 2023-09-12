# FilesDirectEditingInfo200ResponseOcsData


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**editors** | [**Dict[str, FilesDirectEditingInfo200ResponseOcsDataEditorsValue]**](FilesDirectEditingInfo200ResponseOcsDataEditorsValue.md) |  | 
**creators** | [**Dict[str, FilesDirectEditingInfo200ResponseOcsDataCreatorsValue]**](FilesDirectEditingInfo200ResponseOcsDataCreatorsValue.md) |  | 

## Example

```python
from nextcloud_client.models.files_direct_editing_info200_response_ocs_data import FilesDirectEditingInfo200ResponseOcsData

# TODO update the JSON string below
json = "{}"
# create an instance of FilesDirectEditingInfo200ResponseOcsData from a JSON string
files_direct_editing_info200_response_ocs_data_instance = FilesDirectEditingInfo200ResponseOcsData.from_json(json)
# print the JSON string representation of the object
print FilesDirectEditingInfo200ResponseOcsData.to_json()

# convert the object into a dict
files_direct_editing_info200_response_ocs_data_dict = files_direct_editing_info200_response_ocs_data_instance.to_dict()
# create an instance of FilesDirectEditingInfo200ResponseOcsData from a dict
files_direct_editing_info200_response_ocs_data_form_dict = files_direct_editing_info200_response_ocs_data.from_dict(files_direct_editing_info200_response_ocs_data_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


