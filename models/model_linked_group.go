package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the LinkedGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkedGroup{}

// LinkedGroup The referenced group
type LinkedGroup struct {
	GroupId GroupId `json:"groupId"`
}

type _LinkedGroup LinkedGroup

// NewLinkedGroup instantiates a new LinkedGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkedGroup(groupId GroupId) *LinkedGroup {
	this := LinkedGroup{}
	this.GroupId = groupId
	return &this
}

// NewLinkedGroupWithDefaults instantiates a new LinkedGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkedGroupWithDefaults() *LinkedGroup {
	this := LinkedGroup{}
	return &this
}

// GetGroupId returns the GroupId field value
func (o *LinkedGroup) GetGroupId() GroupId {
	if o == nil {
		var ret GroupId
		return ret
	}

	return o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value
// and a boolean to check if the value has been set.
func (o *LinkedGroup) GetGroupIdOk() (*GroupId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GroupId, true
}

// SetGroupId sets field value
func (o *LinkedGroup) SetGroupId(v GroupId) {
	o.GroupId = v
}

func (o LinkedGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LinkedGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupId"] = o.GroupId
	return toSerialize, nil
}

func (o *LinkedGroup) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"groupId",
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

	varLinkedGroup := _LinkedGroup{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varLinkedGroup)

	if err != nil {
		return err
	}

	*o = LinkedGroup(varLinkedGroup)

	return err
}

type NullableLinkedGroup struct {
	value *LinkedGroup
	isSet bool
}

func (v NullableLinkedGroup) Get() *LinkedGroup {
	return v.value
}

func (v *NullableLinkedGroup) Set(val *LinkedGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkedGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkedGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkedGroup(val *LinkedGroup) *NullableLinkedGroup {
	return &NullableLinkedGroup{value: val, isSet: true}
}

func (v NullableLinkedGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinkedGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
