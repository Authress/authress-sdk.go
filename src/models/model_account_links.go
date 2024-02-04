package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccountLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountLinks{}

// AccountLinks struct for AccountLinks
type AccountLinks struct {
	Self NullableLink `json:"self"`
}

type _AccountLinks AccountLinks

// NewAccountLinks instantiates a new AccountLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountLinks(self NullableLink) *AccountLinks {
	this := AccountLinks{}
	this.Self = self
	return &this
}

// NewAccountLinksWithDefaults instantiates a new AccountLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountLinksWithDefaults() *AccountLinks {
	this := AccountLinks{}
	return &this
}

// GetSelf returns the Self field value
// If the value is explicit nil, the zero value for Link will be returned
func (o *AccountLinks) GetSelf() Link {
	if o == nil || o.Self.Get() == nil {
		var ret Link
		return ret
	}

	return *o.Self.Get()
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AccountLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return o.Self.Get(), o.Self.IsSet()
}

// SetSelf sets field value
func (o *AccountLinks) SetSelf(v Link) {
	o.Self.Set(&v)
}

func (o AccountLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self.Get()
	return toSerialize, nil
}

func (o *AccountLinks) UnmarshalJSON(data []byte) (err error) {
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

	varAccountLinks := _AccountLinks{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccountLinks)

	if err != nil {
		return err
	}

	*o = AccountLinks(varAccountLinks)

	return err
}

type NullableAccountLinks struct {
	value *AccountLinks
	isSet bool
}

func (v NullableAccountLinks) Get() *AccountLinks {
	return v.value
}

func (v *NullableAccountLinks) Set(val *AccountLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountLinks(val *AccountLinks) *NullableAccountLinks {
	return &NullableAccountLinks{value: val, isSet: true}
}

func (v NullableAccountLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


