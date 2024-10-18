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

// checks if the SolverStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SolverStep{}

// SolverStep struct for SolverStep
type SolverStep struct {
	// Sub-steps, supplementing how this step is calculated to reach the next step
	Children []SolverStep `json:"children,omitempty"`
	// Description of the current step
	Description *SolverDescription `json:"description,omitempty"`
	// If the current step is drawing, this URL shows the drawing result
	ImageUrl *string `json:"image_url,omitempty"`
	// Latex display of the current step
	Latex *string `json:"latex,omitempty"`
}

// NewSolverStep instantiates a new SolverStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSolverStep() *SolverStep {
	this := SolverStep{}
	return &this
}

// NewSolverStepWithDefaults instantiates a new SolverStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSolverStepWithDefaults() *SolverStep {
	this := SolverStep{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *SolverStep) GetChildren() []SolverStep {
	if o == nil || IsNil(o.Children) {
		var ret []SolverStep
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverStep) GetChildrenOk() ([]SolverStep, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *SolverStep) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []SolverStep and assigns it to the Children field.
func (o *SolverStep) SetChildren(v []SolverStep) {
	o.Children = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SolverStep) GetDescription() SolverDescription {
	if o == nil || IsNil(o.Description) {
		var ret SolverDescription
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverStep) GetDescriptionOk() (*SolverDescription, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SolverStep) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given SolverDescription and assigns it to the Description field.
func (o *SolverStep) SetDescription(v SolverDescription) {
	o.Description = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *SolverStep) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverStep) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *SolverStep) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *SolverStep) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetLatex returns the Latex field value if set, zero value otherwise.
func (o *SolverStep) GetLatex() string {
	if o == nil || IsNil(o.Latex) {
		var ret string
		return ret
	}
	return *o.Latex
}

// GetLatexOk returns a tuple with the Latex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SolverStep) GetLatexOk() (*string, bool) {
	if o == nil || IsNil(o.Latex) {
		return nil, false
	}
	return o.Latex, true
}

// HasLatex returns a boolean if a field has been set.
func (o *SolverStep) HasLatex() bool {
	if o != nil && !IsNil(o.Latex) {
		return true
	}

	return false
}

// SetLatex gets a reference to the given string and assigns it to the Latex field.
func (o *SolverStep) SetLatex(v string) {
	o.Latex = &v
}

func (o SolverStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SolverStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.Latex) {
		toSerialize["latex"] = o.Latex
	}
	return toSerialize, nil
}

type NullableSolverStep struct {
	value *SolverStep
	isSet bool
}

func (v NullableSolverStep) Get() *SolverStep {
	return v.value
}

func (v *NullableSolverStep) Set(val *SolverStep) {
	v.value = val
	v.isSet = true
}

func (v NullableSolverStep) IsSet() bool {
	return v.isSet
}

func (v *NullableSolverStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSolverStep(val *SolverStep) *NullableSolverStep {
	return &NullableSolverStep{value: val, isSet: true}
}

func (v NullableSolverStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSolverStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

