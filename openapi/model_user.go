/*
Keystone API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// User struct for User
type User struct {
	Id *string `json:"id,omitempty"`
	// The user name. Must be unique within the owning domain.
	Name *string `json:"name,omitempty"`
	// The ID of the default project for the user. A user’s default project must not be a domain. Setting this attribute does not grant any actual authorization on the project, and is merely provided for convenience. Therefore, the referenced project does not need to exist within the user domain. (Since v3.1) If the user does not have authorization to their default project, the default project is ignored at token creation. (Since v3.1) Additionally, if your default project is not valid, a token is issued without an explicit scope of authorization.
	DefaultProjectId NullableString `json:"default_project_id,omitempty"`
	// The ID of the domain of the user. If the domain ID is not provided in the request, the Identity service will attempt to pull the domain ID from the token used in the request. Note that this requires the use of a domain-scoped token.
	DomainId  *string                   `json:"domain_id,omitempty"`
	Federated *[]map[string]interface{} `json:"federated,omitempty"`
	// If the user is enabled, this value is `true`. If the user is disabled, this value is `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The email of the user.
	Email NullableString `json:"email,omitempty"`
	// Description of the user.
	Description *string      `json:"description,omitempty"`
	Options     *UserOptions `json:"options,omitempty"`
	Links       *SelfLink    `json:"links,omitempty"`
	// The date and time when the password expires. The time zone is UTC.  This is a response object attribute; not valid for requests. A `null` value indicates that the password never expires.  New in version 3.7
	PasswordExpiresAt    NullableTime   `json:"password_expires_at,omitempty"`
	UserType             NullableString `json:"user_type,omitempty"`
	UserRole             NullableString `json:"user_role,omitempty"`
	FailedAuthLoginCount *int32         `json:"failed_auth_login_count,omitempty"`
	// Date time in `YYYYMMDDHHmm` format.
	AuthLoginLockedEndTime *string `json:"auth_login_locked_end_time,omitempty"`
	AuthLoginStartTime     *string `json:"auth_login_start_time,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *User) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *User) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetDefaultProjectId returns the DefaultProjectId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetDefaultProjectId() string {
	if o == nil || o.DefaultProjectId.Get() == nil {
		var ret string
		return ret
	}
	return *o.DefaultProjectId.Get()
}

// GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetDefaultProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultProjectId.Get(), o.DefaultProjectId.IsSet()
}

// HasDefaultProjectId returns a boolean if a field has been set.
func (o *User) HasDefaultProjectId() bool {
	if o != nil && o.DefaultProjectId.IsSet() {
		return true
	}

	return false
}

// SetDefaultProjectId gets a reference to the given NullableString and assigns it to the DefaultProjectId field.
func (o *User) SetDefaultProjectId(v string) {
	o.DefaultProjectId.Set(&v)
}

// SetDefaultProjectIdNil sets the value for DefaultProjectId to be an explicit nil
func (o *User) SetDefaultProjectIdNil() {
	o.DefaultProjectId.Set(nil)
}

// UnsetDefaultProjectId ensures that no value is present for DefaultProjectId, not even an explicit nil
func (o *User) UnsetDefaultProjectId() {
	o.DefaultProjectId.Unset()
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *User) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *User) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *User) SetDomainId(v string) {
	o.DomainId = &v
}

// GetFederated returns the Federated field value if set, zero value otherwise.
func (o *User) GetFederated() []map[string]interface{} {
	if o == nil || o.Federated == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Federated
}

// GetFederatedOk returns a tuple with the Federated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFederatedOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Federated == nil {
		return nil, false
	}
	return o.Federated, true
}

// HasFederated returns a boolean if a field has been set.
func (o *User) HasFederated() bool {
	if o != nil && o.Federated != nil {
		return true
	}

	return false
}

// SetFederated gets a reference to the given []map[string]interface{} and assigns it to the Federated field.
func (o *User) SetFederated(v []map[string]interface{}) {
	o.Federated = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *User) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *User) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *User) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetEmail() string {
	if o == nil || o.Email.Get() == nil {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *User) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *User) SetEmail(v string) {
	o.Email.Set(&v)
}

// SetEmailNil sets the value for Email to be an explicit nil
func (o *User) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *User) UnsetEmail() {
	o.Email.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *User) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *User) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *User) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *User) GetOptions() UserOptions {
	if o == nil || o.Options == nil {
		var ret UserOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOptionsOk() (*UserOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *User) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UserOptions and assigns it to the Options field.
func (o *User) SetOptions(v UserOptions) {
	o.Options = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *User) GetLinks() SelfLink {
	if o == nil || o.Links == nil {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLinksOk() (*SelfLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *User) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *User) SetLinks(v SelfLink) {
	o.Links = &v
}

// GetPasswordExpiresAt returns the PasswordExpiresAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetPasswordExpiresAt() time.Time {
	if o == nil || o.PasswordExpiresAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PasswordExpiresAt.Get()
}

// GetPasswordExpiresAtOk returns a tuple with the PasswordExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPasswordExpiresAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PasswordExpiresAt.Get(), o.PasswordExpiresAt.IsSet()
}

// HasPasswordExpiresAt returns a boolean if a field has been set.
func (o *User) HasPasswordExpiresAt() bool {
	if o != nil && o.PasswordExpiresAt.IsSet() {
		return true
	}

	return false
}

// SetPasswordExpiresAt gets a reference to the given NullableTime and assigns it to the PasswordExpiresAt field.
func (o *User) SetPasswordExpiresAt(v time.Time) {
	o.PasswordExpiresAt.Set(&v)
}

// SetPasswordExpiresAtNil sets the value for PasswordExpiresAt to be an explicit nil
func (o *User) SetPasswordExpiresAtNil() {
	o.PasswordExpiresAt.Set(nil)
}

// UnsetPasswordExpiresAt ensures that no value is present for PasswordExpiresAt, not even an explicit nil
func (o *User) UnsetPasswordExpiresAt() {
	o.PasswordExpiresAt.Unset()
}

// GetUserType returns the UserType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUserType() string {
	if o == nil || o.UserType.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserType.Get()
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUserTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserType.Get(), o.UserType.IsSet()
}

// HasUserType returns a boolean if a field has been set.
func (o *User) HasUserType() bool {
	if o != nil && o.UserType.IsSet() {
		return true
	}

	return false
}

// SetUserType gets a reference to the given NullableString and assigns it to the UserType field.
func (o *User) SetUserType(v string) {
	o.UserType.Set(&v)
}

// SetUserTypeNil sets the value for UserType to be an explicit nil
func (o *User) SetUserTypeNil() {
	o.UserType.Set(nil)
}

// UnsetUserType ensures that no value is present for UserType, not even an explicit nil
func (o *User) UnsetUserType() {
	o.UserType.Unset()
}

// GetUserRole returns the UserRole field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *User) GetUserRole() string {
	if o == nil || o.UserRole.Get() == nil {
		var ret string
		return ret
	}
	return *o.UserRole.Get()
}

// GetUserRoleOk returns a tuple with the UserRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetUserRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserRole.Get(), o.UserRole.IsSet()
}

// HasUserRole returns a boolean if a field has been set.
func (o *User) HasUserRole() bool {
	if o != nil && o.UserRole.IsSet() {
		return true
	}

	return false
}

// SetUserRole gets a reference to the given NullableString and assigns it to the UserRole field.
func (o *User) SetUserRole(v string) {
	o.UserRole.Set(&v)
}

// SetUserRoleNil sets the value for UserRole to be an explicit nil
func (o *User) SetUserRoleNil() {
	o.UserRole.Set(nil)
}

// UnsetUserRole ensures that no value is present for UserRole, not even an explicit nil
func (o *User) UnsetUserRole() {
	o.UserRole.Unset()
}

// GetFailedAuthLoginCount returns the FailedAuthLoginCount field value if set, zero value otherwise.
func (o *User) GetFailedAuthLoginCount() int32 {
	if o == nil || o.FailedAuthLoginCount == nil {
		var ret int32
		return ret
	}
	return *o.FailedAuthLoginCount
}

// GetFailedAuthLoginCountOk returns a tuple with the FailedAuthLoginCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetFailedAuthLoginCountOk() (*int32, bool) {
	if o == nil || o.FailedAuthLoginCount == nil {
		return nil, false
	}
	return o.FailedAuthLoginCount, true
}

// HasFailedAuthLoginCount returns a boolean if a field has been set.
func (o *User) HasFailedAuthLoginCount() bool {
	if o != nil && o.FailedAuthLoginCount != nil {
		return true
	}

	return false
}

// SetFailedAuthLoginCount gets a reference to the given int32 and assigns it to the FailedAuthLoginCount field.
func (o *User) SetFailedAuthLoginCount(v int32) {
	o.FailedAuthLoginCount = &v
}

// GetAuthLoginLockedEndTime returns the AuthLoginLockedEndTime field value if set, zero value otherwise.
func (o *User) GetAuthLoginLockedEndTime() string {
	if o == nil || o.AuthLoginLockedEndTime == nil {
		var ret string
		return ret
	}
	return *o.AuthLoginLockedEndTime
}

// GetAuthLoginLockedEndTimeOk returns a tuple with the AuthLoginLockedEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAuthLoginLockedEndTimeOk() (*string, bool) {
	if o == nil || o.AuthLoginLockedEndTime == nil {
		return nil, false
	}
	return o.AuthLoginLockedEndTime, true
}

// HasAuthLoginLockedEndTime returns a boolean if a field has been set.
func (o *User) HasAuthLoginLockedEndTime() bool {
	if o != nil && o.AuthLoginLockedEndTime != nil {
		return true
	}

	return false
}

// SetAuthLoginLockedEndTime gets a reference to the given string and assigns it to the AuthLoginLockedEndTime field.
func (o *User) SetAuthLoginLockedEndTime(v string) {
	o.AuthLoginLockedEndTime = &v
}

// GetAuthLoginStartTime returns the AuthLoginStartTime field value if set, zero value otherwise.
func (o *User) GetAuthLoginStartTime() string {
	if o == nil || o.AuthLoginStartTime == nil {
		var ret string
		return ret
	}
	return *o.AuthLoginStartTime
}

// GetAuthLoginStartTimeOk returns a tuple with the AuthLoginStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAuthLoginStartTimeOk() (*string, bool) {
	if o == nil || o.AuthLoginStartTime == nil {
		return nil, false
	}
	return o.AuthLoginStartTime, true
}

// HasAuthLoginStartTime returns a boolean if a field has been set.
func (o *User) HasAuthLoginStartTime() bool {
	if o != nil && o.AuthLoginStartTime != nil {
		return true
	}

	return false
}

// SetAuthLoginStartTime gets a reference to the given string and assigns it to the AuthLoginStartTime field.
func (o *User) SetAuthLoginStartTime(v string) {
	o.AuthLoginStartTime = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultProjectId.IsSet() {
		toSerialize["default_project_id"] = o.DefaultProjectId.Get()
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.Federated != nil {
		toSerialize["federated"] = o.Federated
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.PasswordExpiresAt.IsSet() {
		toSerialize["password_expires_at"] = o.PasswordExpiresAt.Get()
	}
	if o.UserType.IsSet() {
		toSerialize["user_type"] = o.UserType.Get()
	}
	if o.UserRole.IsSet() {
		toSerialize["user_role"] = o.UserRole.Get()
	}
	if o.FailedAuthLoginCount != nil {
		toSerialize["failed_auth_login_count"] = o.FailedAuthLoginCount
	}
	if o.AuthLoginLockedEndTime != nil {
		toSerialize["auth_login_locked_end_time"] = o.AuthLoginLockedEndTime
	}
	if o.AuthLoginStartTime != nil {
		toSerialize["auth_login_start_time"] = o.AuthLoginStartTime
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
