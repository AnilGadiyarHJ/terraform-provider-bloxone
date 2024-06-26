/*
DNS Data API

The DNS Data is a BloxOne DDI service providing primary authoritative zone support. DNS Data is authoritative for all DNS resource records and is acting as a primary DNS server. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsdata

import (
	"encoding/json"
)

// checks if the ListRecordResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRecordResponse{}

// ListRecordResponse The response format to retrieve __Record__ objects.
type ListRecordResponse struct {
	// The list of Record objects.
	Results              []Record `json:"results,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ListRecordResponse ListRecordResponse

// NewListRecordResponse instantiates a new ListRecordResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRecordResponse() *ListRecordResponse {
	this := ListRecordResponse{}
	return &this
}

// NewListRecordResponseWithDefaults instantiates a new ListRecordResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRecordResponseWithDefaults() *ListRecordResponse {
	this := ListRecordResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ListRecordResponse) GetResults() []Record {
	if o == nil || IsNil(o.Results) {
		var ret []Record
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRecordResponse) GetResultsOk() ([]Record, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ListRecordResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []Record and assigns it to the Results field.
func (o *ListRecordResponse) SetResults(v []Record) {
	o.Results = v
}

func (o ListRecordResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRecordResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ListRecordResponse) UnmarshalJSON(data []byte) (err error) {
	varListRecordResponse := _ListRecordResponse{}

	err = json.Unmarshal(data, &varListRecordResponse)

	if err != nil {
		return err
	}

	*o = ListRecordResponse(varListRecordResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableListRecordResponse struct {
	value *ListRecordResponse
	isSet bool
}

func (v NullableListRecordResponse) Get() *ListRecordResponse {
	return v.value
}

func (v *NullableListRecordResponse) Set(val *ListRecordResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListRecordResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListRecordResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRecordResponse(val *ListRecordResponse) *NullableListRecordResponse {
	return &NullableListRecordResponse{value: val, isSet: true}
}

func (v NullableListRecordResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRecordResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
