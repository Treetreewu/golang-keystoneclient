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

// DomainListResponse struct for DomainListResponse
type DomainListResponse struct {
	Domains *[]Domain `json:"domains,omitempty"`
	Links   *Links    `json:"links,omitempty"`
}

// NewDomainListResponse instantiates a new DomainListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainListResponse() *DomainListResponse {
	this := DomainListResponse{}
	return &this
}

// NewDomainListResponseWithDefaults instantiates a new DomainListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainListResponseWithDefaults() *DomainListResponse {
	this := DomainListResponse{}
	return &this
}

// GetDomains returns the Domains field value if set, zero value otherwise.
func (o *DomainListResponse) GetDomains() []Domain {
	if o == nil || o.Domains == nil {
		var ret []Domain
		return ret
	}
	return *o.Domains
}

// GetDomainsOk returns a tuple with the Domains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainListResponse) GetDomainsOk() (*[]Domain, bool) {
	if o == nil || o.Domains == nil {
		return nil, false
	}
	return o.Domains, true
}

// HasDomains returns a boolean if a field has been set.
func (o *DomainListResponse) HasDomains() bool {
	if o != nil && o.Domains != nil {
		return true
	}

	return false
}

// SetDomains gets a reference to the given []Domain and assigns it to the Domains field.
func (o *DomainListResponse) SetDomains(v []Domain) {
	o.Domains = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DomainListResponse) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainListResponse) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DomainListResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *DomainListResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o DomainListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Domains != nil {
		toSerialize["domains"] = o.Domains
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableDomainListResponse struct {
	value *DomainListResponse
	isSet bool
}

func (v NullableDomainListResponse) Get() *DomainListResponse {
	return v.value
}

func (v *NullableDomainListResponse) Set(val *DomainListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainListResponse(val *DomainListResponse) *NullableDomainListResponse {
	return &NullableDomainListResponse{value: val, isSet: true}
}

func (v NullableDomainListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
