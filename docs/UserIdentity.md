# UserIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The user identifier. | 
**Name** | Pointer to **string** | The user&#39;s formatted display name. | [optional] 
**Picture** | Pointer to **string** | A url that resolves to a picture that can be rendered. | [optional] 
**Email** | Pointer to **string** | The user&#39;s verified email address sourced from their SSO IdP. | [optional] 

## Methods

### NewUserIdentity

`func NewUserIdentity(userId string, ) *UserIdentity`

NewUserIdentity instantiates a new UserIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserIdentityWithDefaults

`func NewUserIdentityWithDefaults() *UserIdentity`

NewUserIdentityWithDefaults instantiates a new UserIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *UserIdentity) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserIdentity) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserIdentity) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *UserIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPicture

`func (o *UserIdentity) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *UserIdentity) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *UserIdentity) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *UserIdentity) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetEmail

`func (o *UserIdentity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserIdentity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserIdentity) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserIdentity) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


