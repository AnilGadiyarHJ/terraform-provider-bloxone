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

// checks if the ListHostResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHostResponse{}

// ListHostResponse The DNS Host object list response format.
type ListHostResponse struct {
	// List of DNS Host objects.
	Results              []Host `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListHostResponse ListHostResponse

// NewListHostResponse instantiates a new ListHostResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHostResponse() *ListHostResponse {
	this := ListHostResponse{}
	return &this
}

// NewListHostResponseWithDefaults instantiates a new ListHostResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHostResponseWithDefaults() *ListHostResponse {
	this := ListHostResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListHostResponse) GetResults() []Host {
	if o == nil || IsNil(o.Results) {
		var ret []Host
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHostResponse) GetResultsOk() ([]Host, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListHostResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Host and assigns it to the Results field.
func (o *ListHostResponse) SetResults(v []Host) {
	o.Results = v
}

func (o ListHostResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHostResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListHostResponse) UnmarshalJSON(data []byte) (err error) {
	varListHostResponse := _ListHostResponse{}

	err = json.Unmarshal(data, &varListHostResponse)

	if err != nil {
		return err
	}

	*o = ListHostResponse(varListHostResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListHostResponse struct {
	value *ListHostResponse
	isSet bool
}

func (v NullableListHostResponse) Get() *ListHostResponse {
	return v.value
}

func (v *NullableListHostResponse) Set(val *ListHostResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListHostResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListHostResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListHostResponse(val *ListHostResponse) *NullableListHostResponse {
	return &NullableListHostResponse{value: val, isSet: true}
}

func (v NullableListHostResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListHostResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
