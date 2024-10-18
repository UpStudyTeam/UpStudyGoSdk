# SolverSelfBriefResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | Pointer to **string** | User&#39;s Latex input | [optional] 
**Solutions** | Pointer to [**[]SolverSolution**](SolverSolution.md) | All solutions we can provide | [optional] 

## Methods

### NewSolverSelfBriefResponse

`func NewSolverSelfBriefResponse() *SolverSelfBriefResponse`

NewSolverSelfBriefResponse instantiates a new SolverSelfBriefResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSolverSelfBriefResponseWithDefaults

`func NewSolverSelfBriefResponseWithDefaults() *SolverSelfBriefResponse`

NewSolverSelfBriefResponseWithDefaults instantiates a new SolverSelfBriefResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *SolverSelfBriefResponse) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *SolverSelfBriefResponse) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *SolverSelfBriefResponse) SetInput(v string)`

SetInput sets Input field to given value.

### HasInput

`func (o *SolverSelfBriefResponse) HasInput() bool`

HasInput returns a boolean if a field has been set.

### GetSolutions

`func (o *SolverSelfBriefResponse) GetSolutions() []SolverSolution`

GetSolutions returns the Solutions field if non-nil, zero value otherwise.

### GetSolutionsOk

`func (o *SolverSelfBriefResponse) GetSolutionsOk() (*[]SolverSolution, bool)`

GetSolutionsOk returns a tuple with the Solutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutions

`func (o *SolverSelfBriefResponse) SetSolutions(v []SolverSolution)`

SetSolutions sets Solutions field to given value.

### HasSolutions

`func (o *SolverSelfBriefResponse) HasSolutions() bool`

HasSolutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


