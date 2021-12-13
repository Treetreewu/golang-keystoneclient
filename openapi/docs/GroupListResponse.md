# GroupListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]Group**](Group.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewGroupListResponse

`func NewGroupListResponse() *GroupListResponse`

NewGroupListResponse instantiates a new GroupListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupListResponseWithDefaults

`func NewGroupListResponseWithDefaults() *GroupListResponse`

NewGroupListResponseWithDefaults instantiates a new GroupListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *GroupListResponse) GetGroups() []Group`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *GroupListResponse) GetGroupsOk() (*[]Group, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *GroupListResponse) SetGroups(v []Group)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *GroupListResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetLinks

`func (o *GroupListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GroupListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GroupListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GroupListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


