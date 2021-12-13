# RoleAssignmentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**RoleAssignments** | Pointer to [**[]RoleAssignment**](RoleAssignment.md) |  | [optional] 

## Methods

### NewRoleAssignmentListResponse

`func NewRoleAssignmentListResponse() *RoleAssignmentListResponse`

NewRoleAssignmentListResponse instantiates a new RoleAssignmentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleAssignmentListResponseWithDefaults

`func NewRoleAssignmentListResponseWithDefaults() *RoleAssignmentListResponse`

NewRoleAssignmentListResponseWithDefaults instantiates a new RoleAssignmentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RoleAssignmentListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RoleAssignmentListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RoleAssignmentListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RoleAssignmentListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRoleAssignments

`func (o *RoleAssignmentListResponse) GetRoleAssignments() []RoleAssignment`

GetRoleAssignments returns the RoleAssignments field if non-nil, zero value otherwise.

### GetRoleAssignmentsOk

`func (o *RoleAssignmentListResponse) GetRoleAssignmentsOk() (*[]RoleAssignment, bool)`

GetRoleAssignmentsOk returns a tuple with the RoleAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAssignments

`func (o *RoleAssignmentListResponse) SetRoleAssignments(v []RoleAssignment)`

SetRoleAssignments sets RoleAssignments field to given value.

### HasRoleAssignments

`func (o *RoleAssignmentListResponse) HasRoleAssignments() bool`

HasRoleAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


