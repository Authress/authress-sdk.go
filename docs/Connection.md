# Connection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] [default to "OAUTH2"]
**ConnectionId** | Pointer to **NullableString** |  | [optional] 
**AuthenticationUrl** | Pointer to **string** | Authorization URL of the provider (where the user logs in). For OAuth this is the authorization URL, for SAML, this is the **SSO URL**. | [optional] 
**TokenUrl** | Pointer to **NullableString** | The token exchange url (where we validate the token). | [optional] 
**IssuerUrl** | Pointer to **NullableString** | The unique identifier tied to the provider&#39;s domain used for TLS verification. In OAuth, this is the JWT **iss** property. in SAML this is the **Entity ID**. | [optional] 
**ProviderCertificate** | Pointer to **NullableString** | The Provider&#39;s SAML public certificate which can be used to verify the signature in signed SAML assertions from the provider. | [optional] 
**ClientId** | Pointer to **NullableString** | Provider&#39;s client ID used to verify the token | [optional] 
**ClientSecret** | Pointer to **NullableString** | Provider&#39;s client secret used to verify the token | [optional] 
**UserDataConfiguration** | Pointer to [**NullableConnectionUserDataConfiguration**](ConnectionUserDataConfiguration.md) |  | [optional] 
**Data** | Pointer to [**NullableConnectionData**](ConnectionData.md) |  | [optional] 
**DefaultConnectionProperties** | Pointer to [**NullableConnectionDefaultConnectionProperties**](ConnectionDefaultConnectionProperties.md) |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] [readonly] 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewConnection

`func NewConnection() *Connection`

NewConnection instantiates a new Connection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionWithDefaults

`func NewConnectionWithDefaults() *Connection`

NewConnectionWithDefaults instantiates a new Connection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Connection) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Connection) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Connection) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Connection) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConnectionId

`func (o *Connection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Connection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Connection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *Connection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### SetConnectionIdNil

`func (o *Connection) SetConnectionIdNil(b bool)`

 SetConnectionIdNil sets the value for ConnectionId to be an explicit nil

### UnsetConnectionId
`func (o *Connection) UnsetConnectionId()`

UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
### GetAuthenticationUrl

`func (o *Connection) GetAuthenticationUrl() string`

GetAuthenticationUrl returns the AuthenticationUrl field if non-nil, zero value otherwise.

### GetAuthenticationUrlOk

`func (o *Connection) GetAuthenticationUrlOk() (*string, bool)`

GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationUrl

`func (o *Connection) SetAuthenticationUrl(v string)`

SetAuthenticationUrl sets AuthenticationUrl field to given value.

### HasAuthenticationUrl

`func (o *Connection) HasAuthenticationUrl() bool`

HasAuthenticationUrl returns a boolean if a field has been set.

### GetTokenUrl

`func (o *Connection) GetTokenUrl() string`

GetTokenUrl returns the TokenUrl field if non-nil, zero value otherwise.

### GetTokenUrlOk

`func (o *Connection) GetTokenUrlOk() (*string, bool)`

GetTokenUrlOk returns a tuple with the TokenUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenUrl

`func (o *Connection) SetTokenUrl(v string)`

SetTokenUrl sets TokenUrl field to given value.

### HasTokenUrl

`func (o *Connection) HasTokenUrl() bool`

HasTokenUrl returns a boolean if a field has been set.

### SetTokenUrlNil

`func (o *Connection) SetTokenUrlNil(b bool)`

 SetTokenUrlNil sets the value for TokenUrl to be an explicit nil

### UnsetTokenUrl
`func (o *Connection) UnsetTokenUrl()`

UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
### GetIssuerUrl

`func (o *Connection) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *Connection) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *Connection) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.

### HasIssuerUrl

`func (o *Connection) HasIssuerUrl() bool`

HasIssuerUrl returns a boolean if a field has been set.

### SetIssuerUrlNil

`func (o *Connection) SetIssuerUrlNil(b bool)`

 SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil

### UnsetIssuerUrl
`func (o *Connection) UnsetIssuerUrl()`

UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
### GetProviderCertificate

`func (o *Connection) GetProviderCertificate() string`

GetProviderCertificate returns the ProviderCertificate field if non-nil, zero value otherwise.

### GetProviderCertificateOk

