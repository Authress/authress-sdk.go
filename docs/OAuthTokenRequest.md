# OAuthTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The client identifier to constrain the token to. | 
**ClientSecret** | Pointer to **NullableString** | The secret associated with the client that authorizes the generation of token it&#39;s behalf. (Either the &#x60;client_secret&#x60; or the &#x60;code_verifier&#x60; is required) | [optional] 
**CodeVerifier** | Pointer to **string** | The code verifier is the the value used in the generation of the OAuth authorization request &#x60;code_challenge&#x60; property. (Either the &#x60;client_secret&#x60; or the &#x60;code_verifier&#x60; is required) | [optional] 
**GrantType** | Pointer to **string** | A suggestion to the token generation which type of credentials are being provided. | [optional] 
**Username** | Pointer to **NullableString** | When using the user password grant_type, specify the username. Authress recommends this should always be an email address. | [optional] 
**Password** | Pointer to **NullableString** | When using the user password grant_type, specify the user&#39;s password. | [optional] 
**Type** | Pointer to **NullableString** | Enables additional configuration of the grant_type. In the case of user password grant_type, set this to **signup**, to enable the creation of users. Set this to **update**, force update the password associated with the user. | [optional] 

## Methods

### NewOAuthTokenRequest

`func NewOAuthTokenRequest(clientId string, ) *OAuthTokenRequest`

NewOAuthTokenRequest instantiates a new OAuthTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthTokenRequestWithDefaults

`func NewOAuthTokenRequestWithDefaults() *OAuthTokenRequest`

NewOAuthTokenRequestWithDefaults instantiates a new OAuthTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *OAuthTokenRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *OAuthTokenRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *OAuthTokenRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *OAuthTokenRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *OAuthTokenRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *OAuthTokenRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *OAuthTokenRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *OAuthTokenRequest) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *OAuthTokenRequest) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetCodeVerifier

`func (o *OAuthTokenRequest) GetCodeVerifier() string`

GetCodeVerifier returns the CodeVerifier field if non-nil, zero value otherwise.

### GetCodeVerifierOk

`func (o *OAuthTokenRequest) GetCodeVerifierOk() (*string, bool)`

GetCodeVerifierOk returns a tuple with the CodeVerifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeVerifier

`func (o *OAuthTokenRequest) SetCodeVerifier(v string)`

SetCodeVerifier sets CodeVerifier field to given value.

### HasCodeVerifier

`func (o *OAuthTokenRequest) HasCodeVerifier() bool`

HasCodeVerifier returns a boolean if a field has been set.

### GetGrantType

`func (o *OAuthTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *OAuthTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *OAuthTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *OAuthTokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetUsername

`func (o *OAuthTokenRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *OAuthTokenRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *OAuthTokenRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *OAuthTokenRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *OAuthTokenRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *OAuthTokenRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *OAuthTokenRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OAuthTokenRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OAuthTokenRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OAuthTokenRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *OAuthTokenRequest) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *OAuthTokenRequest) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetType

`func (o *OAuthTokenRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OAuthTokenRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OAuthTokenRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OAuthTokenRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *OAuthTokenRequest) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *OAuthTokenRequest) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


