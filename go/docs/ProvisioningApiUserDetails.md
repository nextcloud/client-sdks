# ProvisioningApiUserDetails

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

### NewProvisioningApiUserDetails

`func NewProvisioningApiUserDetails(additionalMail []string, additionalMailScope []string, address string, addressScope NullableString, avatarScope NullableString, backend string, backendCapabilities ProvisioningApiUserDetailsBackendCapabilities, biography string, biographyScope NullableString, displayname string, displayName string, displaynameScope NullableString, email NullableString, emailScope NullableString, enabled NullableBool, fediverse NullableString, fediverseScope NullableString, groups []string, headline string, headlineScope NullableString, id string, language string, lastLogin int64, locale string, notifyEmail NullableString, organisation string, organisationScope NullableString, phone string, phoneScope NullableString, profileEnabled string, profileEnabledScope NullableString, quota ProvisioningApiUserDetailsQuota, role string, roleScope NullableString, storageLocation NullableString, subadmin []string, twitter string, twitterScope NullableString, website string, websiteScope NullableString, ) *ProvisioningApiUserDetails`

NewProvisioningApiUserDetails instantiates a new ProvisioningApiUserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningApiUserDetailsWithDefaults

`func NewProvisioningApiUserDetailsWithDefaults() *ProvisioningApiUserDetails`

NewProvisioningApiUserDetailsWithDefaults instantiates a new ProvisioningApiUserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalMail

`func (o *ProvisioningApiUserDetails) GetAdditionalMail() []string`

GetAdditionalMail returns the AdditionalMail field if non-nil, zero value otherwise.

### GetAdditionalMailOk

`func (o *ProvisioningApiUserDetails) GetAdditionalMailOk() (*[]string, bool)`

GetAdditionalMailOk returns a tuple with the AdditionalMail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMail

`func (o *ProvisioningApiUserDetails) SetAdditionalMail(v []string)`

SetAdditionalMail sets AdditionalMail field to given value.


### GetAdditionalMailScope

`func (o *ProvisioningApiUserDetails) GetAdditionalMailScope() []string`

GetAdditionalMailScope returns the AdditionalMailScope field if non-nil, zero value otherwise.

### GetAdditionalMailScopeOk

`func (o *ProvisioningApiUserDetails) GetAdditionalMailScopeOk() (*[]string, bool)`

GetAdditionalMailScopeOk returns a tuple with the AdditionalMailScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalMailScope

`func (o *ProvisioningApiUserDetails) SetAdditionalMailScope(v []string)`

SetAdditionalMailScope sets AdditionalMailScope field to given value.


### SetAdditionalMailScopeNil

`func (o *ProvisioningApiUserDetails) SetAdditionalMailScopeNil(b bool)`

 SetAdditionalMailScopeNil sets the value for AdditionalMailScope to be an explicit nil

### UnsetAdditionalMailScope
`func (o *ProvisioningApiUserDetails) UnsetAdditionalMailScope()`

UnsetAdditionalMailScope ensures that no value is present for AdditionalMailScope, not even an explicit nil
### GetAddress

`func (o *ProvisioningApiUserDetails) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ProvisioningApiUserDetails) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ProvisioningApiUserDetails) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAddressScope

`func (o *ProvisioningApiUserDetails) GetAddressScope() string`

GetAddressScope returns the AddressScope field if non-nil, zero value otherwise.

### GetAddressScopeOk

`func (o *ProvisioningApiUserDetails) GetAddressScopeOk() (*string, bool)`

GetAddressScopeOk returns a tuple with the AddressScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressScope

`func (o *ProvisioningApiUserDetails) SetAddressScope(v string)`

SetAddressScope sets AddressScope field to given value.


### SetAddressScopeNil

`func (o *ProvisioningApiUserDetails) SetAddressScopeNil(b bool)`

 SetAddressScopeNil sets the value for AddressScope to be an explicit nil

### UnsetAddressScope
`func (o *ProvisioningApiUserDetails) UnsetAddressScope()`

UnsetAddressScope ensures that no value is present for AddressScope, not even an explicit nil
### GetAvatarScope

`func (o *ProvisioningApiUserDetails) GetAvatarScope() string`

GetAvatarScope returns the AvatarScope field if non-nil, zero value otherwise.

### GetAvatarScopeOk

`func (o *ProvisioningApiUserDetails) GetAvatarScopeOk() (*string, bool)`

GetAvatarScopeOk returns a tuple with the AvatarScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarScope

`func (o *ProvisioningApiUserDetails) SetAvatarScope(v string)`

SetAvatarScope sets AvatarScope field to given value.


### SetAvatarScopeNil

`func (o *ProvisioningApiUserDetails) SetAvatarScopeNil(b bool)`

 SetAvatarScopeNil sets the value for AvatarScope to be an explicit nil

### UnsetAvatarScope
`func (o *ProvisioningApiUserDetails) UnsetAvatarScope()`

UnsetAvatarScope ensures that no value is present for AvatarScope, not even an explicit nil
### GetBackend

`func (o *ProvisioningApiUserDetails) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *ProvisioningApiUserDetails) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *ProvisioningApiUserDetails) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetBackendCapabilities

