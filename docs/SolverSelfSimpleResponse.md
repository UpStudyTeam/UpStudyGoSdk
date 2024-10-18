# SolverSelfSimpleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | User&#39;s Latex input | [optional] 
**Solution** | Pointer to [**SolverSolution**](SolverSolution.md) | The solution we believe the user most likely wants | [optional] 

## Methods

### NewSolverSelfSimpleResponse

`func NewSolverSelfSimpleResponse() *SolverSelfSimpleResponse`

NewSolverSelfSimpleResponse instantiates a new SolverSelfSimpleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverSelfSimpleResponseWithDefaults

`func NewSolverSelfSimpleResponseWithDefaults() *SolverSelfSimpleResponse`

NewSolverSelfSimpleResponseWithDefaults instantiates a new SolverSelfSimpleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *SolverSelfSimpleResponse) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *SolverSelfSimpleResponse) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *SolverSelfSimpleResponse) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *SolverSelfSimpleResponse) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetSolution

`func (o *SolverSelfSimpleResponse) GetSolution() SolverSolution`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *SolverSelfSimpleResponse) GetSolutionOk() (*SolverSolution, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *SolverSelfSimpleResponse) SetSolution(v SolverSolution)`

SetSolution sets Solution field to given value.

### HasSolution

`func (o *SolverSelfSimpleResponse) HasSolution() bool`

HasSolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


