# ClientAccessKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyId** | Pointer to **string** | The unique ID of the client. | [optional] [readonly] 
**ClientId** | **string** | The unique ID of the client. | [readonly] 
**PublicKey** | Pointer to **NullableString** | Specify a public key on access key creation to bring your own private key. When left blank, Authress will automatically generate a private and public key pair and then return the private key as part of the request. This property holds the public key. | [optional] 
**GenerationDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**ClientSecret** | Pointer to **string** | The unencoded OAuth client secret used with the OAuth endpoints to request a JWT using the &#x60;client_credentials&#x60; grant type. Pass the clientId and the clientSecret to the documented /tokens endpoint. | [optional] [readonly] 
**AccessKey** | Pointer to **string** | An encoded access key which contains identifying information for client access token creation. For direct use with the Authress SDKs. This private access key must be saved on first creation as it is discarded afterwards. Authress only saves the corresponding public key to verify the private access key. | [optional] [readonly] 

## Methods

### NewClientAccessKey

`func NewClientAccessKey(clientId string, ) *ClientAccessKey`

NewClientAccessKey instantiates a new ClientAccessKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAccessKeyWithDefaults

`func NewClientAccessKeyWithDefaults() *ClientAccessKey`

NewClientAccessKeyWithDefaults instantiates a new ClientAccessKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyId

`func (o *ClientAccessKey) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *ClientAccessKey) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *ClientAccessKey) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.

### HasKeyId

`func (o *ClientAccessKey) HasKeyId() bool`

HasKeyId returns a boolean if a field has been set.

### GetClientId

`func (o *ClientAccessKey) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ClientAccessKey) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ClientAccessKey) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetPublicKey

`func (o *ClientAccessKey) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ClientAccessKey) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ClientAccessKey) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ClientAccessKey) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *ClientAccessKey) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *ClientAccessKey) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetGenerationDate

`func (o *ClientAccessKey) GetGenerationDate() time.Time`

GetGenerationDate returns the GenerationDate field if non-nil, zero value otherwise.

### GetGenerationDateOk

`func (o *ClientAccessKey) GetGenerationDateOk() (*time.Time, bool)`

GetGenerationDateOk returns a tuple with the GenerationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerationDate

`func (o *ClientAccessKey) SetGenerationDate(v time.Time)`

SetGenerationDate sets GenerationDate field to given value.

### HasGenerationDate

`func (o *ClientAccessKey) HasGenerationDate() bool`

HasGenerationDate returns a boolean if a field has been set.

### GetClientSecret

`func (o *ClientAccessKey) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ClientAccessKey) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ClientAccessKey) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *ClientAccessKey) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetAccessKey

`func (o *ClientAccessKey) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *ClientAccessKey) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *ClientAccessKey) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *ClientAccessKey) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.


[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


