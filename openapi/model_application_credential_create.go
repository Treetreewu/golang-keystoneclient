// Copyright 2017 EasyStack, Inc.

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

// ApplicationCredentialCreate struct for ApplicationCredentialCreate
type ApplicationCredentialCreate struct {
	// The name of the application credential. Must be unique to a user.
	Name string `json:"name"`
	// The secret that the application credential will be created with. If not provided, one will be generated.
	Secret *string `json:"secret,omitempty"`
	// A description of the application credential’s purpose.
	Description *string `json:"description,omitempty"`
	// An optional expiry time for the application credential. If unset, the application credential does not expire.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
	// An optional list of role objects, identified by ID or name. The list may only contain roles that the user has assigned on the project. If not provided, the roles assigned to the application credential will be the same as the roles in the current token.
	Roles *[]IdAndName `json:"roles,omitempty"`
	// An optional flag to restrict whether the application credential may be used for the creation or destruction of other application credentials or trusts. Defaults to false.
	Unrestricted *bool `json:"unrestricted,omitempty"`
	// A list of `access_rules` objects
	AccessRules *[]AccessRuleRequest `json:"access_rules,omitempty"`
}

// NewApplicationCredentialCreate instantiates a new ApplicationCredentialCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialCreate(name string) *ApplicationCredentialCreate {
	this := ApplicationCredentialCreate{}
	this.Name = name
	return &this
}

// NewApplicationCredentialCreateWithDefaults instantiates a new ApplicationCredentialCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialCreateWithDefaults() *ApplicationCredentialCreate {
	this := ApplicationCredentialCreate{}
	return &this
}

// GetName returns the Name field value
func (o *ApplicationCredentialCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApplicationCredentialCreate) SetName(v string) {
	o.Name = v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetSecret() string {
	if o == nil || o.Secret == nil {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetSecretOk() (*string, bool) {
	if o == nil || o.Secret == nil {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasSecret() bool {
	if o != nil && o.Secret != nil {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *ApplicationCredentialCreate) SetSecret(v string) {
	o.Secret = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationCredentialCreate) SetDescription(v string) {
	o.Description = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetExpiresAt() time.Time {
	if o == nil || o.ExpiresAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *ApplicationCredentialCreate) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetRoles() []IdAndName {
	if o == nil || o.Roles == nil {
		var ret []IdAndName
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetRolesOk() (*[]IdAndName, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []IdAndName and assigns it to the Roles field.
func (o *ApplicationCredentialCreate) SetRoles(v []IdAndName) {
	o.Roles = &v
}

// GetUnrestricted returns the Unrestricted field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetUnrestricted() bool {
	if o == nil || o.Unrestricted == nil {
		var ret bool
		return ret
	}
	return *o.Unrestricted
}

// GetUnrestrictedOk returns a tuple with the Unrestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetUnrestrictedOk() (*bool, bool) {
	if o == nil || o.Unrestricted == nil {
		return nil, false
	}
	return o.Unrestricted, true
}

// HasUnrestricted returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasUnrestricted() bool {
	if o != nil && o.Unrestricted != nil {
		return true
	}

	return false
}

// SetUnrestricted gets a reference to the given bool and assigns it to the Unrestricted field.
func (o *ApplicationCredentialCreate) SetUnrestricted(v bool) {
	o.Unrestricted = &v
}

// GetAccessRules returns the AccessRules field value if set, zero value otherwise.
func (o *ApplicationCredentialCreate) GetAccessRules() []AccessRuleRequest {
	if o == nil || o.AccessRules == nil {
		var ret []AccessRuleRequest
		return ret
	}
	return *o.AccessRules
}

// GetAccessRulesOk returns a tuple with the AccessRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialCreate) GetAccessRulesOk() (*[]AccessRuleRequest, bool) {
	if o == nil || o.AccessRules == nil {
		return nil, false
	}
	return o.AccessRules, true
}

// HasAccessRules returns a boolean if a field has been set.
func (o *ApplicationCredentialCreate) HasAccessRules() bool {
	if o != nil && o.AccessRules != nil {
		return true
	}

	return false
}

// SetAccessRules gets a reference to the given []AccessRuleRequest and assigns it to the AccessRules field.
func (o *ApplicationCredentialCreate) SetAccessRules(v []AccessRuleRequest) {
	o.AccessRules = &v
}

func (o ApplicationCredentialCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Secret != nil {
		toSerialize["secret"] = o.Secret
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if o.Unrestricted != nil {
		toSerialize["unrestricted"] = o.Unrestricted
	}
	if o.AccessRules != nil {
		toSerialize["access_rules"] = o.AccessRules
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationCredentialCreate struct {
	value *ApplicationCredentialCreate
	isSet bool
}

func (v NullableApplicationCredentialCreate) Get() *ApplicationCredentialCreate {
	return v.value
}

func (v *NullableApplicationCredentialCreate) Set(val *ApplicationCredentialCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentialCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentialCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentialCreate(val *ApplicationCredentialCreate) *NullableApplicationCredentialCreate {
	return &NullableApplicationCredentialCreate{value: val, isSet: true}
}

func (v NullableApplicationCredentialCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentialCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}