# SolverSelfFullResponseText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | User&#39;s Latex input | [optional] 
**Solutions** | Pointer to [**[]SolverSolutionWithSolvingSteps**](SolverSolutionWithSolvingSteps.md) | All solutions we can provide along with the steps to obtain the solutions | [optional] 

## Methods

### NewSolverSelfFullResponseText

`func NewSolverSelfFullResponseText() *SolverSelfFullResponseText`

NewSolverSelfFullResponseText instantiates a new SolverSelfFullResponseText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverSelfFullResponseTextWithDefaults

`func NewSolverSelfFullResponseTextWithDefaults() *SolverSelfFullResponseText`

NewSolverSelfFullResponseTextWithDefaults instantiates a new SolverSelfFullResponseText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *SolverSelfFullResponseText) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *SolverSelfFullResponseText) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *SolverSelfFullResponseText) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *SolverSelfFullResponseText) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetSolutions

`func (o *SolverSelfFullResponseText) GetSolutions() []SolverSolutionWithSolvingSteps`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *SolverSelfFullResponseText) GetSolutionsOk() (*[]SolverSolutionWithSolvingSteps, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *SolverSelfFullResponseText) SetSolutions(v []SolverSolutionWithSolvingSteps)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *SolverSelfFullResponseText) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


