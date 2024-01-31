# UserToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**PermissionCollectionAccount**](PermissionCollectionAccount.md) |  | [optional] 
**UserId** | [**UserId**](UserId.md) |  | 
**TokenId** | **string** | The unique identifier for the token | 
**Token** | **string** | The access token | 
**Links** | Pointer to [**AccountLinks**](AccountLinks.md) |  | [optional] 

## Methods

### NewUserToken

`func NewUserToken(userId UserId, tokenId string, token string, ) *UserToken`

NewUserToken instantiates a new UserToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTokenWithDefaults

`func NewUserTokenWithDefaults() *UserToken`

NewUserTokenWithDefaults instantiates a new UserToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *UserToken) GetAccount() PermissionCollectionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UserToken) GetAccountOk() (*PermissionCollectionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UserToken) SetAccount(v PermissionCollectionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UserToken) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUserId

`func (o *UserToken) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserToken) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserToken) SetUserId(v UserId)`

SetUserId sets UserId field to given value.


### GetTokenId

`func (o *UserToken) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UserToken) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UserToken) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetToken

`func (o *UserToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UserToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UserToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetLinks

`func (o *UserToken) GetLinks() AccountLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserToken) GetLinksOk() (*AccountLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserToken) SetLinks(v AccountLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserToken) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


