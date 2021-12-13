# ProjectUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The new name of the project, which must be unique within the owning domain. A project can have the same name as its domain. | [optional] 
**IsDomain** | Pointer to **bool** | Indicates whether the project also acts as a domain. If set to &#x60;true&#x60;, this project acts as both a project and domain. As a domain, the project provides a name space in which you can create users, groups, and other projects. If set to &#x60;false&#x60;, this project behaves as a regular project that contains only resources. Default is &#x60;false&#x60;. You cannot update this parameter after you create the project. | [optional] 
**Description** | Pointer to **string** | The new description of the project. | [optional] 
**Enabled** | Pointer to **bool** | If set to true, project is enabled. If set to false, project is disabled. The default is true. | [optional] 

## Methods

### NewProjectUpdate

`func NewProjectUpdate() *ProjectUpdate`

NewProjectUpdate instantiates a new ProjectUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectUpdateWithDefaults

`func NewProjectUpdateWithDefaults() *ProjectUpdate`

NewProjectUpdateWithDefaults instantiates a new ProjectUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ProjectUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetIsDomain

`func (o *ProjectUpdate) GetIsDomain() bool`

GetIsDomain returns the IsDomain field if non-nil, zero value otherwise.

### GetIsDomainOk

`func (o *ProjectUpdate) GetIsDomainOk() (*bool, bool)`

GetIsDomainOk returns a tuple with the IsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomain

`func (o *ProjectUpdate) SetIsDomain(v bool)`

SetIsDomain sets IsDomain field to given value.

### HasIsDomain

`func (o *ProjectUpdate) HasIsDomain() bool`

HasIsDomain returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ProjectUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProjectUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProjectUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProjectUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


