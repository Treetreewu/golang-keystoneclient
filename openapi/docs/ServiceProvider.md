# ServiceProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthUrl** | Pointer to **string** |  | [optional] [default to "null"]
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**RelayStatePrefix** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceProvider

`func NewServiceProvider() *ServiceProvider`

NewServiceProvider instantiates a new ServiceProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderWithDefaults

`func NewServiceProviderWithDefaults() *ServiceProvider`

NewServiceProviderWithDefaults instantiates a new ServiceProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthUrl

`func (o *ServiceProvider) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *ServiceProvider) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *ServiceProvider) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *ServiceProvider) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceProvider) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceProvider) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceProvider) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceProvider) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ServiceProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceProvider) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ServiceProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRelayStatePrefix

`func (o *ServiceProvider) GetRelayStatePrefix() string`

GetRelayStatePrefix returns the RelayStatePrefix field if non-nil, zero value otherwise.

### GetRelayStatePrefixOk

`func (o *ServiceProvider) GetRelayStatePrefixOk() (*string, bool)`

GetRelayStatePrefixOk returns a tuple with the RelayStatePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayStatePrefix

`func (o *ServiceProvider) SetRelayStatePrefix(v string)`

SetRelayStatePrefix sets RelayStatePrefix field to given value.

### HasRelayStatePrefix

`func (o *ServiceProvider) HasRelayStatePrefix() bool`

HasRelayStatePrefix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


