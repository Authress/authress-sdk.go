# ExtensionClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | The unique ID of the client. | [readonly] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewExtensionClient

`func NewExtensionClient(clientId string, ) *ExtensionClient`

NewExtensionClient instantiates a new ExtensionClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionClientWithDefaults

`func NewExtensionClientWithDefaults() *ExtensionClient`

NewExtensionClientWithDefaults instantiates a new ExtensionClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ExtensionClient) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ExtensionClient) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ExtensionClient) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetLinks

`func (o *ExtensionClient) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionClient) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionClient) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExtensionClient) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


