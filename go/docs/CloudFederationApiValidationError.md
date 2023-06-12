# CloudFederationApiValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ValidationErrors** | [**[]CloudFederationApiValidationErrorAllOfValidationErrors**](CloudFederationApiValidationErrorAllOfValidationErrors.md) |  | 

## Methods

### NewCloudFederationApiValidationError

`func NewCloudFederationApiValidationError(message string, validationErrors []CloudFederationApiValidationErrorAllOfValidationErrors, ) *CloudFederationApiValidationError`

NewCloudFederationApiValidationError instantiates a new CloudFederationApiValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFederationApiValidationErrorWithDefaults

`func NewCloudFederationApiValidationErrorWithDefaults() *CloudFederationApiValidationError`

NewCloudFederationApiValidationErrorWithDefaults instantiates a new CloudFederationApiValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CloudFederationApiValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CloudFederationApiValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CloudFederationApiValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetValidationErrors

`func (o *CloudFederationApiValidationError) GetValidationErrors() []CloudFederationApiValidationErrorAllOfValidationErrors`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *CloudFederationApiValidationError) GetValidationErrorsOk() (*[]CloudFederationApiValidationErrorAllOfValidationErrors, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *CloudFederationApiValidationError) SetValidationErrors(v []CloudFederationApiValidationErrorAllOfValidationErrors)`

SetValidationErrors sets ValidationErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


