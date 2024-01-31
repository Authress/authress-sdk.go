# ResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**Allow** | **bool** |  | 

## Methods

### NewResourcePermission

`func NewResourcePermission(action string, allow bool, ) *ResourcePermission`

NewResourcePermission instantiates a new ResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePermissionWithDefaults

`func NewResourcePermissionWithDefaults() *ResourcePermission`

NewResourcePermissionWithDefaults instantiates a new ResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ResourcePermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ResourcePermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ResourcePermission) SetAction(v string)`

SetAction sets Action field to given value.


### GetAllow

`func (o *ResourcePermission) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *ResourcePermission) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *ResourcePermission) SetAllow(v bool)`

SetAllow sets Allow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


