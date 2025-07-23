# AccessRecords API

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClaim**](#CreateClaim) | **Post** /v1/claims | Create resource Claim
[**CreateRecord**](#CreateRecord) | **Post** /v1/records | Create access record
[**CreateRequest**](#CreateRequest) | **Post** /v1/requests | Create access request
[**DeleteRecord**](#DeleteRecord) | **Delete** /v1/records/{recordId} | Deletes access record
[**DeleteRequest**](#DeleteRequest) | **Delete** /v1/requests/{requestId} | Deletes access request
[**GetRecord**](#GetRecord) | **Get** /v1/records/{recordId} | Retrieve access record
[**GetRecords**](#GetRecords) | **Get** /v1/records | List access records
[**GetRequest**](#GetRequest) | **Get** /v1/requests/{requestId} | Retrieve access request
[**GetRequests**](#GetRequests) | **Get** /v1/requests | List access requests
[**RespondToAccessRequest**](#RespondToAccessRequest) | **Patch** /v1/requests/{requestId} | Approve or deny access request
[**UpdateRecord**](#UpdateRecord) | **Put** /v1/records/{recordId} | Update access record



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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	claimRequest := *models.NewClaimRequest("CollectionResource_example", "ResourceId_example") // ClaimRequest | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.CreateClaim(context.Background()).ClaimRequest(claimRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.CreateClaim``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateClaim`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.CreateClaim`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	accessRecord := *models.NewAccessRecord("Name_example", []authress.Statement{*authress.NewStatement([]string{"Roles_example"}, []authress.Resource{*authress.NewResource("/organizations/org_a/documents/doc_1")})}) // AccessRecord | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.CreateRecord(context.Background()).AccessRecord(accessRecord).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.CreateRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRecord`: AccessRecord
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.CreateRecord`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	accessRequest := *models.NewAccessRequest("RequestId_example", *authress.NewAccessTemplate([]authress.User{*authress.NewUser("oauth|userId")}, []authress.Statement{*authress.NewStatement([]string{"Roles_example"}, []authress.Resource{*authress.NewResource("/organizations/org_a/documents/doc_1")})})) // AccessRequest | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.CreateRequest(context.Background()).AccessRequest(accessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.CreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.CreateRequest`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.AccessRecords.DeleteRecord(context.Background(), recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.DeleteRecord``: %v\n", err)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.AccessRecords.DeleteRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.DeleteRequest``: %v\n", err)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.GetRecord(context.Background(), recordId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.GetRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecord`: AccessRecord
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.GetRecord`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)
	filter := "filter_example" // string | Filter to search records by. This is a case insensitive search through every text field. (optional)
	status := "status_example" // string | Filter records by their current status. (optional)

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.GetRecords(context.Background()).Limit(limit).Cursor(cursor).Filter(filter).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.GetRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRecords`: AccessRecordCollection
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.GetRecords`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.GetRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.GetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.GetRequest`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	limit := int32(56) // int32 | Max number of results to return (optional) (default to 20)
	cursor := "cursor_example" // string | Continuation cursor for paging (optional)
	status := "status_example" // string | Filter requests by their current status. (optional)

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.GetRequests(context.Background()).Limit(limit).Cursor(cursor).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.GetRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequests`: AccessRequestCollection
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.GetRequests`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	requestId := "requestId_example" // string | The identifier of the access request.
	accessRequestResponse := *models.NewAccessRequestResponse("Status_example") // AccessRequestResponse | 

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	resp, r, err := authressClient.AccessRecords.RespondToAccessRequest(context.Background(), requestId).AccessRequestResponse(accessRequestResponse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.RespondToAccessRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RespondToAccessRequest`: AccessRequest
	fmt.Fprintf(os.Stdout, "Response from `AccessRecords.RespondToAccessRequest`: %v\n", resp)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


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
	authress "github.com/authress/authress-sdk.go"
)

func main() {
	recordId := "recordId_example" // string | The identifier of the access record.
	accessRecord := *models.NewAccessRecord("Name_example", []authress.Statement{*authress.NewStatement([]string{"Roles_example"}, []authress.Resource{*authress.NewResource("/organizations/org_a/documents/doc_1")})}) // AccessRecord | 
	ifUnmodifiedSince := time.Now() // time.Time | The expected last time the record was modified. (optional)

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
		ServiceClientAccessKey: serviceClientAccessKey,
	})
	r, err := authressClient.AccessRecords.UpdateRecord(context.Background(), recordId).AccessRecord(accessRecord).IfUnmodifiedSince(ifUnmodifiedSince).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccessRecords.UpdateRecord``: %v\n", err)
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

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

