package authress

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the TokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenRequest{}

// TokenRequest struct for TokenRequest
type TokenRequest struct {
	// A list of statements which match roles to resources. The token will have all statements apply to it.
	Statements []Statement `json:"statements"`
	// The ISO8601 datetime when the token will expire. Default is 24 hours from now.
	Expires time.Time `json:"expires"`
}

type _TokenRequest TokenRequest

// NewTokenRequest instantiates a new TokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenRequest(statements []Statement, expires time.Time) *TokenRequest {
	this := TokenRequest{}
	this.Statements = statements
	this.Expires = expires
	return &this
}

// NewTokenRequestWithDefaults instantiates a new TokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenRequestWithDefaults() *TokenRequest {
	this := TokenRequest{}
	return &this
}

// GetStatements returns the Statements field value
func (o *TokenRequest) GetStatements() []Statement {
	if o == nil {
		var ret []Statement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetStatementsOk() ([]Statement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *TokenRequest) SetStatements(v []Statement) {
	o.Statements = v
}

// GetExpires returns the Expires field value
func (o *TokenRequest) GetExpires() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value
// and a boolean to check if the value has been set.
func (o *TokenRequest) GetExpiresOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expires, true
}

// SetExpires sets field value
func (o *TokenRequest) SetExpires(v time.Time) {
	o.Expires = v
}

func (o TokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statements"] = o.Statements
	toSerialize["expires"] = o.Expires
	return toSerialize, nil
}

func (o *TokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statements",
		"expires",
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

	varTokenRequest := _TokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTokenRequest)

	if err != nil {
		return err
	}

	*o = TokenRequest(varTokenRequest)

	return err
}

type NullableTokenRequest struct {
	value *TokenRequest
	isSet bool
}

func (v NullableTokenRequest) Get() *TokenRequest {
	return v.value
}

func (v *NullableTokenRequest) Set(val *TokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenRequest(val *TokenRequest) *NullableTokenRequest {
	return &NullableTokenRequest{value: val, isSet: true}
}

func (v NullableTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


