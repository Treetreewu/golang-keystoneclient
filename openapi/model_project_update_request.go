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
	Project ProjectUpdate `json:"project"`
}

// NewProjectUpdateRequest instantiates a new ProjectUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectUpdateRequest(project ProjectUpdate) *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	this.Project = project
	return &this
}

// NewProjectUpdateRequestWithDefaults instantiates a new ProjectUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectUpdateRequestWithDefaults() *ProjectUpdateRequest {
	this := ProjectUpdateRequest{}
	return &this
}

// GetProject returns the Project field value
func (o *ProjectUpdateRequest) GetProject() ProjectUpdate {
	if o == nil {
		var ret ProjectUpdate
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *ProjectUpdateRequest) GetProjectOk() (*ProjectUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *ProjectUpdateRequest) SetProject(v ProjectUpdate) {
	o.Project = v
}

func (o ProjectUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["project"] = o.Project
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
