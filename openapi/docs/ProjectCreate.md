# ProjectCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | Pointer to **string** | The ID of the parent of the project.  If specified on project creation, this places the project within a hierarchy and implicitly defines the owning domain, which will be the same domain as the parent specified. If &#x60;parent_id&#x60; is not specified and &#x60;is_domain&#x60; is &#x60;false&#x60;, then the project will use its owning domain as its parent. If &#x60;is_domain&#x60; is &#x60;true&#x60; (i.e. the project is acting as a domain), then &#x60;parent_id&#x60; must not specified (or if it is, it must be null) since domains have no parents.  &#x60;parent_id&#x60; is immutable, and can’t be updated after the project is created - hence a project cannot be moved within the hierarchy.  | [optional] 

## Methods

### NewProjectCreate

`func NewProjectCreate() *ProjectCreate`

NewProjectCreate instantiates a new ProjectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectCreateWithDefaults

`func NewProjectCreateWithDefaults() *ProjectCreate`

NewProjectCreateWithDefaults instantiates a new ProjectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *ProjectCreate) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ProjectCreate) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ProjectCreate) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ProjectCreate) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


