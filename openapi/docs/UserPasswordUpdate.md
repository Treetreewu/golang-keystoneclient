# UserPasswordUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OriginalPassword** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewUserPasswordUpdate

`func NewUserPasswordUpdate(originalPassword string, password string, ) *UserPasswordUpdate`

NewUserPasswordUpdate instantiates a new UserPasswordUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordUpdateWithDefaults

`func NewUserPasswordUpdateWithDefaults() *UserPasswordUpdate`

NewUserPasswordUpdateWithDefaults instantiates a new UserPasswordUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOriginalPassword

`func (o *UserPasswordUpdate) GetOriginalPassword() string`

GetOriginalPassword returns the OriginalPassword field if non-nil, zero value otherwise.

### GetOriginalPasswordOk

`func (o *UserPasswordUpdate) GetOriginalPasswordOk() (*string, bool)`

GetOriginalPasswordOk returns a tuple with the OriginalPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPassword

`func (o *UserPasswordUpdate) SetOriginalPassword(v string)`

SetOriginalPassword sets OriginalPassword field to given value.


### GetPassword

`func (o *UserPasswordUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPasswordUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPasswordUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


