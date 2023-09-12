# CommentsCapabilities


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**files** | [**CommentsCapabilitiesFiles**](CommentsCapabilitiesFiles.md) |  | 

## Example

```python
from nextcloud_client.models.comments_capabilities import CommentsCapabilities

# TODO update the JSON string below
json = "{}"
# create an instance of CommentsCapabilities from a JSON string
comments_capabilities_instance = CommentsCapabilities.from_json(json)
# print the JSON string representation of the object
print CommentsCapabilities.to_json()

# convert the object into a dict
comments_capabilities_dict = comments_capabilities_instance.to_dict()
# create an instance of CommentsCapabilities from a dict
comments_capabilities_form_dict = comments_capabilities.from_dict(comments_capabilities_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


