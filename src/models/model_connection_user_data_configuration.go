package authress

import (
	"encoding/json"
)

// checks if the ConnectionUserDataConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionUserDataConfiguration{}

// ConnectionUserDataConfiguration struct for ConnectionUserDataConfiguration
type ConnectionUserDataConfiguration struct {
	// User data residency - The data residency to store the user specific data in. To ensure high performance and reliability, set to **null**, or to store the user's data only in one specific region, set the region here. Specifying the region reduces reliability, durability, and performance at the benefit of controlling the locality. 
	Location NullableString `json:"location,omitempty"`
}

// NewConnectionUserDataConfiguration instantiates a new ConnectionUserDataConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionUserDataConfiguration() *ConnectionUserDataConfiguration {
	this := ConnectionUserDataConfiguration{}
	return &this
}

// NewConnectionUserDataConfigurationWithDefaults instantiates a new ConnectionUserDataConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionUserDataConfigurationWithDefaults() *ConnectionUserDataConfiguration {
	this := ConnectionUserDataConfiguration{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionUserDataConfiguration) GetLocation() string {
	if o == nil || IsNil(o.Location.Get()) {
		var ret string
		return ret
	}
	return *o.Location.Get()
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionUserDataConfiguration) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Location.Get(), o.Location.IsSet()
}

// HasLocation returns a boolean if a field has been set.
func (o *ConnectionUserDataConfiguration) HasLocation() bool {
	if o != nil && o.Location.IsSet() {
		return true
	}

	return false
}

// SetLocation gets a reference to the given NullableString and assigns it to the Location field.
func (o *ConnectionUserDataConfiguration) SetLocation(v string) {
	o.Location.Set(&v)
}
// SetLocationNil sets the value for Location to be an explicit nil
func (o *ConnectionUserDataConfiguration) SetLocationNil() {
	o.Location.Set(nil)
}

// UnsetLocation ensures that no value is present for Location, not even an explicit nil
func (o *ConnectionUserDataConfiguration) UnsetLocation() {
	o.Location.Unset()
}

func (o ConnectionUserDataConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionUserDataConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Location.IsSet() {
		toSerialize["location"] = o.Location.Get()
	}
	return toSerialize, nil
}

type NullableConnectionUserDataConfiguration struct {
	value *ConnectionUserDataConfiguration
	isSet bool
}

func (v NullableConnectionUserDataConfiguration) Get() *ConnectionUserDataConfiguration {
	return v.value
}

func (v *NullableConnectionUserDataConfiguration) Set(val *ConnectionUserDataConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionUserDataConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionUserDataConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionUserDataConfiguration(val *ConnectionUserDataConfiguration) *NullableConnectionUserDataConfiguration {
	return &NullableConnectionUserDataConfiguration{value: val, isSet: true}
}

func (v NullableConnectionUserDataConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionUserDataConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


