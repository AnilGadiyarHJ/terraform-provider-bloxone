/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dns_config

import (
	"encoding/json"
)

// checks if the ConfigForwarder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfigForwarder{}

// ConfigForwarder External DNS server to forward to.
type ConfigForwarder struct {
	// Server IP address.
	Address string `json:"address"`
	// Server FQDN.
	Fqdn string `json:"fqdn"`
	// Server FQDN in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
}

// NewConfigForwarder instantiates a new ConfigForwarder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfigForwarder(address string, fqdn string) *ConfigForwarder {
	this := ConfigForwarder{}
	this.Address = address
	this.Fqdn = fqdn
	return &this
}

// NewConfigForwarderWithDefaults instantiates a new ConfigForwarder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigForwarderWithDefaults() *ConfigForwarder {
	this := ConfigForwarder{}
	return &this
}

// GetAddress returns the Address field value
func (o *ConfigForwarder) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *ConfigForwarder) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *ConfigForwarder) SetAddress(v string) {
	o.Address = v
}

// GetFqdn returns the Fqdn field value
func (o *ConfigForwarder) GetFqdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value
// and a boolean to check if the value has been set.
func (o *ConfigForwarder) GetFqdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fqdn, true
}

// SetFqdn sets field value
func (o *ConfigForwarder) SetFqdn(v string) {
	o.Fqdn = v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ConfigForwarder) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConfigForwarder) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ConfigForwarder) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ConfigForwarder) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

func (o ConfigForwarder) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigForwarder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["fqdn"] = o.Fqdn
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	return toSerialize, nil
}

type NullableConfigForwarder struct {
	value *ConfigForwarder
	isSet bool
}

func (v NullableConfigForwarder) Get() *ConfigForwarder {
	return v.value
}

func (v *NullableConfigForwarder) Set(val *ConfigForwarder) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigForwarder) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigForwarder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigForwarder(val *ConfigForwarder) *NullableConfigForwarder {
	return &NullableConfigForwarder{value: val, isSet: true}
}

func (v NullableConfigForwarder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigForwarder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}