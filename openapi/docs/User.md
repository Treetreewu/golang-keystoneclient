# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The user name. Must be unique within the owning domain. | [optional] 
**DefaultProjectId** | Pointer to **NullableString** | The ID of the default project for the user. A user’s default project must not be a domain. Setting this attribute does not grant any actual authorization on the project, and is merely provided for convenience. Therefore, the referenced project does not need to exist within the user domain. (Since v3.1) If the user does not have authorization to their default project, the default project is ignored at token creation. (Since v3.1) Additionally, if your default project is not valid, a token is issued without an explicit scope of authorization. | [optional] 
**DomainId** | Pointer to **string** | The ID of the domain of the user. If the domain ID is not provided in the request, the Identity service will attempt to pull the domain ID from the token used in the request. Note that this requires the use of a domain-scoped token. | [optional] 
**Federated** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Enabled** | Pointer to **bool** | If the user is enabled, this value is &#x60;true&#x60;. If the user is disabled, this value is &#x60;false&#x60;. | [optional] 
**Email** | Pointer to **NullableString** | The email of the user. | [optional] 
**Description** | Pointer to **string** | Description of the user. | [optional] 
**Options** | Pointer to [**UserOptions**](UserOptions.md) |  | [optional] 
**Links** | Pointer to [**SelfLink**](SelfLink.md) |  | [optional] 
**PasswordExpiresAt** | Pointer to **NullableTime** | The date and time when the password expires. The time zone is UTC.  This is a response object attribute; not valid for requests. A &#x60;null&#x60; value indicates that the password never expires.  New in version 3.7  | [optional] 
**UserType** | Pointer to **NullableString** |  | [optional] 
**UserRole** | Pointer to **NullableString** |  | [optional] 
**FailedAuthLoginCount** | Pointer to **int32** |  | [optional] 
**AuthLoginLockedEndTime** | Pointer to **string** | Date time in &#x60;YYYYMMDDHHmm&#x60; format. | [optional] 
**AuthLoginStartTime** | Pointer to **string** |  | [optional] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultProjectId

`func (o *User) GetDefaultProjectId() string`

GetDefaultProjectId returns the DefaultProjectId field if non-nil, zero value otherwise.

### GetDefaultProjectIdOk

`func (o *User) GetDefaultProjectIdOk() (*string, bool)`

GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProjectId

`func (o *User) SetDefaultProjectId(v string)`

SetDefaultProjectId sets DefaultProjectId field to given value.

### HasDefaultProjectId

`func (o *User) HasDefaultProjectId() bool`

HasDefaultProjectId returns a boolean if a field has been set.

### SetDefaultProjectIdNil

`func (o *User) SetDefaultProjectIdNil(b bool)`

 SetDefaultProjectIdNil sets the value for DefaultProjectId to be an explicit nil

### UnsetDefaultProjectId
`func (o *User) UnsetDefaultProjectId()`

UnsetDefaultProjectId ensures that no value is present for DefaultProjectId, not even an explicit nil
### GetDomainId

`func (o *User) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *User) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *User) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *User) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetFederated

`func (o *User) GetFederated() []map[string]interface{}`

GetFederated returns the Federated field if non-nil, zero value otherwise.

### GetFederatedOk

`func (o *User) GetFederatedOk() (*[]map[string]interface{}, bool)`

GetFederatedOk returns a tuple with the Federated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederated

`func (o *User) SetFederated(v []map[string]interface{})`

SetFederated sets Federated field to given value.

### HasFederated

`func (o *User) HasFederated() bool`

HasFederated returns a boolean if a field has been set.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *User) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *User) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptions

`func (o *User) GetOptions() UserOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *User) GetOptionsOk() (*UserOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *User) SetOptions(v UserOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *User) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetLinks

