package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the PermissionedResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PermissionedResource{}

// PermissionedResource struct for PermissionedResource
type PermissionedResource struct {
	Permissions []ResourcePermission `json:"permissions"`
}

type _PermissionedResource PermissionedResource

// NewPermissionedResource instantiates a new PermissionedResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPermissionedResource(permissions []ResourcePermission) *PermissionedResource {
	this := PermissionedResource{}
	this.Permissions = permissions
	return &this
}

// NewPermissionedResourceWithDefaults instantiates a new PermissionedResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPermissionedResourceWithDefaults() *PermissionedResource {
	this := PermissionedResource{}
	return &this
}

// GetPermissions returns the Permissions field value
func (o *PermissionedResource) GetPermissions() []ResourcePermission {
	if o == nil {
		var ret []ResourcePermission
		return ret
	}

	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value
// and a boolean to check if the value has been set.
func (o *PermissionedResource) GetPermissionsOk() ([]ResourcePermission, bool) {
	if o == nil {
		return nil, false
	}
	return o.Permissions, true
}

// SetPermissions sets field value
func (o *PermissionedResource) SetPermissions(v []ResourcePermission) {
	o.Permissions = v
}

func (o PermissionedResource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PermissionedResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["permissions"] = o.Permissions
	return toSerialize, nil
}

func (o *PermissionedResource) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"permissions",
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

	varPermissionedResource := _PermissionedResource{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varPermissionedResource)

	if err != nil {
		return err
	}

	*o = PermissionedResource(varPermissionedResource)

	return err
}

type NullablePermissionedResource struct {
	value *PermissionedResource
	isSet bool
}

func (v NullablePermissionedResource) Get() *PermissionedResource {
	return v.value
}

func (v *NullablePermissionedResource) Set(val *PermissionedResource) {
	v.value = val
	v.isSet = true
}

func (v NullablePermissionedResource) IsSet() bool {
	return v.isSet
}

func (v *NullablePermissionedResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePermissionedResource(val *PermissionedResource) *NullablePermissionedResource {
	return &NullablePermissionedResource{value: val, isSet: true}
}

func (v NullablePermissionedResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePermissionedResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
