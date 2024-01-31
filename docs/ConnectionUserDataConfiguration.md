# ConnectionUserDataConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to **NullableString** | User data residency - The data residency to store the user specific data in. To ensure high performance and reliability, set to **null**, or to store the user&#39;s data only in one specific region, set the region here. Specifying the region reduces reliability, durability, and performance at the benefit of controlling the locality.  | [optional] 

## Methods

### NewConnectionUserDataConfiguration

`func NewConnectionUserDataConfiguration() *ConnectionUserDataConfiguration`

NewConnectionUserDataConfiguration instantiates a new ConnectionUserDataConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionUserDataConfigurationWithDefaults

`func NewConnectionUserDataConfigurationWithDefaults() *ConnectionUserDataConfiguration`

NewConnectionUserDataConfigurationWithDefaults instantiates a new ConnectionUserDataConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ConnectionUserDataConfiguration) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ConnectionUserDataConfiguration) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ConnectionUserDataConfiguration) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ConnectionUserDataConfiguration) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *ConnectionUserDataConfiguration) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *ConnectionUserDataConfiguration) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


