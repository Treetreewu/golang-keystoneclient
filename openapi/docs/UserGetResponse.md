# UserGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**User**](User.md) |  | [optional] 

## Methods

### NewUserGetResponse

`func NewUserGetResponse() *UserGetResponse`

NewUserGetResponse instantiates a new UserGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserGetResponseWithDefaults

`func NewUserGetResponseWithDefaults() *UserGetResponse`

NewUserGetResponseWithDefaults instantiates a new UserGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UserGetResponse) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserGetResponse) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserGetResponse) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *UserGetResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


