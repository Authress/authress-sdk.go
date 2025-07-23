package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the AccessRecordAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessRecordAccount{}

// AccessRecordAccount struct for AccessRecordAccount
type AccessRecordAccount struct {
	AccountId string `json:"accountId"`
}

type _AccessRecordAccount AccessRecordAccount

// NewAccessRecordAccount instantiates a new AccessRecordAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRecordAccount(accountId string) *AccessRecordAccount {
	this := AccessRecordAccount{}
	this.AccountId = accountId
	return &this
}

// NewAccessRecordAccountWithDefaults instantiates a new AccessRecordAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRecordAccountWithDefaults() *AccessRecordAccount {
	this := AccessRecordAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *AccessRecordAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *AccessRecordAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *AccessRecordAccount) SetAccountId(v string) {
	o.AccountId = v
}

func (o AccessRecordAccount) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessRecordAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accountId"] = o.AccountId
	return toSerialize, nil
}

func (o *AccessRecordAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accountId",
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

	varAccessRecordAccount := _AccessRecordAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varAccessRecordAccount)

	if err != nil {
		return err
	}

	*o = AccessRecordAccount(varAccessRecordAccount)

	return err
}

type NullableAccessRecordAccount struct {
	value *AccessRecordAccount
	isSet bool
}

func (v NullableAccessRecordAccount) Get() *AccessRecordAccount {
	return v.value
}

func (v *NullableAccessRecordAccount) Set(val *AccessRecordAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRecordAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRecordAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRecordAccount(val *AccessRecordAccount) *NullableAccessRecordAccount {
	return &NullableAccessRecordAccount{value: val, isSet: true}
}

func (v NullableAccessRecordAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRecordAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
