# ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalMail** | **[]string** |  | 
**AdditionalMailScope** | **[]string** |  | 
**Address** | **string** |  | 
**AddressScope** | **NullableString** |  | 
**AvatarScope** | **NullableString** |  | 
**Backend** | **string** |  | 
**BackendCapabilities** | [**ProvisioningApiUserDetailsBackendCapabilities**](ProvisioningApiUserDetailsBackendCapabilities.md) |  | 
**Biography** | **string** |  | 
**BiographyScope** | **NullableString** |  | 
**Displayname** | **string** |  | 
**DisplayName** | **string** |  | 
**DisplaynameScope** | **NullableString** |  | 
**Email** | **NullableString** |  | 
**EmailScope** | **NullableString** |  | 
**Enabled** | **NullableBool** |  | 
**Fediverse** | **NullableString** |  | 
**FediverseScope** | **NullableString** |  | 
**Groups** | **[]string** |  | 
**Headline** | **string** |  | 
**HeadlineScope** | **NullableString** |  | 
**Id** | **string** |  | 
**Language** | **string** |  | 
**LastLogin** | **int64** |  | 
**Locale** | **string** |  | 
**NotifyEmail** | **NullableString** |  | 
**Organisation** | **string** |  | 
**OrganisationScope** | **NullableString** |  | 
**Phone** | **string** |  | 
**PhoneScope** | **NullableString** |  | 
**ProfileEnabled** | **string** |  | 
**ProfileEnabledScope** | **NullableString** |  | 
**Quota** | [**ProvisioningApiUserDetailsQuota**](ProvisioningApiUserDetailsQuota.md) |  | 
**Role** | **string** |  | 
**RoleScope** | **NullableString** |  | 
**StorageLocation** | **NullableString** |  | 
**Subadmin** | **[]string** |  | 
**Twitter** | **string** |  | 
**TwitterScope** | **NullableString** |  | 
**Website** | **string** |  | 
**WebsiteScope** | **NullableString** |  | 

## Methods

### NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue

`func NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue(additionalMail []string, additionalMailScope []string, address string, addressScope NullableString, avatarScope NullableString, backend string, backendCapabilities ProvisioningApiUserDetailsBackendCapabilities, biography string, biographyScope NullableString, displayname string, displayName string, displaynameScope NullableString, email NullableString, emailScope NullableString, enabled NullableBool, fediverse NullableString, fediverseScope NullableString, groups []string, headline string, headlineScope NullableString, id string, language string, lastLogin int64, locale string, notifyEmail NullableString, organisation string, organisationScope NullableString, phone string, phoneScope NullableString, profileEnabled string, profileEnabledScope NullableString, quota ProvisioningApiUserDetailsQuota, role string, roleScope NullableString, storageLocation NullableString, subadmin []string, twitter string, twitterScope NullableString, website string, websiteScope NullableString, ) *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue`

NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue instantiates a new ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueWithDefaults

`func NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueWithDefaults() *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue`

NewProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueWithDefaults instantiates a new ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAdditionalMail() []string`

GetAdditionalMail returns the AdditionalMail field if non-nil, zero value otherwise.

### GetAdditionalMailOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAdditionalMailOk() (*[]string, bool)`

GetAdditionalMailOk returns a tuple with the AdditionalMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAdditionalMail(v []string)`

SetAdditionalMail sets AdditionalMail field to given value.


### GetAdditionalMailScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAdditionalMailScope() []string`

GetAdditionalMailScope returns the AdditionalMailScope field if non-nil, zero value otherwise.

### GetAdditionalMailScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAdditionalMailScopeOk() (*[]string, bool)`

GetAdditionalMailScopeOk returns a tuple with the AdditionalMailScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMailScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAdditionalMailScope(v []string)`

SetAdditionalMailScope sets AdditionalMailScope field to given value.


### SetAdditionalMailScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAdditionalMailScopeNil(b bool)`

 SetAdditionalMailScopeNil sets the value for AdditionalMailScope to be an explicit nil

### UnsetAdditionalMailScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetAdditionalMailScope()`

