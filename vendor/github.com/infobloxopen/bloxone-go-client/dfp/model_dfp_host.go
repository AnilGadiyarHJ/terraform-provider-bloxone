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

// checks if the DfpHost type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DfpHost{}

// DfpHost struct for DfpHost
type DfpHost struct {
	// // The DNS Forwarding Proxy legacy ID object identifier.
	LegacyHostId *int32 `json:"legacy_host_id,omitempty"`
	// The name of the DNS Forwarding Proxy.
	Name *string `json:"name,omitempty"`
	// The On-Prem Host identifier.
	Ophid *string `json:"ophid,omitempty"`
	// The DNS Forwarding Proxy site identifier that is appended to DNS queries originating from this DNS Forwarding Proxy and subsequently used for policy lookup purposes.
	SiteId               *string `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DfpHost DfpHost

// NewDfpHost instantiates a new DfpHost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDfpHost() *DfpHost {
	this := DfpHost{}
	return &this
}

// NewDfpHostWithDefaults instantiates a new DfpHost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDfpHostWithDefaults() *DfpHost {
	this := DfpHost{}
	return &this
}

// GetLegacyHostId returns the LegacyHostId field value if set, zero value otherwise.
func (o *DfpHost) GetLegacyHostId() int32 {
	if o == nil || IsNil(o.LegacyHostId) {
		var ret int32
		return ret
	}
	return *o.LegacyHostId
}

// GetLegacyHostIdOk returns a tuple with the LegacyHostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfpHost) GetLegacyHostIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LegacyHostId) {
		return nil, false
	}
	return o.LegacyHostId, true
}

// HasLegacyHostId returns a boolean if a field has been set.
func (o *DfpHost) HasLegacyHostId() bool {
	if o != nil && !IsNil(o.LegacyHostId) {
		return true
	}

	return false
}

// SetLegacyHostId gets a reference to the given int32 and assigns it to the LegacyHostId field.
func (o *DfpHost) SetLegacyHostId(v int32) {
	o.LegacyHostId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DfpHost) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfpHost) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DfpHost) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DfpHost) SetName(v string) {
	o.Name = &v
}

// GetOphid returns the Ophid field value if set, zero value otherwise.
func (o *DfpHost) GetOphid() string {
	if o == nil || IsNil(o.Ophid) {
		var ret string
		return ret
	}
	return *o.Ophid
}

// GetOphidOk returns a tuple with the Ophid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfpHost) GetOphidOk() (*string, bool) {
	if o == nil || IsNil(o.Ophid) {
		return nil, false
	}
	return o.Ophid, true
}

// HasOphid returns a boolean if a field has been set.
func (o *DfpHost) HasOphid() bool {
	if o != nil && !IsNil(o.Ophid) {
		return true
	}

	return false
}

// SetOphid gets a reference to the given string and assigns it to the Ophid field.
func (o *DfpHost) SetOphid(v string) {
	o.Ophid = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *DfpHost) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DfpHost) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *DfpHost) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *DfpHost) SetSiteId(v string) {
	o.SiteId = &v
}

func (o DfpHost) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DfpHost) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LegacyHostId) {
		toSerialize["legacy_host_id"] = o.LegacyHostId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ophid) {
		toSerialize["ophid"] = o.Ophid
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DfpHost) UnmarshalJSON(data []byte) (err error) {
	varDfpHost := _DfpHost{}

	err = json.Unmarshal(data, &varDfpHost)

	if err != nil {
		return err
	}

	*o = DfpHost(varDfpHost)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "legacy_host_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "ophid")
		delete(additionalProperties, "site_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableDfpHost struct {
	value *DfpHost
	isSet bool
}

func (v NullableDfpHost) Get() *DfpHost {
	return v.value
}

func (v *NullableDfpHost) Set(val *DfpHost) {
	v.value = val
	v.isSet = true
}

func (v NullableDfpHost) IsSet() bool {
	return v.isSet
}

func (v *NullableDfpHost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDfpHost(val *DfpHost) *NullableDfpHost {
	return &NullableDfpHost{value: val, isSet: true}
}

func (v NullableDfpHost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDfpHost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
