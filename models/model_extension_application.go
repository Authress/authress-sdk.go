package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the ExtensionApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionApplication{}

// ExtensionApplication The extension's application configuration. The application contains the necessary information for users to log in to the extension.
type ExtensionApplication struct {
	// The unique ID of the application.
	ApplicationId string   `json:"applicationId"`
	RedirectUrls  []string `json:"redirectUrls,omitempty"`
	Links         *Links   `json:"links,omitempty"`
}

type _ExtensionApplication ExtensionApplication

// NewExtensionApplication instantiates a new ExtensionApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionApplication(applicationId string) *ExtensionApplication {
	this := ExtensionApplication{}
	this.ApplicationId = applicationId
	return &this
}

// NewExtensionApplicationWithDefaults instantiates a new ExtensionApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionApplicationWithDefaults() *ExtensionApplication {
	this := ExtensionApplication{}
	return &this
}

// GetApplicationId returns the ApplicationId field value
func (o *ExtensionApplication) GetApplicationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value
// and a boolean to check if the value has been set.
func (o *ExtensionApplication) GetApplicationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApplicationId, true
}

// SetApplicationId sets field value
func (o *ExtensionApplication) SetApplicationId(v string) {
	o.ApplicationId = v
}

// GetRedirectUrls returns the RedirectUrls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExtensionApplication) GetRedirectUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.RedirectUrls
}

// GetRedirectUrlsOk returns a tuple with the RedirectUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExtensionApplication) GetRedirectUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.RedirectUrls) {
		return nil, false
	}
	return o.RedirectUrls, true
}

// HasRedirectUrls returns a boolean if a field has been set.
func (o *ExtensionApplication) HasRedirectUrls() bool {
	if o != nil && IsNil(o.RedirectUrls) {
		return true
	}

	return false
}

// SetRedirectUrls gets a reference to the given []string and assigns it to the RedirectUrls field.
func (o *ExtensionApplication) SetRedirectUrls(v []string) {
	o.RedirectUrls = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ExtensionApplication) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionApplication) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ExtensionApplication) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *ExtensionApplication) SetLinks(v Links) {
	o.Links = &v
}

func (o ExtensionApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["applicationId"] = o.ApplicationId
	if o.RedirectUrls != nil {
		toSerialize["redirectUrls"] = o.RedirectUrls
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *ExtensionApplication) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"applicationId",
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

	varExtensionApplication := _ExtensionApplication{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtensionApplication)

	if err != nil {
		return err
	}

	*o = ExtensionApplication(varExtensionApplication)

	return err
}

type NullableExtensionApplication struct {
	value *ExtensionApplication
	isSet bool
}

func (v NullableExtensionApplication) Get() *ExtensionApplication {
	return v.value
}

func (v *NullableExtensionApplication) Set(val *ExtensionApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionApplication(val *ExtensionApplication) *NullableExtensionApplication {
	return &NullableExtensionApplication{value: val, isSet: true}
}

func (v NullableExtensionApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
