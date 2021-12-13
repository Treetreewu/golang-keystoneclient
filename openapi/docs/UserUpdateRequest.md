# UserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**UserUpdate**](UserUpdate.md) |  | 

## Methods

### NewUserUpdateRequest

`func NewUserUpdateRequest(user UserUpdate, ) *UserUpdateRequest`

NewUserUpdateRequest instantiates a new UserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateRequestWithDefaults

`func NewUserUpdateRequestWithDefaults() *UserUpdateRequest`

NewUserUpdateRequestWithDefaults instantiates a new UserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserUpdateRequest) GetUser() UserUpdate`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserUpdateRequest) GetUserOk() (*UserUpdate, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserUpdateRequest) SetUser(v UserUpdate)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


