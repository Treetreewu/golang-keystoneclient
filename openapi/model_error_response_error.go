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

// ErrorResponseError struct for ErrorResponseError
type ErrorResponseError struct {
	Code    *int32  `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Title   *string `json:"title,omitempty"`
}

// NewErrorResponseError instantiates a new ErrorResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseError() *ErrorResponseError {
	this := ErrorResponseError{}
	return &this
}

// NewErrorResponseErrorWithDefaults instantiates a new ErrorResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorWithDefaults() *ErrorResponseError {
	this := ErrorResponseError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ErrorResponseError) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ErrorResponseError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ErrorResponseError) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ErrorResponseError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorResponseError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ErrorResponseError) SetMessage(v string) {
	o.Message = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ErrorResponseError) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorResponseError) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ErrorResponseError) SetTitle(v string) {
	o.Title = &v
}

func (o ErrorResponseError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableErrorResponseError struct {
	value *ErrorResponseError
	isSet bool
}

func (v NullableErrorResponseError) Get() *ErrorResponseError {
	return v.value
}

func (v *NullableErrorResponseError) Set(val *ErrorResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseError(val *ErrorResponseError) *NullableErrorResponseError {
	return &NullableErrorResponseError{value: val, isSet: true}
}

func (v NullableErrorResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}