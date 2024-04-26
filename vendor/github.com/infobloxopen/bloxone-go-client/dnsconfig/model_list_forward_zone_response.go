/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
)

// checks if the ListForwardZoneResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListForwardZoneResponse{}

// ListForwardZoneResponse The Forward Zone objects list response format.
type ListForwardZoneResponse struct {
	// List of Forward Zone objects.
	Results              []ForwardZone `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListForwardZoneResponse ListForwardZoneResponse

// NewListForwardZoneResponse instantiates a new ListForwardZoneResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListForwardZoneResponse() *ListForwardZoneResponse {
	this := ListForwardZoneResponse{}
	return &this
}

// NewListForwardZoneResponseWithDefaults instantiates a new ListForwardZoneResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListForwardZoneResponseWithDefaults() *ListForwardZoneResponse {
	this := ListForwardZoneResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListForwardZoneResponse) GetResults() []ForwardZone {
	if o == nil || IsNil(o.Results) {
		var ret []ForwardZone
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListForwardZoneResponse) GetResultsOk() ([]ForwardZone, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListForwardZoneResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ForwardZone and assigns it to the Results field.
func (o *ListForwardZoneResponse) SetResults(v []ForwardZone) {
	o.Results = v
}

func (o ListForwardZoneResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListForwardZoneResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListForwardZoneResponse) UnmarshalJSON(data []byte) (err error) {
	varListForwardZoneResponse := _ListForwardZoneResponse{}

	err = json.Unmarshal(data, &varListForwardZoneResponse)

	if err != nil {
		return err
	}

	*o = ListForwardZoneResponse(varListForwardZoneResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListForwardZoneResponse struct {
	value *ListForwardZoneResponse
	isSet bool
}

func (v NullableListForwardZoneResponse) Get() *ListForwardZoneResponse {
	return v.value
}

func (v *NullableListForwardZoneResponse) Set(val *ListForwardZoneResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListForwardZoneResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListForwardZoneResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListForwardZoneResponse(val *ListForwardZoneResponse) *NullableListForwardZoneResponse {
	return &NullableListForwardZoneResponse{value: val, isSet: true}
}

func (v NullableListForwardZoneResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListForwardZoneResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}