`func (o *Connection) GetProviderCertificateOk() (*string, bool)`

GetProviderCertificateOk returns a tuple with the ProviderCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderCertificate

`func (o *Connection) SetProviderCertificate(v string)`

SetProviderCertificate sets ProviderCertificate field to given value.

### HasProviderCertificate

`func (o *Connection) HasProviderCertificate() bool`

HasProviderCertificate returns a boolean if a field has been set.

### SetProviderCertificateNil

`func (o *Connection) SetProviderCertificateNil(b bool)`

 SetProviderCertificateNil sets the value for ProviderCertificate to be an explicit nil

### UnsetProviderCertificate
`func (o *Connection) UnsetProviderCertificate()`

UnsetProviderCertificate ensures that no value is present for ProviderCertificate, not even an explicit nil
### GetClientId

`func (o *Connection) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Connection) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Connection) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Connection) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *Connection) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *Connection) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetClientSecret

`func (o *Connection) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Connection) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Connection) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Connection) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### SetClientSecretNil

`func (o *Connection) SetClientSecretNil(b bool)`

 SetClientSecretNil sets the value for ClientSecret to be an explicit nil

### UnsetClientSecret
`func (o *Connection) UnsetClientSecret()`

UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
### GetUserDataConfiguration

`func (o *Connection) GetUserDataConfiguration() ConnectionUserDataConfiguration`

GetUserDataConfiguration returns the UserDataConfiguration field if non-nil, zero value otherwise.

### GetUserDataConfigurationOk

`func (o *Connection) GetUserDataConfigurationOk() (*ConnectionUserDataConfiguration, bool)`

GetUserDataConfigurationOk returns a tuple with the UserDataConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDataConfiguration

`func (o *Connection) SetUserDataConfiguration(v ConnectionUserDataConfiguration)`

SetUserDataConfiguration sets UserDataConfiguration field to given value.

### HasUserDataConfiguration

`func (o *Connection) HasUserDataConfiguration() bool`

HasUserDataConfiguration returns a boolean if a field has been set.

### SetUserDataConfigurationNil

`func (o *Connection) SetUserDataConfigurationNil(b bool)`

 SetUserDataConfigurationNil sets the value for UserDataConfiguration to be an explicit nil

### UnsetUserDataConfiguration
`func (o *Connection) UnsetUserDataConfiguration()`

UnsetUserDataConfiguration ensures that no value is present for UserDataConfiguration, not even an explicit nil
### GetData

`func (o *Connection) GetData() ConnectionData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Connection) GetDataOk() (*ConnectionData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Connection) SetData(v ConnectionData)`

SetData sets Data field to given value.

### HasData

`func (o *Connection) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *Connection) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *Connection) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetDefaultConnectionProperties

`func (o *Connection) GetDefaultConnectionProperties() ConnectionDefaultConnectionProperties`

GetDefaultConnectionProperties returns the DefaultConnectionProperties field if non-nil, zero value otherwise.

### GetDefaultConnectionPropertiesOk

`func (o *Connection) GetDefaultConnectionPropertiesOk() (*ConnectionDefaultConnectionProperties, bool)`

GetDefaultConnectionPropertiesOk returns a tuple with the DefaultConnectionProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultConnectionProperties

`func (o *Connection) SetDefaultConnectionProperties(v ConnectionDefaultConnectionProperties)`

SetDefaultConnectionProperties sets DefaultConnectionProperties field to given value.

### HasDefaultConnectionProperties

`func (o *Connection) HasDefaultConnectionProperties() bool`

HasDefaultConnectionProperties returns a boolean if a field has been set.

### SetDefaultConnectionPropertiesNil

`func (o *Connection) SetDefaultConnectionPropertiesNil(b bool)`

 SetDefaultConnectionPropertiesNil sets the value for DefaultConnectionProperties to be an explicit nil

### UnsetDefaultConnectionProperties
`func (o *Connection) UnsetDefaultConnectionProperties()`

UnsetDefaultConnectionProperties ensures that no value is present for DefaultConnectionProperties, not even an explicit nil
### GetCreatedTime

`func (o *Connection) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Connection) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Connection) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Connection) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetTags

`func (o *Connection) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Connection) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Connection) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Connection) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Connection) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Connection) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