`func (o *ProvisioningApiUserDetails) GetBackendCapabilities() ProvisioningApiUserDetailsBackendCapabilities`

GetBackendCapabilities returns the BackendCapabilities field if non-nil, zero value otherwise.

### GetBackendCapabilitiesOk

`func (o *ProvisioningApiUserDetails) GetBackendCapabilitiesOk() (*ProvisioningApiUserDetailsBackendCapabilities, bool)`

GetBackendCapabilitiesOk returns a tuple with the BackendCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendCapabilities

`func (o *ProvisioningApiUserDetails) SetBackendCapabilities(v ProvisioningApiUserDetailsBackendCapabilities)`

SetBackendCapabilities sets BackendCapabilities field to given value.


### GetBiography

`func (o *ProvisioningApiUserDetails) GetBiography() string`

GetBiography returns the Biography field if non-nil, zero value otherwise.

### GetBiographyOk

`func (o *ProvisioningApiUserDetails) GetBiographyOk() (*string, bool)`

GetBiographyOk returns a tuple with the Biography field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiography

`func (o *ProvisioningApiUserDetails) SetBiography(v string)`

SetBiography sets Biography field to given value.


### GetBiographyScope

`func (o *ProvisioningApiUserDetails) GetBiographyScope() string`

GetBiographyScope returns the BiographyScope field if non-nil, zero value otherwise.

### GetBiographyScopeOk

`func (o *ProvisioningApiUserDetails) GetBiographyScopeOk() (*string, bool)`

GetBiographyScopeOk returns a tuple with the BiographyScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiographyScope

`func (o *ProvisioningApiUserDetails) SetBiographyScope(v string)`

SetBiographyScope sets BiographyScope field to given value.


### SetBiographyScopeNil

`func (o *ProvisioningApiUserDetails) SetBiographyScopeNil(b bool)`

 SetBiographyScopeNil sets the value for BiographyScope to be an explicit nil

### UnsetBiographyScope
`func (o *ProvisioningApiUserDetails) UnsetBiographyScope()`

UnsetBiographyScope ensures that no value is present for BiographyScope, not even an explicit nil
### GetDisplayname

`func (o *ProvisioningApiUserDetails) GetDisplayname() string`

GetDisplayname returns the Displayname field if non-nil, zero value otherwise.

### GetDisplaynameOk

`func (o *ProvisioningApiUserDetails) GetDisplaynameOk() (*string, bool)`

GetDisplaynameOk returns a tuple with the Displayname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayname

`func (o *ProvisioningApiUserDetails) SetDisplayname(v string)`

SetDisplayname sets Displayname field to given value.


### GetDisplayName

