package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Client type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Client{}

// Client A client configuration.
type Client struct {
	// The unique ID of the client.
	ClientId    string    `json:"clientId"`
	CreatedTime time.Time `json:"createdTime"`
	// The name of the client
	Name    NullableString `json:"name,omitempty"`
	Options *ClientOptions `json:"options,omitempty"`
	// A list of the service client access keys.
	VerificationKeys []ClientAccessKey `json:"verificationKeys,omitempty"`
	// The tags associated with this resource, this property is an map. { key1: value1, key2: value2 }
	Tags map[string]string `json:"tags,omitempty"`
}

type _Client Client

// NewClient instantiates a new Client object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClient(clientId string, createdTime time.Time) *Client {
	this := Client{}
	this.ClientId = clientId
	this.CreatedTime = createdTime
	return &this
}

// NewClientWithDefaults instantiates a new Client object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientWithDefaults() *Client {
	this := Client{}
	return &this
}

// GetClientId returns the ClientId field value
func (o *Client) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *Client) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *Client) SetClientId(v string) {
	o.ClientId = v
}

// GetCreatedTime returns the CreatedTime field value
func (o *Client) GetCreatedTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value
// and a boolean to check if the value has been set.
func (o *Client) GetCreatedTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedTime, true
}

// SetCreatedTime sets field value
func (o *Client) SetCreatedTime(v time.Time) {
	o.CreatedTime = v
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *Client) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *Client) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *Client) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *Client) UnsetName() {
	o.Name.Unset()
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Client) GetOptions() ClientOptions {
	if o == nil || IsNil(o.Options) {
		var ret ClientOptions
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetOptionsOk() (*ClientOptions, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Client) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given ClientOptions and assigns it to the Options field.
func (o *Client) SetOptions(v ClientOptions) {
	o.Options = &v
}

// GetVerificationKeys returns the VerificationKeys field value if set, zero value otherwise.
func (o *Client) GetVerificationKeys() []ClientAccessKey {
	if o == nil || IsNil(o.VerificationKeys) {
		var ret []ClientAccessKey
		return ret
	}
	return o.VerificationKeys
}

// GetVerificationKeysOk returns a tuple with the VerificationKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Client) GetVerificationKeysOk() ([]ClientAccessKey, bool) {
	if o == nil || IsNil(o.VerificationKeys) {
		return nil, false
	}
	return o.VerificationKeys, true
}

// HasVerificationKeys returns a boolean if a field has been set.
func (o *Client) HasVerificationKeys() bool {
	if o != nil && !IsNil(o.VerificationKeys) {
		return true
	}

	return false
}

// SetVerificationKeys gets a reference to the given []ClientAccessKey and assigns it to the VerificationKeys field.
func (o *Client) SetVerificationKeys(v []ClientAccessKey) {
	o.VerificationKeys = v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Client) GetTags() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Client) GetTagsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Client) HasTags() bool {
	if o != nil && IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *Client) SetTags(v map[string]string) {
	o.Tags = v
}

func (o Client) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Client) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientId"] = o.ClientId
	toSerialize["createdTime"] = o.CreatedTime
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.VerificationKeys) {
		toSerialize["verificationKeys"] = o.VerificationKeys
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *Client) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clientId",
		"createdTime",
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

	varClient := _Client{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	err = decoder.Decode(&varClient)

	if err != nil {
		return err
	}

	*o = Client(varClient)

	return err
}

type NullableClient struct {
	value *Client
	isSet bool
}

func (v NullableClient) Get() *Client {
	return v.value
}

func (v *NullableClient) Set(val *Client) {
	v.value = val
	v.isSet = true
}

func (v NullableClient) IsSet() bool {
	return v.isSet
}

func (v *NullableClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClient(val *Client) *NullableClient {
	return &NullableClient{value: val, isSet: true}
}

func (v NullableClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
