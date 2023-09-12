# ProvisioningApiAppInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **NullableBool** |  | 
**Activity** | **map[string]interface{}** |  | 
**Author** | **map[string]interface{}** |  | 
**BackgroundJobs** | **map[string]interface{}** |  | 
**Bugs** | **map[string]interface{}** |  | 
**Category** | **map[string]interface{}** |  | 
**Collaboration** | **map[string]interface{}** |  | 
**Commands** | **map[string]interface{}** |  | 
**DefaultEnable** | **map[string]interface{}** |  | 
**Dependencies** | **map[string]interface{}** |  | 
**Description** | **string** |  | 
**Discussion** | **map[string]interface{}** |  | 
**Documentation** | **map[string]interface{}** |  | 
**Groups** | **map[string]interface{}** |  | 
**Id** | **string** |  | 
**Info** | **map[string]interface{}** |  | 
**Internal** | **NullableBool** |  | 
**Level** | **NullableInt64** |  | 
**Licence** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Namespace** | **map[string]interface{}** |  | 
**Navigations** | **map[string]interface{}** |  | 
**Preview** | **map[string]interface{}** |  | 
**PreviewAsIcon** | **NullableBool** |  | 
**Public** | **map[string]interface{}** |  | 
**Remote** | **map[string]interface{}** |  | 
**Removable** | **NullableBool** |  | 
**RepairSteps** | **map[string]interface{}** |  | 
**Repository** | **map[string]interface{}** |  | 
**Sabre** | **map[string]interface{}** |  | 
**Screenshot** | **map[string]interface{}** |  | 
**Settings** | **map[string]interface{}** |  | 
**Summary** | **string** |  | 
**Trash** | **map[string]interface{}** |  | 
**TwoFactorProviders** | **map[string]interface{}** |  | 
**Types** | **map[string]interface{}** |  | 
**Version** | **string** |  | 
**Versions** | **map[string]interface{}** |  | 
**Website** | **map[string]interface{}** |  | 

## Methods

### NewProvisioningApiAppInfo

`func NewProvisioningApiAppInfo(active NullableBool, activity map[string]interface{}, author map[string]interface{}, backgroundJobs map[string]interface{}, bugs map[string]interface{}, category map[string]interface{}, collaboration map[string]interface{}, commands map[string]interface{}, defaultEnable map[string]interface{}, dependencies map[string]interface{}, description string, discussion map[string]interface{}, documentation map[string]interface{}, groups map[string]interface{}, id string, info map[string]interface{}, internal NullableBool, level NullableInt64, licence map[string]interface{}, name string, namespace map[string]interface{}, navigations map[string]interface{}, preview map[string]interface{}, previewAsIcon NullableBool, public map[string]interface{}, remote map[string]interface{}, removable NullableBool, repairSteps map[string]interface{}, repository map[string]interface{}, sabre map[string]interface{}, screenshot map[string]interface{}, settings map[string]interface{}, summary string, trash map[string]interface{}, twoFactorProviders map[string]interface{}, types map[string]interface{}, version string, versions map[string]interface{}, website map[string]interface{}, ) *ProvisioningApiAppInfo`

NewProvisioningApiAppInfo instantiates a new ProvisioningApiAppInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningApiAppInfoWithDefaults

`func NewProvisioningApiAppInfoWithDefaults() *ProvisioningApiAppInfo`

NewProvisioningApiAppInfoWithDefaults instantiates a new ProvisioningApiAppInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ProvisioningApiAppInfo) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ProvisioningApiAppInfo) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ProvisioningApiAppInfo) SetActive(v bool)`

SetActive sets Active field to given value.


### SetActiveNil

`func (o *ProvisioningApiAppInfo) SetActiveNil(b bool)`

 SetActiveNil sets the value for Active to be an explicit nil

### UnsetActive
`func (o *ProvisioningApiAppInfo) UnsetActive()`

UnsetActive ensures that no value is present for Active, not even an explicit nil
### GetActivity

`func (o *ProvisioningApiAppInfo) GetActivity() map[string]interface{}`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *ProvisioningApiAppInfo) GetActivityOk() (*map[string]interface{}, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *ProvisioningApiAppInfo) SetActivity(v map[string]interface{})`

SetActivity sets Activity field to given value.


