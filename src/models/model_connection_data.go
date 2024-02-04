package authress

import (
	"encoding/json"
)

// checks if the ConnectionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectionData{}

// ConnectionData struct for ConnectionData
type ConnectionData struct {
	TenantId *TenantId `json:"tenantId,omitempty"`
	// The name of this connection when displayed in the Authress management portal
	Name NullableString `json:"name,omitempty"`
	// URL encode OAuth token parameters - Some authentication APIs don't support JSON, in these cases enable the url encoded form data parameters.
	SupportedContentType NullableString `json:"supportedContentType,omitempty"`
	// By default, the **sub** claim of the JWT is used to identify the user from this provider. However, not all providers are OpenID compliant. Here you can provide an optional user data endpoint to fetch additional user profile information and an expression to identify a new user ID from available properties.
	OidcUserEndpointUrl NullableString `json:"oidcUserEndpointUrl,omitempty"`
	// By default, the **sub** claim of the JWT is used to identify the user from this provider. However, not all providers are OpenID compliant. Here you can provide an optional user expression to identify a new user ID from available properties found from the oidcUserEndpointUrl. (The default is **{sub}**, any claims may be used.)
	UserIdExpression NullableString `json:"userIdExpression,omitempty"`
	// Authress generates unique user IDs for every user, however if you trust your identity provider to handle unique ID generate across **ALL customers**, then it is safe to reuse the user ID from the provider.
	TrustIdentityUserId NullableBool `json:"trustIdentityUserId,omitempty"`
}

// NewConnectionData instantiates a new ConnectionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionData() *ConnectionData {
	this := ConnectionData{}
	var supportedContentType SUPPORTED_CONTENT_TYPE = "application/json"
	this.SupportedContentType = *NewNullableString(&supportedContentType)
	var userIdExpression string = "{sub}"
	this.UserIdExpression = *NewNullableString(&userIdExpression)
	var trustIdentityUserId bool = false
	this.TrustIdentityUserId = *NewNullableBool(&trustIdentityUserId)
	return &this
}

// NewConnectionDataWithDefaults instantiates a new ConnectionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionDataWithDefaults() *ConnectionData {
	this := ConnectionData{}
	var supportedContentType SUPPORTED_CONTENT_TYPE = "application/json"
	this.SupportedContentType = *NewNullableString(&supportedContentType)
	var userIdExpression string = "{sub}"
	this.UserIdExpression = *NewNullableString(&userIdExpression)
	var trustIdentityUserId bool = false
	this.TrustIdentityUserId = *NewNullableBool(&trustIdentityUserId)
	return &this
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *ConnectionData) GetTenantId() TenantId {
	if o == nil || IsNil(o.TenantId) {
		var ret TenantId
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionData) GetTenantIdOk() (*TenantId, bool) {
	if o == nil || IsNil(o.TenantId) {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *ConnectionData) HasTenantId() bool {
	if o != nil && !IsNil(o.TenantId) {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given TenantId and assigns it to the TenantId field.
func (o *ConnectionData) SetTenantId(v TenantId) {
	o.TenantId = &v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionData) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *ConnectionData) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *ConnectionData) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *ConnectionData) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *ConnectionData) UnsetName() {
	o.Name.Unset()
}

// GetSupportedContentType returns the SupportedContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionData) GetSupportedContentType() string {
	if o == nil || IsNil(o.SupportedContentType.Get()) {
		var ret string
		return ret
	}
	return *o.SupportedContentType.Get()
}

// GetSupportedContentTypeOk returns a tuple with the SupportedContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionData) GetSupportedContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SupportedContentType.Get(), o.SupportedContentType.IsSet()
}

// HasSupportedContentType returns a boolean if a field has been set.
func (o *ConnectionData) HasSupportedContentType() bool {
	if o != nil && o.SupportedContentType.IsSet() {
		return true
	}

	return false
}

// SetSupportedContentType gets a reference to the given NullableString and assigns it to the SupportedContentType field.
func (o *ConnectionData) SetSupportedContentType(v string) {
	o.SupportedContentType.Set(&v)
}
// SetSupportedContentTypeNil sets the value for SupportedContentType to be an explicit nil
func (o *ConnectionData) SetSupportedContentTypeNil() {
	o.SupportedContentType.Set(nil)
}

// UnsetSupportedContentType ensures that no value is present for SupportedContentType, not even an explicit nil
func (o *ConnectionData) UnsetSupportedContentType() {
	o.SupportedContentType.Unset()
}

// GetOidcUserEndpointUrl returns the OidcUserEndpointUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionData) GetOidcUserEndpointUrl() string {
	if o == nil || IsNil(o.OidcUserEndpointUrl.Get()) {
		var ret string
		return ret
	}
	return *o.OidcUserEndpointUrl.Get()
}

// GetOidcUserEndpointUrlOk returns a tuple with the OidcUserEndpointUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionData) GetOidcUserEndpointUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OidcUserEndpointUrl.Get(), o.OidcUserEndpointUrl.IsSet()
}

