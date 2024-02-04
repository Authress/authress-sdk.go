package models

import (
	"encoding/json"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Connection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Connection{}

// Connection struct for Connection
type Connection struct {
	Type         *string        `json:"type,omitempty"`
	ConnectionId NullableString `json:"connectionId,omitempty"`
	// Authorization URL of the provider (where the user logs in). For OAuth this is the authorization URL, for SAML, this is the **SSO URL**.
	AuthenticationUrl *string `json:"authenticationUrl,omitempty"`
	// The token exchange url (where we validate the token).
	TokenUrl NullableString `json:"tokenUrl,omitempty"`
	// The unique identifier tied to the provider's domain used for TLS verification. In OAuth, this is the JWT **iss** property. in SAML this is the **Entity ID**.
	IssuerUrl NullableString `json:"issuerUrl,omitempty"`
	// The Provider's SAML public certificate which can be used to verify the signature in signed SAML assertions from the provider.
	ProviderCertificate NullableString `json:"providerCertificate,omitempty"`
	// Provider's client ID used to verify the token
	ClientId NullableString `json:"clientId,omitempty"`
	// Provider's client secret used to verify the token
	ClientSecret                NullableString                                `json:"clientSecret,omitempty"`
	UserDataConfiguration       NullableConnectionUserDataConfiguration       `json:"userDataConfiguration,omitempty"`
	Data                        NullableConnectionData                        `json:"data,omitempty"`
	DefaultConnectionProperties NullableConnectionDefaultConnectionProperties `json:"defaultConnectionProperties,omitempty"`
	CreatedTime                 *time.Time                                    `json:"createdTime,omitempty"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

// NewConnection instantiates a new Connection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnection() *Connection {
	this := Connection{}
	var type_ string = "OAUTH2"
	this.Type = &type_
	return &this
}

// NewConnectionWithDefaults instantiates a new Connection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionWithDefaults() *Connection {
	this := Connection{}
	var type_ string = "OAUTH2"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Connection) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Connection) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Connection) SetType(v string) {
	o.Type = &v
}

// GetConnectionId returns the ConnectionId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetConnectionId() string {
	if o == nil || IsNil(o.ConnectionId.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionId.Get()
}

// GetConnectionIdOk returns a tuple with the ConnectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetConnectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionId.Get(), o.ConnectionId.IsSet()
}

// HasConnectionId returns a boolean if a field has been set.
func (o *Connection) HasConnectionId() bool {
	if o != nil && o.ConnectionId.IsSet() {
		return true
	}

	return false
}

// SetConnectionId gets a reference to the given NullableString and assigns it to the ConnectionId field.
func (o *Connection) SetConnectionId(v string) {
	o.ConnectionId.Set(&v)
}

// SetConnectionIdNil sets the value for ConnectionId to be an explicit nil
func (o *Connection) SetConnectionIdNil() {
	o.ConnectionId.Set(nil)
}

// UnsetConnectionId ensures that no value is present for ConnectionId, not even an explicit nil
func (o *Connection) UnsetConnectionId() {
	o.ConnectionId.Unset()
}

// GetAuthenticationUrl returns the AuthenticationUrl field value if set, zero value otherwise.
func (o *Connection) GetAuthenticationUrl() string {
	if o == nil || IsNil(o.AuthenticationUrl) {
		var ret string
		return ret
	}
	return *o.AuthenticationUrl
}

// GetAuthenticationUrlOk returns a tuple with the AuthenticationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetAuthenticationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AuthenticationUrl) {
		return nil, false
	}
	return o.AuthenticationUrl, true
}

// HasAuthenticationUrl returns a boolean if a field has been set.
func (o *Connection) HasAuthenticationUrl() bool {
	if o != nil && !IsNil(o.AuthenticationUrl) {
		return true
	}

	return false
}

// SetAuthenticationUrl gets a reference to the given string and assigns it to the AuthenticationUrl field.
func (o *Connection) SetAuthenticationUrl(v string) {
	o.AuthenticationUrl = &v
}

// GetTokenUrl returns the TokenUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetTokenUrl() string {
	if o == nil || IsNil(o.TokenUrl.Get()) {
		var ret string
		return ret
	}
	return *o.TokenUrl.Get()
}

// GetTokenUrlOk returns a tuple with the TokenUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetTokenUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenUrl.Get(), o.TokenUrl.IsSet()
}

// HasTokenUrl returns a boolean if a field has been set.
func (o *Connection) HasTokenUrl() bool {
	if o != nil && o.TokenUrl.IsSet() {
		return true
	}

	return false
}

// SetTokenUrl gets a reference to the given NullableString and assigns it to the TokenUrl field.
func (o *Connection) SetTokenUrl(v string) {
	o.TokenUrl.Set(&v)
}

// SetTokenUrlNil sets the value for TokenUrl to be an explicit nil
func (o *Connection) SetTokenUrlNil() {
	o.TokenUrl.Set(nil)
}

// UnsetTokenUrl ensures that no value is present for TokenUrl, not even an explicit nil
func (o *Connection) UnsetTokenUrl() {
	o.TokenUrl.Unset()
}

// GetIssuerUrl returns the IssuerUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetIssuerUrl() string {
	if o == nil || IsNil(o.IssuerUrl.Get()) {
		var ret string
		return ret
	}
	return *o.IssuerUrl.Get()
}

// GetIssuerUrlOk returns a tuple with the IssuerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetIssuerUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IssuerUrl.Get(), o.IssuerUrl.IsSet()
}

// HasIssuerUrl returns a boolean if a field has been set.
func (o *Connection) HasIssuerUrl() bool {
	if o != nil && o.IssuerUrl.IsSet() {
		return true
	}

	return false
}

// SetIssuerUrl gets a reference to the given NullableString and assigns it to the IssuerUrl field.
func (o *Connection) SetIssuerUrl(v string) {
	o.IssuerUrl.Set(&v)
}

// SetIssuerUrlNil sets the value for IssuerUrl to be an explicit nil
func (o *Connection) SetIssuerUrlNil() {
	o.IssuerUrl.Set(nil)
}

// UnsetIssuerUrl ensures that no value is present for IssuerUrl, not even an explicit nil
func (o *Connection) UnsetIssuerUrl() {
	o.IssuerUrl.Unset()
}

// GetProviderCertificate returns the ProviderCertificate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetProviderCertificate() string {
	if o == nil || IsNil(o.ProviderCertificate.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderCertificate.Get()
}

// GetProviderCertificateOk returns a tuple with the ProviderCertificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetProviderCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderCertificate.Get(), o.ProviderCertificate.IsSet()
}

// HasProviderCertificate returns a boolean if a field has been set.
func (o *Connection) HasProviderCertificate() bool {
	if o != nil && o.ProviderCertificate.IsSet() {
		return true
	}

	return false
}

// SetProviderCertificate gets a reference to the given NullableString and assigns it to the ProviderCertificate field.
func (o *Connection) SetProviderCertificate(v string) {
	o.ProviderCertificate.Set(&v)
}

// SetProviderCertificateNil sets the value for ProviderCertificate to be an explicit nil
func (o *Connection) SetProviderCertificateNil() {
	o.ProviderCertificate.Set(nil)
}

// UnsetProviderCertificate ensures that no value is present for ProviderCertificate, not even an explicit nil
func (o *Connection) UnsetProviderCertificate() {
	o.ProviderCertificate.Unset()
}

// GetClientId returns the ClientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetClientId() string {
	if o == nil || IsNil(o.ClientId.Get()) {
		var ret string
		return ret
	}
	return *o.ClientId.Get()
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientId.Get(), o.ClientId.IsSet()
}

// HasClientId returns a boolean if a field has been set.
func (o *Connection) HasClientId() bool {
	if o != nil && o.ClientId.IsSet() {
		return true
	}

	return false
}

// SetClientId gets a reference to the given NullableString and assigns it to the ClientId field.
func (o *Connection) SetClientId(v string) {
	o.ClientId.Set(&v)
}

// SetClientIdNil sets the value for ClientId to be an explicit nil
func (o *Connection) SetClientIdNil() {
	o.ClientId.Set(nil)
}

// UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
func (o *Connection) UnsetClientId() {
	o.ClientId.Unset()
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret.Get()) {
		var ret string
		return ret
	}
	return *o.ClientSecret.Get()
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientSecret.Get(), o.ClientSecret.IsSet()
}

// HasClientSecret returns a boolean if a field has been set.
func (o *Connection) HasClientSecret() bool {
	if o != nil && o.ClientSecret.IsSet() {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given NullableString and assigns it to the ClientSecret field.
func (o *Connection) SetClientSecret(v string) {
	o.ClientSecret.Set(&v)
}

// SetClientSecretNil sets the value for ClientSecret to be an explicit nil
func (o *Connection) SetClientSecretNil() {
	o.ClientSecret.Set(nil)
}

// UnsetClientSecret ensures that no value is present for ClientSecret, not even an explicit nil
func (o *Connection) UnsetClientSecret() {
	o.ClientSecret.Unset()
}

// GetUserDataConfiguration returns the UserDataConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetUserDataConfiguration() ConnectionUserDataConfiguration {
	if o == nil || IsNil(o.UserDataConfiguration.Get()) {
		var ret ConnectionUserDataConfiguration
		return ret
	}
	return *o.UserDataConfiguration.Get()
}

// GetUserDataConfigurationOk returns a tuple with the UserDataConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetUserDataConfigurationOk() (*ConnectionUserDataConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserDataConfiguration.Get(), o.UserDataConfiguration.IsSet()
}

// HasUserDataConfiguration returns a boolean if a field has been set.
func (o *Connection) HasUserDataConfiguration() bool {
	if o != nil && o.UserDataConfiguration.IsSet() {
		return true
	}

	return false
}

// SetUserDataConfiguration gets a reference to the given NullableConnectionUserDataConfiguration and assigns it to the UserDataConfiguration field.
func (o *Connection) SetUserDataConfiguration(v ConnectionUserDataConfiguration) {
	o.UserDataConfiguration.Set(&v)
}

// SetUserDataConfigurationNil sets the value for UserDataConfiguration to be an explicit nil
func (o *Connection) SetUserDataConfigurationNil() {
	o.UserDataConfiguration.Set(nil)
}

// UnsetUserDataConfiguration ensures that no value is present for UserDataConfiguration, not even an explicit nil
func (o *Connection) UnsetUserDataConfiguration() {
	o.UserDataConfiguration.Unset()
}

// GetData returns the Data field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetData() ConnectionData {
	if o == nil || IsNil(o.Data.Get()) {
		var ret ConnectionData
		return ret
	}
	return *o.Data.Get()
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetDataOk() (*ConnectionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data.Get(), o.Data.IsSet()
}

// HasData returns a boolean if a field has been set.
func (o *Connection) HasData() bool {
	if o != nil && o.Data.IsSet() {
		return true
	}

	return false
}

// SetData gets a reference to the given NullableConnectionData and assigns it to the Data field.
func (o *Connection) SetData(v ConnectionData) {
	o.Data.Set(&v)
}

// SetDataNil sets the value for Data to be an explicit nil
func (o *Connection) SetDataNil() {
	o.Data.Set(nil)
}

// UnsetData ensures that no value is present for Data, not even an explicit nil
func (o *Connection) UnsetData() {
	o.Data.Unset()
}

// GetDefaultConnectionProperties returns the DefaultConnectionProperties field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetDefaultConnectionProperties() ConnectionDefaultConnectionProperties {
	if o == nil || IsNil(o.DefaultConnectionProperties.Get()) {
		var ret ConnectionDefaultConnectionProperties
		return ret
	}
	return *o.DefaultConnectionProperties.Get()
}

// GetDefaultConnectionPropertiesOk returns a tuple with the DefaultConnectionProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetDefaultConnectionPropertiesOk() (*ConnectionDefaultConnectionProperties, bool) {
	if o == nil {
		return nil, false
	}
	return o.DefaultConnectionProperties.Get(), o.DefaultConnectionProperties.IsSet()
}

// HasDefaultConnectionProperties returns a boolean if a field has been set.
func (o *Connection) HasDefaultConnectionProperties() bool {
	if o != nil && o.DefaultConnectionProperties.IsSet() {
		return true
	}

	return false
}

// SetDefaultConnectionProperties gets a reference to the given NullableConnectionDefaultConnectionProperties and assigns it to the DefaultConnectionProperties field.
func (o *Connection) SetDefaultConnectionProperties(v ConnectionDefaultConnectionProperties) {
	o.DefaultConnectionProperties.Set(&v)
}

// SetDefaultConnectionPropertiesNil sets the value for DefaultConnectionProperties to be an explicit nil
func (o *Connection) SetDefaultConnectionPropertiesNil() {
	o.DefaultConnectionProperties.Set(nil)
}

// UnsetDefaultConnectionProperties ensures that no value is present for DefaultConnectionProperties, not even an explicit nil
func (o *Connection) UnsetDefaultConnectionProperties() {
	o.DefaultConnectionProperties.Unset()
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Connection) GetCreatedTime() time.Time {
	if o == nil || IsNil(o.CreatedTime) {
		var ret time.Time
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Connection) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Connection) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given time.Time and assigns it to the CreatedTime field.
func (o *Connection) SetCreatedTime(v time.Time) {
	o.CreatedTime = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Connection) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Connection) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Connection) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Connection) SetTags(v map[string]string) {
	o.Tags = v
}

func (o Connection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Connection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.ConnectionId.IsSet() {
		toSerialize["connectionId"] = o.ConnectionId.Get()
	}
	if !IsNil(o.AuthenticationUrl) {
		toSerialize["authenticationUrl"] = o.AuthenticationUrl
	}
	if o.TokenUrl.IsSet() {
		toSerialize["tokenUrl"] = o.TokenUrl.Get()
	}
	if o.IssuerUrl.IsSet() {
		toSerialize["issuerUrl"] = o.IssuerUrl.Get()
	}
	if o.ProviderCertificate.IsSet() {
		toSerialize["providerCertificate"] = o.ProviderCertificate.Get()
	}
	if o.ClientId.IsSet() {
		toSerialize["clientId"] = o.ClientId.Get()
	}
	if o.ClientSecret.IsSet() {
		toSerialize["clientSecret"] = o.ClientSecret.Get()
	}
	if o.UserDataConfiguration.IsSet() {
		toSerialize["userDataConfiguration"] = o.UserDataConfiguration.Get()
	}
	if o.Data.IsSet() {
		toSerialize["data"] = o.Data.Get()
	}
	if o.DefaultConnectionProperties.IsSet() {
		toSerialize["defaultConnectionProperties"] = o.DefaultConnectionProperties.Get()
	}
	if !IsNil(o.CreatedTime) {
		toSerialize["createdTime"] = o.CreatedTime
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableConnection struct {
	value *Connection
	isSet bool
}

func (v NullableConnection) Get() *Connection {
	return v.value
}

func (v *NullableConnection) Set(val *Connection) {
	v.value = val
	v.isSet = true
}

func (v NullableConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnection(val *Connection) *NullableConnection {
	return &NullableConnection{value: val, isSet: true}
}

func (v NullableConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
