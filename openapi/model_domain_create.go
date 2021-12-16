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

// DomainCreate struct for DomainCreate
type DomainCreate struct {
	// The description of the domain.
	Description *string `json:"description,omitempty"`
	// If set to true, domain is created enabled. If set to false, domain is created disabled. The default is true. Users can only authorize against an enabled domain (and any of its projects). In addition, users can only authenticate if the domain that owns them is also enabled. Disabling a domain prevents both of these things.
	Enabled *bool `json:"enabled,omitempty"`
	// The ID of the domain. A domain created this way will not use an auto-generated ID, but will use the ID passed in instead. Identifiers passed in this way must conform to the existing ID generation scheme: UUID4 without dashes.
	ExplicitDomainId *string `json:"explicit_domain_id,omitempty"`
	// The name of the domain.
	Name string `json:"name"`
}

// NewDomainCreate instantiates a new DomainCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainCreate(name string) *DomainCreate {
	this := DomainCreate{}
	this.Name = name
	return &this
}

// NewDomainCreateWithDefaults instantiates a new DomainCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainCreateWithDefaults() *DomainCreate {
	this := DomainCreate{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DomainCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DomainCreate) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DomainCreate) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *DomainCreate) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCreate) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *DomainCreate) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *DomainCreate) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExplicitDomainId returns the ExplicitDomainId field value if set, zero value otherwise.
func (o *DomainCreate) GetExplicitDomainId() string {
	if o == nil || o.ExplicitDomainId == nil {
		var ret string
		return ret
	}
	return *o.ExplicitDomainId
}

// GetExplicitDomainIdOk returns a tuple with the ExplicitDomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainCreate) GetExplicitDomainIdOk() (*string, bool) {
	if o == nil || o.ExplicitDomainId == nil {
		return nil, false
	}
	return o.ExplicitDomainId, true
}

// HasExplicitDomainId returns a boolean if a field has been set.
func (o *DomainCreate) HasExplicitDomainId() bool {
	if o != nil && o.ExplicitDomainId != nil {
		return true
	}

	return false
}

// SetExplicitDomainId gets a reference to the given string and assigns it to the ExplicitDomainId field.
func (o *DomainCreate) SetExplicitDomainId(v string) {
	o.ExplicitDomainId = &v
}

// GetName returns the Name field value
func (o *DomainCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DomainCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DomainCreate) SetName(v string) {
	o.Name = v
}

func (o DomainCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.ExplicitDomainId != nil {
		toSerialize["explicit_domain_id"] = o.ExplicitDomainId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableDomainCreate struct {
	value *DomainCreate
	isSet bool
}

func (v NullableDomainCreate) Get() *DomainCreate {
	return v.value
}

func (v *NullableDomainCreate) Set(val *DomainCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainCreate(val *DomainCreate) *NullableDomainCreate {
	return &NullableDomainCreate{value: val, isSet: true}
}

func (v NullableDomainCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
