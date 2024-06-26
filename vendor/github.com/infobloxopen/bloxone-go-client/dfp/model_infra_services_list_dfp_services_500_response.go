/*
DFP API

BloxOne Cloud is a SaaS offering designed to provide protection to devices on and off-premises, including roaming, remote, and branch offices. It provides visibility into infected and compromised devices, prevents DNS-based data exfiltration, and automatically stops device communications with command-and-control servers (C&Cs) and botnets, in addition to providing recursive DNS services in the cloud. You can access the services by deploying the BloxOne Endpoint agent or the DNS forwarding proxy.  For remote office deployments or in cases where installing an endpoint agent is not desirable or possible, you can use the DNS forwarding proxy. It is a software that runs on bare-metal, VM infrastructures, or Infoblox NIOS appliances; and it embeds the client IPs in DNS queries before forwarding them to BloxOne Cloud. The communications are encrypted and client visibility is maintained. The proxy also provides DNS resolution to local DNS zones when you configure local resolvers. Once you set up a DNS forwarding proxy, it becomes the main DNS server for your remote site. It will also cache responses to speed resolution of future queries.  By implementing the DNS forwarding proxy, you can rest assured that BloxOne Cloud effectively enforces DNS client-based security policies at your remote sites. On-premises devices that send DNS queries reveal their actual client IP addresses (instead of their NAT IP address), which allows BloxOne Cloud to apply the security policies applicable to the respective endpoints and identify infected clients.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dfp

import (
	"encoding/json"
)

// checks if the InfraServicesListDfpServices500Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InfraServicesListDfpServices500Response{}

// InfraServicesListDfpServices500Response struct for InfraServicesListDfpServices500Response
type InfraServicesListDfpServices500Response struct {
	Error                *InfraServicesListDfpServices500ResponseError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _InfraServicesListDfpServices500Response InfraServicesListDfpServices500Response

// NewInfraServicesListDfpServices500Response instantiates a new InfraServicesListDfpServices500Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInfraServicesListDfpServices500Response() *InfraServicesListDfpServices500Response {
	this := InfraServicesListDfpServices500Response{}
	return &this
}

// NewInfraServicesListDfpServices500ResponseWithDefaults instantiates a new InfraServicesListDfpServices500Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInfraServicesListDfpServices500ResponseWithDefaults() *InfraServicesListDfpServices500Response {
	this := InfraServicesListDfpServices500Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *InfraServicesListDfpServices500Response) GetError() InfraServicesListDfpServices500ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret InfraServicesListDfpServices500ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InfraServicesListDfpServices500Response) GetErrorOk() (*InfraServicesListDfpServices500ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *InfraServicesListDfpServices500Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given InfraServicesListDfpServices500ResponseError and assigns it to the Error field.
func (o *InfraServicesListDfpServices500Response) SetError(v InfraServicesListDfpServices500ResponseError) {
	o.Error = &v
}

func (o InfraServicesListDfpServices500Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InfraServicesListDfpServices500Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *InfraServicesListDfpServices500Response) UnmarshalJSON(data []byte) (err error) {
	varInfraServicesListDfpServices500Response := _InfraServicesListDfpServices500Response{}

	err = json.Unmarshal(data, &varInfraServicesListDfpServices500Response)

	if err != nil {
		return err
	}

	*o = InfraServicesListDfpServices500Response(varInfraServicesListDfpServices500Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableInfraServicesListDfpServices500Response struct {
	value *InfraServicesListDfpServices500Response
	isSet bool
}

func (v NullableInfraServicesListDfpServices500Response) Get() *InfraServicesListDfpServices500Response {
	return v.value
}

func (v *NullableInfraServicesListDfpServices500Response) Set(val *InfraServicesListDfpServices500Response) {
	v.value = val
	v.isSet = true
}

func (v NullableInfraServicesListDfpServices500Response) IsSet() bool {
	return v.isSet
}

func (v *NullableInfraServicesListDfpServices500Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInfraServicesListDfpServices500Response(val *InfraServicesListDfpServices500Response) *NullableInfraServicesListDfpServices500Response {
	return &NullableInfraServicesListDfpServices500Response{value: val, isSet: true}
}

func (v NullableInfraServicesListDfpServices500Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInfraServicesListDfpServices500Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
