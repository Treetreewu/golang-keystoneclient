# DomainConfigLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The LDAP URL. | [optional] 
**UserTreeDn** | Pointer to **string** | The base distinguished name (DN) of LDAP, from where all users can be reached. | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**QueryScope** | Pointer to **string** | 定义了在搜索基中搜索的深度。值“one”表示对下面紧接基准对象的对象搜索，但不包括基准对象本身。 值“sub”表示同时搜索基准对象本身和它下面的整个子树。 | [optional] 
**UserObjectclass** | Pointer to **string** |  | [optional] [default to "organizationalPerson"]
**UserIdAttribute** | Pointer to **string** |  | [optional] [default to "cn"]
**UserNameAttribute** | Pointer to **string** |  | [optional] [default to "cn"]
**UserMailAttribute** | Pointer to **string** |  | [optional] [default to "mail"]
**PageSize** | Pointer to **int32** | set 0 to disable pagination | [optional] 
**UserFilter** | Pointer to **string** |  | [optional] 

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

### GetUser

`func (o *DomainConfigLDAP) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DomainConfigLDAP) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DomainConfigLDAP) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DomainConfigLDAP) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *DomainConfigLDAP) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DomainConfigLDAP) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DomainConfigLDAP) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DomainConfigLDAP) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetQueryScope

`func (o *DomainConfigLDAP) GetQueryScope() string`

GetQueryScope returns the QueryScope field if non-nil, zero value otherwise.

### GetQueryScopeOk

`func (o *DomainConfigLDAP) GetQueryScopeOk() (*string, bool)`

GetQueryScopeOk returns a tuple with the QueryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryScope

`func (o *DomainConfigLDAP) SetQueryScope(v string)`

SetQueryScope sets QueryScope field to given value.

### HasQueryScope

`func (o *DomainConfigLDAP) HasQueryScope() bool`

HasQueryScope returns a boolean if a field has been set.

### GetUserObjectclass

`func (o *DomainConfigLDAP) GetUserObjectclass() string`

GetUserObjectclass returns the UserObjectclass field if non-nil, zero value otherwise.

### GetUserObjectclassOk

`func (o *DomainConfigLDAP) GetUserObjectclassOk() (*string, bool)`

GetUserObjectclassOk returns a tuple with the UserObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectclass

`func (o *DomainConfigLDAP) SetUserObjectclass(v string)`

SetUserObjectclass sets UserObjectclass field to given value.

### HasUserObjectclass

`func (o *DomainConfigLDAP) HasUserObjectclass() bool`

HasUserObjectclass returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *DomainConfigLDAP) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *DomainConfigLDAP) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *DomainConfigLDAP) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *DomainConfigLDAP) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserNameAttribute

`func (o *DomainConfigLDAP) GetUserNameAttribute() string`

GetUserNameAttribute returns the UserNameAttribute field if non-nil, zero value otherwise.

### GetUserNameAttributeOk

`func (o *DomainConfigLDAP) GetUserNameAttributeOk() (*string, bool)`

GetUserNameAttributeOk returns a tuple with the UserNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameAttribute

`func (o *DomainConfigLDAP) SetUserNameAttribute(v string)`

SetUserNameAttribute sets UserNameAttribute field to given value.

### HasUserNameAttribute

`func (o *DomainConfigLDAP) HasUserNameAttribute() bool`

HasUserNameAttribute returns a boolean if a field has been set.

### GetUserMailAttribute

`func (o *DomainConfigLDAP) GetUserMailAttribute() string`

GetUserMailAttribute returns the UserMailAttribute field if non-nil, zero value otherwise.

### GetUserMailAttributeOk

`func (o *DomainConfigLDAP) GetUserMailAttributeOk() (*string, bool)`

GetUserMailAttributeOk returns a tuple with the UserMailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMailAttribute

`func (o *DomainConfigLDAP) SetUserMailAttribute(v string)`

SetUserMailAttribute sets UserMailAttribute field to given value.

### HasUserMailAttribute

`func (o *DomainConfigLDAP) HasUserMailAttribute() bool`

HasUserMailAttribute returns a boolean if a field has been set.

### GetPageSize

`func (o *DomainConfigLDAP) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DomainConfigLDAP) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DomainConfigLDAP) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DomainConfigLDAP) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetUserFilter

`func (o *DomainConfigLDAP) GetUserFilter() string`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *DomainConfigLDAP) GetUserFilterOk() (*string, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *DomainConfigLDAP) SetUserFilter(v string)`

SetUserFilter sets UserFilter field to given value.

### HasUserFilter

`func (o *DomainConfigLDAP) HasUserFilter() bool`

HasUserFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


