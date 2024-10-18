# SolverSolution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockName** | Pointer to [**SolverDescription**](SolverDescription.md) | Block name: e.g., Function/Solve the equation | [optional] 
**Results** | Pointer to [**[]SolverStep**](SolverStep.md) | Results, possibly multiple, e.g., {\\frac{1}{4},0.25} | [optional] 
**SolutionName** | Pointer to [**SolverDescription**](SolverDescription.md) | Specific solution name under the block classification: e.g., Block: Function, Solution: Find the slope | [optional] 

## Methods

### NewSolverSolution

`func NewSolverSolution() *SolverSolution`

NewSolverSolution instantiates a new SolverSolution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverSolutionWithDefaults

`func NewSolverSolutionWithDefaults() *SolverSolution`

NewSolverSolutionWithDefaults instantiates a new SolverSolution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockName

`func (o *SolverSolution) GetBlockName() SolverDescription`

GetBlockName returns the BlockName field if non-nil, zero value otherwise.

### GetBlockNameOk

`func (o *SolverSolution) GetBlockNameOk() (*SolverDescription, bool)`

GetBlockNameOk returns a tuple with the BlockName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockName

`func (o *SolverSolution) SetBlockName(v SolverDescription)`

SetBlockName sets BlockName field to given value.

### HasBlockName

`func (o *SolverSolution) HasBlockName() bool`

HasBlockName returns a boolean if a field has been set.

### GetResults

`func (o *SolverSolution) GetResults() []SolverStep`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *SolverSolution) GetResultsOk() (*[]SolverStep, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *SolverSolution) SetResults(v []SolverStep)`

SetResults sets Results field to given value.

### HasResults

`func (o *SolverSolution) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetSolutionName

`func (o *SolverSolution) GetSolutionName() SolverDescription`

GetSolutionName returns the SolutionName field if non-nil, zero value otherwise.

### GetSolutionNameOk

`func (o *SolverSolution) GetSolutionNameOk() (*SolverDescription, bool)`

GetSolutionNameOk returns a tuple with the SolutionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionName

`func (o *SolverSolution) SetSolutionName(v SolverDescription)`

SetSolutionName sets SolutionName field to given value.

### HasSolutionName

`func (o *SolverSolution) HasSolutionName() bool`

HasSolutionName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


