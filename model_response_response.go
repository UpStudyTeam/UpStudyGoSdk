/*
api

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package UpStudyGoSdk

import (
	"encoding/json"
)

// checks if the ResponseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResponseResponse{}

// ResponseResponse struct for ResponseResponse
type ResponseResponse struct {
	Data map[string]interface{} `json:"data,omitempty"`
	ErrMsg *string `json:"err_msg,omitempty"`
}

// NewResponseResponse instantiates a new ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResponseResponse() *ResponseResponse {
	this := ResponseResponse{}
	return &this
}

// NewResponseResponseWithDefaults instantiates a new ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResponseResponseWithDefaults() *ResponseResponse {
	this := ResponseResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ResponseResponse) GetData() map[string]interface{} {
	if o == nil || IsNil(o.Data) {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Data) {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ResponseResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *ResponseResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetErrMsg returns the ErrMsg field value if set, zero value otherwise.
func (o *ResponseResponse) GetErrMsg() string {
	if o == nil || IsNil(o.ErrMsg) {
		var ret string
		return ret
	}
	return *o.ErrMsg
}

// GetErrMsgOk returns a tuple with the ErrMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResponseResponse) GetErrMsgOk() (*string, bool) {
	if o == nil || IsNil(o.ErrMsg) {
		return nil, false
	}
	return o.ErrMsg, true
}

// HasErrMsg returns a boolean if a field has been set.
func (o *ResponseResponse) HasErrMsg() bool {
	if o != nil && !IsNil(o.ErrMsg) {
		return true
	}

	return false
}

// SetErrMsg gets a reference to the given string and assigns it to the ErrMsg field.
func (o *ResponseResponse) SetErrMsg(v string) {
	o.ErrMsg = &v
}

func (o ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.ErrMsg) {
		toSerialize["err_msg"] = o.ErrMsg
	}
	return toSerialize, nil
}

type NullableResponseResponse struct {
	value *ResponseResponse
	isSet bool
}

func (v NullableResponseResponse) Get() *ResponseResponse {
	return v.value
}

func (v *NullableResponseResponse) Set(val *ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResponseResponse(val *ResponseResponse) *NullableResponseResponse {
	return &NullableResponseResponse{value: val, isSet: true}
}

func (v NullableResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


