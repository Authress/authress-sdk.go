# PermissionCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**PermissionCollectionAccount**](PermissionCollectionAccount.md) |  | [optional] 
**UserId** | [**UserId**](UserId.md) |  | 
**Permissions** | [**[]PermissionObject**](PermissionObject.md) | A list of the permissions | 

## Methods

### NewPermissionCollection

`func NewPermissionCollection(userId UserId, permissions []PermissionObject, ) *PermissionCollection`

NewPermissionCollection instantiates a new PermissionCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionCollectionWithDefaults

`func NewPermissionCollectionWithDefaults() *PermissionCollection`

NewPermissionCollectionWithDefaults instantiates a new PermissionCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PermissionCollection) GetAccount() PermissionCollectionAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PermissionCollection) GetAccountOk() (*PermissionCollectionAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PermissionCollection) SetAccount(v PermissionCollectionAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PermissionCollection) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetUserId

`func (o *PermissionCollection) GetUserId() UserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PermissionCollection) GetUserIdOk() (*UserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PermissionCollection) SetUserId(v UserId)`

SetUserId sets UserId field to given value.


### GetPermissions

`func (o *PermissionCollection) GetPermissions() []PermissionObject`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionCollection) GetPermissionsOk() (*[]PermissionObject, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionCollection) SetPermissions(v []PermissionObject)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