### SetActivityNil

`func (o *ProvisioningApiAppInfo) SetActivityNil(b bool)`

 SetActivityNil sets the value for Activity to be an explicit nil

### UnsetActivity
`func (o *ProvisioningApiAppInfo) UnsetActivity()`

UnsetActivity ensures that no value is present for Activity, not even an explicit nil
### GetAuthor

`func (o *ProvisioningApiAppInfo) GetAuthor() map[string]interface{}`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ProvisioningApiAppInfo) GetAuthorOk() (*map[string]interface{}, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ProvisioningApiAppInfo) SetAuthor(v map[string]interface{})`

SetAuthor sets Author field to given value.


### SetAuthorNil

`func (o *ProvisioningApiAppInfo) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *ProvisioningApiAppInfo) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetBackgroundJobs

`func (o *ProvisioningApiAppInfo) GetBackgroundJobs() map[string]interface{}`

GetBackgroundJobs returns the BackgroundJobs field if non-nil, zero value otherwise.

### GetBackgroundJobsOk

`func (o *ProvisioningApiAppInfo) GetBackgroundJobsOk() (*map[string]interface{}, bool)`

GetBackgroundJobsOk returns a tuple with the BackgroundJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundJobs

`func (o *ProvisioningApiAppInfo) SetBackgroundJobs(v map[string]interface{})`

SetBackgroundJobs sets BackgroundJobs field to given value.


### SetBackgroundJobsNil

`func (o *ProvisioningApiAppInfo) SetBackgroundJobsNil(b bool)`

 SetBackgroundJobsNil sets the value for BackgroundJobs to be an explicit nil

### UnsetBackgroundJobs
`func (o *ProvisioningApiAppInfo) UnsetBackgroundJobs()`

UnsetBackgroundJobs ensures that no value is present for BackgroundJobs, not even an explicit nil
### GetBugs

`func (o *ProvisioningApiAppInfo) GetBugs() map[string]interface{}`

GetBugs returns the Bugs field if non-nil, zero value otherwise.

### GetBugsOk

`func (o *ProvisioningApiAppInfo) GetBugsOk() (*map[string]interface{}, bool)`

GetBugsOk returns a tuple with the Bugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugs

`func (o *ProvisioningApiAppInfo) SetBugs(v map[string]interface{})`

SetBugs sets Bugs field to given value.


### SetBugsNil

`func (o *ProvisioningApiAppInfo) SetBugsNil(b bool)`

 SetBugsNil sets the value for Bugs to be an explicit nil

### UnsetBugs
`func (o *ProvisioningApiAppInfo) UnsetBugs()`

UnsetBugs ensures that no value is present for Bugs, not even an explicit nil
### GetCategory

`func (o *ProvisioningApiAppInfo) GetCategory() map[string]interface{}`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ProvisioningApiAppInfo) GetCategoryOk() (*map[string]interface{}, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ProvisioningApiAppInfo) SetCategory(v map[string]interface{})`

SetCategory sets Category field to given value.


### SetCategoryNil

`func (o *ProvisioningApiAppInfo) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ProvisioningApiAppInfo) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCollaboration

`func (o *ProvisioningApiAppInfo) GetCollaboration() map[string]interface{}`

GetCollaboration returns the Collaboration field if non-nil, zero value otherwise.

### GetCollaborationOk

`func (o *ProvisioningApiAppInfo) GetCollaborationOk() (*map[string]interface{}, bool)`

GetCollaborationOk returns a tuple with the Collaboration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaboration

`func (o *ProvisioningApiAppInfo) SetCollaboration(v map[string]interface{})`

SetCollaboration sets Collaboration field to given value.


### SetCollaborationNil

`func (o *ProvisioningApiAppInfo) SetCollaborationNil(b bool)`

 SetCollaborationNil sets the value for Collaboration to be an explicit nil

### UnsetCollaboration
`func (o *ProvisioningApiAppInfo) UnsetCollaboration()`

UnsetCollaboration ensures that no value is present for Collaboration, not even an explicit nil
### GetCommands

`func (o *ProvisioningApiAppInfo) GetCommands() map[string]interface{}`

GetCommands returns the Commands field if non-nil, zero value otherwise.

### GetCommandsOk

`func (o *ProvisioningApiAppInfo) GetCommandsOk() (*map[string]interface{}, bool)`

