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

// RoleAssignmentScope Only one type of scope is valid, others will be null.
type RoleAssignmentScope struct {
	Domain  *RoleAssignmentListResponseRole `json:"domain,omitempty"`
	Project *RoleAssignmentListResponseRole `json:"project,omitempty"`
}

// NewRoleAssignmentScope instantiates a new RoleAssignmentScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleAssignmentScope() *RoleAssignmentScope {
	this := RoleAssignmentScope{}
	return &this
}

// NewRoleAssignmentScopeWithDefaults instantiates a new RoleAssignmentScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleAssignmentScopeWithDefaults() *RoleAssignmentScope {
	this := RoleAssignmentScope{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *RoleAssignmentScope) GetDomain() RoleAssignmentListResponseRole {
	if o == nil || o.Domain == nil {
		var ret RoleAssignmentListResponseRole
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentScope) GetDomainOk() (*RoleAssignmentListResponseRole, bool) {
	if o == nil || o.Domain == nil {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *RoleAssignmentScope) HasDomain() bool {
	if o != nil && o.Domain != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given RoleAssignmentListResponseRole and assigns it to the Domain field.
func (o *RoleAssignmentScope) SetDomain(v RoleAssignmentListResponseRole) {
	o.Domain = &v
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *RoleAssignmentScope) GetProject() RoleAssignmentListResponseRole {
	if o == nil || o.Project == nil {
		var ret RoleAssignmentListResponseRole
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleAssignmentScope) GetProjectOk() (*RoleAssignmentListResponseRole, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *RoleAssignmentScope) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given RoleAssignmentListResponseRole and assigns it to the Project field.
func (o *RoleAssignmentScope) SetProject(v RoleAssignmentListResponseRole) {
	o.Project = &v
}

func (o RoleAssignmentScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domain != nil {
		toSerialize["domain"] = o.Domain
	}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableRoleAssignmentScope struct {
	value *RoleAssignmentScope
	isSet bool
}

func (v NullableRoleAssignmentScope) Get() *RoleAssignmentScope {
	return v.value
}

func (v *NullableRoleAssignmentScope) Set(val *RoleAssignmentScope) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleAssignmentScope) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleAssignmentScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleAssignmentScope(val *RoleAssignmentScope) *NullableRoleAssignmentScope {
	return &NullableRoleAssignmentScope{value: val, isSet: true}
}

func (v NullableRoleAssignmentScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleAssignmentScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
