package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ExtensionCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionCollection{}

// ExtensionCollection A collection of platform extensions.
type ExtensionCollection struct {
	Extensions []Extension `json:"extensions"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type _ExtensionCollection ExtensionCollection

// NewExtensionCollection instantiates a new ExtensionCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionCollection(extensions []Extension) *ExtensionCollection {
	this := ExtensionCollection{}
	this.Extensions = extensions
	return &this
}

// NewExtensionCollectionWithDefaults instantiates a new ExtensionCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionCollectionWithDefaults() *ExtensionCollection {
	this := ExtensionCollection{}
	return &this
}

// GetExtensions returns the Extensions field value
func (o *ExtensionCollection) GetExtensions() []Extension {
	if o == nil {
		var ret []Extension
		return ret
	}

	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value
// and a boolean to check if the value has been set.
func (o *ExtensionCollection) GetExtensionsOk() ([]Extension, bool) {
	if o == nil {
		return nil, false
	}
	return o.Extensions, true
}

// SetExtensions sets field value
func (o *ExtensionCollection) SetExtensions(v []Extension) {
	o.Extensions = v
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *ExtensionCollection) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionCollection) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *ExtensionCollection) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *ExtensionCollection) SetPagination(v Pagination) {
	o.Pagination = &v
}

func (o ExtensionCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["extensions"] = o.Extensions
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	return toSerialize, nil
}

func (o *ExtensionCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"extensions",
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

	varExtensionCollection := _ExtensionCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varExtensionCollection)

	if err != nil {
		return err
	}

	*o = ExtensionCollection(varExtensionCollection)

	return err
}

type NullableExtensionCollection struct {
	value *ExtensionCollection
	isSet bool
}

func (v NullableExtensionCollection) Get() *ExtensionCollection {
	return v.value
}

func (v *NullableExtensionCollection) Set(val *ExtensionCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionCollection(val *ExtensionCollection) *NullableExtensionCollection {
	return &NullableExtensionCollection{value: val, isSet: true}
}

func (v NullableExtensionCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
