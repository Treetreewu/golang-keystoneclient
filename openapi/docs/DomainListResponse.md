# DomainListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domains** | Pointer to [**[]Domain**](Domain.md) |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewDomainListResponse

`func NewDomainListResponse() *DomainListResponse`

NewDomainListResponse instantiates a new DomainListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainListResponseWithDefaults

`func NewDomainListResponseWithDefaults() *DomainListResponse`

NewDomainListResponseWithDefaults instantiates a new DomainListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomains

`func (o *DomainListResponse) GetDomains() []Domain`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *DomainListResponse) GetDomainsOk() (*[]Domain, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *DomainListResponse) SetDomains(v []Domain)`

SetDomains sets Domains field to given value.

### HasDomains

`func (o *DomainListResponse) HasDomains() bool`

HasDomains returns a boolean if a field has been set.

### GetLinks

`func (o *DomainListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DomainListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DomainListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DomainListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


