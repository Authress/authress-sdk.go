package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserConnectionCredentials type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserConnectionCredentials{}

// UserConnectionCredentials The user credentials for this connection which can be used to access the connection provider APIs.
type UserConnectionCredentials struct {
	// The access token.
	AccessToken string `json:"accessToken"`
}

type _UserConnectionCredentials UserConnectionCredentials

// NewUserConnectionCredentials instantiates a new UserConnectionCredentials object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserConnectionCredentials(accessToken string) *UserConnectionCredentials {
	this := UserConnectionCredentials{}
	this.AccessToken = accessToken
	return &this
}

// NewUserConnectionCredentialsWithDefaults instantiates a new UserConnectionCredentials object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserConnectionCredentialsWithDefaults() *UserConnectionCredentials {
	this := UserConnectionCredentials{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *UserConnectionCredentials) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *UserConnectionCredentials) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *UserConnectionCredentials) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o UserConnectionCredentials) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserConnectionCredentials) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["accessToken"] = o.AccessToken
	return toSerialize, nil
}

func (o *UserConnectionCredentials) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"accessToken",
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

	varUserConnectionCredentials := _UserConnectionCredentials{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserConnectionCredentials)

	if err != nil {
		return err
	}

	*o = UserConnectionCredentials(varUserConnectionCredentials)

	return err
}

type NullableUserConnectionCredentials struct {
	value *UserConnectionCredentials
	isSet bool
}

func (v NullableUserConnectionCredentials) Get() *UserConnectionCredentials {
	return v.value
}

func (v *NullableUserConnectionCredentials) Set(val *UserConnectionCredentials) {
	v.value = val
	v.isSet = true
}

func (v NullableUserConnectionCredentials) IsSet() bool {
	return v.isSet
}

func (v *NullableUserConnectionCredentials) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserConnectionCredentials(val *UserConnectionCredentials) *NullableUserConnectionCredentials {
	return &NullableUserConnectionCredentials{value: val, isSet: true}
}

func (v NullableUserConnectionCredentials) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserConnectionCredentials) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


