package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OAuthTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthTokenResponse{}

// OAuthTokenResponse struct for OAuthTokenResponse
type OAuthTokenResponse struct {
	// An expiring access token that can be used to access either Authress or any platform service.
	AccessToken string `json:"access_token"`
}

type _OAuthTokenResponse OAuthTokenResponse

// NewOAuthTokenResponse instantiates a new OAuthTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthTokenResponse(accessToken string) *OAuthTokenResponse {
	this := OAuthTokenResponse{}
	this.AccessToken = accessToken
	return &this
}

// NewOAuthTokenResponseWithDefaults instantiates a new OAuthTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthTokenResponseWithDefaults() *OAuthTokenResponse {
	this := OAuthTokenResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *OAuthTokenResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *OAuthTokenResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

func (o OAuthTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	return toSerialize, nil
}

func (o *OAuthTokenResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
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

	varOAuthTokenResponse := _OAuthTokenResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOAuthTokenResponse)

	if err != nil {
		return err
	}

	*o = OAuthTokenResponse(varOAuthTokenResponse)

	return err
}

type NullableOAuthTokenResponse struct {
	value *OAuthTokenResponse
	isSet bool
}

func (v NullableOAuthTokenResponse) Get() *OAuthTokenResponse {
	return v.value
}

func (v *NullableOAuthTokenResponse) Set(val *OAuthTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthTokenResponse(val *OAuthTokenResponse) *NullableOAuthTokenResponse {
	return &NullableOAuthTokenResponse{value: val, isSet: true}
}

func (v NullableOAuthTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


