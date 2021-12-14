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

// DomainUpdateRequest struct for DomainUpdateRequest
type DomainUpdateRequest struct {
	Domain DomainUpdate `json:"domain"`
}

// NewDomainUpdateRequest instantiates a new DomainUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainUpdateRequest(domain DomainUpdate) *DomainUpdateRequest {
	this := DomainUpdateRequest{}
	this.Domain = domain
	return &this
}

// NewDomainUpdateRequestWithDefaults instantiates a new DomainUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainUpdateRequestWithDefaults() *DomainUpdateRequest {
	this := DomainUpdateRequest{}
	return &this
}

// GetDomain returns the Domain field value
func (o *DomainUpdateRequest) GetDomain() DomainUpdate {
	if o == nil {
		var ret DomainUpdate
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *DomainUpdateRequest) GetDomainOk() (*DomainUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *DomainUpdateRequest) SetDomain(v DomainUpdate) {
	o.Domain = v
}

func (o DomainUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	return json.Marshal(toSerialize)
}

type NullableDomainUpdateRequest struct {
	value *DomainUpdateRequest
	isSet bool
}

func (v NullableDomainUpdateRequest) Get() *DomainUpdateRequest {
	return v.value
}

func (v *NullableDomainUpdateRequest) Set(val *DomainUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainUpdateRequest(val *DomainUpdateRequest) *NullableDomainUpdateRequest {
	return &NullableDomainUpdateRequest{value: val, isSet: true}
}

func (v NullableDomainUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}