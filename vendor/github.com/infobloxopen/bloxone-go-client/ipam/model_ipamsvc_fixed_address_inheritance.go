/*
IP Address Management API

The IPAM/DHCP Application is a BloxOne DDI service providing IP address management and DHCP protocol features. The IPAM component provides visibility into and provisioning tools to manage networking spaces, monitoring and reporting of entire IP address infrastructures, and integration with DNS and DHCP protocols. The DHCP component provides DHCP protocol configuration service with on-prem host serving DHCP protocol. It is part of the full-featured, DDI cloud solution that enables customers to deploy large numbers of protocol servers to deliver DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ipam

import (
	"encoding/json"
)

// checks if the IpamsvcFixedAddressInheritance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcFixedAddressInheritance{}

// IpamsvcFixedAddressInheritance The __FixedAddressInheritance__ object specifies how and which fields _FixedAddress_ object inherits from the parent.
type IpamsvcFixedAddressInheritance struct {
	DhcpOptions               *IpamsvcInheritedDHCPOptionList `json:"dhcp_options,omitempty"`
	HeaderOptionFilename      *InheritanceInheritedString     `json:"header_option_filename,omitempty"`
	HeaderOptionServerAddress *InheritanceInheritedString     `json:"header_option_server_address,omitempty"`
	HeaderOptionServerName    *InheritanceInheritedString     `json:"header_option_server_name,omitempty"`
}

// NewIpamsvcFixedAddressInheritance instantiates a new IpamsvcFixedAddressInheritance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcFixedAddressInheritance() *IpamsvcFixedAddressInheritance {
	this := IpamsvcFixedAddressInheritance{}
	return &this
}

// NewIpamsvcFixedAddressInheritanceWithDefaults instantiates a new IpamsvcFixedAddressInheritance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcFixedAddressInheritanceWithDefaults() *IpamsvcFixedAddressInheritance {
	this := IpamsvcFixedAddressInheritance{}
	return &this
}

// GetDhcpOptions returns the DhcpOptions field value if set, zero value otherwise.
func (o *IpamsvcFixedAddressInheritance) GetDhcpOptions() IpamsvcInheritedDHCPOptionList {
	if o == nil || IsNil(o.DhcpOptions) {
		var ret IpamsvcInheritedDHCPOptionList
		return ret
	}
	return *o.DhcpOptions
}

// GetDhcpOptionsOk returns a tuple with the DhcpOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcFixedAddressInheritance) GetDhcpOptionsOk() (*IpamsvcInheritedDHCPOptionList, bool) {
	if o == nil || IsNil(o.DhcpOptions) {
		return nil, false
	}
	return o.DhcpOptions, true
}

// HasDhcpOptions returns a boolean if a field has been set.
func (o *IpamsvcFixedAddressInheritance) HasDhcpOptions() bool {
	if o != nil && !IsNil(o.DhcpOptions) {
		return true
	}

	return false
}

// SetDhcpOptions gets a reference to the given IpamsvcInheritedDHCPOptionList and assigns it to the DhcpOptions field.
func (o *IpamsvcFixedAddressInheritance) SetDhcpOptions(v IpamsvcInheritedDHCPOptionList) {
	o.DhcpOptions = &v
}

// GetHeaderOptionFilename returns the HeaderOptionFilename field value if set, zero value otherwise.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionFilename() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionFilename
}

// GetHeaderOptionFilenameOk returns a tuple with the HeaderOptionFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionFilenameOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionFilename) {
		return nil, false
	}
	return o.HeaderOptionFilename, true
}

// HasHeaderOptionFilename returns a boolean if a field has been set.
func (o *IpamsvcFixedAddressInheritance) HasHeaderOptionFilename() bool {
	if o != nil && !IsNil(o.HeaderOptionFilename) {
		return true
	}

	return false
}

// SetHeaderOptionFilename gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionFilename field.
func (o *IpamsvcFixedAddressInheritance) SetHeaderOptionFilename(v InheritanceInheritedString) {
	o.HeaderOptionFilename = &v
}

// GetHeaderOptionServerAddress returns the HeaderOptionServerAddress field value if set, zero value otherwise.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionServerAddress() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionServerAddress
}

// GetHeaderOptionServerAddressOk returns a tuple with the HeaderOptionServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionServerAddressOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionServerAddress) {
		return nil, false
	}
	return o.HeaderOptionServerAddress, true
}

// HasHeaderOptionServerAddress returns a boolean if a field has been set.
func (o *IpamsvcFixedAddressInheritance) HasHeaderOptionServerAddress() bool {
	if o != nil && !IsNil(o.HeaderOptionServerAddress) {
		return true
	}

	return false
}

// SetHeaderOptionServerAddress gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionServerAddress field.
func (o *IpamsvcFixedAddressInheritance) SetHeaderOptionServerAddress(v InheritanceInheritedString) {
	o.HeaderOptionServerAddress = &v
}

// GetHeaderOptionServerName returns the HeaderOptionServerName field value if set, zero value otherwise.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionServerName() InheritanceInheritedString {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		var ret InheritanceInheritedString
		return ret
	}
	return *o.HeaderOptionServerName
}

// GetHeaderOptionServerNameOk returns a tuple with the HeaderOptionServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamsvcFixedAddressInheritance) GetHeaderOptionServerNameOk() (*InheritanceInheritedString, bool) {
	if o == nil || IsNil(o.HeaderOptionServerName) {
		return nil, false
	}
	return o.HeaderOptionServerName, true
}

// HasHeaderOptionServerName returns a boolean if a field has been set.
func (o *IpamsvcFixedAddressInheritance) HasHeaderOptionServerName() bool {
	if o != nil && !IsNil(o.HeaderOptionServerName) {
		return true
	}

	return false
}

// SetHeaderOptionServerName gets a reference to the given InheritanceInheritedString and assigns it to the HeaderOptionServerName field.
func (o *IpamsvcFixedAddressInheritance) SetHeaderOptionServerName(v InheritanceInheritedString) {
	o.HeaderOptionServerName = &v
}

func (o IpamsvcFixedAddressInheritance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcFixedAddressInheritance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpOptions) {
		toSerialize["dhcp_options"] = o.DhcpOptions
	}
	if !IsNil(o.HeaderOptionFilename) {
		toSerialize["header_option_filename"] = o.HeaderOptionFilename
	}
	if !IsNil(o.HeaderOptionServerAddress) {
		toSerialize["header_option_server_address"] = o.HeaderOptionServerAddress
	}
	if !IsNil(o.HeaderOptionServerName) {
		toSerialize["header_option_server_name"] = o.HeaderOptionServerName
	}
	return toSerialize, nil
}

type NullableIpamsvcFixedAddressInheritance struct {
	value *IpamsvcFixedAddressInheritance
	isSet bool
}

func (v NullableIpamsvcFixedAddressInheritance) Get() *IpamsvcFixedAddressInheritance {
	return v.value
}

func (v *NullableIpamsvcFixedAddressInheritance) Set(val *IpamsvcFixedAddressInheritance) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcFixedAddressInheritance) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcFixedAddressInheritance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcFixedAddressInheritance(val *IpamsvcFixedAddressInheritance) *NullableIpamsvcFixedAddressInheritance {
	return &NullableIpamsvcFixedAddressInheritance{value: val, isSet: true}
}

func (v NullableIpamsvcFixedAddressInheritance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcFixedAddressInheritance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}