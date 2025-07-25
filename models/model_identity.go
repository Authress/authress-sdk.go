package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Identity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Identity{}

// Identity The identifying information which uniquely links a JWT OIDC token to an identity provider
type Identity struct {
	// The issuer of the JWT token. This can be any OIDC compliant provider.
	Issuer string `json:"issuer"`
	// The audience of the JWT token. This can be either an audience for your entire app, or one particular audience for it. If there is more than one audience, multiple identities can be created.
	Audience string `json:"audience"`
	// By default, the **sub** claim of the JWT is used to identify the user from this provider. To use an alternate claim or a compound userId resolution, specify an expression. The resolved userId must be the same one that is later used in Authress access records.
	UserIdExpression NullableString `json:"userIdExpression,omitempty"`
}

type _Identity Identity

// NewIdentity instantiates a new Identity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentity(issuer string, audience string) *Identity {
	this := Identity{}
	this.Issuer = issuer
	this.Audience = audience
	var userIdExpression string = "{sub}"
	this.UserIdExpression = *NewNullableString(&userIdExpression)
	return &this
}

// NewIdentityWithDefaults instantiates a new Identity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityWithDefaults() *Identity {
	this := Identity{}
	var userIdExpression string = "{sub}"
	this.UserIdExpression = *NewNullableString(&userIdExpression)
	return &this
}

// GetIssuer returns the Issuer field value
func (o *Identity) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *Identity) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *Identity) SetIssuer(v string) {
	o.Issuer = v
}

// GetAudience returns the Audience field value
func (o *Identity) GetAudience() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Audience
}

// GetAudienceOk returns a tuple with the Audience field value
// and a boolean to check if the value has been set.
func (o *Identity) GetAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Audience, true
}

// SetAudience sets field value
func (o *Identity) SetAudience(v string) {
	o.Audience = v
}

// GetUserIdExpression returns the UserIdExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Identity) GetUserIdExpression() string {
	if o == nil || IsNil(o.UserIdExpression.Get()) {
		var ret string
		return ret
	}
	return *o.UserIdExpression.Get()
}

// GetUserIdExpressionOk returns a tuple with the UserIdExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Identity) GetUserIdExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIdExpression.Get(), o.UserIdExpression.IsSet()
}

// HasUserIdExpression returns a boolean if a field has been set.
func (o *Identity) HasUserIdExpression() bool {
	if o != nil && o.UserIdExpression.IsSet() {
		return true
	}

	return false
}

// SetUserIdExpression gets a reference to the given NullableString and assigns it to the UserIdExpression field.
func (o *Identity) SetUserIdExpression(v string) {
	o.UserIdExpression.Set(&v)
}

// SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil
func (o *Identity) SetUserIdExpressionNil() {
	o.UserIdExpression.Set(nil)
}

// UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil
func (o *Identity) UnsetUserIdExpression() {
	o.UserIdExpression.Unset()
}

func (o Identity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Identity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["issuer"] = o.Issuer
	toSerialize["audience"] = o.Audience
	if o.UserIdExpression.IsSet() {
		toSerialize["userIdExpression"] = o.UserIdExpression.Get()
	}
	return toSerialize, nil
}

func (o *Identity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"issuer",
		"audience",
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

	varIdentity := _Identity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varIdentity)

	if err != nil {
		return err
	}

	*o = Identity(varIdentity)

	return err
}

type NullableIdentity struct {
	value *Identity
	isSet bool
}

func (v NullableIdentity) Get() *Identity {
	return v.value
}

func (v *NullableIdentity) Set(val *Identity) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentity(val *Identity) *NullableIdentity {
	return &NullableIdentity{value: val, isSet: true}
}

func (v NullableIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