GetCommandsOk returns a tuple with the Commands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommands

`func (o *ProvisioningApiAppInfo) SetCommands(v map[string]interface{})`

SetCommands sets Commands field to given value.


### SetCommandsNil

`func (o *ProvisioningApiAppInfo) SetCommandsNil(b bool)`

 SetCommandsNil sets the value for Commands to be an explicit nil

### UnsetCommands
`func (o *ProvisioningApiAppInfo) UnsetCommands()`

UnsetCommands ensures that no value is present for Commands, not even an explicit nil
### GetDefaultEnable

`func (o *ProvisioningApiAppInfo) GetDefaultEnable() map[string]interface{}`

GetDefaultEnable returns the DefaultEnable field if non-nil, zero value otherwise.

### GetDefaultEnableOk

`func (o *ProvisioningApiAppInfo) GetDefaultEnableOk() (*map[string]interface{}, bool)`

GetDefaultEnableOk returns a tuple with the DefaultEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEnable

`func (o *ProvisioningApiAppInfo) SetDefaultEnable(v map[string]interface{})`

SetDefaultEnable sets DefaultEnable field to given value.


### SetDefaultEnableNil

`func (o *ProvisioningApiAppInfo) SetDefaultEnableNil(b bool)`

 SetDefaultEnableNil sets the value for DefaultEnable to be an explicit nil

### UnsetDefaultEnable
`func (o *ProvisioningApiAppInfo) UnsetDefaultEnable()`

UnsetDefaultEnable ensures that no value is present for DefaultEnable, not even an explicit nil
### GetDependencies

`func (o *ProvisioningApiAppInfo) GetDependencies() map[string]interface{}`

GetDependencies returns the Dependencies field if non-nil, zero value otherwise.

### GetDependenciesOk

`func (o *ProvisioningApiAppInfo) GetDependenciesOk() (*map[string]interface{}, bool)`

GetDependenciesOk returns a tuple with the Dependencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencies

`func (o *ProvisioningApiAppInfo) SetDependencies(v map[string]interface{})`

SetDependencies sets Dependencies field to given value.


### SetDependenciesNil

`func (o *ProvisioningApiAppInfo) SetDependenciesNil(b bool)`

 SetDependenciesNil sets the value for Dependencies to be an explicit nil

### UnsetDependencies
`func (o *ProvisioningApiAppInfo) UnsetDependencies()`

UnsetDependencies ensures that no value is present for Dependencies, not even an explicit nil
### GetDescription

`func (o *ProvisioningApiAppInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProvisioningApiAppInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProvisioningApiAppInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDiscussion

`func (o *ProvisioningApiAppInfo) GetDiscussion() map[string]interface{}`

GetDiscussion returns the Discussion field if non-nil, zero value otherwise.

### GetDiscussionOk

`func (o *ProvisioningApiAppInfo) GetDiscussionOk() (*map[string]interface{}, bool)`

GetDiscussionOk returns a tuple with the Discussion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscussion

`func (o *ProvisioningApiAppInfo) SetDiscussion(v map[string]interface{})`

SetDiscussion sets Discussion field to given value.


### SetDiscussionNil

`func (o *ProvisioningApiAppInfo) SetDiscussionNil(b bool)`

 SetDiscussionNil sets the value for Discussion to be an explicit nil

### UnsetDiscussion
`func (o *ProvisioningApiAppInfo) UnsetDiscussion()`

UnsetDiscussion ensures that no value is present for Discussion, not even an explicit nil
### GetDocumentation

`func (o *ProvisioningApiAppInfo) GetDocumentation() map[string]interface{}`

GetDocumentation returns the Documentation field if non-nil, zero value otherwise.

### GetDocumentationOk

`func (o *ProvisioningApiAppInfo) GetDocumentationOk() (*map[string]interface{}, bool)`

GetDocumentationOk returns a tuple with the Documentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentation

`func (o *ProvisioningApiAppInfo) SetDocumentation(v map[string]interface{})`

SetDocumentation sets Documentation field to given value.


### SetDocumentationNil

`func (o *ProvisioningApiAppInfo) SetDocumentationNil(b bool)`

 SetDocumentationNil sets the value for Documentation to be an explicit nil

### UnsetDocumentation
`func (o *ProvisioningApiAppInfo) UnsetDocumentation()`

UnsetDocumentation ensures that no value is present for Documentation, not even an explicit nil
### GetGroups

`func (o *ProvisioningApiAppInfo) GetGroups() map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProvisioningApiAppInfo) GetGroupsOk() (*map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProvisioningApiAppInfo) SetGroups(v map[string]interface{})`

SetGroups sets Groups field to given value.


### SetGroupsNil

`func (o *ProvisioningApiAppInfo) SetGroupsNil(b bool)`

 SetGroupsNil sets the value for Groups to be an explicit nil

### UnsetGroups
`func (o *ProvisioningApiAppInfo) UnsetGroups()`

UnsetGroups ensures that no value is present for Groups, not even an explicit nil
### GetId

`func (o *ProvisioningApiAppInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningApiAppInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningApiAppInfo) SetId(v string)`

SetId sets Id field to given value.


### GetInfo

`func (o *ProvisioningApiAppInfo) GetInfo() map[string]interface{}`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ProvisioningApiAppInfo) GetInfoOk() (*map[string]interface{}, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ProvisioningApiAppInfo) SetInfo(v map[string]interface{})`

