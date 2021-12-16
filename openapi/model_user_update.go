/*
Keystone API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UserUpdate struct for UserUpdate
type UserUpdate struct {
	// The user name. Must be unique within the owning domain.
	Name *string `json:"name,omitempty"`
	// The ID of the default project for the user. A user’s default project must not be a domain. Setting this attribute does not grant any actual authorization on the project, and is merely provided for convenience. Therefore, the referenced project does not need to exist within the user domain. (Since v3.1) If the user does not have authorization to their default project, the default project is ignored at token creation. (Since v3.1) Additionally, if your default project is not valid, a token is issued without an explicit scope of authorization.
	DefaultProjectId *string `json:"default_project_id,omitempty"`
	// The ID of the domain of the user. If the domain ID is not provided in the request, the Identity service will attempt to pull the domain ID from the token used in the request. Note that this requires the use of a domain-scoped token.
	DomainId  *string                   `json:"domain_id,omitempty"`
	Federated *[]map[string]interface{} `json:"federated,omitempty"`
	// If the user is enabled, this value is `true`. If the user is disabled, this value is `false`.
	Enabled *bool `json:"enabled,omitempty"`
	// The password of the user.
	Password *string `json:"password,omitempty"`
	// The email of the user.
	Email *string `json:"email,omitempty"`
	// Description of the user.
	Description *string      `json:"description,omitempty"`
	Options     *UserOptions `json:"options,omitempty"`
}

// NewUserUpdate instantiates a new UserUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdate() *UserUpdate {
	this := UserUpdate{}
	return &this
}

// NewUserUpdateWithDefaults instantiates a new UserUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateWithDefaults() *UserUpdate {
	this := UserUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserUpdate) SetName(v string) {
	o.Name = &v
}

// GetDefaultProjectId returns the DefaultProjectId field value if set, zero value otherwise.
func (o *UserUpdate) GetDefaultProjectId() string {
	if o == nil || o.DefaultProjectId == nil {
		var ret string
		return ret
	}
	return *o.DefaultProjectId
}

// GetDefaultProjectIdOk returns a tuple with the DefaultProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetDefaultProjectIdOk() (*string, bool) {
	if o == nil || o.DefaultProjectId == nil {
		return nil, false
	}
	return o.DefaultProjectId, true
}

// HasDefaultProjectId returns a boolean if a field has been set.
func (o *UserUpdate) HasDefaultProjectId() bool {
	if o != nil && o.DefaultProjectId != nil {
		return true
	}

	return false
}

// SetDefaultProjectId gets a reference to the given string and assigns it to the DefaultProjectId field.
func (o *UserUpdate) SetDefaultProjectId(v string) {
	o.DefaultProjectId = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *UserUpdate) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *UserUpdate) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *UserUpdate) SetDomainId(v string) {
	o.DomainId = &v
}

// GetFederated returns the Federated field value if set, zero value otherwise.
func (o *UserUpdate) GetFederated() []map[string]interface{} {
	if o == nil || o.Federated == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Federated
}

// GetFederatedOk returns a tuple with the Federated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetFederatedOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Federated == nil {
		return nil, false
	}
	return o.Federated, true
}

// HasFederated returns a boolean if a field has been set.
func (o *UserUpdate) HasFederated() bool {
	if o != nil && o.Federated != nil {
		return true
	}

	return false
}

// SetFederated gets a reference to the given []map[string]interface{} and assigns it to the Federated field.
func (o *UserUpdate) SetFederated(v []map[string]interface{}) {
	o.Federated = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserUpdate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserUpdate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserUpdate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UserUpdate) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UserUpdate) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UserUpdate) SetPassword(v string) {
	o.Password = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserUpdate) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserUpdate) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserUpdate) SetEmail(v string) {
	o.Email = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UserUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UserUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UserUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *UserUpdate) GetOptions() UserOptions {
	if o == nil || o.Options == nil {
		var ret UserOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdate) GetOptionsOk() (*UserOptions, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *UserUpdate) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given UserOptions and assigns it to the Options field.
func (o *UserUpdate) SetOptions(v UserOptions) {
	o.Options = &v
}

func (o UserUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DefaultProjectId != nil {
		toSerialize["default_project_id"] = o.DefaultProjectId
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
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	return json.Marshal(toSerialize)
}

type NullableUserUpdate struct {
	value *UserUpdate
	isSet bool
}

func (v NullableUserUpdate) Get() *UserUpdate {
	return v.value
}

func (v *NullableUserUpdate) Set(val *UserUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdate(val *UserUpdate) *NullableUserUpdate {
	return &NullableUserUpdate{value: val, isSet: true}
}

func (v NullableUserUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
