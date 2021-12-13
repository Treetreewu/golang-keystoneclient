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

// UserListResponse struct for UserListResponse
type UserListResponse struct {
	Links *Links  `json:"links,omitempty"`
	Users *[]User `json:"users,omitempty"`
}

// NewUserListResponse instantiates a new UserListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserListResponse() *UserListResponse {
	this := UserListResponse{}
	return &this
}

// NewUserListResponseWithDefaults instantiates a new UserListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserListResponseWithDefaults() *UserListResponse {
	this := UserListResponse{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserListResponse) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserListResponse) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserListResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *UserListResponse) SetLinks(v Links) {
	o.Links = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *UserListResponse) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserListResponse) GetUsersOk() (*[]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *UserListResponse) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *UserListResponse) SetUsers(v []User) {
	o.Users = &v
}

func (o UserListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableUserListResponse struct {
	value *UserListResponse
	isSet bool
}

func (v NullableUserListResponse) Get() *UserListResponse {
	return v.value
}

func (v *NullableUserListResponse) Set(val *UserListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserListResponse(val *UserListResponse) *NullableUserListResponse {
	return &NullableUserListResponse{value: val, isSet: true}
}

func (v NullableUserListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
