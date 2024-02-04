# AccessRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | New result, either approve or deny the request | 

## Methods

### NewAccessRequestResponse

`func NewAccessRequestResponse(status string, ) *AccessRequestResponse`

NewAccessRequestResponse instantiates a new AccessRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessRequestResponseWithDefaults

`func NewAccessRequestResponseWithDefaults() *AccessRequestResponse`

NewAccessRequestResponseWithDefaults instantiates a new AccessRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AccessRequestResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AccessRequestResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AccessRequestResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](./README.md#documentation-for-models) [[Back to API list]](./README.md#documentation-for-api-endpoints) [[Back to README]](./README.md)


