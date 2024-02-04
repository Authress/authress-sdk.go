package models

import (
	"encoding/json"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ConnectionDefaultConnectionProperties type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionDefaultConnectionProperties{}

// ConnectionDefaultConnectionProperties struct for ConnectionDefaultConnectionProperties
type ConnectionDefaultConnectionProperties struct {
	// Default OAuth scopes to use for every request (leave blank to remove scopes).
	Scope                NullableString `json:"scope,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectionDefaultConnectionProperties ConnectionDefaultConnectionProperties

// NewConnectionDefaultConnectionProperties instantiates a new ConnectionDefaultConnectionProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionDefaultConnectionProperties() *ConnectionDefaultConnectionProperties {
	this := ConnectionDefaultConnectionProperties{}
	var scope string = "profile email openid"
	this.Scope = *NewNullableString(&scope)
	return &this
}

// NewConnectionDefaultConnectionPropertiesWithDefaults instantiates a new ConnectionDefaultConnectionProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDefaultConnectionPropertiesWithDefaults() *ConnectionDefaultConnectionProperties {
	this := ConnectionDefaultConnectionProperties{}
	var scope string = "profile email openid"
	this.Scope = *NewNullableString(&scope)
	return &this
}

// GetScope returns the Scope field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionDefaultConnectionProperties) GetScope() string {
	if o == nil || IsNil(o.Scope.Get()) {
		var ret string
		return ret
	}
	return *o.Scope.Get()
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionDefaultConnectionProperties) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scope.Get(), o.Scope.IsSet()
}

// HasScope returns a boolean if a field has been set.
func (o *ConnectionDefaultConnectionProperties) HasScope() bool {
	if o != nil && o.Scope.IsSet() {
		return true
	}

	return false
}

// SetScope gets a reference to the given NullableString and assigns it to the Scope field.
func (o *ConnectionDefaultConnectionProperties) SetScope(v string) {
	o.Scope.Set(&v)
}

// SetScopeNil sets the value for Scope to be an explicit nil
func (o *ConnectionDefaultConnectionProperties) SetScopeNil() {
	o.Scope.Set(nil)
}

// UnsetScope ensures that no value is present for Scope, not even an explicit nil
func (o *ConnectionDefaultConnectionProperties) UnsetScope() {
	o.Scope.Unset()
}

func (o ConnectionDefaultConnectionProperties) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionDefaultConnectionProperties) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Scope.IsSet() {
		toSerialize["scope"] = o.Scope.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ConnectionDefaultConnectionProperties) UnmarshalJSON(data []byte) (err error) {
	varConnectionDefaultConnectionProperties := _ConnectionDefaultConnectionProperties{}

	err = json.Unmarshal(data, &varConnectionDefaultConnectionProperties)

	if err != nil {
		return err
	}

	*o = ConnectionDefaultConnectionProperties(varConnectionDefaultConnectionProperties)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "scope")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableConnectionDefaultConnectionProperties struct {
	value *ConnectionDefaultConnectionProperties
	isSet bool
}

func (v NullableConnectionDefaultConnectionProperties) Get() *ConnectionDefaultConnectionProperties {
	return v.value
}

func (v *NullableConnectionDefaultConnectionProperties) Set(val *ConnectionDefaultConnectionProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionDefaultConnectionProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionDefaultConnectionProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionDefaultConnectionProperties(val *ConnectionDefaultConnectionProperties) *NullableConnectionDefaultConnectionProperties {
	return &NullableConnectionDefaultConnectionProperties{value: val, isSet: true}
}

func (v NullableConnectionDefaultConnectionProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionDefaultConnectionProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
