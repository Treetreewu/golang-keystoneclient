# DomainConfigLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The LDAP URL. | [optional] 
**UserTreeDn** | Pointer to **string** | The base distinguished name (DN) of LDAP, from where all users can be reached. | [optional] 

## Methods

### NewDomainConfigLDAP

`func NewDomainConfigLDAP() *DomainConfigLDAP`

NewDomainConfigLDAP instantiates a new DomainConfigLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainConfigLDAPWithDefaults

`func NewDomainConfigLDAPWithDefaults() *DomainConfigLDAP`

NewDomainConfigLDAPWithDefaults instantiates a new DomainConfigLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DomainConfigLDAP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DomainConfigLDAP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DomainConfigLDAP) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DomainConfigLDAP) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserTreeDn

`func (o *DomainConfigLDAP) GetUserTreeDn() string`

GetUserTreeDn returns the UserTreeDn field if non-nil, zero value otherwise.

### GetUserTreeDnOk

`func (o *DomainConfigLDAP) GetUserTreeDnOk() (*string, bool)`

GetUserTreeDnOk returns a tuple with the UserTreeDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTreeDn

`func (o *DomainConfigLDAP) SetUserTreeDn(v string)`

SetUserTreeDn sets UserTreeDn field to given value.

### HasUserTreeDn

`func (o *DomainConfigLDAP) HasUserTreeDn() bool`

HasUserTreeDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


