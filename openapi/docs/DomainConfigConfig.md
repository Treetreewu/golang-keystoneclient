# DomainConfigConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identity** | Pointer to [**DomainConfigIdentity**](DomainConfigIdentity.md) |  | [optional] 
**Ldap** | Pointer to [**DomainConfigLDAP**](DomainConfigLDAP.md) |  | [optional] 

## Methods

### NewDomainConfigConfig

`func NewDomainConfigConfig() *DomainConfigConfig`

NewDomainConfigConfig instantiates a new DomainConfigConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainConfigConfigWithDefaults

`func NewDomainConfigConfigWithDefaults() *DomainConfigConfig`

NewDomainConfigConfigWithDefaults instantiates a new DomainConfigConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentity

`func (o *DomainConfigConfig) GetIdentity() DomainConfigIdentity`

GetIdentity returns the Identity field if non-nil, zero value otherwise.

### GetIdentityOk

`func (o *DomainConfigConfig) GetIdentityOk() (*DomainConfigIdentity, bool)`

GetIdentityOk returns a tuple with the Identity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentity

`func (o *DomainConfigConfig) SetIdentity(v DomainConfigIdentity)`

SetIdentity sets Identity field to given value.

### HasIdentity

`func (o *DomainConfigConfig) HasIdentity() bool`

HasIdentity returns a boolean if a field has been set.

### GetLdap

`func (o *DomainConfigConfig) GetLdap() DomainConfigLDAP`

GetLdap returns the Ldap field if non-nil, zero value otherwise.

### GetLdapOk

`func (o *DomainConfigConfig) GetLdapOk() (*DomainConfigLDAP, bool)`

GetLdapOk returns a tuple with the Ldap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdap

`func (o *DomainConfigConfig) SetLdap(v DomainConfigLDAP)`

SetLdap sets Ldap field to given value.

### HasLdap

`func (o *DomainConfigConfig) HasLdap() bool`

HasLdap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


