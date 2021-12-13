# RoleAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**RoleAssignmentLinks**](RoleAssignmentLinks.md) |  | [optional] 
**Role** | Pointer to [**RoleAssignmentListResponseRole**](RoleAssignmentListResponseRole.md) |  | [optional] 
**User** | Pointer to [**RoleAssignmentListResponseRole**](RoleAssignmentListResponseRole.md) |  | [optional] 
**Scope** | Pointer to [**RoleAssignmentScope**](RoleAssignmentScope.md) |  | [optional] 

## Methods

### NewRoleAssignment

`func NewRoleAssignment() *RoleAssignment`

NewRoleAssignment instantiates a new RoleAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentWithDefaults

`func NewRoleAssignmentWithDefaults() *RoleAssignment`

NewRoleAssignmentWithDefaults instantiates a new RoleAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RoleAssignment) GetLinks() RoleAssignmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleAssignment) GetLinksOk() (*RoleAssignmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleAssignment) SetLinks(v RoleAssignmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleAssignment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRole

`func (o *RoleAssignment) GetRole() RoleAssignmentListResponseRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RoleAssignment) GetRoleOk() (*RoleAssignmentListResponseRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RoleAssignment) SetRole(v RoleAssignmentListResponseRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *RoleAssignment) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetUser

`func (o *RoleAssignment) GetUser() RoleAssignmentListResponseRole`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RoleAssignment) GetUserOk() (*RoleAssignmentListResponseRole, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RoleAssignment) SetUser(v RoleAssignmentListResponseRole)`

SetUser sets User field to given value.

### HasUser

`func (o *RoleAssignment) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetScope

`func (o *RoleAssignment) GetScope() RoleAssignmentScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RoleAssignment) GetScopeOk() (*RoleAssignmentScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RoleAssignment) SetScope(v RoleAssignmentScope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RoleAssignment) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


