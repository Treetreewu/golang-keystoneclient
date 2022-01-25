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

// DomainConfigConfig struct for DomainConfigConfig
type DomainConfigConfig struct {
	Identity *DomainConfigIdentity `json:"identity,omitempty"`
	Ldap     *DomainConfigLDAP     `json:"ldap,omitempty"`
}

// NewDomainConfigConfig instantiates a new DomainConfigConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDomainConfigConfig() *DomainConfigConfig {
	this := DomainConfigConfig{}
	return &this
}

// NewDomainConfigConfigWithDefaults instantiates a new DomainConfigConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDomainConfigConfigWithDefaults() *DomainConfigConfig {
	this := DomainConfigConfig{}
	return &this
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *DomainConfigConfig) GetIdentity() DomainConfigIdentity {
	if o == nil || o.Identity == nil {
		var ret DomainConfigIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainConfigConfig) GetIdentityOk() (*DomainConfigIdentity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *DomainConfigConfig) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given DomainConfigIdentity and assigns it to the Identity field.
func (o *DomainConfigConfig) SetIdentity(v DomainConfigIdentity) {
	o.Identity = &v
}

// GetLdap returns the Ldap field value if set, zero value otherwise.
func (o *DomainConfigConfig) GetLdap() DomainConfigLDAP {
	if o == nil || o.Ldap == nil {
		var ret DomainConfigLDAP
		return ret
	}
	return *o.Ldap
}

// GetLdapOk returns a tuple with the Ldap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DomainConfigConfig) GetLdapOk() (*DomainConfigLDAP, bool) {
	if o == nil || o.Ldap == nil {
		return nil, false
	}
	return o.Ldap, true
}

// HasLdap returns a boolean if a field has been set.
func (o *DomainConfigConfig) HasLdap() bool {
	if o != nil && o.Ldap != nil {
		return true
	}

	return false
}

// SetLdap gets a reference to the given DomainConfigLDAP and assigns it to the Ldap field.
func (o *DomainConfigConfig) SetLdap(v DomainConfigLDAP) {
	o.Ldap = &v
}

func (o DomainConfigConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Identity != nil {
		toSerialize["identity"] = o.Identity
	}
	if o.Ldap != nil {
		toSerialize["ldap"] = o.Ldap
	}
	return json.Marshal(toSerialize)
}

type NullableDomainConfigConfig struct {
	value *DomainConfigConfig
	isSet bool
}

func (v NullableDomainConfigConfig) Get() *DomainConfigConfig {
	return v.value
}

func (v *NullableDomainConfigConfig) Set(val *DomainConfigConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableDomainConfigConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableDomainConfigConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDomainConfigConfig(val *DomainConfigConfig) *NullableDomainConfigConfig {
	return &NullableDomainConfigConfig{value: val, isSet: true}
}

func (v NullableDomainConfigConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDomainConfigConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}