package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ResourcePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcePermission{}

// ResourcePermission struct for ResourcePermission
type ResourcePermission struct {
	Action string `json:"action"`
	Allow  bool   `json:"allow"`
}

type _ResourcePermission ResourcePermission

// NewResourcePermission instantiates a new ResourcePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePermission(action string, allow bool) *ResourcePermission {
	this := ResourcePermission{}
	this.Action = action
	this.Allow = allow
	return &this
}

// NewResourcePermissionWithDefaults instantiates a new ResourcePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePermissionWithDefaults() *ResourcePermission {
	this := ResourcePermission{}
	return &this
}

// GetAction returns the Action field value
func (o *ResourcePermission) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ResourcePermission) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ResourcePermission) SetAction(v string) {
	o.Action = v
}

// GetAllow returns the Allow field value
func (o *ResourcePermission) GetAllow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *ResourcePermission) GetAllowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *ResourcePermission) SetAllow(v bool) {
	o.Allow = v
}

func (o ResourcePermission) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["allow"] = o.Allow
	return toSerialize, nil
}

func (o *ResourcePermission) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"allow",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varResourcePermission := _ResourcePermission{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResourcePermission)

	if err != nil {
		return err
	}

	*o = ResourcePermission(varResourcePermission)

	return err
}

type NullableResourcePermission struct {
	value *ResourcePermission
	isSet bool
}

func (v NullableResourcePermission) Get() *ResourcePermission {
	return v.value
}

func (v *NullableResourcePermission) Set(val *ResourcePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePermission(val *ResourcePermission) *NullableResourcePermission {
	return &NullableResourcePermission{value: val, isSet: true}
}

func (v NullableResourcePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
