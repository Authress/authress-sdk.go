/*
Authress

<p> <h2>Introduction</h2> <p>Welcome to the Authress Authorization API. <br>The Authress REST API provides the operations and resources necessary to create records, assign permissions, and verify any user in your platform.</p> <p><ul>   <li>Manage multitenant platforms and create user tenants for SSO connections.</li>   <li>Create records to assign roles and resources to grant access for users.</li>   <li>Check user access control by calling the authorization API at the right time.</li>   <li>Configure service clients to securely access services in your platform.</li> </ul></p> <p>For more in-depth scenarios check out the <a href=\"https://authress.io/knowledge-base\" target=\"_blank\">Authress knowledge base</a>.</p> </p>

API version: v1
Contact: support@authress.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authress

import (
	"encoding/json"
)

// checks if the TenantConnection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TenantConnection{}

// TenantConnection struct for TenantConnection
type TenantConnection struct {
	ConnectionId NullableString `json:"connectionId,omitempty"`
}

// NewTenantConnection instantiates a new TenantConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantConnection() *TenantConnection {
	this := TenantConnection{}
	return &this
}

// NewTenantConnectionWithDefaults instantiates a new TenantConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantConnectionWithDefaults() *TenantConnection {
	this := TenantConnection{}
	return &this
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TenantConnection) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TenantConnection) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *TenantConnection) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableString and assigns it to the ConnectionId field.
func (o *TenantConnection) SetConnectionId(v string) {
	o.ConnectionId.Set(&v)
}
// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *TenantConnection) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *TenantConnection) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

func (o TenantConnection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TenantConnection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	return toSerialize, nil
}

type NullableTenantConnection struct {
	value *TenantConnection
	isSet bool
}

func (v NullableTenantConnection) Get() *TenantConnection {
	return v.value
}

func (v *NullableTenantConnection) Set(val *TenantConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantConnection(val *TenantConnection) *NullableTenantConnection {
	return &NullableTenantConnection{value: val, isSet: true}
}

func (v NullableTenantConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


