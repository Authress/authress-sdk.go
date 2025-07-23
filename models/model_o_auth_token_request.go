package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the OAuthTokenRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OAuthTokenRequest{}

// OAuthTokenRequest The OAuth 2.1 token request that follows [RFC 6749](https://www.rfc-editor.org/rfc/rfc6749). The properties in the request must be snake_case to follow the standard.
type OAuthTokenRequest struct {
	// The client identifier to constrain the token to.
	ClientId string `json:"client_id"`
	// The secret associated with the client that authorizes the generation of token it's behalf. (Either the `client_secret` or the `code_verifier` is required)
	ClientSecret NullableString `json:"client_secret,omitempty"`
	// The code verifier is the the value used in the generation of the OAuth authorization request `code_challenge` property. (Either the `client_secret` or the `code_verifier` is required)
	CodeVerifier *string `json:"code_verifier,omitempty"`
	// A suggestion to the token generation which type of credentials are being provided.
	GrantType *string `json:"grant_type,omitempty"`
	// When using the user password grant_type, specify the username. Authress recommends this should always be an email address.
	Username NullableString `json:"username,omitempty"`
	// When using the user password grant_type, specify the user's password.
	Password NullableString `json:"password,omitempty"`
	// Enables additional configuration of the grant_type. In the case of user password grant_type, set this to **signup**, to enable the creation of users. Set this to **update**, force update the password associated with the user.
	LoginType NullableString `json:"loginType,omitempty"`
}

type _OAuthTokenRequest OAuthTokenRequest

// NewOAuthTokenRequest instantiates a new OAuthTokenRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOAuthTokenRequest(clientId string) *OAuthTokenRequest {
	this := OAuthTokenRequest{}
	this.ClientId = clientId
	return &this
}

// NewOAuthTokenRequestWithDefaults instantiates a new OAuthTokenRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOAuthTokenRequestWithDefaults() *OAuthTokenRequest {
	this := OAuthTokenRequest{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *OAuthTokenRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OAuthTokenRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OAuthTokenRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthTokenRequest) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthTokenRequest) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *OAuthTokenRequest) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *OAuthTokenRequest) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *OAuthTokenRequest) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetCodeVerifier returns the CodeVerifier field value if set, zero value otherwise.
func (o *OAuthTokenRequest) GetCodeVerifier() string {
	if o == nil || IsNil(o.CodeVerifier) {
		var ret string
		return ret
	}
	return *o.CodeVerifier
}

// GetCodeVerifierOk returns a tuple with the CodeVerifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthTokenRequest) GetCodeVerifierOk() (*string, bool) {
	if o == nil || IsNil(o.CodeVerifier) {
		return nil, false
	}
	return o.CodeVerifier, true
}

// HasCodeVerifier returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasCodeVerifier() bool {
	if o != nil && !IsNil(o.CodeVerifier) {
		return true
	}

	return false
}

// SetCodeVerifier gets a reference to the given string and assigns it to the CodeVerifier field.
func (o *OAuthTokenRequest) SetCodeVerifier(v string) {
	o.CodeVerifier = &v
}

// GetGrantType returns the GrantType field value if set, zero value otherwise.
func (o *OAuthTokenRequest) GetGrantType() string {
	if o == nil || IsNil(o.GrantType) {
		var ret string
		return ret
	}
	return *o.GrantType
}

// GetGrantTypeOk returns a tuple with the GrantType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OAuthTokenRequest) GetGrantTypeOk() (*string, bool) {
	if o == nil || IsNil(o.GrantType) {
		return nil, false
	}
	return o.GrantType, true
}

// HasGrantType returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasGrantType() bool {
	if o != nil && !IsNil(o.GrantType) {
		return true
	}

	return false
}

// SetGrantType gets a reference to the given string and assigns it to the GrantType field.
func (o *OAuthTokenRequest) SetGrantType(v string) {
	o.GrantType = &v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthTokenRequest) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthTokenRequest) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *OAuthTokenRequest) SetUsername(v string) {
	o.Username.Set(&v)
}

// SetUsernameNil sets the value for Username to be an explicit nil
func (o *OAuthTokenRequest) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *OAuthTokenRequest) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthTokenRequest) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthTokenRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *OAuthTokenRequest) SetPassword(v string) {
	o.Password.Set(&v)
}

// SetPasswordNil sets the value for Password to be an explicit nil
func (o *OAuthTokenRequest) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *OAuthTokenRequest) UnsetPassword() {
	o.Password.Unset()
}

// GetLoginType returns the LoginType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *OAuthTokenRequest) GetLoginType() string {
	if o == nil || IsNil(o.LoginType.Get()) {
		var ret string
		return ret
	}
	return *o.LoginType.Get()
}

// GetLoginTypeOk returns a tuple with the LoginType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OAuthTokenRequest) GetLoginTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LoginType.Get(), o.LoginType.IsSet()
}

// HasLoginType returns a boolean if a field has been set.
func (o *OAuthTokenRequest) HasLoginType() bool {
	if o != nil && o.LoginType.IsSet() {
		return true
	}

	return false
}

// SetLoginType gets a reference to the given NullableString and assigns it to the LoginType field.
func (o *OAuthTokenRequest) SetLoginType(v string) {
	o.LoginType.Set(&v)
}

// SetLoginTypeNil sets the value for LoginType to be an explicit nil
func (o *OAuthTokenRequest) SetLoginTypeNil() {
	o.LoginType.Set(nil)
}

// UnsetLoginType ensures that no value is present for LoginType, not even an explicit nil
func (o *OAuthTokenRequest) UnsetLoginType() {
	o.LoginType.Unset()
}

func (o OAuthTokenRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OAuthTokenRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	if o.ClientSecret.IsSet() {
		toSerialize["client_secret"] = o.ClientSecret.Get()
	}
	if !IsNil(o.CodeVerifier) {
		toSerialize["code_verifier"] = o.CodeVerifier
	}
	if !IsNil(o.GrantType) {
		toSerialize["grant_type"] = o.GrantType
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.LoginType.IsSet() {
		toSerialize["loginType"] = o.LoginType.Get()
	}
	return toSerialize, nil
}

func (o *OAuthTokenRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_id",
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

	varOAuthTokenRequest := _OAuthTokenRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varOAuthTokenRequest)

	if err != nil {
		return err
	}

	*o = OAuthTokenRequest(varOAuthTokenRequest)

	return err
}

type NullableOAuthTokenRequest struct {
	value *OAuthTokenRequest
	isSet bool
}

func (v NullableOAuthTokenRequest) Get() *OAuthTokenRequest {
	return v.value
}

func (v *NullableOAuthTokenRequest) Set(val *OAuthTokenRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOAuthTokenRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOAuthTokenRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOAuthTokenRequest(val *OAuthTokenRequest) *NullableOAuthTokenRequest {
	return &NullableOAuthTokenRequest{value: val, isSet: true}
}

func (v NullableOAuthTokenRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOAuthTokenRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