SetInfo sets Info field to given value.


### SetInfoNil

`func (o *ProvisioningApiAppInfo) SetInfoNil(b bool)`

 SetInfoNil sets the value for Info to be an explicit nil

### UnsetInfo
`func (o *ProvisioningApiAppInfo) UnsetInfo()`

UnsetInfo ensures that no value is present for Info, not even an explicit nil
### GetInternal

`func (o *ProvisioningApiAppInfo) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ProvisioningApiAppInfo) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ProvisioningApiAppInfo) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### SetInternalNil

`func (o *ProvisioningApiAppInfo) SetInternalNil(b bool)`

 SetInternalNil sets the value for Internal to be an explicit nil

### UnsetInternal
`func (o *ProvisioningApiAppInfo) UnsetInternal()`

UnsetInternal ensures that no value is present for Internal, not even an explicit nil
### GetLevel

`func (o *ProvisioningApiAppInfo) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ProvisioningApiAppInfo) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ProvisioningApiAppInfo) SetLevel(v int64)`

SetLevel sets Level field to given value.


### SetLevelNil

`func (o *ProvisioningApiAppInfo) SetLevelNil(b bool)`

 SetLevelNil sets the value for Level to be an explicit nil

### UnsetLevel
`func (o *ProvisioningApiAppInfo) UnsetLevel()`

UnsetLevel ensures that no value is present for Level, not even an explicit nil
### GetLicence

`func (o *ProvisioningApiAppInfo) GetLicence() map[string]interface{}`

GetLicence returns the Licence field if non-nil, zero value otherwise.

### GetLicenceOk

`func (o *ProvisioningApiAppInfo) GetLicenceOk() (*map[string]interface{}, bool)`

GetLicenceOk returns a tuple with the Licence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicence

`func (o *ProvisioningApiAppInfo) SetLicence(v map[string]interface{})`

SetLicence sets Licence field to given value.


### SetLicenceNil

`func (o *ProvisioningApiAppInfo) SetLicenceNil(b bool)`

 SetLicenceNil sets the value for Licence to be an explicit nil

### UnsetLicence
`func (o *ProvisioningApiAppInfo) UnsetLicence()`

UnsetLicence ensures that no value is present for Licence, not even an explicit nil
### GetName

`func (o *ProvisioningApiAppInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProvisioningApiAppInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProvisioningApiAppInfo) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *ProvisioningApiAppInfo) GetNamespace() map[string]interface{}`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ProvisioningApiAppInfo) GetNamespaceOk() (*map[string]interface{}, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ProvisioningApiAppInfo) SetNamespace(v map[string]interface{})`

SetNamespace sets Namespace field to given value.


### SetNamespaceNil

`func (o *ProvisioningApiAppInfo) SetNamespaceNil(b bool)`

 SetNamespaceNil sets the value for Namespace to be an explicit nil

### UnsetNamespace
`func (o *ProvisioningApiAppInfo) UnsetNamespace()`

UnsetNamespace ensures that no value is present for Namespace, not even an explicit nil
### GetNavigations

