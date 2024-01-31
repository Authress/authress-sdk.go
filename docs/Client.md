# Client

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The unique ID of the client. | [readonly] 
**CreatedTime** | **time.Time** |  | [readonly] 
**Name** | Pointer to **NullableString** | The name of the client | [optional] 
**Options** | Pointer to [**ClientOptions**](ClientOptions.md) |  | [optional] 
**VerificationKeys** | Pointer to [**[]ClientAccessKey**](ClientAccessKey.md) | A list of the service client access keys. | [optional] [readonly] 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewClient

`func NewClient(clientId string, createdTime time.Time, ) *Client`

NewClient instantiates a new Client object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientWithDefaults

`func NewClientWithDefaults() *Client`

NewClientWithDefaults instantiates a new Client object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *Client) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Client) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Client) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetCreatedTime

`func (o *Client) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Client) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Client) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetName

`func (o *Client) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Client) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Client) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Client) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Client) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Client) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptions

`func (o *Client) GetOptions() ClientOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Client) GetOptionsOk() (*ClientOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Client) SetOptions(v ClientOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Client) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetVerificationKeys

`func (o *Client) GetVerificationKeys() []ClientAccessKey`

GetVerificationKeys returns the VerificationKeys field if non-nil, zero value otherwise.

### GetVerificationKeysOk

`func (o *Client) GetVerificationKeysOk() (*[]ClientAccessKey, bool)`

GetVerificationKeysOk returns a tuple with the VerificationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationKeys

`func (o *Client) SetVerificationKeys(v []ClientAccessKey)`

SetVerificationKeys sets VerificationKeys field to given value.

### HasVerificationKeys

`func (o *Client) HasVerificationKeys() bool`

HasVerificationKeys returns a boolean if a field has been set.

### GetTags

`func (o *Client) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Client) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Client) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Client) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Client) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Client) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


