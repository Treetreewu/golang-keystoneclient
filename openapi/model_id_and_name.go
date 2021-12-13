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

// IdAndName struct for IdAndName
type IdAndName struct {
	Id   *string        `json:"id,omitempty"`
	Name NullableString `json:"name,omitempty"`
}

// NewIdAndName instantiates a new IdAndName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdAndName() *IdAndName {
	this := IdAndName{}
	return &this
}

// NewIdAndNameWithDefaults instantiates a new IdAndName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdAndNameWithDefaults() *IdAndName {
	this := IdAndName{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdAndName) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdAndName) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdAndName) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdAndName) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdAndName) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdAndName) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *IdAndName) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *IdAndName) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *IdAndName) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *IdAndName) UnsetName() {
	o.Name.Unset()
}

func (o IdAndName) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableIdAndName struct {
	value *IdAndName
	isSet bool
}

func (v NullableIdAndName) Get() *IdAndName {
	return v.value
}

func (v *NullableIdAndName) Set(val *IdAndName) {
	v.value = val
	v.isSet = true
}

func (v NullableIdAndName) IsSet() bool {
	return v.isSet
}

func (v *NullableIdAndName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdAndName(val *IdAndName) *NullableIdAndName {
	return &NullableIdAndName{value: val, isSet: true}
}

func (v NullableIdAndName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdAndName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
