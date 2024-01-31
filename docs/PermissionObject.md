# PermissionObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action the permission grants, can be scoped using &#x60;:&#x60; and parent actions imply sub-resource permissions, action:* or action implies action:sub-action. This property is case-insensitive, it will always be cast to lowercase before comparing actions to user permissions. | 
**Allow** | **bool** | Does this permission grant the user the ability to execute the action? | 
**Grant** | **bool** | Allows the user to give the permission to others without being able to execute the action. | 
**Delegate** | **bool** | Allows delegating or granting the permission to others without being able to execute the action. | 

## Methods

### NewPermissionObject

`func NewPermissionObject(action string, allow bool, grant bool, delegate bool, ) *PermissionObject`

NewPermissionObject instantiates a new PermissionObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionObjectWithDefaults

`func NewPermissionObjectWithDefaults() *PermissionObject`

NewPermissionObjectWithDefaults instantiates a new PermissionObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *PermissionObject) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *PermissionObject) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *PermissionObject) SetAction(v string)`

SetAction sets Action field to given value.


### GetAllow

`func (o *PermissionObject) GetAllow() bool`

GetAllow returns the Allow field if non-nil, zero value otherwise.

### GetAllowOk

`func (o *PermissionObject) GetAllowOk() (*bool, bool)`

GetAllowOk returns a tuple with the Allow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllow

`func (o *PermissionObject) SetAllow(v bool)`

SetAllow sets Allow field to given value.


### GetGrant

`func (o *PermissionObject) GetGrant() bool`

GetGrant returns the Grant field if non-nil, zero value otherwise.

### GetGrantOk

`func (o *PermissionObject) GetGrantOk() (*bool, bool)`

GetGrantOk returns a tuple with the Grant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrant

`func (o *PermissionObject) SetGrant(v bool)`

SetGrant sets Grant field to given value.


### GetDelegate

`func (o *PermissionObject) GetDelegate() bool`

GetDelegate returns the Delegate field if non-nil, zero value otherwise.

### GetDelegateOk

`func (o *PermissionObject) GetDelegateOk() (*bool, bool)`

GetDelegateOk returns a tuple with the Delegate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelegate

`func (o *PermissionObject) SetDelegate(v bool)`

SetDelegate sets Delegate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


