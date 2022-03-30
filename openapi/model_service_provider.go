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

// ServiceProvider struct for ServiceProvider
type ServiceProvider struct {
	AuthUrl          *string   `json:"auth_url,omitempty"`
	Links            *SelfLink `json:"links,omitempty"`
	Description      *string   `json:"description,omitempty"`
	Enabled          *bool     `json:"enabled,omitempty"`
	Id               *string   `json:"id,omitempty"`
	RelayStatePrefix *string   `json:"relay_state_prefix,omitempty"`
}

// NewServiceProvider instantiates a new ServiceProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceProvider() *ServiceProvider {
	this := ServiceProvider{}
	var authUrl string = "null"
	this.AuthUrl = &authUrl
	return &this
}

// NewServiceProviderWithDefaults instantiates a new ServiceProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceProviderWithDefaults() *ServiceProvider {
	this := ServiceProvider{}
	var authUrl string = "null"
	this.AuthUrl = &authUrl
	return &this
}

// GetAuthUrl returns the AuthUrl field value if set, zero value otherwise.
func (o *ServiceProvider) GetAuthUrl() string {
	if o == nil || o.AuthUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthUrl
}

// GetAuthUrlOk returns a tuple with the AuthUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetAuthUrlOk() (*string, bool) {
	if o == nil || o.AuthUrl == nil {
		return nil, false
	}
	return o.AuthUrl, true
}

// HasAuthUrl returns a boolean if a field has been set.
func (o *ServiceProvider) HasAuthUrl() bool {
	if o != nil && o.AuthUrl != nil {
		return true
	}

	return false
}

// SetAuthUrl gets a reference to the given string and assigns it to the AuthUrl field.
func (o *ServiceProvider) SetAuthUrl(v string) {
	o.AuthUrl = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ServiceProvider) GetLinks() SelfLink {
	if o == nil || o.Links == nil {
		var ret SelfLink
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetLinksOk() (*SelfLink, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ServiceProvider) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given SelfLink and assigns it to the Links field.
func (o *ServiceProvider) SetLinks(v SelfLink) {
	o.Links = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceProvider) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceProvider) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceProvider) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ServiceProvider) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ServiceProvider) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ServiceProvider) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceProvider) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceProvider) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceProvider) SetId(v string) {
	o.Id = &v
}

// GetRelayStatePrefix returns the RelayStatePrefix field value if set, zero value otherwise.
func (o *ServiceProvider) GetRelayStatePrefix() string {
	if o == nil || o.RelayStatePrefix == nil {
		var ret string
		return ret
	}
	return *o.RelayStatePrefix
}

// GetRelayStatePrefixOk returns a tuple with the RelayStatePrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceProvider) GetRelayStatePrefixOk() (*string, bool) {
	if o == nil || o.RelayStatePrefix == nil {
		return nil, false
	}
	return o.RelayStatePrefix, true
}

// HasRelayStatePrefix returns a boolean if a field has been set.
func (o *ServiceProvider) HasRelayStatePrefix() bool {
	if o != nil && o.RelayStatePrefix != nil {
		return true
	}

	return false
}

// SetRelayStatePrefix gets a reference to the given string and assigns it to the RelayStatePrefix field.
func (o *ServiceProvider) SetRelayStatePrefix(v string) {
	o.RelayStatePrefix = &v
}

func (o ServiceProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthUrl != nil {
		toSerialize["auth_url"] = o.AuthUrl
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.RelayStatePrefix != nil {
		toSerialize["relay_state_prefix"] = o.RelayStatePrefix
	}
	return json.Marshal(toSerialize)
}

type NullableServiceProvider struct {
	value *ServiceProvider
	isSet bool
}

func (v NullableServiceProvider) Get() *ServiceProvider {
	return v.value
}

func (v *NullableServiceProvider) Set(val *ServiceProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceProvider(val *ServiceProvider) *NullableServiceProvider {
	return &NullableServiceProvider{value: val, isSet: true}
}

func (v NullableServiceProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
