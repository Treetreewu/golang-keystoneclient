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

// ServiceProviderListResponse struct for ServiceProviderListResponse
type ServiceProviderListResponse struct {
	ServiceProviders *[]ServiceProvider `json:"service_providers,omitempty"`
	Links            *Links             `json:"links,omitempty"`
}

// NewServiceProviderListResponse instantiates a new ServiceProviderListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProviderListResponse() *ServiceProviderListResponse {
	this := ServiceProviderListResponse{}
	return &this
}

// NewServiceProviderListResponseWithDefaults instantiates a new ServiceProviderListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProviderListResponseWithDefaults() *ServiceProviderListResponse {
	this := ServiceProviderListResponse{}
	return &this
}

// GetServiceProviders returns the ServiceProviders field value if set, zero value otherwise.
func (o *ServiceProviderListResponse) GetServiceProviders() []ServiceProvider {
	if o == nil || o.ServiceProviders == nil {
		var ret []ServiceProvider
		return ret
	}
	return *o.ServiceProviders
}

// GetServiceProvidersOk returns a tuple with the ServiceProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderListResponse) GetServiceProvidersOk() (*[]ServiceProvider, bool) {
	if o == nil || o.ServiceProviders == nil {
		return nil, false
	}
	return o.ServiceProviders, true
}

// HasServiceProviders returns a boolean if a field has been set.
func (o *ServiceProviderListResponse) HasServiceProviders() bool {
	if o != nil && o.ServiceProviders != nil {
		return true
	}

	return false
}

// SetServiceProviders gets a reference to the given []ServiceProvider and assigns it to the ServiceProviders field.
func (o *ServiceProviderListResponse) SetServiceProviders(v []ServiceProvider) {
	o.ServiceProviders = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceProviderListResponse) GetLinks() Links {
	if o == nil || o.Links == nil {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProviderListResponse) GetLinksOk() (*Links, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceProviderListResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ServiceProviderListResponse) SetLinks(v Links) {
	o.Links = &v
}

func (o ServiceProviderListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceProviders != nil {
		toSerialize["service_providers"] = o.ServiceProviders
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	return json.Marshal(toSerialize)
}

type NullableServiceProviderListResponse struct {
	value *ServiceProviderListResponse
	isSet bool
}

func (v NullableServiceProviderListResponse) Get() *ServiceProviderListResponse {
	return v.value
}

func (v *NullableServiceProviderListResponse) Set(val *ServiceProviderListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProviderListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProviderListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProviderListResponse(val *ServiceProviderListResponse) *NullableServiceProviderListResponse {
	return &NullableServiceProviderListResponse{value: val, isSet: true}
}

func (v NullableServiceProviderListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProviderListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}