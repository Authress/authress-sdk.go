# Accounts API

Method | HTTP request | Description
------------- | ------------- | -------------
[**DelegateAuthentication**](#DelegateAuthentication) | **Post** /v1/identities | Link external provider
[**GetAccount**](#GetAccount) | **Get** /v1/accounts/{accountId} | Retrieve account information
[**GetAccountIdentities**](#GetAccountIdentities) | **Get** /v1/identities | List linked external providers
[**GetAccounts**](#GetAccounts) | **Get** /v1/accounts | List user Authress accounts



## DelegateAuthentication

> DelegateAuthentication(ctx).IdentityRequest(identityRequest).Execute()

Link external provider



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
	identityRequest := *authress.NewIdentityRequest() // IdentityRequest | 

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	r, err := authressClient.Accounts.DelegateAuthentication(context.Background()).IdentityRequest(identityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.DelegateAuthentication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDelegateAuthenticationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityRequest** | [**IdentityRequest**](IdentityRequest.md) |  | 

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetAccount

> Account GetAccount(ctx, accountId).Execute()

Retrieve account information



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
	accountId := "accountId_example" // string | The unique identifier for the account

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := authressClient.Accounts.GetAccount(context.Background(), accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountId** | **string** | The unique identifier for the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Account**](Account.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetAccountIdentities

> IdentityCollection GetAccountIdentities(ctx).Execute()

List linked external providers



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

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := authressClient.Accounts.GetAccountIdentities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetAccountIdentities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountIdentities`: IdentityCollection
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetAccountIdentities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountIdentitiesRequest struct via the builder pattern


### Return type

[**IdentityCollection**](IdentityCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)


## GetAccounts

> AccountCollection GetAccounts(ctx).EarliestCacheTime(earliestCacheTime).Execute()

List user Authress accounts



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
	earliestCacheTime := time.Now() // time.Time | Ensure the accounts list is not cached before this time. (optional)

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})
	resp, r, err := authressClient.Accounts.GetAccounts(context.Background()).EarliestCacheTime(earliestCacheTime).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `Accounts.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: AccountCollection
	fmt.Fprintf(os.Stdout, "Response from `Accounts.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **earliestCacheTime** | **time.Time** | Ensure the accounts list is not cached before this time. | 

### Return type

[**AccountCollection**](AccountCollection.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/links+json

[[Back to top]](#) [[Back to API list]](./README.md#documentation-for-api-endpoints)
[[Back to Model list]](./README.md#documentation-for-models)
[[Back to README]](./README.md)