UnsetAdditionalMailScope ensures that no value is present for AdditionalMailScope, not even an explicit nil
### GetAddress

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAddressScope() string`

GetAddressScope returns the AddressScope field if non-nil, zero value otherwise.

### GetAddressScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAddressScopeOk() (*string, bool)`

GetAddressScopeOk returns a tuple with the AddressScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAddressScope(v string)`

SetAddressScope sets AddressScope field to given value.


### SetAddressScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAddressScopeNil(b bool)`

 SetAddressScopeNil sets the value for AddressScope to be an explicit nil

### UnsetAddressScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetAddressScope()`

UnsetAddressScope ensures that no value is present for AddressScope, not even an explicit nil
### GetAvatarScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAvatarScope() string`

GetAvatarScope returns the AvatarScope field if non-nil, zero value otherwise.

### GetAvatarScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetAvatarScopeOk() (*string, bool)`

GetAvatarScopeOk returns a tuple with the AvatarScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAvatarScope(v string)`

SetAvatarScope sets AvatarScope field to given value.


### SetAvatarScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetAvatarScopeNil(b bool)`

 SetAvatarScopeNil sets the value for AvatarScope to be an explicit nil

### UnsetAvatarScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetAvatarScope()`

UnsetAvatarScope ensures that no value is present for AvatarScope, not even an explicit nil
### GetBackend

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetBackendCapabilities

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBackendCapabilities() ProvisioningApiUserDetailsBackendCapabilities`

GetBackendCapabilities returns the BackendCapabilities field if non-nil, zero value otherwise.

### GetBackendCapabilitiesOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBackendCapabilitiesOk() (*ProvisioningApiUserDetailsBackendCapabilities, bool)`

GetBackendCapabilitiesOk returns a tuple with the BackendCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendCapabilities

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetBackendCapabilities(v ProvisioningApiUserDetailsBackendCapabilities)`

SetBackendCapabilities sets BackendCapabilities field to given value.


### GetBiography

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetBiography(v string)`

SetBiography sets Biography field to given value.


### GetBiographyScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBiographyScope() string`

GetBiographyScope returns the BiographyScope field if non-nil, zero value otherwise.

### GetBiographyScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetBiographyScopeOk() (*string, bool)`

GetBiographyScopeOk returns a tuple with the BiographyScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiographyScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetBiographyScope(v string)`

SetBiographyScope sets BiographyScope field to given value.


### SetBiographyScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetBiographyScopeNil(b bool)`

 SetBiographyScopeNil sets the value for BiographyScope to be an explicit nil

### UnsetBiographyScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetBiographyScope()`

UnsetBiographyScope ensures that no value is present for BiographyScope, not even an explicit nil
### GetDisplayname

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetDisplayName

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplaynameScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplaynameScope() string`

GetDisplaynameScope returns the DisplaynameScope field if non-nil, zero value otherwise.

### GetDisplaynameScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetDisplaynameScopeOk() (*string, bool)`

GetDisplaynameScopeOk returns a tuple with the DisplaynameScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetDisplaynameScope(v string)`

SetDisplaynameScope sets DisplaynameScope field to given value.


### SetDisplaynameScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetDisplaynameScopeNil(b bool)`

 SetDisplaynameScopeNil sets the value for DisplaynameScope to be an explicit nil

### UnsetDisplaynameScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetDisplaynameScope()`

UnsetDisplaynameScope ensures that no value is present for DisplaynameScope, not even an explicit nil
### GetEmail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEmailScope() string`

GetEmailScope returns the EmailScope field if non-nil, zero value otherwise.

### GetEmailScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEmailScopeOk() (*string, bool)`

GetEmailScopeOk returns a tuple with the EmailScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEmailScope(v string)`

SetEmailScope sets EmailScope field to given value.


### SetEmailScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEmailScopeNil(b bool)`

 SetEmailScopeNil sets the value for EmailScope to be an explicit nil

### UnsetEmailScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetEmailScope()`

