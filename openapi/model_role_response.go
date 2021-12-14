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
)

// RoleResponse struct for RoleResponse
type RoleResponse struct {
	Role *Role `json:"role,omitempty"`
}

// NewRoleResponse instantiates a new RoleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResponse() *RoleResponse {
	this := RoleResponse{}
	return &this
}

// NewRoleResponseWithDefaults instantiates a new RoleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResponseWithDefaults() *RoleResponse {
	this := RoleResponse{}
	return &this
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RoleResponse) GetRole() Role {
	if o == nil || o.Role == nil {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponse) GetRoleOk() (*Role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RoleResponse) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *RoleResponse) SetRole(v Role) {
	o.Role = &v
}

func (o RoleResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	return json.Marshal(toSerialize)
}

type NullableRoleResponse struct {
	value *RoleResponse
	isSet bool
}

func (v NullableRoleResponse) Get() *RoleResponse {
	return v.value
}

func (v *NullableRoleResponse) Set(val *RoleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleResponse(val *RoleResponse) *NullableRoleResponse {
	return &NullableRoleResponse{value: val, isSet: true}
}

func (v NullableRoleResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}