`func (o *ProvisioningApiUserDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ProvisioningApiUserDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ProvisioningApiUserDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplaynameScope

`func (o *ProvisioningApiUserDetails) GetDisplaynameScope() string`

GetDisplaynameScope returns the DisplaynameScope field if non-nil, zero value otherwise.

### GetDisplaynameScopeOk

`func (o *ProvisioningApiUserDetails) GetDisplaynameScopeOk() (*string, bool)`

GetDisplaynameScopeOk returns a tuple with the DisplaynameScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplaynameScope

`func (o *ProvisioningApiUserDetails) SetDisplaynameScope(v string)`

SetDisplaynameScope sets DisplaynameScope field to given value.


### SetDisplaynameScopeNil

`func (o *ProvisioningApiUserDetails) SetDisplaynameScopeNil(b bool)`

 SetDisplaynameScopeNil sets the value for DisplaynameScope to be an explicit nil

### UnsetDisplaynameScope
`func (o *ProvisioningApiUserDetails) UnsetDisplaynameScope()`

UnsetDisplaynameScope ensures that no value is present for DisplaynameScope, not even an explicit nil
### GetEmail

`func (o *ProvisioningApiUserDetails) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProvisioningApiUserDetails) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProvisioningApiUserDetails) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *ProvisioningApiUserDetails) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ProvisioningApiUserDetails) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetEmailScope

`func (o *ProvisioningApiUserDetails) GetEmailScope() string`

GetEmailScope returns the EmailScope field if non-nil, zero value otherwise.

### GetEmailScopeOk

`func (o *ProvisioningApiUserDetails) GetEmailScopeOk() (*string, bool)`

GetEmailScopeOk returns a tuple with the EmailScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailScope

`func (o *ProvisioningApiUserDetails) SetEmailScope(v string)`

SetEmailScope sets EmailScope field to given value.


### SetEmailScopeNil

`func (o *ProvisioningApiUserDetails) SetEmailScopeNil(b bool)`

 SetEmailScopeNil sets the value for EmailScope to be an explicit nil

### UnsetEmailScope
`func (o *ProvisioningApiUserDetails) UnsetEmailScope()`

UnsetEmailScope ensures that no value is present for EmailScope, not even an explicit nil
### GetEnabled

`func (o *ProvisioningApiUserDetails) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProvisioningApiUserDetails) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProvisioningApiUserDetails) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### SetEnabledNil

`func (o *ProvisioningApiUserDetails) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *ProvisioningApiUserDetails) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetFediverse

`func (o *ProvisioningApiUserDetails) GetFediverse() string`

GetFediverse returns the Fediverse field if non-nil, zero value otherwise.

### GetFediverseOk

`func (o *ProvisioningApiUserDetails) GetFediverseOk() (*string, bool)`

GetFediverseOk returns a tuple with the Fediverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFediverse

`func (o *ProvisioningApiUserDetails) SetFediverse(v string)`

SetFediverse sets Fediverse field to given value.


### SetFediverseNil

`func (o *ProvisioningApiUserDetails) SetFediverseNil(b bool)`

 SetFediverseNil sets the value for Fediverse to be an explicit nil

### UnsetFediverse
`func (o *ProvisioningApiUserDetails) UnsetFediverse()`

UnsetFediverse ensures that no value is present for Fediverse, not even an explicit nil
### GetFediverseScope

`func (o *ProvisioningApiUserDetails) GetFediverseScope() string`

GetFediverseScope returns the FediverseScope field if non-nil, zero value otherwise.

### GetFediverseScopeOk

`func (o *ProvisioningApiUserDetails) GetFediverseScopeOk() (*string, bool)`

GetFediverseScopeOk returns a tuple with the FediverseScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFediverseScope

`func (o *ProvisioningApiUserDetails) SetFediverseScope(v string)`

SetFediverseScope sets FediverseScope field to given value.


### SetFediverseScopeNil

`func (o *ProvisioningApiUserDetails) SetFediverseScopeNil(b bool)`

 SetFediverseScopeNil sets the value for FediverseScope to be an explicit nil

### UnsetFediverseScope
`func (o *ProvisioningApiUserDetails) UnsetFediverseScope()`

UnsetFediverseScope ensures that no value is present for FediverseScope, not even an explicit nil
### GetGroups

`func (o *ProvisioningApiUserDetails) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *ProvisioningApiUserDetails) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *ProvisioningApiUserDetails) SetGroups(v []string)`

SetGroups sets Groups field to given value.


### GetHeadline

