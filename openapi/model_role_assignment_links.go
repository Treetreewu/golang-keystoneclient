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

// RoleAssignmentLinks struct for RoleAssignmentLinks
type RoleAssignmentLinks struct {
	Assignment NullableString `json:"assignment,omitempty"`
	Membership NullableString `json:"membership,omitempty"`
}

// NewRoleAssignmentLinks instantiates a new RoleAssignmentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentLinks() *RoleAssignmentLinks {
	this := RoleAssignmentLinks{}
	return &this
}

// NewRoleAssignmentLinksWithDefaults instantiates a new RoleAssignmentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentLinksWithDefaults() *RoleAssignmentLinks {
	this := RoleAssignmentLinks{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAssignmentLinks) GetAssignment() string {
	if o == nil || o.Assignment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Assignment.Get()
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAssignmentLinks) GetAssignmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Assignment.Get(), o.Assignment.IsSet()
}

// HasAssignment returns a boolean if a field has been set.
func (o *RoleAssignmentLinks) HasAssignment() bool {
	if o != nil && o.Assignment.IsSet() {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given NullableString and assigns it to the Assignment field.
func (o *RoleAssignmentLinks) SetAssignment(v string) {
	o.Assignment.Set(&v)
}

// SetAssignmentNil sets the value for Assignment to be an explicit nil
func (o *RoleAssignmentLinks) SetAssignmentNil() {
	o.Assignment.Set(nil)
}

// UnsetAssignment ensures that no value is present for Assignment, not even an explicit nil
func (o *RoleAssignmentLinks) UnsetAssignment() {
	o.Assignment.Unset()
}

// GetMembership returns the Membership field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RoleAssignmentLinks) GetMembership() string {
	if o == nil || o.Membership.Get() == nil {
		var ret string
		return ret
	}
	return *o.Membership.Get()
}

// GetMembershipOk returns a tuple with the Membership field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RoleAssignmentLinks) GetMembershipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Membership.Get(), o.Membership.IsSet()
}

// HasMembership returns a boolean if a field has been set.
func (o *RoleAssignmentLinks) HasMembership() bool {
	if o != nil && o.Membership.IsSet() {
		return true
	}

	return false
}

// SetMembership gets a reference to the given NullableString and assigns it to the Membership field.
func (o *RoleAssignmentLinks) SetMembership(v string) {
	o.Membership.Set(&v)
}

// SetMembershipNil sets the value for Membership to be an explicit nil
func (o *RoleAssignmentLinks) SetMembershipNil() {
	o.Membership.Set(nil)
}

// UnsetMembership ensures that no value is present for Membership, not even an explicit nil
func (o *RoleAssignmentLinks) UnsetMembership() {
	o.Membership.Unset()
}

func (o RoleAssignmentLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assignment.IsSet() {
		toSerialize["assignment"] = o.Assignment.Get()
	}
	if o.Membership.IsSet() {
		toSerialize["membership"] = o.Membership.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignmentLinks struct {
	value *RoleAssignmentLinks
	isSet bool
}

func (v NullableRoleAssignmentLinks) Get() *RoleAssignmentLinks {
	return v.value
}

func (v *NullableRoleAssignmentLinks) Set(val *RoleAssignmentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentLinks(val *RoleAssignmentLinks) *NullableRoleAssignmentLinks {
	return &NullableRoleAssignmentLinks{value: val, isSet: true}
}

func (v NullableRoleAssignmentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}