# SolverSolutionWithSolvingSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockName** | Pointer to [**SolverDescription**](SolverDescription.md) | Block name: e.g., Function/Solve the equation | [optional] 
**Results** | Pointer to [**[]SolverStep**](SolverStep.md) | Results, possibly multiple, e.g., {\\frac{1}{4},0.25} | [optional] 
**SolutionName** | Pointer to [**SolverDescription**](SolverDescription.md) | Specific solution name under the block classification: e.g., Block: Function, Solution: Find the slope | [optional] 
**SolvingSteps** | Pointer to [**[]SolverStep**](SolverStep.md) | Steps to obtain the solution | [optional] 

## Methods

### NewSolverSolutionWithSolvingSteps

`func NewSolverSolutionWithSolvingSteps() *SolverSolutionWithSolvingSteps`

NewSolverSolutionWithSolvingSteps instantiates a new SolverSolutionWithSolvingSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverSolutionWithSolvingStepsWithDefaults

`func NewSolverSolutionWithSolvingStepsWithDefaults() *SolverSolutionWithSolvingSteps`

NewSolverSolutionWithSolvingStepsWithDefaults instantiates a new SolverSolutionWithSolvingSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockName

`func (o *SolverSolutionWithSolvingSteps) GetBlockName() SolverDescription`

GetBlockName returns the BlockName field if non-nil, zero value otherwise.

### GetBlockNameOk

`func (o *SolverSolutionWithSolvingSteps) GetBlockNameOk() (*SolverDescription, bool)`

GetBlockNameOk returns a tuple with the BlockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockName

`func (o *SolverSolutionWithSolvingSteps) SetBlockName(v SolverDescription)`

SetBlockName sets BlockName field to given value.

### HasBlockName

`func (o *SolverSolutionWithSolvingSteps) HasBlockName() bool`

HasBlockName returns a boolean if a field has been set.

### GetResults

`func (o *SolverSolutionWithSolvingSteps) GetResults() []SolverStep`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SolverSolutionWithSolvingSteps) GetResultsOk() (*[]SolverStep, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SolverSolutionWithSolvingSteps) SetResults(v []SolverStep)`

SetResults sets Results field to given value.

### HasResults

`func (o *SolverSolutionWithSolvingSteps) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSolutionName

`func (o *SolverSolutionWithSolvingSteps) GetSolutionName() SolverDescription`

GetSolutionName returns the SolutionName field if non-nil, zero value otherwise.

### GetSolutionNameOk

`func (o *SolverSolutionWithSolvingSteps) GetSolutionNameOk() (*SolverDescription, bool)`

GetSolutionNameOk returns a tuple with the SolutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionName

`func (o *SolverSolutionWithSolvingSteps) SetSolutionName(v SolverDescription)`

SetSolutionName sets SolutionName field to given value.

### HasSolutionName

`func (o *SolverSolutionWithSolvingSteps) HasSolutionName() bool`

HasSolutionName returns a boolean if a field has been set.

### GetSolvingSteps

`func (o *SolverSolutionWithSolvingSteps) GetSolvingSteps() []SolverStep`

GetSolvingSteps returns the SolvingSteps field if non-nil, zero value otherwise.

### GetSolvingStepsOk

`func (o *SolverSolutionWithSolvingSteps) GetSolvingStepsOk() (*[]SolverStep, bool)`

GetSolvingStepsOk returns a tuple with the SolvingSteps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolvingSteps

`func (o *SolverSolutionWithSolvingSteps) SetSolvingSteps(v []SolverStep)`

SetSolvingSteps sets SolvingSteps field to given value.

### HasSolvingSteps

`func (o *SolverSolutionWithSolvingSteps) HasSolvingSteps() bool`

HasSolvingSteps returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


