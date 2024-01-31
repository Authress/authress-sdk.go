# ApplicationDelegation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationUrl** | **string** | Redirect the user to this url to automatically log them into a third-party application. | 

## Methods

### NewApplicationDelegation

`func NewApplicationDelegation(authenticationUrl string, ) *ApplicationDelegation`

NewApplicationDelegation instantiates a new ApplicationDelegation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDelegationWithDefaults

`func NewApplicationDelegationWithDefaults() *ApplicationDelegation`

NewApplicationDelegationWithDefaults instantiates a new ApplicationDelegation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationUrl

`func (o *ApplicationDelegation) GetAuthenticationUrl() string`

GetAuthenticationUrl returns the AuthenticationUrl field if non-nil, zero value otherwise.

### GetAuthenticationUrlOk

`func (o *ApplicationDelegation) GetAuthenticationUrlOk() (*string, bool)`

GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationUrl

`func (o *ApplicationDelegation) SetAuthenticationUrl(v string)`

SetAuthenticationUrl sets AuthenticationUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


