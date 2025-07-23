package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the AccountCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCollection{}

// AccountCollection struct for AccountCollection
type AccountCollection struct {
	Accounts Account `json:"accounts"`
}

type _AccountCollection AccountCollection

// NewAccountCollection instantiates a new AccountCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCollection(accounts Account) *AccountCollection {
	this := AccountCollection{}
	this.Accounts = accounts
	return &this
}

// NewAccountCollectionWithDefaults instantiates a new AccountCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCollectionWithDefaults() *AccountCollection {
	this := AccountCollection{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *AccountCollection) GetAccounts() Account {
	if o == nil {
		var ret Account
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *AccountCollection) GetAccountsOk() (*Account, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Accounts, true
}

// SetAccounts sets field value
func (o *AccountCollection) SetAccounts(v Account) {
	o.Accounts = v
}

func (o AccountCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accounts"] = o.Accounts
	return toSerialize, nil
}

func (o *AccountCollection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accounts",
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

	varAccountCollection := _AccountCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAccountCollection)

	if err != nil {
		return err
	}

	*o = AccountCollection(varAccountCollection)

	return err
}

type NullableAccountCollection struct {
	value *AccountCollection
	isSet bool
}

func (v NullableAccountCollection) Get() *AccountCollection {
	return v.value
}

func (v *NullableAccountCollection) Set(val *AccountCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCollection(val *AccountCollection) *NullableAccountCollection {
	return &NullableAccountCollection{value: val, isSet: true}
}

func (v NullableAccountCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
