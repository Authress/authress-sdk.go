# OAuthTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | An expiring access token that can be used to access either Authress or any platform service. | 

## Methods

### NewOAuthTokenResponse

`func NewOAuthTokenResponse(accessToken string, ) *OAuthTokenResponse`

NewOAuthTokenResponse instantiates a new OAuthTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenResponseWithDefaults

`func NewOAuthTokenResponseWithDefaults() *OAuthTokenResponse`

NewOAuthTokenResponseWithDefaults instantiates a new OAuthTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *OAuthTokenResponse) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *OAuthTokenResponse) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *OAuthTokenResponse) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


