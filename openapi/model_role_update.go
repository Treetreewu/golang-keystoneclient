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

// RoleUpdate struct for RoleUpdate
type RoleUpdate struct {
	// The new role name.
	Name *string `json:"name,omitempty"`
	// The new role description.
	Description *string `json:"description,omitempty"`
}

// NewRoleUpdate instantiates a new RoleUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleUpdate() *RoleUpdate {
	this := RoleUpdate{}
	return &this
}

// NewRoleUpdateWithDefaults instantiates a new RoleUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleUpdateWithDefaults() *RoleUpdate {
	this := RoleUpdate{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RoleUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RoleUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RoleUpdate) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RoleUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RoleUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RoleUpdate) SetDescription(v string) {
	o.Description = &v
}

func (o RoleUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableRoleUpdate struct {
	value *RoleUpdate
	isSet bool
}

func (v NullableRoleUpdate) Get() *RoleUpdate {
	return v.value
}

func (v *NullableRoleUpdate) Set(val *RoleUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleUpdate(val *RoleUpdate) *NullableRoleUpdate {
	return &NullableRoleUpdate{value: val, isSet: true}
}

func (v NullableRoleUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
