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

// ProjectUpdateRequest struct for ProjectUpdateRequest
type ProjectUpdateRequest struct {
	Project *ProjectUpdate `json:"project,omitempty"`
}

// NewProjectUpdateRequest instantiates a new ProjectUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUpdateRequest() *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	return &this
}

// NewProjectUpdateRequestWithDefaults instantiates a new ProjectUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUpdateRequestWithDefaults() *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ProjectUpdateRequest) GetDomain() ProjectUpdate {
	if o == nil || o.Project == nil {
		var ret ProjectUpdate
		return ret
	}
	return *o.Project
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectUpdateRequest) GetDomainOk() (*ProjectUpdate, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ProjectUpdateRequest) HasDomain() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetDomain gets a reference to the given ProjectUpdate and assigns it to the Domain field.
func (o *ProjectUpdateRequest) SetDomain(v ProjectUpdate) {
	o.Project = &v
}

func (o ProjectUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["domain"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProjectUpdateRequest struct {
	value *ProjectUpdateRequest
	isSet bool
}

func (v NullableProjectUpdateRequest) Get() *ProjectUpdateRequest {
	return v.value
}

func (v *NullableProjectUpdateRequest) Set(val *ProjectUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectUpdateRequest(val *ProjectUpdateRequest) *NullableProjectUpdateRequest {
	return &NullableProjectUpdateRequest{value: val, isSet: true}
}

func (v NullableProjectUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
