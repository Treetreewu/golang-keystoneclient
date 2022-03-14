# RoleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The role name. | 
**Type** | Pointer to **string** | The role type. | [optional] 
**DomainId** | Pointer to **string** | The ID of the domain of the role. | [optional] 
**Description** | Pointer to **string** | Add description about the role. | [optional] 
**Display** | Pointer to **bool** | Show this role in ECP or not. | [optional] 

## Methods

### NewRoleCreate

`func NewRoleCreate(name string, ) *RoleCreate`

NewRoleCreate instantiates a new RoleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleCreateWithDefaults

`func NewRoleCreateWithDefaults() *RoleCreate`

NewRoleCreateWithDefaults instantiates a new RoleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RoleCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleCreate) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *RoleCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RoleCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RoleCreate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RoleCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDomainId

`func (o *RoleCreate) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *RoleCreate) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *RoleCreate) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *RoleCreate) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDescription

`func (o *RoleCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RoleCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RoleCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RoleCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplay

`func (o *RoleCreate) GetDisplay() bool`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *RoleCreate) GetDisplayOk() (*bool, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *RoleCreate) SetDisplay(v bool)`

SetDisplay sets Display field to given value.

### HasDisplay

`func (o *RoleCreate) HasDisplay() bool`

HasDisplay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


