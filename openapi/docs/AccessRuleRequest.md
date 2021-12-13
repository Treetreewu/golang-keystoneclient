# AccessRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** |  | [optional] 
**Method** | Pointer to **string** |  | [optional] 
**Service** | Pointer to **string** |  | [optional] 

## Methods

### NewAccessRuleRequest

`func NewAccessRuleRequest() *AccessRuleRequest`

NewAccessRuleRequest instantiates a new AccessRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRuleRequestWithDefaults

`func NewAccessRuleRequestWithDefaults() *AccessRuleRequest`

NewAccessRuleRequestWithDefaults instantiates a new AccessRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *AccessRuleRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AccessRuleRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AccessRuleRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AccessRuleRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetMethod

`func (o *AccessRuleRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *AccessRuleRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *AccessRuleRequest) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *AccessRuleRequest) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetService

`func (o *AccessRuleRequest) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *AccessRuleRequest) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *AccessRuleRequest) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *AccessRuleRequest) HasService() bool`

HasService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


