# DomainConfigLdap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | The LDAP URL. | [optional] 
**UserTreeDn** | Pointer to **string** | The base distinguished name (DN) of LDAP, from where all users can be reached. | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**QueryScope** | Pointer to **string** | 定义了在搜索基中搜索的深度。值“one”表示对下面紧接基准对象的对象搜索，但不包括基准对象本身。 值“sub”表示同时搜索基准对象本身和它下面的整个子树。 | [optional] 
**UserObjectclass** | Pointer to **string** | The LDAP object class to use for users. | [optional] [default to "organizationalPerson"]
**UserIdAttribute** | Pointer to **string** | The LDAP attribute mapped to user IDs in keystone. This must NOT be a multivalued attribute. User IDs are expected to be globally unique across keystone domains and URL-safe.  | [optional] [default to "cn"]
**UserNameAttribute** | Pointer to **string** | The LDAP attribute mapped to user names in keystone. User names are expected to be unique only within a keystone domain and are not expected to be URL-safe.  | [optional] [default to "sn"]
**UserMailAttribute** | Pointer to **string** | The LDAP attribute mapped to user emails in keystone. | [optional] [default to "mail"]
**UserPasswordAttribute** | Pointer to **string** | The LDAP attribute mapped to user passwords in keystone. | [optional] [default to "userPassword"]
**UserEnabledMask** | Pointer to **int32** | Bitmask integer to select which bit indicates the enabled value if the LDAP server represents \&quot;enabled\&quot; as a bit on an integer rather than as a discrete boolean. A value of &#x60;0&#x60; indicates that the mask is not used. If this is not set to &#x60;0&#x60; the typical value is &#x60;2&#x60;. This is typically used when &#x60;[ldap] user_enabled_attribute &#x3D; userAccountControl&#x60;. Setting this option causes keystone to ignore the value of &#x60;[ldap] user_enabled_invert&#x60;.  | [optional] [default to 2]
**UserEnabledDefault** | Pointer to **string** | The default value to enable users. This should match an appropriate integer value if the LDAP server uses non-boolean (bitmask) values to indicate if a user is enabled or disabled. If this is not set to &#x60;True&#x60;, then the typical value is &#x60;512&#x60;. This is typically used when &#x60;[ldap] user_enabled_attribute &#x3D; userAccountControl&#x60;.  | [optional] [default to "true"]
**UserAllowCreate** | Pointer to **bool** | If enabled, keystone is allowed to create users in the LDAP server. | [optional] [default to true]
**UserAllowUpdate** | Pointer to **bool** | If enabled, keystone is allowed to update users in the LDAP server. | [optional] [default to true]
**UserAllowDelete** | Pointer to **bool** | If enabled, keystone is allowed to delete users in the LDAP server. | [optional] [default to true]
**PageSize** | Pointer to **int32** | set 0 to disable pagination | [optional] [default to 0]
**UserFilter** | Pointer to **string** |  | [optional] 

## Methods

### NewDomainConfigLdap

`func NewDomainConfigLdap() *DomainConfigLdap`

NewDomainConfigLdap instantiates a new DomainConfigLdap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomainConfigLdapWithDefaults

`func NewDomainConfigLdapWithDefaults() *DomainConfigLdap`

NewDomainConfigLdapWithDefaults instantiates a new DomainConfigLdap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *DomainConfigLdap) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DomainConfigLdap) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DomainConfigLdap) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DomainConfigLdap) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserTreeDn

`func (o *DomainConfigLdap) GetUserTreeDn() string`

GetUserTreeDn returns the UserTreeDn field if non-nil, zero value otherwise.

### GetUserTreeDnOk

`func (o *DomainConfigLdap) GetUserTreeDnOk() (*string, bool)`

GetUserTreeDnOk returns a tuple with the UserTreeDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTreeDn

`func (o *DomainConfigLdap) SetUserTreeDn(v string)`

SetUserTreeDn sets UserTreeDn field to given value.

### HasUserTreeDn

`func (o *DomainConfigLdap) HasUserTreeDn() bool`

HasUserTreeDn returns a boolean if a field has been set.

### GetUser

`func (o *DomainConfigLdap) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DomainConfigLdap) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DomainConfigLdap) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DomainConfigLdap) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetPassword

`func (o *DomainConfigLdap) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DomainConfigLdap) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DomainConfigLdap) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DomainConfigLdap) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSuffix

