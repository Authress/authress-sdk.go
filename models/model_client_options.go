package models

import (
	"encoding/json"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ClientOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientOptions{}

// ClientOptions A set of client specific options
type ClientOptions struct {
	// Grant the client access to verify authorization on behalf of any user.
	GrantUserPermissionsAccess NullableBool `json:"grantUserPermissionsAccess,omitempty"`
	// Grant the client access to generate oauth tokens on behalf of the Authress account. **Security Warning**: This means that this client can impersonate any user, and should only be used when connecting an existing custom Authorization Server to Authress, when that server does not support a standard OAuth connection.
	GrantTokenGeneration NullableBool `json:"grantTokenGeneration,omitempty"`
}

// NewClientOptions instantiates a new ClientOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientOptions() *ClientOptions {
	this := ClientOptions{}
	var grantUserPermissionsAccess bool = false
	this.GrantUserPermissionsAccess = *NewNullableBool(&grantUserPermissionsAccess)
	var grantTokenGeneration bool = false
	this.GrantTokenGeneration = *NewNullableBool(&grantTokenGeneration)
	return &this
}

// NewClientOptionsWithDefaults instantiates a new ClientOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientOptionsWithDefaults() *ClientOptions {
	this := ClientOptions{}
	var grantUserPermissionsAccess bool = false
	this.GrantUserPermissionsAccess = *NewNullableBool(&grantUserPermissionsAccess)
	var grantTokenGeneration bool = false
	this.GrantTokenGeneration = *NewNullableBool(&grantTokenGeneration)
	return &this
}

// GetGrantUserPermissionsAccess returns the GrantUserPermissionsAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientOptions) GetGrantUserPermissionsAccess() bool {
	if o == nil || IsNil(o.GrantUserPermissionsAccess.Get()) {
		var ret bool
		return ret
	}
	return *o.GrantUserPermissionsAccess.Get()
}

// GetGrantUserPermissionsAccessOk returns a tuple with the GrantUserPermissionsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientOptions) GetGrantUserPermissionsAccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantUserPermissionsAccess.Get(), o.GrantUserPermissionsAccess.IsSet()
}

// HasGrantUserPermissionsAccess returns a boolean if a field has been set.
func (o *ClientOptions) HasGrantUserPermissionsAccess() bool {
	if o != nil && o.GrantUserPermissionsAccess.IsSet() {
		return true
	}

	return false
}

// SetGrantUserPermissionsAccess gets a reference to the given NullableBool and assigns it to the GrantUserPermissionsAccess field.
func (o *ClientOptions) SetGrantUserPermissionsAccess(v bool) {
	o.GrantUserPermissionsAccess.Set(&v)
}

// SetGrantUserPermissionsAccessNil sets the value for GrantUserPermissionsAccess to be an explicit nil
func (o *ClientOptions) SetGrantUserPermissionsAccessNil() {
	o.GrantUserPermissionsAccess.Set(nil)
}

// UnsetGrantUserPermissionsAccess ensures that no value is present for GrantUserPermissionsAccess, not even an explicit nil
func (o *ClientOptions) UnsetGrantUserPermissionsAccess() {
	o.GrantUserPermissionsAccess.Unset()
}

// GetGrantTokenGeneration returns the GrantTokenGeneration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientOptions) GetGrantTokenGeneration() bool {
	if o == nil || IsNil(o.GrantTokenGeneration.Get()) {
		var ret bool
		return ret
	}
	return *o.GrantTokenGeneration.Get()
}

// GetGrantTokenGenerationOk returns a tuple with the GrantTokenGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientOptions) GetGrantTokenGenerationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.GrantTokenGeneration.Get(), o.GrantTokenGeneration.IsSet()
}

// HasGrantTokenGeneration returns a boolean if a field has been set.
func (o *ClientOptions) HasGrantTokenGeneration() bool {
	if o != nil && o.GrantTokenGeneration.IsSet() {
		return true
	}

	return false
}

// SetGrantTokenGeneration gets a reference to the given NullableBool and assigns it to the GrantTokenGeneration field.
func (o *ClientOptions) SetGrantTokenGeneration(v bool) {
	o.GrantTokenGeneration.Set(&v)
}

// SetGrantTokenGenerationNil sets the value for GrantTokenGeneration to be an explicit nil
func (o *ClientOptions) SetGrantTokenGenerationNil() {
	o.GrantTokenGeneration.Set(nil)
}

// UnsetGrantTokenGeneration ensures that no value is present for GrantTokenGeneration, not even an explicit nil
func (o *ClientOptions) UnsetGrantTokenGeneration() {
	o.GrantTokenGeneration.Unset()
}

func (o ClientOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.GrantUserPermissionsAccess.IsSet() {
		toSerialize["grantUserPermissionsAccess"] = o.GrantUserPermissionsAccess.Get()
	}
	if o.GrantTokenGeneration.IsSet() {
		toSerialize["grantTokenGeneration"] = o.GrantTokenGeneration.Get()
	}
	return toSerialize, nil
}

type NullableClientOptions struct {
	value *ClientOptions
	isSet bool
}

func (v NullableClientOptions) Get() *ClientOptions {
	return v.value
}

func (v *NullableClientOptions) Set(val *ClientOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableClientOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableClientOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientOptions(val *ClientOptions) *NullableClientOptions {
	return &NullableClientOptions{value: val, isSet: true}
}

func (v NullableClientOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
