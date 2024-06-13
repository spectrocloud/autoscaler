/*
 * CLOUD API
 *
 * IONOS Enterprise-grade Infrastructure as a Service (IaaS) solutions can be managed through the Cloud API, in addition or as an alternative to the \"Data Center Designer\" (DCD) browser-based tool.    Both methods employ consistent concepts and features, deliver similar power and flexibility, and can be used to perform a multitude of management tasks, including adding servers, volumes, configuring networks, and so on.
 *
 * API version: 6.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// User struct for User
type User struct {
	Entities *UsersEntities `json:"entities,omitempty"`
	// URL to the object representation (absolute path).
	Href *string `json:"href,omitempty"`
	// The resource's unique identifier.
	Id         *string         `json:"id,omitempty"`
	Metadata   *UserMetadata   `json:"metadata,omitempty"`
	Properties *UserProperties `json:"properties"`
	// The type of object that has been created.
	Type *Type `json:"type,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser(properties UserProperties) *User {
	this := User{}

	this.Properties = &properties

	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetEntities returns the Entities field value
// If the value is explicit nil, nil is returned
func (o *User) GetEntities() *UsersEntities {
	if o == nil {
		return nil
	}

	return o.Entities

}

// GetEntitiesOk returns a tuple with the Entities field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetEntitiesOk() (*UsersEntities, bool) {
	if o == nil {
		return nil, false
	}

	return o.Entities, true
}

// SetEntities sets field value
func (o *User) SetEntities(v UsersEntities) {

	o.Entities = &v

}

// HasEntities returns a boolean if a field has been set.
func (o *User) HasEntities() bool {
	if o != nil && o.Entities != nil {
		return true
	}

	return false
}

// GetHref returns the Href field value
// If the value is explicit nil, nil is returned
func (o *User) GetHref() *string {
	if o == nil {
		return nil
	}

	return o.Href

}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Href, true
}

// SetHref sets field value
func (o *User) SetHref(v string) {

	o.Href = &v

}

// HasHref returns a boolean if a field has been set.
func (o *User) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// GetId returns the Id field value
// If the value is explicit nil, nil is returned
func (o *User) GetId() *string {
	if o == nil {
		return nil
	}

	return o.Id

}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Id, true
}

// SetId sets field value
func (o *User) SetId(v string) {

	o.Id = &v

}

// HasId returns a boolean if a field has been set.
func (o *User) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, nil is returned
func (o *User) GetMetadata() *UserMetadata {
	if o == nil {
		return nil
	}

	return o.Metadata

}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetMetadataOk() (*UserMetadata, bool) {
	if o == nil {
		return nil, false
	}

	return o.Metadata, true
}

// SetMetadata sets field value
func (o *User) SetMetadata(v UserMetadata) {

	o.Metadata = &v

}

// HasMetadata returns a boolean if a field has been set.
func (o *User) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// GetProperties returns the Properties field value
// If the value is explicit nil, nil is returned
func (o *User) GetProperties() *UserProperties {
	if o == nil {
		return nil
	}

	return o.Properties

}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetPropertiesOk() (*UserProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Properties, true
}

// SetProperties sets field value
func (o *User) SetProperties(v UserProperties) {

	o.Properties = &v

}

// HasProperties returns a boolean if a field has been set.
func (o *User) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// GetType returns the Type field value
// If the value is explicit nil, nil is returned
func (o *User) GetType() *Type {
	if o == nil {
		return nil
	}

	return o.Type

}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *User) GetTypeOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}

	return o.Type, true
}

// SetType sets field value
func (o *User) SetType(v Type) {

	o.Type = &v

}

// HasType returns a boolean if a field has been set.
func (o *User) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Entities != nil {
		toSerialize["entities"] = o.Entities
	}

	if o.Href != nil {
		toSerialize["href"] = o.Href
	}

	if o.Id != nil {
		toSerialize["id"] = o.Id
	}

	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}

	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
