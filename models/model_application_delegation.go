package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ApplicationDelegation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationDelegation{}

// ApplicationDelegation The delegation response.
type ApplicationDelegation struct {
	// Redirect the user to this url to automatically log them into a third-party application.
	AuthenticationUrl string `json:"authenticationUrl"`
}

type _ApplicationDelegation ApplicationDelegation

// NewApplicationDelegation instantiates a new ApplicationDelegation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationDelegation(authenticationUrl string) *ApplicationDelegation {
	this := ApplicationDelegation{}
	this.AuthenticationUrl = authenticationUrl
	return &this
}

// NewApplicationDelegationWithDefaults instantiates a new ApplicationDelegation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationDelegationWithDefaults() *ApplicationDelegation {
	this := ApplicationDelegation{}
	return &this
}

// GetAuthenticationUrl returns the AuthenticationUrl field value
func (o *ApplicationDelegation) GetAuthenticationUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthenticationUrl
}

// GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field value
// and a boolean to check if the value has been set.
func (o *ApplicationDelegation) GetAuthenticationUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthenticationUrl, true
}

// SetAuthenticationUrl sets field value
func (o *ApplicationDelegation) SetAuthenticationUrl(v string) {
	o.AuthenticationUrl = v
}

func (o ApplicationDelegation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationDelegation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authenticationUrl"] = o.AuthenticationUrl
	return toSerialize, nil
}

func (o *ApplicationDelegation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"authenticationUrl",
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

	varApplicationDelegation := _ApplicationDelegation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varApplicationDelegation)

	if err != nil {
		return err
	}

	*o = ApplicationDelegation(varApplicationDelegation)

	return err
}

type NullableApplicationDelegation struct {
	value *ApplicationDelegation
	isSet bool
}

func (v NullableApplicationDelegation) Get() *ApplicationDelegation {
	return v.value
}

func (v *NullableApplicationDelegation) Set(val *ApplicationDelegation) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationDelegation) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationDelegation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationDelegation(val *ApplicationDelegation) *NullableApplicationDelegation {
	return &NullableApplicationDelegation{value: val, isSet: true}
}

func (v NullableApplicationDelegation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationDelegation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
