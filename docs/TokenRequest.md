# TokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statements** | [**[]Statement**](Statement.md) | A list of statements which match roles to resources. The token will have all statements apply to it. | 
**Expires** | **time.Time** | The ISO8601 datetime when the token will expire. Default is 24 hours from now. | 

## Methods

### NewTokenRequest

`func NewTokenRequest(statements []Statement, expires time.Time, ) *TokenRequest`

NewTokenRequest instantiates a new TokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenRequestWithDefaults

`func NewTokenRequestWithDefaults() *TokenRequest`

NewTokenRequestWithDefaults instantiates a new TokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatements

`func (o *TokenRequest) GetStatements() []Statement`

GetStatements returns the Statements field if non-nil, zero value otherwise.

### GetStatementsOk

`func (o *TokenRequest) GetStatementsOk() (*[]Statement, bool)`

GetStatementsOk returns a tuple with the Statements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatements

`func (o *TokenRequest) SetStatements(v []Statement)`

SetStatements sets Statements field to given value.


### GetExpires

`func (o *TokenRequest) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenRequest) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenRequest) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


