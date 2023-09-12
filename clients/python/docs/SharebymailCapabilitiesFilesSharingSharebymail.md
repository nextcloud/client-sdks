# SharebymailCapabilitiesFilesSharingSharebymail


## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**enabled** | **bool** |  | 
**send_password_by_mail** | **bool** |  | 
**upload_files_drop** | [**FilesSharingCapabilitiesFilesSharingUserExpireDate**](FilesSharingCapabilitiesFilesSharingUserExpireDate.md) |  | 
**password** | [**SharebymailCapabilitiesFilesSharingSharebymailPassword**](SharebymailCapabilitiesFilesSharingSharebymailPassword.md) |  | 
**expire_date** | [**SharebymailCapabilitiesFilesSharingSharebymailPassword**](SharebymailCapabilitiesFilesSharingSharebymailPassword.md) |  | 

## Example

```python
from nextcloud_client.models.sharebymail_capabilities_files_sharing_sharebymail import SharebymailCapabilitiesFilesSharingSharebymail

# TODO update the JSON string below
json = "{}"
# create an instance of SharebymailCapabilitiesFilesSharingSharebymail from a JSON string
sharebymail_capabilities_files_sharing_sharebymail_instance = SharebymailCapabilitiesFilesSharingSharebymail.from_json(json)
# print the JSON string representation of the object
print SharebymailCapabilitiesFilesSharingSharebymail.to_json()

# convert the object into a dict
sharebymail_capabilities_files_sharing_sharebymail_dict = sharebymail_capabilities_files_sharing_sharebymail_instance.to_dict()
# create an instance of SharebymailCapabilitiesFilesSharingSharebymail from a dict
sharebymail_capabilities_files_sharing_sharebymail_form_dict = sharebymail_capabilities_files_sharing_sharebymail.from_dict(sharebymail_capabilities_files_sharing_sharebymail_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


