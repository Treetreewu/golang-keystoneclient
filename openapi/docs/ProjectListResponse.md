# ProjectListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewProjectListResponse

`func NewProjectListResponse() *ProjectListResponse`

NewProjectListResponse instantiates a new ProjectListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectListResponseWithDefaults

`func NewProjectListResponseWithDefaults() *ProjectListResponse`

NewProjectListResponseWithDefaults instantiates a new ProjectListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *ProjectListResponse) GetProjects() []Project`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ProjectListResponse) GetProjectsOk() (*[]Project, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ProjectListResponse) SetProjects(v []Project)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ProjectListResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetLinks

`func (o *ProjectListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ProjectListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ProjectListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ProjectListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


