# ProjectGet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDomain** | Pointer to **bool** | Indicates whether the project also acts as a domain. If set to true, this project acts as both a project and domain. As a domain, the project provides a name space in which you can create users, groups, and other projects. If set to false, this project behaves as a regular project that contains only resources. | [optional] 
**Enabled** | Pointer to **bool** | If set to true, project is enabled. If set to false, project is disabled. | [optional] 
**Description** | Pointer to **string** | The description of the project. | [optional] 
**DomainId** | Pointer to **string** | The ID of the domain for the project. | [optional] 
**Id** | Pointer to **string** | The ID of the project. | [optional] 
**Name** | Pointer to **string** | The name of the project. | [optional] 
**ParentId** | Pointer to **string** | The ID of the parent for the project. | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 
**Parents** | Pointer to [**[]OneOfstringprojectResponse**](OneOfstringprojectResponse.md) |  | [optional] 
**Subtree** | Pointer to [**[]OneOfstringprojectResponse**](OneOfstringprojectResponse.md) |  | [optional] 

## Methods

### NewProjectGet

`func NewProjectGet() *ProjectGet`

NewProjectGet instantiates a new ProjectGet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectGetWithDefaults

`func NewProjectGetWithDefaults() *ProjectGet`

NewProjectGetWithDefaults instantiates a new ProjectGet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDomain

`func (o *ProjectGet) GetIsDomain() bool`

GetIsDomain returns the IsDomain field if non-nil, zero value otherwise.

### GetIsDomainOk

`func (o *ProjectGet) GetIsDomainOk() (*bool, bool)`

GetIsDomainOk returns a tuple with the IsDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDomain

`func (o *ProjectGet) SetIsDomain(v bool)`

SetIsDomain sets IsDomain field to given value.

### HasIsDomain

`func (o *ProjectGet) HasIsDomain() bool`

HasIsDomain returns a boolean if a field has been set.

### GetEnabled

`func (o *ProjectGet) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ProjectGet) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ProjectGet) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ProjectGet) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ProjectGet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProjectGet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProjectGet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProjectGet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *ProjectGet) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *ProjectGet) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *ProjectGet) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *ProjectGet) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetId

`func (o *ProjectGet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectGet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectGet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectGet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ProjectGet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProjectGet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProjectGet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ProjectGet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *ProjectGet) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProjectGet) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProjectGet) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProjectGet) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetLinks

`func (o *ProjectGet) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectGet) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectGet) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectGet) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetParents

`func (o *ProjectGet) GetParents() []OneOfstringprojectResponse`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *ProjectGet) GetParentsOk() (*[]OneOfstringprojectResponse, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *ProjectGet) SetParents(v []OneOfstringprojectResponse)`

SetParents sets Parents field to given value.

### HasParents

`func (o *ProjectGet) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetSubtree

`func (o *ProjectGet) GetSubtree() []OneOfstringprojectResponse`

GetSubtree returns the Subtree field if non-nil, zero value otherwise.

### GetSubtreeOk

`func (o *ProjectGet) GetSubtreeOk() (*[]OneOfstringprojectResponse, bool)`

GetSubtreeOk returns a tuple with the Subtree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtree

`func (o *ProjectGet) SetSubtree(v []OneOfstringprojectResponse)`

SetSubtree sets Subtree field to given value.

### HasSubtree

`func (o *ProjectGet) HasSubtree() bool`

HasSubtree returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


