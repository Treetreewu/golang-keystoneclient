# DomainUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The new description of the domain. | [optional] 
**Enabled** | Pointer to **bool** | If set to true, domain is enabled. If set to false, domain is disabled. | [optional] 
**Name** | Pointer to **string** | The new name of the domain. | [optional] 

## Methods

### NewDomainUpdate

`func NewDomainUpdate() *DomainUpdate`

NewDomainUpdate instantiates a new DomainUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainUpdateWithDefaults

`func NewDomainUpdateWithDefaults() *DomainUpdate`

NewDomainUpdateWithDefaults instantiates a new DomainUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DomainUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DomainUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DomainUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DomainUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DomainUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *DomainUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DomainUpdate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


