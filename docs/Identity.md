# Identity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** | The issuer of the JWT token. This can be any OIDC compliant provider. | 
**Audience** | **string** | The audience of the JWT token. This can be either an audience for your entire app, or one particular audience for it. If there is more than one audience, multiple identities can be created. | 
**UserIdExpression** | Pointer to **NullableString** | By default, the **sub** claim of the JWT is used to identify the user from this provider. To use an alternate claim or a compound userId resolution, specify an expression. The resolved userId must be the same one that is later used in Authress access records. | [optional] [default to "{sub}"]

## Methods

### NewIdentity

`func NewIdentity(issuer string, audience string, ) *Identity`

NewIdentity instantiates a new Identity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityWithDefaults

`func NewIdentityWithDefaults() *Identity`

NewIdentityWithDefaults instantiates a new Identity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *Identity) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *Identity) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *Identity) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetAudience

`func (o *Identity) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Identity) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Identity) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetUserIdExpression

`func (o *Identity) GetUserIdExpression() string`

GetUserIdExpression returns the UserIdExpression field if non-nil, zero value otherwise.

### GetUserIdExpressionOk

`func (o *Identity) GetUserIdExpressionOk() (*string, bool)`

GetUserIdExpressionOk returns a tuple with the UserIdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdExpression

`func (o *Identity) SetUserIdExpression(v string)`

SetUserIdExpression sets UserIdExpression field to given value.

### HasUserIdExpression

`func (o *Identity) HasUserIdExpression() bool`

HasUserIdExpression returns a boolean if a field has been set.

### SetUserIdExpressionNil

`func (o *Identity) SetUserIdExpressionNil(b bool)`

 SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil

### UnsetUserIdExpression
`func (o *Identity) UnsetUserIdExpression()`

UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


