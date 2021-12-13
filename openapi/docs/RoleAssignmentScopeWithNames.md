# RoleAssignmentScopeWithNames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to [**IdAndName**](IdAndName.md) |  | [optional] 
**Project** | Pointer to [**RoleAssignmentListResponseWithNamesRole**](RoleAssignmentListResponseWithNamesRole.md) |  | [optional] 

## Methods

### NewRoleAssignmentScopeWithNames

`func NewRoleAssignmentScopeWithNames() *RoleAssignmentScopeWithNames`

NewRoleAssignmentScopeWithNames instantiates a new RoleAssignmentScopeWithNames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentScopeWithNamesWithDefaults

`func NewRoleAssignmentScopeWithNamesWithDefaults() *RoleAssignmentScopeWithNames`

NewRoleAssignmentScopeWithNamesWithDefaults instantiates a new RoleAssignmentScopeWithNames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *RoleAssignmentScopeWithNames) GetDomain() IdAndName`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *RoleAssignmentScopeWithNames) GetDomainOk() (*IdAndName, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *RoleAssignmentScopeWithNames) SetDomain(v IdAndName)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *RoleAssignmentScopeWithNames) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetProject

`func (o *RoleAssignmentScopeWithNames) GetProject() RoleAssignmentListResponseWithNamesRole`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *RoleAssignmentScopeWithNames) GetProjectOk() (*RoleAssignmentListResponseWithNamesRole, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *RoleAssignmentScopeWithNames) SetProject(v RoleAssignmentListResponseWithNamesRole)`

SetProject sets Project field to given value.

### HasProject

`func (o *RoleAssignmentScopeWithNames) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


