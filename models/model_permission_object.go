package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the PermissionObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionObject{}

// PermissionObject The collective action and associate grants on a permission
type PermissionObject struct {
	// The action the permission grants, can be scoped using `:` and parent actions imply sub-resource permissions, action:* or action implies action:sub-action. This property is case-insensitive, it will always be cast to lowercase before comparing actions to user permissions.
	Action string `json:"action"`
	// Does this permission grant the user the ability to execute the action?
	Allow bool `json:"allow"`
	// Allows the user to give the permission to others without being able to execute the action.
	Grant bool `json:"grant"`
	// Allows delegating or granting the permission to others without being able to execute the action.
	Delegate bool `json:"delegate"`
}

type _PermissionObject PermissionObject

// NewPermissionObject instantiates a new PermissionObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionObject(action string, allow bool, grant bool, delegate bool) *PermissionObject {
	this := PermissionObject{}
	this.Action = action
	this.Allow = allow
	this.Grant = grant
	this.Delegate = delegate
	return &this
}

// NewPermissionObjectWithDefaults instantiates a new PermissionObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionObjectWithDefaults() *PermissionObject {
	this := PermissionObject{}
	return &this
}

// GetAction returns the Action field value
func (o *PermissionObject) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PermissionObject) SetAction(v string) {
	o.Action = v
}

// GetAllow returns the Allow field value
func (o *PermissionObject) GetAllow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allow
}

// GetAllowOk returns a tuple with the Allow field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetAllowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allow, true
}

// SetAllow sets field value
func (o *PermissionObject) SetAllow(v bool) {
	o.Allow = v
}

// GetGrant returns the Grant field value
func (o *PermissionObject) GetGrant() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Grant
}

// GetGrantOk returns a tuple with the Grant field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetGrantOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grant, true
}

// SetGrant sets field value
func (o *PermissionObject) SetGrant(v bool) {
	o.Grant = v
}

// GetDelegate returns the Delegate field value
func (o *PermissionObject) GetDelegate() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Delegate
}

// GetDelegateOk returns a tuple with the Delegate field value
// and a boolean to check if the value has been set.
func (o *PermissionObject) GetDelegateOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Delegate, true
}

// SetDelegate sets field value
func (o *PermissionObject) SetDelegate(v bool) {
	o.Delegate = v
}

func (o PermissionObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["allow"] = o.Allow
	toSerialize["grant"] = o.Grant
	toSerialize["delegate"] = o.Delegate
	return toSerialize, nil
}

func (o *PermissionObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"allow",
		"grant",
		"delegate",
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

	varPermissionObject := _PermissionObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPermissionObject)

	if err != nil {
		return err
	}

	*o = PermissionObject(varPermissionObject)

	return err
}

type NullablePermissionObject struct {
	value *PermissionObject
	isSet bool
}

func (v NullablePermissionObject) Get() *PermissionObject {
	return v.value
}

func (v *NullablePermissionObject) Set(val *PermissionObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionObject(val *PermissionObject) *NullablePermissionObject {
	return &NullablePermissionObject{value: val, isSet: true}
}

func (v NullablePermissionObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