`func (o *User) GetLinks() SelfLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *User) GetLinksOk() (*SelfLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *User) SetLinks(v SelfLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *User) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetPasswordExpiresAt

`func (o *User) GetPasswordExpiresAt() time.Time`

GetPasswordExpiresAt returns the PasswordExpiresAt field if non-nil, zero value otherwise.

### GetPasswordExpiresAtOk

`func (o *User) GetPasswordExpiresAtOk() (*time.Time, bool)`

GetPasswordExpiresAtOk returns a tuple with the PasswordExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordExpiresAt

`func (o *User) SetPasswordExpiresAt(v time.Time)`

SetPasswordExpiresAt sets PasswordExpiresAt field to given value.

### HasPasswordExpiresAt

`func (o *User) HasPasswordExpiresAt() bool`

HasPasswordExpiresAt returns a boolean if a field has been set.

### SetPasswordExpiresAtNil

`func (o *User) SetPasswordExpiresAtNil(b bool)`

 SetPasswordExpiresAtNil sets the value for PasswordExpiresAt to be an explicit nil

### UnsetPasswordExpiresAt
`func (o *User) UnsetPasswordExpiresAt()`

UnsetPasswordExpiresAt ensures that no value is present for PasswordExpiresAt, not even an explicit nil
### GetUserType

`func (o *User) GetUserType() string`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *User) GetUserTypeOk() (*string, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *User) SetUserType(v string)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *User) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### SetUserTypeNil

`func (o *User) SetUserTypeNil(b bool)`

 SetUserTypeNil sets the value for UserType to be an explicit nil

### UnsetUserType
`func (o *User) UnsetUserType()`

UnsetUserType ensures that no value is present for UserType, not even an explicit nil
### GetUserRole

`func (o *User) GetUserRole() string`

GetUserRole returns the UserRole field if non-nil, zero value otherwise.

### GetUserRoleOk

`func (o *User) GetUserRoleOk() (*string, bool)`

GetUserRoleOk returns a tuple with the UserRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRole

`func (o *User) SetUserRole(v string)`

SetUserRole sets UserRole field to given value.

### HasUserRole

`func (o *User) HasUserRole() bool`

HasUserRole returns a boolean if a field has been set.

### SetUserRoleNil

`func (o *User) SetUserRoleNil(b bool)`

 SetUserRoleNil sets the value for UserRole to be an explicit nil

### UnsetUserRole
`func (o *User) UnsetUserRole()`

UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
### GetFailedAuthLoginCount

`func (o *User) GetFailedAuthLoginCount() int32`

GetFailedAuthLoginCount returns the FailedAuthLoginCount field if non-nil, zero value otherwise.

### GetFailedAuthLoginCountOk

`func (o *User) GetFailedAuthLoginCountOk() (*int32, bool)`

GetFailedAuthLoginCountOk returns a tuple with the FailedAuthLoginCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAuthLoginCount

`func (o *User) SetFailedAuthLoginCount(v int32)`

SetFailedAuthLoginCount sets FailedAuthLoginCount field to given value.

### HasFailedAuthLoginCount

`func (o *User) HasFailedAuthLoginCount() bool`

HasFailedAuthLoginCount returns a boolean if a field has been set.

### GetAuthLoginLockedEndTime

`func (o *User) GetAuthLoginLockedEndTime() string`

GetAuthLoginLockedEndTime returns the AuthLoginLockedEndTime field if non-nil, zero value otherwise.

### GetAuthLoginLockedEndTimeOk

`func (o *User) GetAuthLoginLockedEndTimeOk() (*string, bool)`

GetAuthLoginLockedEndTimeOk returns a tuple with the AuthLoginLockedEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLoginLockedEndTime

`func (o *User) SetAuthLoginLockedEndTime(v string)`

SetAuthLoginLockedEndTime sets AuthLoginLockedEndTime field to given value.

### HasAuthLoginLockedEndTime

`func (o *User) HasAuthLoginLockedEndTime() bool`

HasAuthLoginLockedEndTime returns a boolean if a field has been set.

### GetAuthLoginStartTime

`func (o *User) GetAuthLoginStartTime() string`

GetAuthLoginStartTime returns the AuthLoginStartTime field if non-nil, zero value otherwise.

### GetAuthLoginStartTimeOk

`func (o *User) GetAuthLoginStartTimeOk() (*string, bool)`

GetAuthLoginStartTimeOk returns a tuple with the AuthLoginStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthLoginStartTime

`func (o *User) SetAuthLoginStartTime(v string)`

SetAuthLoginStartTime sets AuthLoginStartTime field to given value.

### HasAuthLoginStartTime

`func (o *User) HasAuthLoginStartTime() bool`

HasAuthLoginStartTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


