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

// ProjectGetResponse struct for ProjectGetResponse
type ProjectGetResponse struct {
	Project *ProjectGet `json:"project,omitempty"`
}

// NewProjectGetResponse instantiates a new ProjectGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectGetResponse() *ProjectGetResponse {
	this := ProjectGetResponse{}
	return &this
}

// NewProjectGetResponseWithDefaults instantiates a new ProjectGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectGetResponseWithDefaults() *ProjectGetResponse {
	this := ProjectGetResponse{}
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *ProjectGetResponse) GetProject() ProjectGet {
	if o == nil || o.Project == nil {
		var ret ProjectGet
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectGetResponse) GetProjectOk() (*ProjectGet, bool) {
	if o == nil || o.Project == nil {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *ProjectGetResponse) HasProject() bool {
	if o != nil && o.Project != nil {
		return true
	}

	return false
}

// SetProject gets a reference to the given ProjectGet and assigns it to the Project field.
func (o *ProjectGetResponse) SetProject(v ProjectGet) {
	o.Project = &v
}

func (o ProjectGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Project != nil {
		toSerialize["project"] = o.Project
	}
	return json.Marshal(toSerialize)
}

type NullableProjectGetResponse struct {
	value *ProjectGetResponse
	isSet bool
}

func (v NullableProjectGetResponse) Get() *ProjectGetResponse {
	return v.value
}

func (v *NullableProjectGetResponse) Set(val *ProjectGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectGetResponse(val *ProjectGetResponse) *NullableProjectGetResponse {
	return &NullableProjectGetResponse{value: val, isSet: true}
}

func (v NullableProjectGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
