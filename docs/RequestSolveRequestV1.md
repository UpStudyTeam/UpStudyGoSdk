# RequestSolveRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Input** | **string** | Mathematical expression input, must be in Latex format, e.g., \\frac{x+1}{2} \\leq 0 | 
**Lang** | Pointer to [**RequestSolveRequestV1Lang**](RequestSolveRequestV1Lang.md) |  | [optional] 

## Methods

### NewRequestSolveRequestV1

`func NewRequestSolveRequestV1(input string, ) *RequestSolveRequestV1`

NewRequestSolveRequestV1 instantiates a new RequestSolveRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestSolveRequestV1WithDefaults

`func NewRequestSolveRequestV1WithDefaults() *RequestSolveRequestV1`

NewRequestSolveRequestV1WithDefaults instantiates a new RequestSolveRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInput

`func (o *RequestSolveRequestV1) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *RequestSolveRequestV1) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *RequestSolveRequestV1) SetInput(v string)`

SetInput sets Input field to given value.


### GetLang

`func (o *RequestSolveRequestV1) GetLang() RequestSolveRequestV1Lang`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *RequestSolveRequestV1) GetLangOk() (*RequestSolveRequestV1Lang, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *RequestSolveRequestV1) SetLang(v RequestSolveRequestV1Lang)`

SetLang sets Lang field to given value.

### HasLang

`func (o *RequestSolveRequestV1) HasLang() bool`

HasLang returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


