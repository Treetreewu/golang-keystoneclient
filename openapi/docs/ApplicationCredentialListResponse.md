# ApplicationCredentialListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**ApplicationCredentials** | Pointer to [**[]ApplicationCredential**](ApplicationCredential.md) |  | [optional] 

## Methods

### NewApplicationCredentialListResponse

`func NewApplicationCredentialListResponse() *ApplicationCredentialListResponse`

NewApplicationCredentialListResponse instantiates a new ApplicationCredentialListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCredentialListResponseWithDefaults

`func NewApplicationCredentialListResponseWithDefaults() *ApplicationCredentialListResponse`

NewApplicationCredentialListResponseWithDefaults instantiates a new ApplicationCredentialListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationCredentialListResponse) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationCredentialListResponse) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationCredentialListResponse) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationCredentialListResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplicationCredentials

`func (o *ApplicationCredentialListResponse) GetApplicationCredentials() []ApplicationCredential`

GetApplicationCredentials returns the ApplicationCredentials field if non-nil, zero value otherwise.

### GetApplicationCredentialsOk

`func (o *ApplicationCredentialListResponse) GetApplicationCredentialsOk() (*[]ApplicationCredential, bool)`

GetApplicationCredentialsOk returns a tuple with the ApplicationCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCredentials

`func (o *ApplicationCredentialListResponse) SetApplicationCredentials(v []ApplicationCredential)`

SetApplicationCredentials sets ApplicationCredentials field to given value.

### HasApplicationCredentials

`func (o *ApplicationCredentialListResponse) HasApplicationCredentials() bool`

HasApplicationCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


