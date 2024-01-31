/*
Authress

<p> <h2>Introduction</h2> <p>Welcome to the Authress Authorization API. <br>The Authress REST API provides the operations and resources necessary to create records, assign permissions, and verify any user in your platform.</p> <p><ul>   <li>Manage multitenant platforms and create user tenants for SSO connections.</li>   <li>Create records to assign roles and resources to grant access for users.</li>   <li>Check user access control by calling the authorization API at the right time.</li>   <li>Configure service clients to securely access services in your platform.</li> </ul></p> <p>For more in-depth scenarios check out the <a href=\"https://authress.io/knowledge-base\" target=\"_blank\">Authress knowledge base</a>.</p> </p>

API version: v1
Contact: support@authress.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authress

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AccessTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessTemplate{}

// AccessTemplate A logical grouping of access related elements
type AccessTemplate struct {
	// The list of users the access applies to
	Users []User `json:"users"`
	// A list of statements which match roles to resources.
	Statements []Statement `json:"statements"`
}

type _AccessTemplate AccessTemplate

// NewAccessTemplate instantiates a new AccessTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessTemplate(users []User, statements []Statement) *AccessTemplate {
	this := AccessTemplate{}
	this.Users = users
	this.Statements = statements
	return &this
}

// NewAccessTemplateWithDefaults instantiates a new AccessTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessTemplateWithDefaults() *AccessTemplate {
	this := AccessTemplate{}
	return &this
}

// GetUsers returns the Users field value
func (o *AccessTemplate) GetUsers() []User {
	if o == nil {
		var ret []User
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *AccessTemplate) GetUsersOk() ([]User, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *AccessTemplate) SetUsers(v []User) {
	o.Users = v
}

// GetStatements returns the Statements field value
func (o *AccessTemplate) GetStatements() []Statement {
	if o == nil {
		var ret []Statement
		return ret
	}

	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value
// and a boolean to check if the value has been set.
func (o *AccessTemplate) GetStatementsOk() ([]Statement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Statements, true
}

// SetStatements sets field value
func (o *AccessTemplate) SetStatements(v []Statement) {
	o.Statements = v
}

func (o AccessTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["users"] = o.Users
	toSerialize["statements"] = o.Statements
	return toSerialize, nil
}

func (o *AccessTemplate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
		"statements",
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

	varAccessTemplate := _AccessTemplate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAccessTemplate)

	if err != nil {
		return err
	}

	*o = AccessTemplate(varAccessTemplate)

	return err
}

type NullableAccessTemplate struct {
	value *AccessTemplate
	isSet bool
}

func (v NullableAccessTemplate) Get() *AccessTemplate {
	return v.value
}

func (v *NullableAccessTemplate) Set(val *AccessTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTemplate(val *AccessTemplate) *NullableAccessTemplate {
	return &NullableAccessTemplate{value: val, isSet: true}
}

func (v NullableAccessTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


