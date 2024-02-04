# IdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jwt** | Pointer to **NullableString** | A valid JWT OIDC compliant token which will still pass authentication requests to the identity provider. Must contain a unique audience and issuer. | [optional] 
**Issuer** | Pointer to **NullableString** | The issuer of the OAuth OIDC provider&#39;s JWTs. This value should match the &#x60;iss&#x60; claim in the provided tokens exactly. | [optional] 
**PreferredAudience** | Pointer to **NullableString** | If the &#x60;jwt&#x60; token contains more than one valid audience, then the single audience that should associated with Authress. If more than one audience is preferred, repeat this call with each one. | [optional] [default to "*"]
**UserIdExpression** | Pointer to **NullableString** | By default, the **sub** claim of the JWT is used to identify the user from this provider. To use an alternate claim or a compound userId resolution, specify an expression. The resolved userId must be the same one that is later used in Authress access records. | [optional] [default to "{sub}"]

## Methods

### NewIdentityRequest

`func NewIdentityRequest() *IdentityRequest`

NewIdentityRequest instantiates a new IdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityRequestWithDefaults

`func NewIdentityRequestWithDefaults() *IdentityRequest`

NewIdentityRequestWithDefaults instantiates a new IdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwt

`func (o *IdentityRequest) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *IdentityRequest) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *IdentityRequest) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *IdentityRequest) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### SetJwtNil

`func (o *IdentityRequest) SetJwtNil(b bool)`

 SetJwtNil sets the value for Jwt to be an explicit nil

### UnsetJwt
`func (o *IdentityRequest) UnsetJwt()`

UnsetJwt ensures that no value is present for Jwt, not even an explicit nil
### GetIssuer

`func (o *IdentityRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *IdentityRequest) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### SetIssuerNil

`func (o *IdentityRequest) SetIssuerNil(b bool)`

 SetIssuerNil sets the value for Issuer to be an explicit nil

### UnsetIssuer
`func (o *IdentityRequest) UnsetIssuer()`

UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
### GetPreferredAudience

`func (o *IdentityRequest) GetPreferredAudience() string`

GetPreferredAudience returns the PreferredAudience field if non-nil, zero value otherwise.

### GetPreferredAudienceOk

`func (o *IdentityRequest) GetPreferredAudienceOk() (*string, bool)`

GetPreferredAudienceOk returns a tuple with the PreferredAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredAudience

`func (o *IdentityRequest) SetPreferredAudience(v string)`

SetPreferredAudience sets PreferredAudience field to given value.

### HasPreferredAudience

`func (o *IdentityRequest) HasPreferredAudience() bool`

HasPreferredAudience returns a boolean if a field has been set.

### SetPreferredAudienceNil

`func (o *IdentityRequest) SetPreferredAudienceNil(b bool)`

 SetPreferredAudienceNil sets the value for PreferredAudience to be an explicit nil

### UnsetPreferredAudience
`func (o *IdentityRequest) UnsetPreferredAudience()`

UnsetPreferredAudience ensures that no value is present for PreferredAudience, not even an explicit nil
### GetUserIdExpression

`func (o *IdentityRequest) GetUserIdExpression() string`

GetUserIdExpression returns the UserIdExpression field if non-nil, zero value otherwise.

### GetUserIdExpressionOk

`func (o *IdentityRequest) GetUserIdExpressionOk() (*string, bool)`

GetUserIdExpressionOk returns a tuple with the UserIdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdExpression

`func (o *IdentityRequest) SetUserIdExpression(v string)`

SetUserIdExpression sets UserIdExpression field to given value.

### HasUserIdExpression

`func (o *IdentityRequest) HasUserIdExpression() bool`

HasUserIdExpression returns a boolean if a field has been set.

### SetUserIdExpressionNil

`func (o *IdentityRequest) SetUserIdExpressionNil(b bool)`

 SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil

### UnsetUserIdExpression
`func (o *IdentityRequest) UnsetUserIdExpression()`

UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


