package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserIdentity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserIdentity{}

// UserIdentity The composite user identity stored in Authress sourced by the customer SSO/SAML/OAuth IdP.
type UserIdentity struct {
	// The user identifier.
	UserId string `json:"userId"`
	// The user's formatted display name.
	Name *string `json:"name,omitempty"`
	// A url that resolves to a picture that can be rendered.
	Picture *string `json:"picture,omitempty"`
	// The user's verified email address sourced from their SSO IdP.
	Email *string `json:"email,omitempty"`
}

type _UserIdentity UserIdentity

// NewUserIdentity instantiates a new UserIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserIdentity(userId string) *UserIdentity {
	this := UserIdentity{}
	this.UserId = userId
	return &this
}

// NewUserIdentityWithDefaults instantiates a new UserIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserIdentityWithDefaults() *UserIdentity {
	this := UserIdentity{}
	return &this
}

// GetUserId returns the UserId field value
func (o *UserIdentity) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *UserIdentity) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *UserIdentity) SetUserId(v string) {
	o.UserId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UserIdentity) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentity) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UserIdentity) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UserIdentity) SetName(v string) {
	o.Name = &v
}

// GetPicture returns the Picture field value if set, zero value otherwise.
func (o *UserIdentity) GetPicture() string {
	if o == nil || IsNil(o.Picture) {
		var ret string
		return ret
	}
	return *o.Picture
}

// GetPictureOk returns a tuple with the Picture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentity) GetPictureOk() (*string, bool) {
	if o == nil || IsNil(o.Picture) {
		return nil, false
	}
	return o.Picture, true
}

// HasPicture returns a boolean if a field has been set.
func (o *UserIdentity) HasPicture() bool {
	if o != nil && !IsNil(o.Picture) {
		return true
	}

	return false
}

// SetPicture gets a reference to the given string and assigns it to the Picture field.
func (o *UserIdentity) SetPicture(v string) {
	o.Picture = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserIdentity) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserIdentity) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserIdentity) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserIdentity) SetEmail(v string) {
	o.Email = &v
}

func (o UserIdentity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserIdentity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["userId"] = o.UserId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Picture) {
		toSerialize["picture"] = o.Picture
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

func (o *UserIdentity) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"userId",
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

	varUserIdentity := _UserIdentity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserIdentity)

	if err != nil {
		return err
	}

	*o = UserIdentity(varUserIdentity)

	return err
}

type NullableUserIdentity struct {
	value *UserIdentity
	isSet bool
}

func (v NullableUserIdentity) Get() *UserIdentity {
	return v.value
}

func (v *NullableUserIdentity) Set(val *UserIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserIdentity(val *UserIdentity) *NullableUserIdentity {
	return &NullableUserIdentity{value: val, isSet: true}
}

func (v NullableUserIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


