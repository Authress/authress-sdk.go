# Extension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionId** | **string** |  | [readonly] 
**Name** | Pointer to **NullableString** | The name of the extension. This name is visible in the Authress management portal | [optional] 
**CreatedTime** | **time.Time** |  | [readonly] 
**Application** | Pointer to [**ExtensionApplication**](ExtensionApplication.md) |  | [optional] 
**Client** | [**ExtensionClient**](ExtensionClient.md) |  | 
**Tags** | Pointer to **map[string]string** | The tags associated with this resource, this property is an map. { key1: value1, key2: value2 } | [optional] 

## Methods

### NewExtension

`func NewExtension(extensionId string, createdTime time.Time, client ExtensionClient, ) *Extension`

NewExtension instantiates a new Extension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionWithDefaults

`func NewExtensionWithDefaults() *Extension`

NewExtensionWithDefaults instantiates a new Extension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtensionId

`func (o *Extension) GetExtensionId() string`

GetExtensionId returns the ExtensionId field if non-nil, zero value otherwise.

### GetExtensionIdOk

`func (o *Extension) GetExtensionIdOk() (*string, bool)`

GetExtensionIdOk returns a tuple with the ExtensionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionId

`func (o *Extension) SetExtensionId(v string)`

SetExtensionId sets ExtensionId field to given value.


### GetName

`func (o *Extension) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Extension) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Extension) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Extension) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *Extension) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Extension) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedTime

`func (o *Extension) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Extension) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Extension) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.


### GetApplication

`func (o *Extension) GetApplication() ExtensionApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *Extension) GetApplicationOk() (*ExtensionApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *Extension) SetApplication(v ExtensionApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *Extension) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetClient

`func (o *Extension) GetClient() ExtensionClient`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *Extension) GetClientOk() (*ExtensionClient, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *Extension) SetClient(v ExtensionClient)`

SetClient sets Client field to given value.


### GetTags

`func (o *Extension) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Extension) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Extension) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Extension) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *Extension) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *Extension) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


