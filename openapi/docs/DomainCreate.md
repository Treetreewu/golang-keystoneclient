# DomainCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the domain. | [optional] 
**Enabled** | Pointer to **bool** | If set to true, domain is created enabled. If set to false, domain is created disabled. The default is true. Users can only authorize against an enabled domain (and any of its projects). In addition, users can only authenticate if the domain that owns them is also enabled. Disabling a domain prevents both of these things. | [optional] 
**ExplicitDomainId** | Pointer to **string** | The ID of the domain. A domain created this way will not use an auto-generated ID, but will use the ID passed in instead. Identifiers passed in this way must conform to the existing ID generation scheme: UUID4 without dashes. | [optional] 
**Name** | **string** | The name of the domain. | 

## Methods

### NewDomainCreate

`func NewDomainCreate(name string, ) *DomainCreate`

NewDomainCreate instantiates a new DomainCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainCreateWithDefaults

`func NewDomainCreateWithDefaults() *DomainCreate`

NewDomainCreateWithDefaults instantiates a new DomainCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DomainCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DomainCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DomainCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DomainCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DomainCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DomainCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DomainCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DomainCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExplicitDomainId

`func (o *DomainCreate) GetExplicitDomainId() string`

GetExplicitDomainId returns the ExplicitDomainId field if non-nil, zero value otherwise.

### GetExplicitDomainIdOk

`func (o *DomainCreate) GetExplicitDomainIdOk() (*string, bool)`

GetExplicitDomainIdOk returns a tuple with the ExplicitDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitDomainId

`func (o *DomainCreate) SetExplicitDomainId(v string)`

SetExplicitDomainId sets ExplicitDomainId field to given value.

### HasExplicitDomainId

`func (o *DomainCreate) HasExplicitDomainId() bool`

HasExplicitDomainId returns a boolean if a field has been set.

### GetName

`func (o *DomainCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DomainCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DomainCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


