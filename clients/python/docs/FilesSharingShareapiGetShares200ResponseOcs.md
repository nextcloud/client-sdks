# FilesSharingShareapiGetShares200ResponseOcs


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**meta** | [**OCSMeta**](OCSMeta.md) |  | 
**data** | [**List[FilesSharingShare]**](FilesSharingShare.md) |  | 

## Example

```python
from nextcloud_client.models.files_sharing_shareapi_get_shares200_response_ocs import FilesSharingShareapiGetShares200ResponseOcs

# TODO update the JSON string below
json = "{}"
# create an instance of FilesSharingShareapiGetShares200ResponseOcs from a JSON string
files_sharing_shareapi_get_shares200_response_ocs_instance = FilesSharingShareapiGetShares200ResponseOcs.from_json(json)
# print the JSON string representation of the object
print FilesSharingShareapiGetShares200ResponseOcs.to_json()

# convert the object into a dict
files_sharing_shareapi_get_shares200_response_ocs_dict = files_sharing_shareapi_get_shares200_response_ocs_instance.to_dict()
# create an instance of FilesSharingShareapiGetShares200ResponseOcs from a dict
files_sharing_shareapi_get_shares200_response_ocs_form_dict = files_sharing_shareapi_get_shares200_response_ocs.from_dict(files_sharing_shareapi_get_shares200_response_ocs_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


