package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the IdentityCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityCollection{}

// IdentityCollection struct for IdentityCollection
type IdentityCollection struct {
	Identities []Identity `json:"identities"`
}

type _IdentityCollection IdentityCollection

// NewIdentityCollection instantiates a new IdentityCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityCollection(identities []Identity) *IdentityCollection {
	this := IdentityCollection{}
	this.Identities = identities
	return &this
}

// NewIdentityCollectionWithDefaults instantiates a new IdentityCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityCollectionWithDefaults() *IdentityCollection {
	this := IdentityCollection{}
	return &this
}

// GetIdentities returns the Identities field value
func (o *IdentityCollection) GetIdentities() []Identity {
	if o == nil {
		var ret []Identity
		return ret
	}

	return o.Identities
}

// GetIdentitiesOk returns a tuple with the Identities field value
// and a boolean to check if the value has been set.
func (o *IdentityCollection) GetIdentitiesOk() ([]Identity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Identities, true
}

// SetIdentities sets field value
func (o *IdentityCollection) SetIdentities(v []Identity) {
	o.Identities = v
}

func (o IdentityCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["identities"] = o.Identities
	return toSerialize, nil
}

func (o *IdentityCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"identities",
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

	varIdentityCollection := _IdentityCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIdentityCollection)

	if err != nil {
		return err
	}

	*o = IdentityCollection(varIdentityCollection)

	return err
}

type NullableIdentityCollection struct {
	value *IdentityCollection
	isSet bool
}

func (v NullableIdentityCollection) Get() *IdentityCollection {
	return v.value
}

func (v *NullableIdentityCollection) Set(val *IdentityCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityCollection(val *IdentityCollection) *NullableIdentityCollection {
	return &NullableIdentityCollection{value: val, isSet: true}
}

func (v NullableIdentityCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