`func (o *ProvisioningApiUserDetails) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *ProvisioningApiUserDetails) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *ProvisioningApiUserDetails) SetHeadline(v string)`

SetHeadline sets Headline field to given value.


### GetHeadlineScope

`func (o *ProvisioningApiUserDetails) GetHeadlineScope() string`

GetHeadlineScope returns the HeadlineScope field if non-nil, zero value otherwise.

### GetHeadlineScopeOk

`func (o *ProvisioningApiUserDetails) GetHeadlineScopeOk() (*string, bool)`

GetHeadlineScopeOk returns a tuple with the HeadlineScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadlineScope

`func (o *ProvisioningApiUserDetails) SetHeadlineScope(v string)`

SetHeadlineScope sets HeadlineScope field to given value.


### SetHeadlineScopeNil

`func (o *ProvisioningApiUserDetails) SetHeadlineScopeNil(b bool)`

 SetHeadlineScopeNil sets the value for HeadlineScope to be an explicit nil

### UnsetHeadlineScope
`func (o *ProvisioningApiUserDetails) UnsetHeadlineScope()`

UnsetHeadlineScope ensures that no value is present for HeadlineScope, not even an explicit nil
### GetId

`func (o *ProvisioningApiUserDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisioningApiUserDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisioningApiUserDetails) SetId(v string)`

SetId sets Id field to given value.


### GetLanguage

`func (o *ProvisioningApiUserDetails) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ProvisioningApiUserDetails) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ProvisioningApiUserDetails) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLastLogin

`func (o *ProvisioningApiUserDetails) GetLastLogin() int64`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *ProvisioningApiUserDetails) GetLastLoginOk() (*int64, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *ProvisioningApiUserDetails) SetLastLogin(v int64)`

SetLastLogin sets LastLogin field to given value.


### GetLocale

`func (o *ProvisioningApiUserDetails) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *ProvisioningApiUserDetails) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *ProvisioningApiUserDetails) SetLocale(v string)`

SetLocale sets Locale field to given value.


### GetNotifyEmail

`func (o *ProvisioningApiUserDetails) GetNotifyEmail() string`

GetNotifyEmail returns the NotifyEmail field if non-nil, zero value otherwise.

### GetNotifyEmailOk

`func (o *ProvisioningApiUserDetails) GetNotifyEmailOk() (*string, bool)`

GetNotifyEmailOk returns a tuple with the NotifyEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyEmail

`func (o *ProvisioningApiUserDetails) SetNotifyEmail(v string)`

SetNotifyEmail sets NotifyEmail field to given value.


### SetNotifyEmailNil

`func (o *ProvisioningApiUserDetails) SetNotifyEmailNil(b bool)`

 SetNotifyEmailNil sets the value for NotifyEmail to be an explicit nil

### UnsetNotifyEmail
`func (o *ProvisioningApiUserDetails) UnsetNotifyEmail()`

UnsetNotifyEmail ensures that no value is present for NotifyEmail, not even an explicit nil
### GetOrganisation

`func (o *ProvisioningApiUserDetails) GetOrganisation() string`

GetOrganisation returns the Organisation field if non-nil, zero value otherwise.

### GetOrganisationOk

`func (o *ProvisioningApiUserDetails) GetOrganisationOk() (*string, bool)`

GetOrganisationOk returns a tuple with the Organisation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisation

`func (o *ProvisioningApiUserDetails) SetOrganisation(v string)`

SetOrganisation sets Organisation field to given value.


### GetOrganisationScope

`func (o *ProvisioningApiUserDetails) GetOrganisationScope() string`

GetOrganisationScope returns the OrganisationScope field if non-nil, zero value otherwise.

### GetOrganisationScopeOk

`func (o *ProvisioningApiUserDetails) GetOrganisationScopeOk() (*string, bool)`

GetOrganisationScopeOk returns a tuple with the OrganisationScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganisationScope

`func (o *ProvisioningApiUserDetails) SetOrganisationScope(v string)`

SetOrganisationScope sets OrganisationScope field to given value.


### SetOrganisationScopeNil

`func (o *ProvisioningApiUserDetails) SetOrganisationScopeNil(b bool)`

 SetOrganisationScopeNil sets the value for OrganisationScope to be an explicit nil

