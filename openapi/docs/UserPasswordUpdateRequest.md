# UserPasswordUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserPasswordUpdate**](UserPasswordUpdate.md) |  | 

## Methods

### NewUserPasswordUpdateRequest

`func NewUserPasswordUpdateRequest(user UserPasswordUpdate, ) *UserPasswordUpdateRequest`

NewUserPasswordUpdateRequest instantiates a new UserPasswordUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordUpdateRequestWithDefaults

`func NewUserPasswordUpdateRequestWithDefaults() *UserPasswordUpdateRequest`

NewUserPasswordUpdateRequestWithDefaults instantiates a new UserPasswordUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserPasswordUpdateRequest) GetUser() UserPasswordUpdate`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserPasswordUpdateRequest) GetUserOk() (*UserPasswordUpdate, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserPasswordUpdateRequest) SetUser(v UserPasswordUpdate)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


