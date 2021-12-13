# GroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The new description for the group. | [optional] 
**DomainId** | Pointer to **string** | The ID of the new domain for the group. The ability to change the domain of a group is now deprecated, and will be removed in subsequent release. It is already disabled by default in most Identity service implementations. | [optional] 
**Name** | Pointer to **string** | The new name for the group. | [optional] 

## Methods

### NewGroupUpdate

`func NewGroupUpdate() *GroupUpdate`

NewGroupUpdate instantiates a new GroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateWithDefaults

`func NewGroupUpdateWithDefaults() *GroupUpdate`

NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *GroupUpdate) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GroupUpdate) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GroupUpdate) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *GroupUpdate) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetName

`func (o *GroupUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GroupUpdate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