### UnsetOrganisationScope
`func (o *ProvisioningApiUserDetails) UnsetOrganisationScope()`

UnsetOrganisationScope ensures that no value is present for OrganisationScope, not even an explicit nil
### GetPhone

`func (o *ProvisioningApiUserDetails) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ProvisioningApiUserDetails) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ProvisioningApiUserDetails) SetPhone(v string)`

SetPhone sets Phone field to given value.


### GetPhoneScope

`func (o *ProvisioningApiUserDetails) GetPhoneScope() string`

GetPhoneScope returns the PhoneScope field if non-nil, zero value otherwise.

### GetPhoneScopeOk

`func (o *ProvisioningApiUserDetails) GetPhoneScopeOk() (*string, bool)`

GetPhoneScopeOk returns a tuple with the PhoneScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneScope

`func (o *ProvisioningApiUserDetails) SetPhoneScope(v string)`

SetPhoneScope sets PhoneScope field to given value.


### SetPhoneScopeNil

`func (o *ProvisioningApiUserDetails) SetPhoneScopeNil(b bool)`

 SetPhoneScopeNil sets the value for PhoneScope to be an explicit nil

### UnsetPhoneScope
`func (o *ProvisioningApiUserDetails) UnsetPhoneScope()`

UnsetPhoneScope ensures that no value is present for PhoneScope, not even an explicit nil
### GetProfileEnabled

`func (o *ProvisioningApiUserDetails) GetProfileEnabled() string`

GetProfileEnabled returns the ProfileEnabled field if non-nil, zero value otherwise.

### GetProfileEnabledOk

`func (o *ProvisioningApiUserDetails) GetProfileEnabledOk() (*string, bool)`

GetProfileEnabledOk returns a tuple with the ProfileEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnabled

`func (o *ProvisioningApiUserDetails) SetProfileEnabled(v string)`

SetProfileEnabled sets ProfileEnabled field to given value.


### GetProfileEnabledScope

`func (o *ProvisioningApiUserDetails) GetProfileEnabledScope() string`

GetProfileEnabledScope returns the ProfileEnabledScope field if non-nil, zero value otherwise.

### GetProfileEnabledScopeOk

`func (o *ProvisioningApiUserDetails) GetProfileEnabledScopeOk() (*string, bool)`

GetProfileEnabledScopeOk returns a tuple with the ProfileEnabledScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileEnabledScope

`func (o *ProvisioningApiUserDetails) SetProfileEnabledScope(v string)`

SetProfileEnabledScope sets ProfileEnabledScope field to given value.


### SetProfileEnabledScopeNil

`func (o *ProvisioningApiUserDetails) SetProfileEnabledScopeNil(b bool)`

 SetProfileEnabledScopeNil sets the value for ProfileEnabledScope to be an explicit nil

### UnsetProfileEnabledScope
`func (o *ProvisioningApiUserDetails) UnsetProfileEnabledScope()`

UnsetProfileEnabledScope ensures that no value is present for ProfileEnabledScope, not even an explicit nil
### GetQuota

