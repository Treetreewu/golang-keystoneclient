# ApplicationCredentialCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the application credential. Must be unique to a user. | 
**Secret** | Pointer to **string** | The secret that the application credential will be created with. If not provided, one will be generated. | [optional] 
**Description** | Pointer to **string** | A description of the application credential’s purpose. | [optional] 
**ExpiresAt** | Pointer to **time.Time** | An optional expiry time for the application credential. If unset, the application credential does not expire. | [optional] 
**Roles** | Pointer to [**[]IdAndName**](IdAndName.md) | An optional list of role objects, identified by ID or name. The list may only contain roles that the user has assigned on the project. If not provided, the roles assigned to the application credential will be the same as the roles in the current token. | [optional] 
**Unrestricted** | Pointer to **bool** | An optional flag to restrict whether the application credential may be used for the creation or destruction of other application credentials or trusts. Defaults to false. | [optional] 
**AccessRules** | Pointer to [**[]AccessRuleRequest**](AccessRuleRequest.md) | A list of &#x60;access_rules&#x60; objects | [optional] 

## Methods

### NewApplicationCredentialCreate

`func NewApplicationCredentialCreate(name string, ) *ApplicationCredentialCreate`

NewApplicationCredentialCreate instantiates a new ApplicationCredentialCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialCreateWithDefaults

`func NewApplicationCredentialCreateWithDefaults() *ApplicationCredentialCreate`

NewApplicationCredentialCreateWithDefaults instantiates a new ApplicationCredentialCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApplicationCredentialCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationCredentialCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationCredentialCreate) SetName(v string)`

SetName sets Name field to given value.


### GetSecret

`func (o *ApplicationCredentialCreate) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApplicationCredentialCreate) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApplicationCredentialCreate) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApplicationCredentialCreate) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationCredentialCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationCredentialCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationCredentialCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationCredentialCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApplicationCredentialCreate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApplicationCredentialCreate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApplicationCredentialCreate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApplicationCredentialCreate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRoles

`func (o *ApplicationCredentialCreate) GetRoles() []IdAndName`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApplicationCredentialCreate) GetRolesOk() (*[]IdAndName, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApplicationCredentialCreate) SetRoles(v []IdAndName)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApplicationCredentialCreate) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetUnrestricted

`func (o *ApplicationCredentialCreate) GetUnrestricted() bool`

GetUnrestricted returns the Unrestricted field if non-nil, zero value otherwise.

### GetUnrestrictedOk

`func (o *ApplicationCredentialCreate) GetUnrestrictedOk() (*bool, bool)`

GetUnrestrictedOk returns a tuple with the Unrestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnrestricted

`func (o *ApplicationCredentialCreate) SetUnrestricted(v bool)`

SetUnrestricted sets Unrestricted field to given value.

### HasUnrestricted

`func (o *ApplicationCredentialCreate) HasUnrestricted() bool`

HasUnrestricted returns a boolean if a field has been set.

### GetAccessRules

`func (o *ApplicationCredentialCreate) GetAccessRules() []AccessRuleRequest`

GetAccessRules returns the AccessRules field if non-nil, zero value otherwise.

### GetAccessRulesOk

`func (o *ApplicationCredentialCreate) GetAccessRulesOk() (*[]AccessRuleRequest, bool)`

GetAccessRulesOk returns a tuple with the AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessRules

`func (o *ApplicationCredentialCreate) SetAccessRules(v []AccessRuleRequest)`

SetAccessRules sets AccessRules field to given value.

### HasAccessRules

`func (o *ApplicationCredentialCreate) HasAccessRules() bool`

HasAccessRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


