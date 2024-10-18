# SolverDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Descriptive information | [optional] 
**Format** | Pointer to [**SolverDescriptionFormat**](SolverDescriptionFormat.md) |  | [optional] 

## Methods

### NewSolverDescription

`func NewSolverDescription() *SolverDescription`

NewSolverDescription instantiates a new SolverDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverDescriptionWithDefaults

`func NewSolverDescriptionWithDefaults() *SolverDescription`

NewSolverDescriptionWithDefaults instantiates a new SolverDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SolverDescription) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SolverDescription) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SolverDescription) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SolverDescription) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetFormat

`func (o *SolverDescription) GetFormat() SolverDescriptionFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *SolverDescription) GetFormatOk() (*SolverDescriptionFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *SolverDescription) SetFormat(v SolverDescriptionFormat)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *SolverDescription) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


