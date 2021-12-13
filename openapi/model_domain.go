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

// Domain struct for Domain
type Domain struct {
	// The description of the domain.
	Description *string `json:"description,omitempty"`
	// If set to true, domain is enabled. If set to false, domain is disabled.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the domain.
	Id    *string   `json:"id,omitempty"`
	Links *SelfLink `json:"links,omitempty"`
	// The name of the domain.
	Name *string `json:"name,omitempty"`
	// A list of parent ID.
	Parents *[]string `json:"parents,omitempty"`
}

// NewDomain instantiates a new Domain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomain() *Domain {
	this := Domain{}
	return &this
}

// NewDomainWithDefaults instantiates a new Domain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainWithDefaults() *Domain {
	this := Domain{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Domain) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Domain) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Domain) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Domain) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Domain) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Domain) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Domain) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Domain) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Domain) SetId(v string) {
	o.Id = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Domain) GetLinks() SelfLink {
	if o == nil || o.Links == nil {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetLinksOk() (*SelfLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Domain) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *Domain) SetLinks(v SelfLink) {
	o.Links = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Domain) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Domain) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Domain) SetName(v string) {
	o.Name = &v
}

// GetParents returns the Parents field value if set, zero value otherwise.
func (o *Domain) GetParents() []string {
	if o == nil || o.Parents == nil {
		var ret []string
		return ret
	}
	return *o.Parents
}

// GetParentsOk returns a tuple with the Parents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Domain) GetParentsOk() (*[]string, bool) {
	if o == nil || o.Parents == nil {
		return nil, false
	}
	return o.Parents, true
}

// HasParents returns a boolean if a field has been set.
func (o *Domain) HasParents() bool {
	if o != nil && o.Parents != nil {
		return true
	}

	return false
}

// SetParents gets a reference to the given []string and assigns it to the Parents field.
func (o *Domain) SetParents(v []string) {
	o.Parents = &v
}

func (o Domain) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parents != nil {
		toSerialize["parents"] = o.Parents
	}
	return json.Marshal(toSerialize)
}

type NullableDomain struct {
	value *Domain
	isSet bool
}

func (v NullableDomain) Get() *Domain {
	return v.value
}

func (v *NullableDomain) Set(val *Domain) {
	v.value = val
	v.isSet = true
}

func (v NullableDomain) IsSet() bool {
	return v.isSet
}

func (v *NullableDomain) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomain(val *Domain) *NullableDomain {
	return &NullableDomain{value: val, isSet: true}
}

func (v NullableDomain) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomain) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
