# UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The user name. Must be unique within the owning domain. | 
**DefaultProjectId** | Pointer to **string** | The ID of the default project for the user. A user’s default project must not be a domain. Setting this attribute does not grant any actual authorization on the project, and is merely provided for convenience. Therefore, the referenced project does not need to exist within the user domain. (Since v3.1) If the user does not have authorization to their default project, the default project is ignored at token creation. (Since v3.1) Additionally, if your default project is not valid, a token is issued without an explicit scope of authorization. | [optional] 
**DomainId** | Pointer to **string** | The ID of the domain of the user. If the domain ID is not provided in the request, the Identity service will attempt to pull the domain ID from the token used in the request. Note that this requires the use of a domain-scoped token. | [optional] 
**Federated** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** | If the user is enabled, this value is &#x60;true&#x60;. If the user is disabled, this value is &#x60;false&#x60;. | [optional] 
**Password** | Pointer to **string** | The password of the user. | [optional] 
**Email** | Pointer to **string** | The email of the user. | [optional] 
**Description** | Pointer to **string** | Description of the user. | [optional] 
**Options** | Pointer to [**UserOptions**](UserOptions.md) |  | [optional] 

## Methods

### NewUserCreate

`func NewUserCreate(name string, ) *UserCreate`

NewUserCreate instantiates a new UserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateWithDefaults

`func NewUserCreateWithDefaults() *UserCreate`

NewUserCreateWithDefaults instantiates a new UserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UserCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDefaultProjectId

`func (o *UserCreate) GetDefaultProjectId() string`

GetDefaultProjectId returns the DefaultProjectId field if non-nil, zero value otherwise.

### GetDefaultProjectIdOk

`func (o *UserCreate) GetDefaultProjectIdOk() (*string, bool)`

GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectId

`func (o *UserCreate) SetDefaultProjectId(v string)`

SetDefaultProjectId sets DefaultProjectId field to given value.

### HasDefaultProjectId

`func (o *UserCreate) HasDefaultProjectId() bool`

HasDefaultProjectId returns a boolean if a field has been set.

### GetDomainId

`func (o *UserCreate) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *UserCreate) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *UserCreate) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *UserCreate) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetFederated

`func (o *UserCreate) GetFederated() []map[string]interface{}`

GetFederated returns the Federated field if non-nil, zero value otherwise.

### GetFederatedOk

`func (o *UserCreate) GetFederatedOk() (*[]map[string]interface{}, bool)`

GetFederatedOk returns a tuple with the Federated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederated

`func (o *UserCreate) SetFederated(v []map[string]interface{})`

SetFederated sets Federated field to given value.

### HasFederated

`func (o *UserCreate) HasFederated() bool`

HasFederated returns a boolean if a field has been set.

### GetEnabled

`func (o *UserCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPassword

`func (o *UserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEmail

`func (o *UserCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDescription

`func (o *UserCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptions

`func (o *UserCreate) GetOptions() UserOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UserCreate) GetOptionsOk() (*UserOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UserCreate) SetOptions(v UserOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UserCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


