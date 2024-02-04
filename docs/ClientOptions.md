# ClientOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantUserPermissionsAccess** | Pointer to **NullableBool** | Grant the client access to verify authorization on behalf of any user. | [optional] [default to false]
**GrantTokenGeneration** | Pointer to **NullableBool** | Grant the client access to generate oauth tokens on behalf of the Authress account. **Security Warning**: This means that this client can impersonate any user, and should only be used when connecting an existing custom Authorization Server to Authress, when that server does not support a standard OAuth connection. | [optional] [default to false]

## Methods

### NewClientOptions

`func NewClientOptions() *ClientOptions`

NewClientOptions instantiates a new ClientOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientOptionsWithDefaults

`func NewClientOptionsWithDefaults() *ClientOptions`

NewClientOptionsWithDefaults instantiates a new ClientOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantUserPermissionsAccess

`func (o *ClientOptions) GetGrantUserPermissionsAccess() bool`

GetGrantUserPermissionsAccess returns the GrantUserPermissionsAccess field if non-nil, zero value otherwise.

### GetGrantUserPermissionsAccessOk

`func (o *ClientOptions) GetGrantUserPermissionsAccessOk() (*bool, bool)`

GetGrantUserPermissionsAccessOk returns a tuple with the GrantUserPermissionsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantUserPermissionsAccess

`func (o *ClientOptions) SetGrantUserPermissionsAccess(v bool)`

SetGrantUserPermissionsAccess sets GrantUserPermissionsAccess field to given value.

### HasGrantUserPermissionsAccess

`func (o *ClientOptions) HasGrantUserPermissionsAccess() bool`

HasGrantUserPermissionsAccess returns a boolean if a field has been set.

### SetGrantUserPermissionsAccessNil

`func (o *ClientOptions) SetGrantUserPermissionsAccessNil(b bool)`

 SetGrantUserPermissionsAccessNil sets the value for GrantUserPermissionsAccess to be an explicit nil

### UnsetGrantUserPermissionsAccess
`func (o *ClientOptions) UnsetGrantUserPermissionsAccess()`

UnsetGrantUserPermissionsAccess ensures that no value is present for GrantUserPermissionsAccess, not even an explicit nil
### GetGrantTokenGeneration

`func (o *ClientOptions) GetGrantTokenGeneration() bool`

GetGrantTokenGeneration returns the GrantTokenGeneration field if non-nil, zero value otherwise.

### GetGrantTokenGenerationOk

`func (o *ClientOptions) GetGrantTokenGenerationOk() (*bool, bool)`

GetGrantTokenGenerationOk returns a tuple with the GrantTokenGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTokenGeneration

`func (o *ClientOptions) SetGrantTokenGeneration(v bool)`

SetGrantTokenGeneration sets GrantTokenGeneration field to given value.

### HasGrantTokenGeneration

`func (o *ClientOptions) HasGrantTokenGeneration() bool`

HasGrantTokenGeneration returns a boolean if a field has been set.

### SetGrantTokenGenerationNil

`func (o *ClientOptions) SetGrantTokenGenerationNil(b bool)`

 SetGrantTokenGenerationNil sets the value for GrantTokenGeneration to be an explicit nil

### UnsetGrantTokenGeneration
`func (o *ClientOptions) UnsetGrantTokenGeneration()`

UnsetGrantTokenGeneration ensures that no value is present for GrantTokenGeneration, not even an explicit nil

[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


