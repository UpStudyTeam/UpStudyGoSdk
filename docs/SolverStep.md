# SolverStep

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]SolverStep**](SolverStep.md) | Sub-steps, supplementing how this step is calculated to reach the next step | [optional] 
**Description** | Pointer to [**SolverDescription**](SolverDescription.md) | Description of the current step | [optional] 
**ImageUrl** | Pointer to **string** | If the current step is drawing, this URL shows the drawing result | [optional] 
**Latex** | Pointer to **string** | Latex display of the current step | [optional] 

## Methods

### NewSolverStep

`func NewSolverStep() *SolverStep`

NewSolverStep instantiates a new SolverStep object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverStepWithDefaults

`func NewSolverStepWithDefaults() *SolverStep`

NewSolverStepWithDefaults instantiates a new SolverStep object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *SolverStep) GetChildren() []SolverStep`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *SolverStep) GetChildrenOk() (*[]SolverStep, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *SolverStep) SetChildren(v []SolverStep)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *SolverStep) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetDescription

`func (o *SolverStep) GetDescription() SolverDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SolverStep) GetDescriptionOk() (*SolverDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SolverStep) SetDescription(v SolverDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SolverStep) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *SolverStep) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SolverStep) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SolverStep) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SolverStep) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetLatex

`func (o *SolverStep) GetLatex() string`

GetLatex returns the Latex field if non-nil, zero value otherwise.

### GetLatexOk

`func (o *SolverStep) GetLatexOk() (*string, bool)`

GetLatexOk returns a tuple with the Latex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatex

`func (o *SolverStep) SetLatex(v string)`

SetLatex sets Latex field to given value.

### HasLatex

`func (o *SolverStep) HasLatex() bool`

HasLatex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