`func (o *ProvisioningApiUserDetails) GetQuota() ProvisioningApiUserDetailsQuota`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *ProvisioningApiUserDetails) GetQuotaOk() (*ProvisioningApiUserDetailsQuota, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *ProvisioningApiUserDetails) SetQuota(v ProvisioningApiUserDetailsQuota)`

SetQuota sets Quota field to given value.


### GetRole

`func (o *ProvisioningApiUserDetails) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ProvisioningApiUserDetails) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ProvisioningApiUserDetails) SetRole(v string)`

SetRole sets Role field to given value.


### GetRoleScope

`func (o *ProvisioningApiUserDetails) GetRoleScope() string`

GetRoleScope returns the RoleScope field if non-nil, zero value otherwise.

### GetRoleScopeOk

`func (o *ProvisioningApiUserDetails) GetRoleScopeOk() (*string, bool)`

GetRoleScopeOk returns a tuple with the RoleScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleScope

`func (o *ProvisioningApiUserDetails) SetRoleScope(v string)`

SetRoleScope sets RoleScope field to given value.


### SetRoleScopeNil

`func (o *ProvisioningApiUserDetails) SetRoleScopeNil(b bool)`

 SetRoleScopeNil sets the value for RoleScope to be an explicit nil

### UnsetRoleScope
`func (o *ProvisioningApiUserDetails) UnsetRoleScope()`

UnsetRoleScope ensures that no value is present for RoleScope, not even an explicit nil
### GetStorageLocation

`func (o *ProvisioningApiUserDetails) GetStorageLocation() string`

GetStorageLocation returns the StorageLocation field if non-nil, zero value otherwise.

### GetStorageLocationOk

`func (o *ProvisioningApiUserDetails) GetStorageLocationOk() (*string, bool)`

GetStorageLocationOk returns a tuple with the StorageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageLocation

`func (o *ProvisioningApiUserDetails) SetStorageLocation(v string)`

SetStorageLocation sets StorageLocation field to given value.


### SetStorageLocationNil

`func (o *ProvisioningApiUserDetails) SetStorageLocationNil(b bool)`

 SetStorageLocationNil sets the value for StorageLocation to be an explicit nil

### UnsetStorageLocation
`func (o *ProvisioningApiUserDetails) UnsetStorageLocation()`

UnsetStorageLocation ensures that no value is present for StorageLocation, not even an explicit nil
### GetSubadmin

`func (o *ProvisioningApiUserDetails) GetSubadmin() []string`

GetSubadmin returns the Subadmin field if non-nil, zero value otherwise.

### GetSubadminOk

`func (o *ProvisioningApiUserDetails) GetSubadminOk() (*[]string, bool)`

GetSubadminOk returns a tuple with the Subadmin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubadmin

`func (o *ProvisioningApiUserDetails) SetSubadmin(v []string)`

SetSubadmin sets Subadmin field to given value.


### GetTwitter

`func (o *ProvisioningApiUserDetails) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *ProvisioningApiUserDetails) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *ProvisioningApiUserDetails) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.


### GetTwitterScope

`func (o *ProvisioningApiUserDetails) GetTwitterScope() string`

GetTwitterScope returns the TwitterScope field if non-nil, zero value otherwise.

### GetTwitterScopeOk

`func (o *ProvisioningApiUserDetails) GetTwitterScopeOk() (*string, bool)`

GetTwitterScopeOk returns a tuple with the TwitterScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitterScope

`func (o *ProvisioningApiUserDetails) SetTwitterScope(v string)`

SetTwitterScope sets TwitterScope field to given value.


### SetTwitterScopeNil

`func (o *ProvisioningApiUserDetails) SetTwitterScopeNil(b bool)`

 SetTwitterScopeNil sets the value for TwitterScope to be an explicit nil

### UnsetTwitterScope
`func (o *ProvisioningApiUserDetails) UnsetTwitterScope()`

UnsetTwitterScope ensures that no value is present for TwitterScope, not even an explicit nil
### GetWebsite

`func (o *ProvisioningApiUserDetails) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *ProvisioningApiUserDetails) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *ProvisioningApiUserDetails) SetWebsite(v string)`

SetWebsite sets Website field to given value.


### GetWebsiteScope

`func (o *ProvisioningApiUserDetails) GetWebsiteScope() string`

GetWebsiteScope returns the WebsiteScope field if non-nil, zero value otherwise.

### GetWebsiteScopeOk

`func (o *ProvisioningApiUserDetails) GetWebsiteScopeOk() (*string, bool)`

GetWebsiteScopeOk returns a tuple with the WebsiteScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteScope

`func (o *ProvisioningApiUserDetails) SetWebsiteScope(v string)`

SetWebsiteScope sets WebsiteScope field to given value.


### SetWebsiteScopeNil

`func (o *ProvisioningApiUserDetails) SetWebsiteScopeNil(b bool)`

 SetWebsiteScopeNil sets the value for WebsiteScope to be an explicit nil

### UnsetWebsiteScope
`func (o *ProvisioningApiUserDetails) UnsetWebsiteScope()`

UnsetWebsiteScope ensures that no value is present for WebsiteScope, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