`func (o *DomainConfigLdap) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *DomainConfigLdap) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *DomainConfigLdap) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *DomainConfigLdap) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetQueryScope

`func (o *DomainConfigLdap) GetQueryScope() string`

GetQueryScope returns the QueryScope field if non-nil, zero value otherwise.

### GetQueryScopeOk

`func (o *DomainConfigLdap) GetQueryScopeOk() (*string, bool)`

GetQueryScopeOk returns a tuple with the QueryScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryScope

`func (o *DomainConfigLdap) SetQueryScope(v string)`

SetQueryScope sets QueryScope field to given value.

### HasQueryScope

`func (o *DomainConfigLdap) HasQueryScope() bool`

HasQueryScope returns a boolean if a field has been set.

### GetUserObjectclass

`func (o *DomainConfigLdap) GetUserObjectclass() string`

GetUserObjectclass returns the UserObjectclass field if non-nil, zero value otherwise.

### GetUserObjectclassOk

`func (o *DomainConfigLdap) GetUserObjectclassOk() (*string, bool)`

GetUserObjectclassOk returns a tuple with the UserObjectclass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserObjectclass

`func (o *DomainConfigLdap) SetUserObjectclass(v string)`

SetUserObjectclass sets UserObjectclass field to given value.

### HasUserObjectclass

`func (o *DomainConfigLdap) HasUserObjectclass() bool`

HasUserObjectclass returns a boolean if a field has been set.

### GetUserIdAttribute

`func (o *DomainConfigLdap) GetUserIdAttribute() string`

GetUserIdAttribute returns the UserIdAttribute field if non-nil, zero value otherwise.

### GetUserIdAttributeOk

`func (o *DomainConfigLdap) GetUserIdAttributeOk() (*string, bool)`

GetUserIdAttributeOk returns a tuple with the UserIdAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdAttribute

`func (o *DomainConfigLdap) SetUserIdAttribute(v string)`

SetUserIdAttribute sets UserIdAttribute field to given value.

### HasUserIdAttribute

`func (o *DomainConfigLdap) HasUserIdAttribute() bool`

HasUserIdAttribute returns a boolean if a field has been set.

### GetUserNameAttribute

`func (o *DomainConfigLdap) GetUserNameAttribute() string`

GetUserNameAttribute returns the UserNameAttribute field if non-nil, zero value otherwise.

### GetUserNameAttributeOk

`func (o *DomainConfigLdap) GetUserNameAttributeOk() (*string, bool)`

GetUserNameAttributeOk returns a tuple with the UserNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserNameAttribute

`func (o *DomainConfigLdap) SetUserNameAttribute(v string)`

SetUserNameAttribute sets UserNameAttribute field to given value.

### HasUserNameAttribute

`func (o *DomainConfigLdap) HasUserNameAttribute() bool`

HasUserNameAttribute returns a boolean if a field has been set.

### GetUserMailAttribute

`func (o *DomainConfigLdap) GetUserMailAttribute() string`

GetUserMailAttribute returns the UserMailAttribute field if non-nil, zero value otherwise.

### GetUserMailAttributeOk

`func (o *DomainConfigLdap) GetUserMailAttributeOk() (*string, bool)`

GetUserMailAttributeOk returns a tuple with the UserMailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserMailAttribute

`func (o *DomainConfigLdap) SetUserMailAttribute(v string)`

SetUserMailAttribute sets UserMailAttribute field to given value.

### HasUserMailAttribute

`func (o *DomainConfigLdap) HasUserMailAttribute() bool`

HasUserMailAttribute returns a boolean if a field has been set.

### GetUserPasswordAttribute

`func (o *DomainConfigLdap) GetUserPasswordAttribute() string`

GetUserPasswordAttribute returns the UserPasswordAttribute field if non-nil, zero value otherwise.

### GetUserPasswordAttributeOk

`func (o *DomainConfigLdap) GetUserPasswordAttributeOk() (*string, bool)`

GetUserPasswordAttributeOk returns a tuple with the UserPasswordAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserPasswordAttribute

`func (o *DomainConfigLdap) SetUserPasswordAttribute(v string)`

SetUserPasswordAttribute sets UserPasswordAttribute field to given value.

### HasUserPasswordAttribute

`func (o *DomainConfigLdap) HasUserPasswordAttribute() bool`

HasUserPasswordAttribute returns a boolean if a field has been set.

### GetUserEnabledMask

`func (o *DomainConfigLdap) GetUserEnabledMask() int32`

GetUserEnabledMask returns the UserEnabledMask field if non-nil, zero value otherwise.

### GetUserEnabledMaskOk

`func (o *DomainConfigLdap) GetUserEnabledMaskOk() (*int32, bool)`

GetUserEnabledMaskOk returns a tuple with the UserEnabledMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabledMask

`func (o *DomainConfigLdap) SetUserEnabledMask(v int32)`

SetUserEnabledMask sets UserEnabledMask field to given value.

### HasUserEnabledMask

`func (o *DomainConfigLdap) HasUserEnabledMask() bool`

HasUserEnabledMask returns a boolean if a field has been set.

### GetUserEnabledDefault

`func (o *DomainConfigLdap) GetUserEnabledDefault() string`

GetUserEnabledDefault returns the UserEnabledDefault field if non-nil, zero value otherwise.

### GetUserEnabledDefaultOk

`func (o *DomainConfigLdap) GetUserEnabledDefaultOk() (*string, bool)`

GetUserEnabledDefaultOk returns a tuple with the UserEnabledDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserEnabledDefault

`func (o *DomainConfigLdap) SetUserEnabledDefault(v string)`

SetUserEnabledDefault sets UserEnabledDefault field to given value.

### HasUserEnabledDefault

`func (o *DomainConfigLdap) HasUserEnabledDefault() bool`

HasUserEnabledDefault returns a boolean if a field has been set.

### GetUserAllowCreate

`func (o *DomainConfigLdap) GetUserAllowCreate() bool`

GetUserAllowCreate returns the UserAllowCreate field if non-nil, zero value otherwise.

### GetUserAllowCreateOk

`func (o *DomainConfigLdap) GetUserAllowCreateOk() (*bool, bool)`

GetUserAllowCreateOk returns a tuple with the UserAllowCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAllowCreate

`func (o *DomainConfigLdap) SetUserAllowCreate(v bool)`

SetUserAllowCreate sets UserAllowCreate field to given value.

### HasUserAllowCreate

`func (o *DomainConfigLdap) HasUserAllowCreate() bool`

HasUserAllowCreate returns a boolean if a field has been set.

### GetUserAllowUpdate

`func (o *DomainConfigLdap) GetUserAllowUpdate() bool`

GetUserAllowUpdate returns the UserAllowUpdate field if non-nil, zero value otherwise.

### GetUserAllowUpdateOk

`func (o *DomainConfigLdap) GetUserAllowUpdateOk() (*bool, bool)`

GetUserAllowUpdateOk returns a tuple with the UserAllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAllowUpdate

`func (o *DomainConfigLdap) SetUserAllowUpdate(v bool)`

SetUserAllowUpdate sets UserAllowUpdate field to given value.

### HasUserAllowUpdate

`func (o *DomainConfigLdap) HasUserAllowUpdate() bool`

HasUserAllowUpdate returns a boolean if a field has been set.

### GetUserAllowDelete

`func (o *DomainConfigLdap) GetUserAllowDelete() bool`

GetUserAllowDelete returns the UserAllowDelete field if non-nil, zero value otherwise.

### GetUserAllowDeleteOk

`func (o *DomainConfigLdap) GetUserAllowDeleteOk() (*bool, bool)`

GetUserAllowDeleteOk returns a tuple with the UserAllowDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAllowDelete

`func (o *DomainConfigLdap) SetUserAllowDelete(v bool)`

SetUserAllowDelete sets UserAllowDelete field to given value.

### HasUserAllowDelete

`func (o *DomainConfigLdap) HasUserAllowDelete() bool`

HasUserAllowDelete returns a boolean if a field has been set.

### GetPageSize

`func (o *DomainConfigLdap) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DomainConfigLdap) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DomainConfigLdap) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DomainConfigLdap) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetUserFilter

`func (o *DomainConfigLdap) GetUserFilter() string`

GetUserFilter returns the UserFilter field if non-nil, zero value otherwise.

### GetUserFilterOk

`func (o *DomainConfigLdap) GetUserFilterOk() (*string, bool)`

GetUserFilterOk returns a tuple with the UserFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFilter

`func (o *DomainConfigLdap) SetUserFilter(v string)`

SetUserFilter sets UserFilter field to given value.

### HasUserFilter

`func (o *DomainConfigLdap) HasUserFilter() bool`

HasUserFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