// HasOidcUserEndpointUrl returns a boolean if a field has been set.
func (o *ConnectionData) HasOidcUserEndpointUrl() bool {
	if o != nil && o.OidcUserEndpointUrl.IsSet() {
		return true
	}

	return false
}

// SetOidcUserEndpointUrl gets a reference to the given NullableString and assigns it to the OidcUserEndpointUrl field.
func (o *ConnectionData) SetOidcUserEndpointUrl(v string) {
	o.OidcUserEndpointUrl.Set(&v)
}
// SetOidcUserEndpointUrlNil sets the value for OidcUserEndpointUrl to be an explicit nil
func (o *ConnectionData) SetOidcUserEndpointUrlNil() {
	o.OidcUserEndpointUrl.Set(nil)
}

// UnsetOidcUserEndpointUrl ensures that no value is present for OidcUserEndpointUrl, not even an explicit nil
func (o *ConnectionData) UnsetOidcUserEndpointUrl() {
	o.OidcUserEndpointUrl.Unset()
}

// GetUserIdExpression returns the UserIdExpression field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionData) GetUserIdExpression() string {
	if o == nil || IsNil(o.UserIdExpression.Get()) {
		var ret string
		return ret
	}
	return *o.UserIdExpression.Get()
}

// GetUserIdExpressionOk returns a tuple with the UserIdExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionData) GetUserIdExpressionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIdExpression.Get(), o.UserIdExpression.IsSet()
}

// HasUserIdExpression returns a boolean if a field has been set.
func (o *ConnectionData) HasUserIdExpression() bool {
	if o != nil && o.UserIdExpression.IsSet() {
		return true
	}

	return false
}

// SetUserIdExpression gets a reference to the given NullableString and assigns it to the UserIdExpression field.
func (o *ConnectionData) SetUserIdExpression(v string) {
	o.UserIdExpression.Set(&v)
}
// SetUserIdExpressionNil sets the value for UserIdExpression to be an explicit nil
func (o *ConnectionData) SetUserIdExpressionNil() {
	o.UserIdExpression.Set(nil)
}

// UnsetUserIdExpression ensures that no value is present for UserIdExpression, not even an explicit nil
func (o *ConnectionData) UnsetUserIdExpression() {
	o.UserIdExpression.Unset()
}

// GetTrustIdentityUserId returns the TrustIdentityUserId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ConnectionData) GetTrustIdentityUserId() bool {
	if o == nil || IsNil(o.TrustIdentityUserId.Get()) {
		var ret bool
		return ret
	}
	return *o.TrustIdentityUserId.Get()
}

// GetTrustIdentityUserIdOk returns a tuple with the TrustIdentityUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ConnectionData) GetTrustIdentityUserIdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrustIdentityUserId.Get(), o.TrustIdentityUserId.IsSet()
}

// HasTrustIdentityUserId returns a boolean if a field has been set.
func (o *ConnectionData) HasTrustIdentityUserId() bool {
	if o != nil && o.TrustIdentityUserId.IsSet() {
		return true
	}

	return false
}

// SetTrustIdentityUserId gets a reference to the given NullableBool and assigns it to the TrustIdentityUserId field.
func (o *ConnectionData) SetTrustIdentityUserId(v bool) {
	o.TrustIdentityUserId.Set(&v)
}
// SetTrustIdentityUserIdNil sets the value for TrustIdentityUserId to be an explicit nil
func (o *ConnectionData) SetTrustIdentityUserIdNil() {
	o.TrustIdentityUserId.Set(nil)
}

// UnsetTrustIdentityUserId ensures that no value is present for TrustIdentityUserId, not even an explicit nil
func (o *ConnectionData) UnsetTrustIdentityUserId() {
	o.TrustIdentityUserId.Unset()
}

func (o ConnectionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TenantId) {
		toSerialize["tenantId"] = o.TenantId
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.SupportedContentType.IsSet() {
		toSerialize["supportedContentType"] = o.SupportedContentType.Get()
	}
	if o.OidcUserEndpointUrl.IsSet() {
		toSerialize["oidcUserEndpointUrl"] = o.OidcUserEndpointUrl.Get()
	}
	if o.UserIdExpression.IsSet() {
		toSerialize["userIdExpression"] = o.UserIdExpression.Get()
	}
	if o.TrustIdentityUserId.IsSet() {
		toSerialize["trustIdentityUserId"] = o.TrustIdentityUserId.Get()
	}
	return toSerialize, nil
}

type NullableConnectionData struct {
	value *ConnectionData
	isSet bool
}

func (v NullableConnectionData) Get() *ConnectionData {
	return v.value
}

func (v *NullableConnectionData) Set(val *ConnectionData) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionData) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionData(val *ConnectionData) *NullableConnectionData {
	return &NullableConnectionData{value: val, isSet: true}
}

func (v NullableConnectionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


