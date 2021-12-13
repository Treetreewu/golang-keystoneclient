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

// ApplicationCredentialRole struct for ApplicationCredentialRole
type ApplicationCredentialRole struct {
	Id       *string `json:"id,omitempty"`
	Name     *string `json:"name,omitempty"`
	DomainId *string `json:"domain_id,omitempty"`
}

// NewApplicationCredentialRole instantiates a new ApplicationCredentialRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCredentialRole() *ApplicationCredentialRole {
	this := ApplicationCredentialRole{}
	return &this
}

// NewApplicationCredentialRoleWithDefaults instantiates a new ApplicationCredentialRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCredentialRoleWithDefaults() *ApplicationCredentialRole {
	this := ApplicationCredentialRole{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationCredentialRole) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialRole) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationCredentialRole) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationCredentialRole) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationCredentialRole) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialRole) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationCredentialRole) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationCredentialRole) SetName(v string) {
	o.Name = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *ApplicationCredentialRole) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCredentialRole) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *ApplicationCredentialRole) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *ApplicationCredentialRole) SetDomainId(v string) {
	o.DomainId = &v
}

func (o ApplicationCredentialRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationCredentialRole struct {
	value *ApplicationCredentialRole
	isSet bool
}

func (v NullableApplicationCredentialRole) Get() *ApplicationCredentialRole {
	return v.value
}

func (v *NullableApplicationCredentialRole) Set(val *ApplicationCredentialRole) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCredentialRole) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCredentialRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCredentialRole(val *ApplicationCredentialRole) *NullableApplicationCredentialRole {
	return &NullableApplicationCredentialRole{value: val, isSet: true}
}

func (v NullableApplicationCredentialRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCredentialRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