`func (o *ProvisioningApiAppInfo) GetNavigations() map[string]interface{}`

GetNavigations returns the Navigations field if non-nil, zero value otherwise.

### GetNavigationsOk

`func (o *ProvisioningApiAppInfo) GetNavigationsOk() (*map[string]interface{}, bool)`

GetNavigationsOk returns a tuple with the Navigations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigations

`func (o *ProvisioningApiAppInfo) SetNavigations(v map[string]interface{})`

SetNavigations sets Navigations field to given value.


### SetNavigationsNil

`func (o *ProvisioningApiAppInfo) SetNavigationsNil(b bool)`

 SetNavigationsNil sets the value for Navigations to be an explicit nil

### UnsetNavigations
`func (o *ProvisioningApiAppInfo) UnsetNavigations()`

UnsetNavigations ensures that no value is present for Navigations, not even an explicit nil
### GetPreview

`func (o *ProvisioningApiAppInfo) GetPreview() map[string]interface{}`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *ProvisioningApiAppInfo) GetPreviewOk() (*map[string]interface{}, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *ProvisioningApiAppInfo) SetPreview(v map[string]interface{})`

SetPreview sets Preview field to given value.


### SetPreviewNil

`func (o *ProvisioningApiAppInfo) SetPreviewNil(b bool)`

 SetPreviewNil sets the value for Preview to be an explicit nil

### UnsetPreview
`func (o *ProvisioningApiAppInfo) UnsetPreview()`

UnsetPreview ensures that no value is present for Preview, not even an explicit nil
### GetPreviewAsIcon

`func (o *ProvisioningApiAppInfo) GetPreviewAsIcon() bool`

GetPreviewAsIcon returns the PreviewAsIcon field if non-nil, zero value otherwise.

### GetPreviewAsIconOk

`func (o *ProvisioningApiAppInfo) GetPreviewAsIconOk() (*bool, bool)`

GetPreviewAsIconOk returns a tuple with the PreviewAsIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewAsIcon

`func (o *ProvisioningApiAppInfo) SetPreviewAsIcon(v bool)`

SetPreviewAsIcon sets PreviewAsIcon field to given value.


### SetPreviewAsIconNil

`func (o *ProvisioningApiAppInfo) SetPreviewAsIconNil(b bool)`

 SetPreviewAsIconNil sets the value for PreviewAsIcon to be an explicit nil

### UnsetPreviewAsIcon
`func (o *ProvisioningApiAppInfo) UnsetPreviewAsIcon()`

UnsetPreviewAsIcon ensures that no value is present for PreviewAsIcon, not even an explicit nil
### GetPublic

`func (o *ProvisioningApiAppInfo) GetPublic() map[string]interface{}`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *ProvisioningApiAppInfo) GetPublicOk() (*map[string]interface{}, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *ProvisioningApiAppInfo) SetPublic(v map[string]interface{})`

SetPublic sets Public field to given value.


### SetPublicNil

`func (o *ProvisioningApiAppInfo) SetPublicNil(b bool)`

 SetPublicNil sets the value for Public to be an explicit nil

### UnsetPublic
`func (o *ProvisioningApiAppInfo) UnsetPublic()`

UnsetPublic ensures that no value is present for Public, not even an explicit nil
### GetRemote

