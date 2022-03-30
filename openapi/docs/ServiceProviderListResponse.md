# ServiceProviderListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceProviders** | Pointer to [**[]ServiceProvider**](ServiceProvider.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewServiceProviderListResponse

`func NewServiceProviderListResponse() *ServiceProviderListResponse`

NewServiceProviderListResponse instantiates a new ServiceProviderListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceProviderListResponseWithDefaults

`func NewServiceProviderListResponseWithDefaults() *ServiceProviderListResponse`

NewServiceProviderListResponseWithDefaults instantiates a new ServiceProviderListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceProviders

`func (o *ServiceProviderListResponse) GetServiceProviders() []ServiceProvider`

GetServiceProviders returns the ServiceProviders field if non-nil, zero value otherwise.

### GetServiceProvidersOk

`func (o *ServiceProviderListResponse) GetServiceProvidersOk() (*[]ServiceProvider, bool)`

GetServiceProvidersOk returns a tuple with the ServiceProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceProviders

`func (o *ServiceProviderListResponse) SetServiceProviders(v []ServiceProvider)`

SetServiceProviders sets ServiceProviders field to given value.

### HasServiceProviders

`func (o *ServiceProviderListResponse) HasServiceProviders() bool`

HasServiceProviders returns a boolean if a field has been set.

### GetLinks

`func (o *ServiceProviderListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ServiceProviderListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ServiceProviderListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ServiceProviderListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


