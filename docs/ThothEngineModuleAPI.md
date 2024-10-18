# \ThothEngineModuleAPI

All URIs are relative to *https://api.cameramath.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1BriefAnswersPost**](ThothEngineModuleAPI.md#V1BriefAnswersPost) | **Post** /v1/brief-answers | Answers to all lists of each block
[**V1ShowStepsPost**](ThothEngineModuleAPI.md#V1ShowStepsPost) | **Post** /v1/show-steps | All lists and all steps of all blocks
[**V1SingleAnswerPost**](ThothEngineModuleAPI.md#V1SingleAnswerPost) | **Post** /v1/single-answer | First answer of the first block, text version



## V1BriefAnswersPost

> V1BriefAnswersPost200Response V1BriefAnswersPost(ctx).Data(data).Execute()

Answers to all lists of each block

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UpStudyTeam/UpStudyGoSdk"
)

func main() {
	data := *openapiclient.NewRequestSolveRequestV1("x+1=2") // RequestSolveRequestV1 | Solution request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThothEngineModuleAPI.V1BriefAnswersPost(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThothEngineModuleAPI.V1BriefAnswersPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1BriefAnswersPost`: V1BriefAnswersPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ThothEngineModuleAPI.V1BriefAnswersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1BriefAnswersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**RequestSolveRequestV1**](RequestSolveRequestV1.md) | Solution request parameters | 

### Return type

[**V1BriefAnswersPost200Response**](V1BriefAnswersPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1ShowStepsPost

> V1ShowStepsPost200Response V1ShowStepsPost(ctx).Data(data).Execute()

All lists and all steps of all blocks

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UpStudyTeam/UpStudyGoSdk"
)

func main() {
	data := *openapiclient.NewRequestSolveRequestV1("x+1=2") // RequestSolveRequestV1 | Solution request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThothEngineModuleAPI.V1ShowStepsPost(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThothEngineModuleAPI.V1ShowStepsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1ShowStepsPost`: V1ShowStepsPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ThothEngineModuleAPI.V1ShowStepsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1ShowStepsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**RequestSolveRequestV1**](RequestSolveRequestV1.md) | Solution request parameters | 

### Return type

[**V1ShowStepsPost200Response**](V1ShowStepsPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1SingleAnswerPost

> V1SingleAnswerPost200Response V1SingleAnswerPost(ctx).Data(data).Execute()

First answer of the first block, text version

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/UpStudyTeam/UpStudyGoSdk"
)

func main() {
	data := *openapiclient.NewRequestSolveRequestV1("x+1=2") // RequestSolveRequestV1 | Solution request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ThothEngineModuleAPI.V1SingleAnswerPost(context.Background()).Data(data).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ThothEngineModuleAPI.V1SingleAnswerPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `V1SingleAnswerPost`: V1SingleAnswerPost200Response
	fmt.Fprintf(os.Stdout, "Response from `ThothEngineModuleAPI.V1SingleAnswerPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiV1SingleAnswerPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **data** | [**RequestSolveRequestV1**](RequestSolveRequestV1.md) | Solution request parameters | 

### Return type

[**V1SingleAnswerPost200Response**](V1SingleAnswerPost200Response.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

