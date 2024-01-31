# ExtensionApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | **string** | The unique ID of the application. | [readonly] 
**RedirectUrls** | Pointer to **[]string** |  | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 

## Methods

### NewExtensionApplication

`func NewExtensionApplication(applicationId string, ) *ExtensionApplication`

NewExtensionApplication instantiates a new ExtensionApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtensionApplicationWithDefaults

`func NewExtensionApplicationWithDefaults() *ExtensionApplication`

NewExtensionApplicationWithDefaults instantiates a new ExtensionApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ExtensionApplication) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ExtensionApplication) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ExtensionApplication) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.


### GetRedirectUrls

`func (o *ExtensionApplication) GetRedirectUrls() []string`

GetRedirectUrls returns the RedirectUrls field if non-nil, zero value otherwise.

### GetRedirectUrlsOk

`func (o *ExtensionApplication) GetRedirectUrlsOk() (*[]string, bool)`

GetRedirectUrlsOk returns a tuple with the RedirectUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrls

`func (o *ExtensionApplication) SetRedirectUrls(v []string)`

SetRedirectUrls sets RedirectUrls field to given value.

### HasRedirectUrls

`func (o *ExtensionApplication) HasRedirectUrls() bool`

HasRedirectUrls returns a boolean if a field has been set.

### SetRedirectUrlsNil

`func (o *ExtensionApplication) SetRedirectUrlsNil(b bool)`

 SetRedirectUrlsNil sets the value for RedirectUrls to be an explicit nil

### UnsetRedirectUrls
`func (o *ExtensionApplication) UnsetRedirectUrls()`

UnsetRedirectUrls ensures that no value is present for RedirectUrls, not even an explicit nil
### GetLinks

`func (o *ExtensionApplication) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExtensionApplication) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExtensionApplication) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ExtensionApplication) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


