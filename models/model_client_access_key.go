package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ClientAccessKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientAccessKey{}

// ClientAccessKey A client access key configuration. The configuration contains information about the key. On first creation the configuration also contains the raw `clientSecret` and `accessKey` for use with OAuth and the Authress SDKs.
type ClientAccessKey struct {
	// The unique ID of the client.
	KeyId *string `json:"keyId,omitempty"`
	// The unique ID of the client.
	ClientId string `json:"clientId"`
	// Specify a public key on access key creation to bring your own private key. When left blank, Authress will automatically generate a private and public key pair and then return the private key as part of the request. This property holds the public key.
	PublicKey      NullableString `json:"publicKey,omitempty"`
	GenerationDate *time.Time     `json:"generationDate,omitempty"`
	// The unencoded OAuth client secret used with the OAuth endpoints to request a JWT using the `client_credentials` grant type. Pass the clientId and the clientSecret to the documented /tokens endpoint.
	ClientSecret *string `json:"clientSecret,omitempty"`
	// An encoded access key which contains identifying information for client access token creation. For direct use with the Authress SDKs. This private access key must be saved on first creation as it is discarded afterwards. Authress only saves the corresponding public key to verify the private access key.
	AccessKey *string `json:"accessKey,omitempty"`
}

type _ClientAccessKey ClientAccessKey

// NewClientAccessKey instantiates a new ClientAccessKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientAccessKey(clientId string) *ClientAccessKey {
	this := ClientAccessKey{}
	this.ClientId = clientId
	return &this
}

// NewClientAccessKeyWithDefaults instantiates a new ClientAccessKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientAccessKeyWithDefaults() *ClientAccessKey {
	this := ClientAccessKey{}
	return &this
}

// GetKeyId returns the KeyId field value if set, zero value otherwise.
func (o *ClientAccessKey) GetKeyId() string {
	if o == nil || IsNil(o.KeyId) {
		var ret string
		return ret
	}
	return *o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAccessKey) GetKeyIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeyId) {
		return nil, false
	}
	return o.KeyId, true
}

// HasKeyId returns a boolean if a field has been set.
func (o *ClientAccessKey) HasKeyId() bool {
	if o != nil && !IsNil(o.KeyId) {
		return true
	}

	return false
}

// SetKeyId gets a reference to the given string and assigns it to the KeyId field.
func (o *ClientAccessKey) SetKeyId(v string) {
	o.KeyId = &v
}

// GetClientId returns the ClientId field value
func (o *ClientAccessKey) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *ClientAccessKey) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *ClientAccessKey) SetClientId(v string) {
	o.ClientId = v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ClientAccessKey) GetPublicKey() string {
	if o == nil || IsNil(o.PublicKey.Get()) {
		var ret string
		return ret
	}
	return *o.PublicKey.Get()
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClientAccessKey) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublicKey.Get(), o.PublicKey.IsSet()
}

// HasPublicKey returns a boolean if a field has been set.
func (o *ClientAccessKey) HasPublicKey() bool {
	if o != nil && o.PublicKey.IsSet() {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given NullableString and assigns it to the PublicKey field.
func (o *ClientAccessKey) SetPublicKey(v string) {
	o.PublicKey.Set(&v)
}

// SetPublicKeyNil sets the value for PublicKey to be an explicit nil
func (o *ClientAccessKey) SetPublicKeyNil() {
	o.PublicKey.Set(nil)
}

// UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
func (o *ClientAccessKey) UnsetPublicKey() {
	o.PublicKey.Unset()
}

// GetGenerationDate returns the GenerationDate field value if set, zero value otherwise.
func (o *ClientAccessKey) GetGenerationDate() time.Time {
	if o == nil || IsNil(o.GenerationDate) {
		var ret time.Time
		return ret
	}
	return *o.GenerationDate
}

// GetGenerationDateOk returns a tuple with the GenerationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAccessKey) GetGenerationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.GenerationDate) {
		return nil, false
	}
	return o.GenerationDate, true
}

// HasGenerationDate returns a boolean if a field has been set.
func (o *ClientAccessKey) HasGenerationDate() bool {
	if o != nil && !IsNil(o.GenerationDate) {
		return true
	}

	return false
}

// SetGenerationDate gets a reference to the given time.Time and assigns it to the GenerationDate field.
func (o *ClientAccessKey) SetGenerationDate(v time.Time) {
	o.GenerationDate = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *ClientAccessKey) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAccessKey) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// HasClientSecret returns a boolean if a field has been set.
func (o *ClientAccessKey) HasClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *ClientAccessKey) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetAccessKey returns the AccessKey field value if set, zero value otherwise.
func (o *ClientAccessKey) GetAccessKey() string {
	if o == nil || IsNil(o.AccessKey) {
		var ret string
		return ret
	}
	return *o.AccessKey
}

// GetAccessKeyOk returns a tuple with the AccessKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientAccessKey) GetAccessKeyOk() (*string, bool) {
	if o == nil || IsNil(o.AccessKey) {
		return nil, false
	}
	return o.AccessKey, true
}

// HasAccessKey returns a boolean if a field has been set.
func (o *ClientAccessKey) HasAccessKey() bool {
	if o != nil && !IsNil(o.AccessKey) {
		return true
	}

	return false
}

// SetAccessKey gets a reference to the given string and assigns it to the AccessKey field.
func (o *ClientAccessKey) SetAccessKey(v string) {
	o.AccessKey = &v
}

func (o ClientAccessKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientAccessKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeyId) {
		toSerialize["keyId"] = o.KeyId
	}
	toSerialize["clientId"] = o.ClientId
	if o.PublicKey.IsSet() {
		toSerialize["publicKey"] = o.PublicKey.Get()
	}
	if !IsNil(o.GenerationDate) {
		toSerialize["generationDate"] = o.GenerationDate
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.AccessKey) {
		toSerialize["accessKey"] = o.AccessKey
	}
	return toSerialize, nil
}

func (o *ClientAccessKey) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
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

	varClientAccessKey := _ClientAccessKey{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varClientAccessKey)

	if err != nil {
		return err
	}

	*o = ClientAccessKey(varClientAccessKey)

	return err
}

type NullableClientAccessKey struct {
	value *ClientAccessKey
	isSet bool
}

func (v NullableClientAccessKey) Get() *ClientAccessKey {
	return v.value
}

func (v *NullableClientAccessKey) Set(val *ClientAccessKey) {
	v.value = val
	v.isSet = true
}

func (v NullableClientAccessKey) IsSet() bool {
	return v.isSet
}

func (v *NullableClientAccessKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientAccessKey(val *ClientAccessKey) *NullableClientAccessKey {
	return &NullableClientAccessKey{value: val, isSet: true}
}

func (v NullableClientAccessKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientAccessKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
