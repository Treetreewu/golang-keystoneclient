# RoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**RoleUpdate**](RoleUpdate.md) |  | [optional] 

## Methods

### NewRoleUpdateRequest

`func NewRoleUpdateRequest() *RoleUpdateRequest`

NewRoleUpdateRequest instantiates a new RoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleUpdateRequestWithDefaults

`func NewRoleUpdateRequestWithDefaults() *RoleUpdateRequest`

NewRoleUpdateRequestWithDefaults instantiates a new RoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *RoleUpdateRequest) GetRole() RoleUpdate`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleUpdateRequest) GetRoleOk() (*RoleUpdate, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleUpdateRequest) SetRole(v RoleUpdate)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleUpdateRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


