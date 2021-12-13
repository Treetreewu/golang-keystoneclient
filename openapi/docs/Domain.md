# Domain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the domain. | [optional] 
**Enabled** | Pointer to **bool** | If set to true, domain is enabled. If set to false, domain is disabled. | [optional] 
**Id** | Pointer to **string** | The ID of the domain. | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the domain. | [optional] 
**Parents** | Pointer to **[]string** | A list of parent ID. | [optional] 

## Methods

### NewDomain

`func NewDomain() *Domain`

NewDomain instantiates a new Domain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainWithDefaults

`func NewDomainWithDefaults() *Domain`

NewDomainWithDefaults instantiates a new Domain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Domain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Domain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Domain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Domain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Domain) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Domain) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Domain) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Domain) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *Domain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Domain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *Domain) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Domain) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Domain) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Domain) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *Domain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Domain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Domain) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Domain) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParents

`func (o *Domain) GetParents() []string`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *Domain) GetParentsOk() (*[]string, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *Domain) SetParents(v []string)`

SetParents sets Parents field to given value.

### HasParents

`func (o *Domain) HasParents() bool`

HasParents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


