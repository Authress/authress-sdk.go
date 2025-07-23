package models

import (
	"bytes"
	"encoding/json"
	"fmt"

	. "github.com/authress/authress-sdk.go/utilities"
)

// checks if the Invite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Invite{}

// Invite The user invite used to invite users to your application or to Authress as an admin.
type Invite struct {
	// The unique identifier for the invite. Use this ID to accept the invite. This parameter is ignored during invite creation.
	InviteId string    `json:"inviteId"`
	DefaultLoginTenantId string `json:"string,omitempty"`
	// Specify a User ID that logging in user should receive when login completes. This ID is used to automatically assign a user ID to the user rather than a dynamically generated Authress User ID when using the Authress Login UI SDK. This parameter is ignored when accepting invites directly. Note: If the user logging in has already signed up, then this parameter is ignored.
	LinkedUserId string `json:"string,omitempty"`
	// A list of statements which match roles to resources. The invited user will all statements apply to them when the invite is accepted.
	Statements []InviteStatement `json:"statements"`
	Links *AccountLinks `json:"links,omitempty"`
}

type _Invite Invite

// NewInvite instantiates a new Invite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInvite(inviteId string, statements []InviteStatement) *Invite {
	this := Invite{}
	this.InviteId = inviteId
	this.Statements = statements
	return &this
}

// NewInviteWithDefaults instantiates a new Invite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInviteWithDefaults() *Invite {
	this := Invite{}
	return &this
}

// GetInviteId returns the InviteId field value
func (o *Invite) GetInviteId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InviteId
}

// GetInviteIdOk returns a tuple with the InviteId field value
// and a boolean to check if the value has been set.
func (o *Invite) GetInviteIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InviteId, true
}

// SetInviteId sets field value
func (o *Invite) SetInviteId(v string) {
	o.InviteId = v
}

// GetDefaultLoginTenantId returns the DefaultLoginTenantId field value if set, zero value otherwise.
func (o *Invite) GetDefaultLoginTenantId() string {
	if o == nil || IsNil(o.DefaultLoginTenantId) {
		var ret string
		return ret
	}
	return o.DefaultLoginTenantId
}

// GetDefaultLoginTenantIdOk returns a tuple with the DefaultLoginTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetDefaultLoginTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultLoginTenantId) {
		return nil, false
	}
	return &o.DefaultLoginTenantId, true
}

// HasDefaultLoginTenantId returns a boolean if a field has been set.
func (o *Invite) HasDefaultLoginTenantId() bool {
	if o != nil && !IsNil(o.DefaultLoginTenantId) {
		return true
	}

	return false
}

// SetDefaultLoginTenantId gets a reference to the given DefaultLoginTenantId and assigns it to the DefaultLoginTenantId field.
func (o *Invite) SetDefaultLoginTenantId(v string) {
	o.DefaultLoginTenantId = v
}

// GetLinkedUserId returns the LinkedUserId field value if set, zero value otherwise.
func (o *Invite) GetLinkedUserId() string {
	if o == nil || IsNil(o.LinkedUserId) {
		var ret string
		return ret
	}
	return o.LinkedUserId
}

// GetLinkedUserIdOk returns a tuple with the LinkedUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetLinkedUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkedUserId) {
		return nil, false
	}
	return &o.LinkedUserId, true
}

// HasLinkedUserId returns a boolean if a field has been set.
func (o *Invite) HasLinkedUserId() bool {
	if o != nil && !IsNil(o.LinkedUserId) {
		return true
	}

	return false
}

// SetLinkedUserId gets a reference to the given LinkedUserId and assigns it to the LinkedUserId field.
func (o *Invite) SetLinkedUserId(v string) {
	o.LinkedUserId = v
}

// GetStatements returns the Statements field value
func (o *Invite) GetStatements() []InviteStatement {
	if o == nil {
		var ret []InviteStatement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *Invite) GetStatementsOk() ([]InviteStatement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *Invite) SetStatements(v []InviteStatement) {
	o.Statements = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Invite) GetLinks() AccountLinks {
	if o == nil || IsNil(o.Links) {
		var ret AccountLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Invite) GetLinksOk() (*AccountLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Invite) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AccountLinks and assigns it to the Links field.
func (o *Invite) SetLinks(v AccountLinks) {
	o.Links = &v
}

func (o Invite) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Invite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inviteId"] = o.InviteId
	if !IsNil(o.DefaultLoginTenantId) {
		toSerialize["defaultLoginTenantId"] = o.DefaultLoginTenantId
	}
	if !IsNil(o.LinkedUserId) {
		toSerialize["linkedUserId"] = o.LinkedUserId
	}
	toSerialize["statements"] = o.Statements
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

func (o *Invite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inviteId",
		"statements",
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

	varInvite := _Invite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInvite)

	if err != nil {
		return err
	}

	*o = Invite(varInvite)

	return err
}

type NullableInvite struct {
	value *Invite
	isSet bool
}

func (v NullableInvite) Get() *Invite {
	return v.value
}

func (v *NullableInvite) Set(val *Invite) {
	v.value = val
	v.isSet = true
}

func (v NullableInvite) IsSet() bool {
	return v.isSet
}

func (v *NullableInvite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvite(val *Invite) *NullableInvite {
	return &NullableInvite{value: val, isSet: true}
}

func (v NullableInvite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
