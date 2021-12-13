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

// GroupUpdate struct for GroupUpdate
type GroupUpdate struct {
	// The new description for the group.
	Description *string `json:"description,omitempty"`
	// The ID of the new domain for the group. The ability to change the domain of a group is now deprecated, and will be removed in subsequent release. It is already disabled by default in most Identity service implementations.
	// Deprecated
	DomainId *string `json:"domain_id,omitempty"`
	// The new name for the group.
	Name *string `json:"name,omitempty"`
}

// NewGroupUpdate instantiates a new GroupUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdate() *GroupUpdate {
	this := GroupUpdate{}
	return &this
}

// NewGroupUpdateWithDefaults instantiates a new GroupUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateWithDefaults() *GroupUpdate {
	this := GroupUpdate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupUpdate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupUpdate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupUpdate) SetDescription(v string) {
	o.Description = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
// Deprecated
func (o *GroupUpdate) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *GroupUpdate) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *GroupUpdate) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
// Deprecated
func (o *GroupUpdate) SetDomainId(v string) {
	o.DomainId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupUpdate) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdate) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupUpdate) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupUpdate) SetName(v string) {
	o.Name = &v
}

func (o GroupUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUpdate struct {
	value *GroupUpdate
	isSet bool
}

func (v NullableGroupUpdate) Get() *GroupUpdate {
	return v.value
}

func (v *NullableGroupUpdate) Set(val *GroupUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdate(val *GroupUpdate) *NullableGroupUpdate {
	return &NullableGroupUpdate{value: val, isSet: true}
}

func (v NullableGroupUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
