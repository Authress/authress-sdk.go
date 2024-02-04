package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Links type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Links{}

// Links struct for Links
type Links struct {
	Self NullableLink `json:"self"`
}

type _Links Links

// NewLinks instantiates a new Links object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinks(self NullableLink) *Links {
	this := Links{}
	this.Self = self
	return &this
}

// NewLinksWithDefaults instantiates a new Links object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinksWithDefaults() *Links {
	this := Links{}
	return &this
}

// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for Link will be returned
func (o *Links) GetSelf() Link {
	if o == nil || o.Self.Get() == nil {
		var ret Link
		return ret
	}

	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Links) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// SetSelf sets field value
func (o *Links) SetSelf(v Link) {
	o.Self.Set(&v)
}

func (o Links) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Links) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self.Get()
	return toSerialize, nil
}

func (o *Links) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"self",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varLinks := _Links{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLinks)

	if err != nil {
		return err
	}

	*o = Links(varLinks)

	return err
}

type NullableLinks struct {
	value *Links
	isSet bool
}

func (v NullableLinks) Get() *Links {
	return v.value
}

func (v *NullableLinks) Set(val *Links) {
	v.value = val
	v.isSet = true
}

func (v NullableLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinks(val *Links) *NullableLinks {
	return &NullableLinks{value: val, isSet: true}
}

func (v NullableLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


