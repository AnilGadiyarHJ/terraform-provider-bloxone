/*
Host Activation Service

Host activation service provides a RESTful interface to manage cert and join token object. Join tokens are essentially a password that allows on-prem hosts to auto-associate themselves to a customer's account and receive a signed cert.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package infraprovision

import (
	"encoding/json"
)

// checks if the DeleteJoinTokensRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteJoinTokensRequest{}

// DeleteJoinTokensRequest struct for DeleteJoinTokensRequest
type DeleteJoinTokensRequest struct {
	// The resource identifier.
	Ids                  []string `json:"ids,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DeleteJoinTokensRequest DeleteJoinTokensRequest

// NewDeleteJoinTokensRequest instantiates a new DeleteJoinTokensRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteJoinTokensRequest() *DeleteJoinTokensRequest {
	this := DeleteJoinTokensRequest{}
	return &this
}

// NewDeleteJoinTokensRequestWithDefaults instantiates a new DeleteJoinTokensRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteJoinTokensRequestWithDefaults() *DeleteJoinTokensRequest {
	this := DeleteJoinTokensRequest{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteJoinTokensRequest) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteJoinTokensRequest) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteJoinTokensRequest) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteJoinTokensRequest) SetIds(v []string) {
	o.Ids = v
}

func (o DeleteJoinTokensRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteJoinTokensRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DeleteJoinTokensRequest) UnmarshalJSON(data []byte) (err error) {
	varDeleteJoinTokensRequest := _DeleteJoinTokensRequest{}

	err = json.Unmarshal(data, &varDeleteJoinTokensRequest)

	if err != nil {
		return err
	}

	*o = DeleteJoinTokensRequest(varDeleteJoinTokensRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "ids")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDeleteJoinTokensRequest struct {
	value *DeleteJoinTokensRequest
	isSet bool
}

func (v NullableDeleteJoinTokensRequest) Get() *DeleteJoinTokensRequest {
	return v.value
}

func (v *NullableDeleteJoinTokensRequest) Set(val *DeleteJoinTokensRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteJoinTokensRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteJoinTokensRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteJoinTokensRequest(val *DeleteJoinTokensRequest) *NullableDeleteJoinTokensRequest {
	return &NullableDeleteJoinTokensRequest{value: val, isSet: true}
}

func (v NullableDeleteJoinTokensRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteJoinTokensRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
