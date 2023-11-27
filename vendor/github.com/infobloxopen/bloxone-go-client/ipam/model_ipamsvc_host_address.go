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

// checks if the IpamsvcHostAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamsvcHostAddress{}

// IpamsvcHostAddress The __HostAddress__ object represents addresses associated with a Host object.
type IpamsvcHostAddress struct {
	// Field usage depends on the operation:  * For read operation, _address_ of the _Address_ corresponding to the _ref_ resource.  * For write operation, _address_ to be created if the _Address_ does not exist. Required if _ref_ is not set on write:     * If the _Address_ already exists and is already pointing to the right _Host_, the operation proceeds.     * If the _Address_ already exists and is pointing to a different _Host, the operation must abort.     * If the _Address_ already exists and is not pointing to any _Host_, it is linked to the _Host_.
	Address string `json:"address"`
	// The resource identifier.
	Ref string `json:"ref"`
	// The resource identifier.
	Space string `json:"space"`
}

// NewIpamsvcHostAddress instantiates a new IpamsvcHostAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamsvcHostAddress(address string, ref string, space string) *IpamsvcHostAddress {
	this := IpamsvcHostAddress{}
	this.Address = address
	this.Ref = ref
	this.Space = space
	return &this
}

// NewIpamsvcHostAddressWithDefaults instantiates a new IpamsvcHostAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamsvcHostAddressWithDefaults() *IpamsvcHostAddress {
	this := IpamsvcHostAddress{}
	return &this
}

// GetAddress returns the Address field value
func (o *IpamsvcHostAddress) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHostAddress) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *IpamsvcHostAddress) SetAddress(v string) {
	o.Address = v
}

// GetRef returns the Ref field value
func (o *IpamsvcHostAddress) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHostAddress) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *IpamsvcHostAddress) SetRef(v string) {
	o.Ref = v
}

// GetSpace returns the Space field value
func (o *IpamsvcHostAddress) GetSpace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Space
}

// GetSpaceOk returns a tuple with the Space field value
// and a boolean to check if the value has been set.
func (o *IpamsvcHostAddress) GetSpaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Space, true
}

// SetSpace sets field value
func (o *IpamsvcHostAddress) SetSpace(v string) {
	o.Space = v
}

func (o IpamsvcHostAddress) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamsvcHostAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["address"] = o.Address
	toSerialize["ref"] = o.Ref
	toSerialize["space"] = o.Space
	return toSerialize, nil
}

type NullableIpamsvcHostAddress struct {
	value *IpamsvcHostAddress
	isSet bool
}

func (v NullableIpamsvcHostAddress) Get() *IpamsvcHostAddress {
	return v.value
}

func (v *NullableIpamsvcHostAddress) Set(val *IpamsvcHostAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamsvcHostAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamsvcHostAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamsvcHostAddress(val *IpamsvcHostAddress) *NullableIpamsvcHostAddress {
	return &NullableIpamsvcHostAddress{value: val, isSet: true}
}

func (v NullableIpamsvcHostAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamsvcHostAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}