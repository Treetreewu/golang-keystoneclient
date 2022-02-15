# GroupCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the group. | [optional] 
**DomainId** | Pointer to **string** | The ID of the domain of the group. If the domain ID is not provided in the request, the Identity service will attempt to pull the domain ID from the token used in the request. Note that this requires the use of a domain-scoped token. | [optional] 
**Name** | **string** | The name of the group. | 
**CreatedAt** | Pointer to **time.Time** | The creation time of the group. Stored in &#x60;keystone.group.extra&#x60;. | [optional] 

## Methods

### NewGroupCreate

`func NewGroupCreate(name string, ) *GroupCreate`

NewGroupCreate instantiates a new GroupCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCreateWithDefaults

`func NewGroupCreateWithDefaults() *GroupCreate`

NewGroupCreateWithDefaults instantiates a new GroupCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomainId

`func (o *GroupCreate) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *GroupCreate) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *GroupCreate) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *GroupCreate) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetName

`func (o *GroupCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GroupCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GroupCreate) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *GroupCreate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GroupCreate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GroupCreate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GroupCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