UnsetEmailScope ensures that no value is present for EmailScope, not even an explicit nil
### GetEnabled

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFediverse

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetFediverse() string`

GetFediverse returns the Fediverse field if non-nil, zero value otherwise.

### GetFediverseOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetFediverseOk() (*string, bool)`

GetFediverseOk returns a tuple with the Fediverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFediverse

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetFediverse(v string)`

SetFediverse sets Fediverse field to given value.


### SetFediverseNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetFediverseNil(b bool)`

 SetFediverseNil sets the value for Fediverse to be an explicit nil

### UnsetFediverse
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetFediverse()`

UnsetFediverse ensures that no value is present for Fediverse, not even an explicit nil
### GetFediverseScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetFediverseScope() string`

GetFediverseScope returns the FediverseScope field if non-nil, zero value otherwise.

### GetFediverseScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetFediverseScopeOk() (*string, bool)`

GetFediverseScopeOk returns a tuple with the FediverseScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFediverseScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetFediverseScope(v string)`

SetFediverseScope sets FediverseScope field to given value.


### SetFediverseScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetFediverseScopeNil(b bool)`

 SetFediverseScopeNil sets the value for FediverseScope to be an explicit nil

### UnsetFediverseScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetFediverseScope()`

UnsetFediverseScope ensures that no value is present for FediverseScope, not even an explicit nil
### GetGroups

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetHeadline

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetHeadline(v string)`

SetHeadline sets Headline field to given value.


### GetHeadlineScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetHeadlineScope() string`

GetHeadlineScope returns the HeadlineScope field if non-nil, zero value otherwise.

### GetHeadlineScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetHeadlineScopeOk() (*string, bool)`

GetHeadlineScopeOk returns a tuple with the HeadlineScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetHeadlineScope(v string)`

SetHeadlineScope sets HeadlineScope field to given value.


### SetHeadlineScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetHeadlineScopeNil(b bool)`

 SetHeadlineScopeNil sets the value for HeadlineScope to be an explicit nil

### UnsetHeadlineScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetHeadlineScope()`

UnsetHeadlineScope ensures that no value is present for HeadlineScope, not even an explicit nil
### GetId

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastLogin

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLastLogin() int64`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLastLoginOk() (*int64, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetLastLogin(v int64)`

SetLastLogin sets LastLogin field to given value.


