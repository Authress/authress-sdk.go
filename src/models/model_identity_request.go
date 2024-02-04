package authress

import (
	"encoding/json"
)

// checks if the IdentityRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityRequest{}

// IdentityRequest Request to link an identity provider's audience and your app's audience with Authress.
type IdentityRequest struct {
	// A valid JWT OIDC compliant token which will still pass authentication requests to the identity provider. Must contain a unique audience and issuer.
	Jwt NullableString `json:"jwt,omitempty"`
	// The issuer of the OAuth OIDC provider's JWTs. This value should match the `iss` claim in the provided tokens exactly.
	Issuer NullableString `json:"issuer,omitempty"`
	// If the `jwt` token contains more than one valid audience, then the single audience that should associated with Authress. If more than one audience is preferred, repeat this call with each one.
	PreferredAudience NullableString `json:"preferredAudience,omitempty"`
	// By default, the **sub** claim of the JWT is used to identify the user from this provider. To use an alternate claim or a compound userId resolution, specify an expression. The resolved userId must be the same one that is later used in Authress access records.
	UserIdExpression NullableString `json:"userIdExpression,omitempty"`
}

// NewIdentityRequest instantiates a new IdentityRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityRequest() *IdentityRequest {
	this := IdentityRequest{}
	return &this
}

// NewIdentityRequestWithDefaults instantiates a new IdentityRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityRequestWithDefaults() *IdentityRequest {
	this := IdentityRequest{}
	var preferredAudience string = "*"
	this.PreferredAudience = *NewNullableString(&preferredAudience)
	var userIdExpression string = "{sub}"
	this.UserIdExpression = *NewNullableString(&userIdExpression)
	return &this
}

// GetJwt returns the Jwt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRequest) GetJwt() string {
	if o == nil || IsNil(o.Jwt.Get()) {
		var ret string
		return ret
	}
	return *o.Jwt.Get()
}

// GetJwtOk returns a tuple with the Jwt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRequest) GetJwtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Jwt.Get(), o.Jwt.IsSet()
}

// HasJwt returns a boolean if a field has been set.
func (o *IdentityRequest) HasJwt() bool {
	if o != nil && o.Jwt.IsSet() {
		return true
	}

	return false
}

// SetJwt gets a reference to the given NullableString and assigns it to the Jwt field.
func (o *IdentityRequest) SetJwt(v string) {
	o.Jwt.Set(&v)
}
// SetJwtNil sets the value for Jwt to be an explicit nil
func (o *IdentityRequest) SetJwtNil() {
	o.Jwt.Set(nil)
}

// UnsetJwt ensures that no value is present for Jwt, not even an explicit nil
func (o *IdentityRequest) UnsetJwt() {
	o.Jwt.Unset()
}

// GetIssuer returns the Issuer field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRequest) GetIssuer() string {
	if o == nil || IsNil(o.Issuer.Get()) {
		var ret string
		return ret
	}
	return *o.Issuer.Get()
}

// GetIssuerOk returns a tuple with the Issuer field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRequest) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Issuer.Get(), o.Issuer.IsSet()
}

// HasIssuer returns a boolean if a field has been set.
func (o *IdentityRequest) HasIssuer() bool {
	if o != nil && o.Issuer.IsSet() {
		return true
	}

	return false
}

// SetIssuer gets a reference to the given NullableString and assigns it to the Issuer field.
func (o *IdentityRequest) SetIssuer(v string) {
	o.Issuer.Set(&v)
}
// SetIssuerNil sets the value for Issuer to be an explicit nil
func (o *IdentityRequest) SetIssuerNil() {
	o.Issuer.Set(nil)
}

// UnsetIssuer ensures that no value is present for Issuer, not even an explicit nil
func (o *IdentityRequest) UnsetIssuer() {
	o.Issuer.Unset()
}

// GetPreferredAudience returns the PreferredAudience field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRequest) GetPreferredAudience() string {
	if o == nil || IsNil(o.PreferredAudience.Get()) {
		var ret string
		return ret
	}
	return *o.PreferredAudience.Get()
}

// GetPreferredAudienceOk returns a tuple with the PreferredAudience field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRequest) GetPreferredAudienceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreferredAudience.Get(), o.PreferredAudience.IsSet()
}

// HasPreferredAudience returns a boolean if a field has been set.
func (o *IdentityRequest) HasPreferredAudience() bool {
	if o != nil && o.PreferredAudience.IsSet() {
		return true
	}

	return false
}

// SetPreferredAudience gets a reference to the given NullableString and assigns it to the PreferredAudience field.
func (o *IdentityRequest) SetPreferredAudience(v string) {
	o.PreferredAudience.Set(&v)
}
// SetPreferredAudienceNil sets the value for PreferredAudience to be an explicit nil
func (o *IdentityRequest) SetPreferredAudienceNil() {
	o.PreferredAudience.Set(nil)
}

// UnsetPreferredAudience ensures that no value is present for PreferredAudience, not even an explicit nil
func (o *IdentityRequest) UnsetPreferredAudience() {
	o.PreferredAudience.Unset()
}

// GetUserIdExpression returns the UserIdExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IdentityRequest) GetUserIdExpression() string {
	if o == nil || IsNil(o.UserIdExpression.Get()) {
		var ret string
		return ret
	}
	return *o.UserIdExpression.Get()
}

// GetUserIdExpressionOk returns a tuple with the UserIdExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IdentityRequest) GetUserIdExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIdExpression.Get(), o.UserIdExpression.IsSet()
}

// HasUserIdExpression returns a boolean if a field has been set.
func (o *IdentityRequest) HasUserIdExpression() bool {
	if o != nil && o.UserIdExpression.IsSet() {
		return true
	}

	return false
}

// SetUserIdExpression gets a reference to the given NullableString and assigns it to the UserIdExpression field.
func (o *IdentityRequest) SetUserIdExpression(v string) {
	o.UserIdExpression.Set(&v)
}
// SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil
func (o *IdentityRequest) SetUserIdExpressionNil() {
	o.UserIdExpression.Set(nil)
}

// UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil
func (o *IdentityRequest) UnsetUserIdExpression() {
	o.UserIdExpression.Unset()
}

func (o IdentityRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Jwt.IsSet() {
		toSerialize["jwt"] = o.Jwt.Get()
	}
	if o.Issuer.IsSet() {
		toSerialize["issuer"] = o.Issuer.Get()
	}
	if o.PreferredAudience.IsSet() {
		toSerialize["preferredAudience"] = o.PreferredAudience.Get()
	}
	if o.UserIdExpression.IsSet() {
		toSerialize["userIdExpression"] = o.UserIdExpression.Get()
	}
	return toSerialize, nil
}

type NullableIdentityRequest struct {
	value *IdentityRequest
	isSet bool
}

func (v NullableIdentityRequest) Get() *IdentityRequest {
	return v.value
}

func (v *NullableIdentityRequest) Set(val *IdentityRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRequest(val *IdentityRequest) *NullableIdentityRequest {
	return &NullableIdentityRequest{value: val, isSet: true}
}

func (v NullableIdentityRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