`func (o *ProvisioningApiAppInfo) GetRemote() map[string]interface{}`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ProvisioningApiAppInfo) GetRemoteOk() (*map[string]interface{}, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ProvisioningApiAppInfo) SetRemote(v map[string]interface{})`

SetRemote sets Remote field to given value.


### SetRemoteNil

`func (o *ProvisioningApiAppInfo) SetRemoteNil(b bool)`

 SetRemoteNil sets the value for Remote to be an explicit nil

### UnsetRemote
`func (o *ProvisioningApiAppInfo) UnsetRemote()`

UnsetRemote ensures that no value is present for Remote, not even an explicit nil
### GetRemovable

`func (o *ProvisioningApiAppInfo) GetRemovable() bool`

GetRemovable returns the Removable field if non-nil, zero value otherwise.

### GetRemovableOk

`func (o *ProvisioningApiAppInfo) GetRemovableOk() (*bool, bool)`

GetRemovableOk returns a tuple with the Removable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovable

`func (o *ProvisioningApiAppInfo) SetRemovable(v bool)`

SetRemovable sets Removable field to given value.


### SetRemovableNil

`func (o *ProvisioningApiAppInfo) SetRemovableNil(b bool)`

 SetRemovableNil sets the value for Removable to be an explicit nil

### UnsetRemovable
`func (o *ProvisioningApiAppInfo) UnsetRemovable()`

UnsetRemovable ensures that no value is present for Removable, not even an explicit nil
### GetRepairSteps

`func (o *ProvisioningApiAppInfo) GetRepairSteps() map[string]interface{}`

GetRepairSteps returns the RepairSteps field if non-nil, zero value otherwise.

### GetRepairStepsOk

`func (o *ProvisioningApiAppInfo) GetRepairStepsOk() (*map[string]interface{}, bool)`

GetRepairStepsOk returns a tuple with the RepairSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairSteps

`func (o *ProvisioningApiAppInfo) SetRepairSteps(v map[string]interface{})`

SetRepairSteps sets RepairSteps field to given value.


### SetRepairStepsNil

`func (o *ProvisioningApiAppInfo) SetRepairStepsNil(b bool)`

 SetRepairStepsNil sets the value for RepairSteps to be an explicit nil

### UnsetRepairSteps
`func (o *ProvisioningApiAppInfo) UnsetRepairSteps()`

UnsetRepairSteps ensures that no value is present for RepairSteps, not even an explicit nil
### GetRepository

`func (o *ProvisioningApiAppInfo) GetRepository() map[string]interface{}`

GetRepository returns the Repository field if non-nil, zero value otherwise.

### GetRepositoryOk

`func (o *ProvisioningApiAppInfo) GetRepositoryOk() (*map[string]interface{}, bool)`

GetRepositoryOk returns a tuple with the Repository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepository

`func (o *ProvisioningApiAppInfo) SetRepository(v map[string]interface{})`

SetRepository sets Repository field to given value.


### SetRepositoryNil

`func (o *ProvisioningApiAppInfo) SetRepositoryNil(b bool)`

 SetRepositoryNil sets the value for Repository to be an explicit nil

### UnsetRepository
`func (o *ProvisioningApiAppInfo) UnsetRepository()`

UnsetRepository ensures that no value is present for Repository, not even an explicit nil
### GetSabre

`func (o *ProvisioningApiAppInfo) GetSabre() map[string]interface{}`

GetSabre returns the Sabre field if non-nil, zero value otherwise.

### GetSabreOk

`func (o *ProvisioningApiAppInfo) GetSabreOk() (*map[string]interface{}, bool)`

GetSabreOk returns a tuple with the Sabre field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSabre

`func (o *ProvisioningApiAppInfo) SetSabre(v map[string]interface{})`

SetSabre sets Sabre field to given value.


### SetSabreNil

`func (o *ProvisioningApiAppInfo) SetSabreNil(b bool)`

 SetSabreNil sets the value for Sabre to be an explicit nil

### UnsetSabre
`func (o *ProvisioningApiAppInfo) UnsetSabre()`

UnsetSabre ensures that no value is present for Sabre, not even an explicit nil
### GetScreenshot

`func (o *ProvisioningApiAppInfo) GetScreenshot() map[string]interface{}`

GetScreenshot returns the Screenshot field if non-nil, zero value otherwise.

### GetScreenshotOk

`func (o *ProvisioningApiAppInfo) GetScreenshotOk() (*map[string]interface{}, bool)`

GetScreenshotOk returns a tuple with the Screenshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshot

`func (o *ProvisioningApiAppInfo) SetScreenshot(v map[string]interface{})`

SetScreenshot sets Screenshot field to given value.


### SetScreenshotNil

`func (o *ProvisioningApiAppInfo) SetScreenshotNil(b bool)`

 SetScreenshotNil sets the value for Screenshot to be an explicit nil

### UnsetScreenshot
`func (o *ProvisioningApiAppInfo) UnsetScreenshot()`

UnsetScreenshot ensures that no value is present for Screenshot, not even an explicit nil
### GetSettings

`func (o *ProvisioningApiAppInfo) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *ProvisioningApiAppInfo) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *ProvisioningApiAppInfo) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### SetSettingsNil

