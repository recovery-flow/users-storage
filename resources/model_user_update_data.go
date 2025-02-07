/*
User storage service

User storage service for recovery flow

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserUpdateData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserUpdateData{}

// UserUpdateData struct for UserUpdateData
type UserUpdateData struct {
	// User ID
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes UserUpdateDataAttributes `json:"attributes"`
}

type _UserUpdateData UserUpdateData

// NewUserUpdateData instantiates a new UserUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateData(id string, type_ string, attributes UserUpdateDataAttributes) *UserUpdateData {
	this := UserUpdateData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewUserUpdateDataWithDefaults instantiates a new UserUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateDataWithDefaults() *UserUpdateData {
	this := UserUpdateData{}
	return &this
}

// GetId returns the Id field value
func (o *UserUpdateData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserUpdateData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *UserUpdateData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *UserUpdateData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *UserUpdateData) GetAttributes() UserUpdateDataAttributes {
	if o == nil {
		var ret UserUpdateDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetAttributesOk() (*UserUpdateDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *UserUpdateData) SetAttributes(v UserUpdateDataAttributes) {
	o.Attributes = v
}

func (o UserUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserUpdateData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *UserUpdateData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
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

	varUserUpdateData := _UserUpdateData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserUpdateData)

	if err != nil {
		return err
	}

	*o = UserUpdateData(varUserUpdateData)

	return err
}

type NullableUserUpdateData struct {
	value *UserUpdateData
	isSet bool
}

func (v NullableUserUpdateData) Get() *UserUpdateData {
	return v.value
}

func (v *NullableUserUpdateData) Set(val *UserUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateData(val *UserUpdateData) *NullableUserUpdateData {
	return &NullableUserUpdateData{value: val, isSet: true}
}

func (v NullableUserUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


