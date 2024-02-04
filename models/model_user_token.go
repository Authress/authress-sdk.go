package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the UserToken type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserToken{}

// UserToken A JWT token with represents the user.
type UserToken struct {
	Account *PermissionCollectionAccount `json:"account,omitempty"`
	UserId  UserId                       `json:"userId"`
	// The unique identifier for the token
	TokenId string `json:"tokenId"`
	// The access token
	Token string        `json:"token"`
	Links *AccountLinks `json:"links,omitempty"`
}

type _UserToken UserToken

// NewUserToken instantiates a new UserToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserToken(userId UserId, tokenId string, token string) *UserToken {
	this := UserToken{}
	this.UserId = userId
	this.TokenId = tokenId
	this.Token = token
	return &this
}

// NewUserTokenWithDefaults instantiates a new UserToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserTokenWithDefaults() *UserToken {
	this := UserToken{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UserToken) GetAccount() PermissionCollectionAccount {
	if o == nil || IsNil(o.Account) {
		var ret PermissionCollectionAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserToken) GetAccountOk() (*PermissionCollectionAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *UserToken) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given PermissionCollectionAccount and assigns it to the Account field.
func (o *UserToken) SetAccount(v PermissionCollectionAccount) {
	o.Account = &v
}

// GetUserId returns the UserId field value
func (o *UserToken) GetUserId() UserId {
	if o == nil {
		var ret UserId
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserToken) GetUserIdOk() (*UserId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserToken) SetUserId(v UserId) {
	o.UserId = v
}

// GetTokenId returns the TokenId field value
func (o *UserToken) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *UserToken) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *UserToken) SetTokenId(v string) {
	o.TokenId = v
}

// GetToken returns the Token field value
func (o *UserToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *UserToken) GetTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *UserToken) SetToken(v string) {
	o.Token = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *UserToken) GetLinks() AccountLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserToken) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *UserToken) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *UserToken) SetLinks(v AccountLinks) {
	o.Links = &v
}

func (o UserToken) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserToken) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	toSerialize["userId"] = o.UserId
	toSerialize["tokenId"] = o.TokenId
	toSerialize["token"] = o.Token
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *UserToken) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
		"tokenId",
		"token",
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

	varUserToken := _UserToken{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserToken)

	if err != nil {
		return err
	}

	*o = UserToken(varUserToken)

	return err
}

type NullableUserToken struct {
	value *UserToken
	isSet bool
}

func (v NullableUserToken) Get() *UserToken {
	return v.value
}

func (v *NullableUserToken) Set(val *UserToken) {
	v.value = val
	v.isSet = true
}

func (v NullableUserToken) IsSet() bool {
	return v.isSet
}

func (v *NullableUserToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserToken(val *UserToken) *NullableUserToken {
	return &NullableUserToken{value: val, isSet: true}
}

func (v NullableUserToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