`func (o *ProvisioningApiAppInfo) SetSettingsNil(b bool)`

 SetSettingsNil sets the value for Settings to be an explicit nil

### UnsetSettings
`func (o *ProvisioningApiAppInfo) UnsetSettings()`

UnsetSettings ensures that no value is present for Settings, not even an explicit nil
### GetSummary

`func (o *ProvisioningApiAppInfo) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ProvisioningApiAppInfo) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ProvisioningApiAppInfo) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetTrash

`func (o *ProvisioningApiAppInfo) GetTrash() map[string]interface{}`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *ProvisioningApiAppInfo) GetTrashOk() (*map[string]interface{}, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *ProvisioningApiAppInfo) SetTrash(v map[string]interface{})`

SetTrash sets Trash field to given value.


### SetTrashNil

`func (o *ProvisioningApiAppInfo) SetTrashNil(b bool)`

 SetTrashNil sets the value for Trash to be an explicit nil

### UnsetTrash
`func (o *ProvisioningApiAppInfo) UnsetTrash()`

UnsetTrash ensures that no value is present for Trash, not even an explicit nil
### GetTwoFactorProviders

`func (o *ProvisioningApiAppInfo) GetTwoFactorProviders() map[string]interface{}`

GetTwoFactorProviders returns the TwoFactorProviders field if non-nil, zero value otherwise.

### GetTwoFactorProvidersOk

`func (o *ProvisioningApiAppInfo) GetTwoFactorProvidersOk() (*map[string]interface{}, bool)`

GetTwoFactorProvidersOk returns a tuple with the TwoFactorProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwoFactorProviders

`func (o *ProvisioningApiAppInfo) SetTwoFactorProviders(v map[string]interface{})`

SetTwoFactorProviders sets TwoFactorProviders field to given value.


### SetTwoFactorProvidersNil

`func (o *ProvisioningApiAppInfo) SetTwoFactorProvidersNil(b bool)`

 SetTwoFactorProvidersNil sets the value for TwoFactorProviders to be an explicit nil

### UnsetTwoFactorProviders
`func (o *ProvisioningApiAppInfo) UnsetTwoFactorProviders()`

UnsetTwoFactorProviders ensures that no value is present for TwoFactorProviders, not even an explicit nil
### GetTypes

`func (o *ProvisioningApiAppInfo) GetTypes() map[string]interface{}`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *ProvisioningApiAppInfo) GetTypesOk() (*map[string]interface{}, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *ProvisioningApiAppInfo) SetTypes(v map[string]interface{})`

SetTypes sets Types field to given value.


### SetTypesNil

`func (o *ProvisioningApiAppInfo) SetTypesNil(b bool)`

 SetTypesNil sets the value for Types to be an explicit nil

### UnsetTypes
`func (o *ProvisioningApiAppInfo) UnsetTypes()`

UnsetTypes ensures that no value is present for Types, not even an explicit nil
### GetVersion

`func (o *ProvisioningApiAppInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisioningApiAppInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisioningApiAppInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVersions

`func (o *ProvisioningApiAppInfo) GetVersions() map[string]interface{}`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ProvisioningApiAppInfo) GetVersionsOk() (*map[string]interface{}, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ProvisioningApiAppInfo) SetVersions(v map[string]interface{})`

SetVersions sets Versions field to given value.


### SetVersionsNil

`func (o *ProvisioningApiAppInfo) SetVersionsNil(b bool)`

 SetVersionsNil sets the value for Versions to be an explicit nil

### UnsetVersions
`func (o *ProvisioningApiAppInfo) UnsetVersions()`

UnsetVersions ensures that no value is present for Versions, not even an explicit nil
### GetWebsite

`func (o *ProvisioningApiAppInfo) GetWebsite() map[string]interface{}`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProvisioningApiAppInfo) GetWebsiteOk() (*map[string]interface{}, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProvisioningApiAppInfo) SetWebsite(v map[string]interface{})`

SetWebsite sets Website field to given value.


### SetWebsiteNil

`func (o *ProvisioningApiAppInfo) SetWebsiteNil(b bool)`

 SetWebsiteNil sets the value for Website to be an explicit nil

### UnsetWebsite
`func (o *ProvisioningApiAppInfo) UnsetWebsite()`

UnsetWebsite ensures that no value is present for Website, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


