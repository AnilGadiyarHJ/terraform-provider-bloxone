/*
BloxOne FW API

BloxOne Threat Defense Cloud is an extension of the BloxOne Suite that provides visibility into infected and compromised off-premises devices, roaming users, remote sites, and branch offices. You can subscribe to BloxOne Cloud and use its functionality to mitigate and control malware as well as provide unprecedented insight into your network security posture and enable timely action. BloxOne Cloud also offers unified policy management, reporting, and threat analytics across the entire spectrum. Using automated and high-quality threat intelligence feeds and unique behavioral analytics, it automatically stops device communications with C&Cs/botnets and prevents DNS based data exfiltration.  The mission-critical DNS infrastructure can become a vulnerable component in your network when it is inadequately protected by traditional security solutions and consequently used as an attack surface. Compromised DNS services can result in catastrophic network and system failures. To fully protect your network in today’s cyber security threat environment, Infoblox sets a new DNS security standard by offering scalable, enterprise-grade, and integrated protection for your DNS infrastructure.  Through the Infoblox Cloud Services Portal, you can view the status of your subscription and threat intelligence feeds, manage your network scope and roaming end users, and learn more about threats on your networks through the Infoblox Threat Lookup tool and predefined reports.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package fw

import (
	"encoding/json"
)

// checks if the AtcfwReadPoPRegionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AtcfwReadPoPRegionResponse{}

// AtcfwReadPoPRegionResponse struct for AtcfwReadPoPRegionResponse
type AtcfwReadPoPRegionResponse struct {
	Result *AtcfwPoPRegion `json:"result,omitempty"`
}

// NewAtcfwReadPoPRegionResponse instantiates a new AtcfwReadPoPRegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAtcfwReadPoPRegionResponse() *AtcfwReadPoPRegionResponse {
	this := AtcfwReadPoPRegionResponse{}
	return &this
}

// NewAtcfwReadPoPRegionResponseWithDefaults instantiates a new AtcfwReadPoPRegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtcfwReadPoPRegionResponseWithDefaults() *AtcfwReadPoPRegionResponse {
	this := AtcfwReadPoPRegionResponse{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *AtcfwReadPoPRegionResponse) GetResult() AtcfwPoPRegion {
	if o == nil || IsNil(o.Result) {
		var ret AtcfwPoPRegion
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AtcfwReadPoPRegionResponse) GetResultOk() (*AtcfwPoPRegion, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *AtcfwReadPoPRegionResponse) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given AtcfwPoPRegion and assigns it to the Result field.
func (o *AtcfwReadPoPRegionResponse) SetResult(v AtcfwPoPRegion) {
	o.Result = &v
}

func (o AtcfwReadPoPRegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AtcfwReadPoPRegionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	return toSerialize, nil
}

type NullableAtcfwReadPoPRegionResponse struct {
	value *AtcfwReadPoPRegionResponse
	isSet bool
}

func (v NullableAtcfwReadPoPRegionResponse) Get() *AtcfwReadPoPRegionResponse {
	return v.value
}

func (v *NullableAtcfwReadPoPRegionResponse) Set(val *AtcfwReadPoPRegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAtcfwReadPoPRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAtcfwReadPoPRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAtcfwReadPoPRegionResponse(val *AtcfwReadPoPRegionResponse) *NullableAtcfwReadPoPRegionResponse {
	return &NullableAtcfwReadPoPRegionResponse{value: val, isSet: true}
}

func (v NullableAtcfwReadPoPRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAtcfwReadPoPRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
