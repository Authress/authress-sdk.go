package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the OAuthAuthorizeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthAuthorizeResponse{}

// OAuthAuthorizeResponse struct for OAuthAuthorizeResponse
type OAuthAuthorizeResponse struct {
	// The authorization code to be used with the /tokens endpoint to retrieve an access_token.
	Code string `json:"code"`
}

type _OAuthAuthorizeResponse OAuthAuthorizeResponse

// NewOAuthAuthorizeResponse instantiates a new OAuthAuthorizeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthAuthorizeResponse(code string) *OAuthAuthorizeResponse {
	this := OAuthAuthorizeResponse{}
	this.Code = code
	return &this
}

// NewOAuthAuthorizeResponseWithDefaults instantiates a new OAuthAuthorizeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthAuthorizeResponseWithDefaults() *OAuthAuthorizeResponse {
	this := OAuthAuthorizeResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *OAuthAuthorizeResponse) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *OAuthAuthorizeResponse) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *OAuthAuthorizeResponse) SetCode(v string) {
	o.Code = v
}

func (o OAuthAuthorizeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthAuthorizeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	return toSerialize, nil
}

func (o *OAuthAuthorizeResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
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

	varOAuthAuthorizeResponse := _OAuthAuthorizeResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuthAuthorizeResponse)

	if err != nil {
		return err
	}

	*o = OAuthAuthorizeResponse(varOAuthAuthorizeResponse)

	return err
}

type NullableOAuthAuthorizeResponse struct {
	value *OAuthAuthorizeResponse
	isSet bool
}

func (v NullableOAuthAuthorizeResponse) Get() *OAuthAuthorizeResponse {
	return v.value
}

func (v *NullableOAuthAuthorizeResponse) Set(val *OAuthAuthorizeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthAuthorizeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthAuthorizeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthAuthorizeResponse(val *OAuthAuthorizeResponse) *NullableOAuthAuthorizeResponse {
	return &NullableOAuthAuthorizeResponse{value: val, isSet: true}
}

func (v NullableOAuthAuthorizeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthAuthorizeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