### GetLocale

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetNotifyEmail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetNotifyEmail() string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetNotifyEmailOk() (*string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetNotifyEmail(v string)`

SetNotifyEmail sets NotifyEmail field to given value.


### SetNotifyEmailNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetNotifyEmailNil(b bool)`

 SetNotifyEmailNil sets the value for NotifyEmail to be an explicit nil

### UnsetNotifyEmail
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetNotifyEmail()`

UnsetNotifyEmail ensures that no value is present for NotifyEmail, not even an explicit nil
### GetOrganisation

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetOrganisationScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetOrganisationScope() string`

GetOrganisationScope returns the OrganisationScope field if non-nil, zero value otherwise.

### GetOrganisationScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetOrganisationScopeOk() (*string, bool)`

GetOrganisationScopeOk returns a tuple with the OrganisationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetOrganisationScope(v string)`

SetOrganisationScope sets OrganisationScope field to given value.


### SetOrganisationScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetOrganisationScopeNil(b bool)`

 SetOrganisationScopeNil sets the value for OrganisationScope to be an explicit nil

### UnsetOrganisationScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetOrganisationScope()`

UnsetOrganisationScope ensures that no value is present for OrganisationScope, not even an explicit nil
### GetPhone

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPhoneScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetPhoneScope() string`

GetPhoneScope returns the PhoneScope field if non-nil, zero value otherwise.

### GetPhoneScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetPhoneScopeOk() (*string, bool)`

GetPhoneScopeOk returns a tuple with the PhoneScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetPhoneScope(v string)`

SetPhoneScope sets PhoneScope field to given value.


### SetPhoneScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetPhoneScopeNil(b bool)`

 SetPhoneScopeNil sets the value for PhoneScope to be an explicit nil

### UnsetPhoneScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetPhoneScope()`

UnsetPhoneScope ensures that no value is present for PhoneScope, not even an explicit nil
### GetProfileEnabled

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetProfileEnabled() string`

GetProfileEnabled returns the ProfileEnabled field if non-nil, zero value otherwise.

### GetProfileEnabledOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetProfileEnabledOk() (*string, bool)`

GetProfileEnabledOk returns a tuple with the ProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnabled

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetProfileEnabled(v string)`

SetProfileEnabled sets ProfileEnabled field to given value.


### GetProfileEnabledScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetProfileEnabledScope() string`

GetProfileEnabledScope returns the ProfileEnabledScope field if non-nil, zero value otherwise.

### GetProfileEnabledScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetProfileEnabledScopeOk() (*string, bool)`

GetProfileEnabledScopeOk returns a tuple with the ProfileEnabledScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnabledScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetProfileEnabledScope(v string)`

SetProfileEnabledScope sets ProfileEnabledScope field to given value.


### SetProfileEnabledScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetProfileEnabledScopeNil(b bool)`

 SetProfileEnabledScopeNil sets the value for ProfileEnabledScope to be an explicit nil

### UnsetProfileEnabledScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetProfileEnabledScope()`

UnsetProfileEnabledScope ensures that no value is present for ProfileEnabledScope, not even an explicit nil
### GetQuota

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetQuota() ProvisioningApiUserDetailsQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetQuotaOk() (*ProvisioningApiUserDetailsQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetQuota(v ProvisioningApiUserDetailsQuota)`

SetQuota sets Quota field to given value.


### GetRole

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetRole(v string)`

SetRole sets Role field to given value.


### GetRoleScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetRoleScope() string`

GetRoleScope returns the RoleScope field if non-nil, zero value otherwise.

### GetRoleScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetRoleScopeOk() (*string, bool)`

GetRoleScopeOk returns a tuple with the RoleScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetRoleScope(v string)`

SetRoleScope sets RoleScope field to given value.


### SetRoleScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetRoleScopeNil(b bool)`

 SetRoleScopeNil sets the value for RoleScope to be an explicit nil

### UnsetRoleScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetRoleScope()`

UnsetRoleScope ensures that no value is present for RoleScope, not even an explicit nil
### GetStorageLocation

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### SetStorageLocationNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetSubadmin

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetSubadmin() []string`

GetSubadmin returns the Subadmin field if non-nil, zero value otherwise.

### GetSubadminOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetSubadminOk() (*[]string, bool)`

GetSubadminOk returns a tuple with the Subadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubadmin

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetSubadmin(v []string)`

SetSubadmin sets Subadmin field to given value.


### GetTwitter

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.


### GetTwitterScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetTwitterScope() string`

GetTwitterScope returns the TwitterScope field if non-nil, zero value otherwise.

### GetTwitterScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetTwitterScopeOk() (*string, bool)`

GetTwitterScopeOk returns a tuple with the TwitterScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetTwitterScope(v string)`

SetTwitterScope sets TwitterScope field to given value.


### SetTwitterScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetTwitterScopeNil(b bool)`

 SetTwitterScopeNil sets the value for TwitterScope to be an explicit nil

### UnsetTwitterScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetTwitterScope()`

UnsetTwitterScope ensures that no value is present for TwitterScope, not even an explicit nil
### GetWebsite

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetWebsiteScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetWebsiteScope() string`

GetWebsiteScope returns the WebsiteScope field if non-nil, zero value otherwise.

### GetWebsiteScopeOk

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) GetWebsiteScopeOk() (*string, bool)`

GetWebsiteScopeOk returns a tuple with the WebsiteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteScope

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetWebsiteScope(v string)`

SetWebsiteScope sets WebsiteScope field to given value.


### SetWebsiteScopeNil

`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) SetWebsiteScopeNil(b bool)`

 SetWebsiteScopeNil sets the value for WebsiteScope to be an explicit nil

### UnsetWebsiteScope
`func (o *ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue) UnsetWebsiteScope()`

UnsetWebsiteScope ensures that no value is present for WebsiteScope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


