# ConnectionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** |  | [optional] 
**Name** | Pointer to **NullableString** | The name of this connection when displayed in the Authress management portal | [optional] 
**SupportedContentType** | Pointer to **NullableString** | URL encode OAuth token parameters - Some authentication APIs don&#39;t support JSON, in these cases enable the url encoded form data parameters. | [optional] [default to "application/json"]
**OidcUserEndpointUrl** | Pointer to **NullableString** | By default, the **sub** claim of the JWT is used to identify the user from this provider. However, not all providers are OpenID compliant. Here you can provide an optional user data endpoint to fetch additional user profile information and an expression to identify a new user ID from available properties. | [optional] 
**UserIdExpression** | Pointer to **NullableString** | By default, the **sub** claim of the JWT is used to identify the user from this provider. However, not all providers are OpenID compliant. Here you can provide an optional user expression to identify a new user ID from available properties found from the oidcUserEndpointUrl. (The default is **{sub}**, any claims may be used.) | [optional] [default to "{sub}"]
**TrustIdentityUserId** | Pointer to **NullableBool** | Authress generates unique user IDs for every user, however if you trust your identity provider to handle unique ID generate across **ALL customers**, then it is safe to reuse the user ID from the provider. | [optional] [default to false]

## Methods

### NewConnectionData

`func NewConnectionData() *ConnectionData`

NewConnectionData instantiates a new ConnectionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDataWithDefaults

`func NewConnectionDataWithDefaults() *ConnectionData`

NewConnectionDataWithDefaults instantiates a new ConnectionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *ConnectionData) GetTenantId() TenantId`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ConnectionData) GetTenantIdOk() (*TenantId, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ConnectionData) SetTenantId(v TenantId)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ConnectionData) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetName

`func (o *ConnectionData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionData) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ConnectionData) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ConnectionData) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSupportedContentType

`func (o *ConnectionData) GetSupportedContentType() string`

GetSupportedContentType returns the SupportedContentType field if non-nil, zero value otherwise.

### GetSupportedContentTypeOk

`func (o *ConnectionData) GetSupportedContentTypeOk() (*string, bool)`

GetSupportedContentTypeOk returns a tuple with the SupportedContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedContentType

`func (o *ConnectionData) SetSupportedContentType(v string)`

SetSupportedContentType sets SupportedContentType field to given value.

### HasSupportedContentType

`func (o *ConnectionData) HasSupportedContentType() bool`

HasSupportedContentType returns a boolean if a field has been set.

### SetSupportedContentTypeNil

`func (o *ConnectionData) SetSupportedContentTypeNil(b bool)`

 SetSupportedContentTypeNil sets the value for SupportedContentType to be an explicit nil

### UnsetSupportedContentType
`func (o *ConnectionData) UnsetSupportedContentType()`

UnsetSupportedContentType ensures that no value is present for SupportedContentType, not even an explicit nil
### GetOidcUserEndpointUrl

`func (o *ConnectionData) GetOidcUserEndpointUrl() string`

GetOidcUserEndpointUrl returns the OidcUserEndpointUrl field if non-nil, zero value otherwise.

### GetOidcUserEndpointUrlOk

`func (o *ConnectionData) GetOidcUserEndpointUrlOk() (*string, bool)`

GetOidcUserEndpointUrlOk returns a tuple with the OidcUserEndpointUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcUserEndpointUrl

`func (o *ConnectionData) SetOidcUserEndpointUrl(v string)`

SetOidcUserEndpointUrl sets OidcUserEndpointUrl field to given value.

### HasOidcUserEndpointUrl

`func (o *ConnectionData) HasOidcUserEndpointUrl() bool`

HasOidcUserEndpointUrl returns a boolean if a field has been set.

### SetOidcUserEndpointUrlNil

`func (o *ConnectionData) SetOidcUserEndpointUrlNil(b bool)`

 SetOidcUserEndpointUrlNil sets the value for OidcUserEndpointUrl to be an explicit nil

### UnsetOidcUserEndpointUrl
`func (o *ConnectionData) UnsetOidcUserEndpointUrl()`

UnsetOidcUserEndpointUrl ensures that no value is present for OidcUserEndpointUrl, not even an explicit nil
### GetUserIdExpression

`func (o *ConnectionData) GetUserIdExpression() string`

GetUserIdExpression returns the UserIdExpression field if non-nil, zero value otherwise.

### GetUserIdExpressionOk

`func (o *ConnectionData) GetUserIdExpressionOk() (*string, bool)`

GetUserIdExpressionOk returns a tuple with the UserIdExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIdExpression

`func (o *ConnectionData) SetUserIdExpression(v string)`

SetUserIdExpression sets UserIdExpression field to given value.

### HasUserIdExpression

`func (o *ConnectionData) HasUserIdExpression() bool`

HasUserIdExpression returns a boolean if a field has been set.

### SetUserIdExpressionNil

`func (o *ConnectionData) SetUserIdExpressionNil(b bool)`

 SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil

### UnsetUserIdExpression
`func (o *ConnectionData) UnsetUserIdExpression()`

UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil
### GetTrustIdentityUserId

`func (o *ConnectionData) GetTrustIdentityUserId() bool`

GetTrustIdentityUserId returns the TrustIdentityUserId field if non-nil, zero value otherwise.

### GetTrustIdentityUserIdOk

`func (o *ConnectionData) GetTrustIdentityUserIdOk() (*bool, bool)`

GetTrustIdentityUserIdOk returns a tuple with the TrustIdentityUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustIdentityUserId

`func (o *ConnectionData) SetTrustIdentityUserId(v bool)`

SetTrustIdentityUserId sets TrustIdentityUserId field to given value.

### HasTrustIdentityUserId

`func (o *ConnectionData) HasTrustIdentityUserId() bool`

HasTrustIdentityUserId returns a boolean if a field has been set.

### SetTrustIdentityUserIdNil

`func (o *ConnectionData) SetTrustIdentityUserIdNil(b bool)`

 SetTrustIdentityUserIdNil sets the value for TrustIdentityUserId to be an explicit nil

### UnsetTrustIdentityUserId
`func (o *ConnectionData) UnsetTrustIdentityUserId()`

UnsetTrustIdentityUserId ensures that no value is present for TrustIdentityUserId, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


