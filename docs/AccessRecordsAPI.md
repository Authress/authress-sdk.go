# \AccessRecordsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClaim**](AccessRecordsAPI.md#CreateClaim) | **Post** /v1/claims | Create resource Claim
[**CreateRecord**](AccessRecordsAPI.md#CreateRecord) | **Post** /v1/records | Create access record
[**CreateRequest**](AccessRecordsAPI.md#CreateRequest) | **Post** /v1/requests | Create access request
[**DeleteRecord**](AccessRecordsAPI.md#DeleteRecord) | **Delete** /v1/records/{recordId} | Deletes access record
[**DeleteRequest**](AccessRecordsAPI.md#DeleteRequest) | **Delete** /v1/requests/{requestId} | Deletes access request
[**GetRecord**](AccessRecordsAPI.md#GetRecord) | **Get** /v1/records/{recordId} | Retrieve access record
[**GetRecords**](AccessRecordsAPI.md#GetRecords) | **Get** /v1/records | List access records
[**GetRequest**](AccessRecordsAPI.md#GetRequest) | **Get** /v1/requests/{requestId} | Retrieve access request
[**GetRequests**](AccessRecordsAPI.md#GetRequests) | **Get** /v1/requests | List access requests
[**RespondToAccessRequest**](AccessRecordsAPI.md#RespondToAccessRequest) | **Patch** /v1/requests/{requestId} | Approve or deny access request
[**UpdateRecord**](AccessRecordsAPI.md#UpdateRecord) | **Put** /v1/records/{recordId} | Update access record



## CreateClaim

> map[string]interface{} CreateClaim(ctx).ClaimRequest(claimRequest).Execute()

Create resource Claim



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	claimRequest := *openapiclient.NewClaimRequest("CollectionResource_example", "ResourceId_example") // ClaimRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.CreateClaim(context.Background()).ClaimRequest(claimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.CreateClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClaim`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.CreateClaim`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **claimRequest** | [**ClaimRequest**](ClaimRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRecord

> AccessRecord CreateRecord(ctx).AccessRecord(accessRecord).Execute()

Create access record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	accessRecord := *openapiclient.NewAccessRecord("Name_example", []openapiclient.Statement{*openapiclient.NewStatement([]string{"Roles_example"}, []openapiclient.Resource{*openapiclient.NewResource("/organizations/org_a/documents/doc_1")})}) // AccessRecord | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.CreateRecord(context.Background()).AccessRecord(accessRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.CreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecord`: AccessRecord
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.CreateRecord`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRecord** | [**AccessRecord**](AccessRecord.md) |  | 

### Return type

[**AccessRecord**](AccessRecord.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequest

> AccessRequest CreateRequest(ctx).AccessRequest(accessRequest).Execute()

Create access request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	accessRequest := *openapiclient.NewAccessRequest("RequestId_example", *openapiclient.NewAccessTemplate([]openapiclient.User{*openapiclient.NewUser("oauth|userId")}, []openapiclient.Statement{*openapiclient.NewStatement([]string{"Roles_example"}, []openapiclient.Resource{*openapiclient.NewResource("/organizations/org_a/documents/doc_1")})})) // AccessRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.CreateRequest(context.Background()).AccessRequest(accessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.CreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessRequest** | [**AccessRequest**](AccessRequest.md) |  | 

### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRecord

> DeleteRecord(ctx, recordId).Execute()

Deletes access record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessRecordsAPI.DeleteRecord(context.Background(), recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.DeleteRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** | The identifier of the access record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequest

> DeleteRequest(ctx, requestId).Execute()

Deletes access request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessRecordsAPI.DeleteRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.DeleteRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The identifier of the access request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecord

> AccessRecord GetRecord(ctx, recordId).Execute()

Retrieve access record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.GetRecord(context.Background(), recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.GetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecord`: AccessRecord
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.GetRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** | The identifier of the access record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRecord**](AccessRecord.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecords

> AccessRecordCollection GetRecords(ctx).Limit(limit).Cursor(cursor).Filter(filter).Status(status).Execute()

List access records



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)
	filter := "filter_example" // string | Filter to search records by. This is a case insensitive search through every text field. (optional)
	status := "status_example" // string | Filter records by their current status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.GetRecords(context.Background()).Limit(limit).Cursor(cursor).Filter(filter).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.GetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecords`: AccessRecordCollection
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.GetRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 
 **filter** | **string** | Filter to search records by. This is a case insensitive search through every text field. | 
 **status** | **string** | Filter records by their current status. | 

### Return type

[**AccessRecordCollection**](AccessRecordCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequest

> AccessRequest GetRequest(ctx, requestId).Execute()

Retrieve access request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.GetRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.GetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.GetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The identifier of the access request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequests

> AccessRequestCollection GetRequests(ctx).Limit(limit).Cursor(cursor).Status(status).Execute()

List access requests



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)
	status := "status_example" // string | Filter requests by their current status. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.GetRequests(context.Background()).Limit(limit).Cursor(cursor).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.GetRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequests`: AccessRequestCollection
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.GetRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Max number of results to return | [default to 20]
 **cursor** | **string** | Continuation cursor for paging | 
 **status** | **string** | Filter requests by their current status. | 

### Return type

[**AccessRequestCollection**](AccessRequestCollection.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RespondToAccessRequest

> AccessRequest RespondToAccessRequest(ctx, requestId).AccessRequestResponse(accessRequestResponse).Execute()

Approve or deny access request



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "//"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.
	accessRequestResponse := *openapiclient.NewAccessRequestResponse("Status_example") // AccessRequestResponse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccessRecordsAPI.RespondToAccessRequest(context.Background(), requestId).AccessRequestResponse(accessRequestResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.RespondToAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondToAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecordsAPI.RespondToAccessRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The identifier of the access request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRespondToAccessRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRequestResponse** | [**AccessRequestResponse**](AccessRequestResponse.md) |  | 

### Return type

[**AccessRequest**](AccessRequest.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRecord

> UpdateRecord(ctx, recordId).AccessRecord(accessRecord).IfUnmodifiedSince(ifUnmodifiedSince).Execute()

Update access record



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "//"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.
	accessRecord := *openapiclient.NewAccessRecord("Name_example", []openapiclient.Statement{*openapiclient.NewStatement([]string{"Roles_example"}, []openapiclient.Resource{*openapiclient.NewResource("/organizations/org_a/documents/doc_1")})}) // AccessRecord | 
	ifUnmodifiedSince := time.Now() // time.Time | The expected last time the record was modified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AccessRecordsAPI.UpdateRecord(context.Background(), recordId).AccessRecord(accessRecord).IfUnmodifiedSince(ifUnmodifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecordsAPI.UpdateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recordId** | **string** | The identifier of the access record. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessRecord** | [**AccessRecord**](AccessRecord.md) |  | 
 **ifUnmodifiedSince** | **time.Time** | The expected last time the record was modified. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

