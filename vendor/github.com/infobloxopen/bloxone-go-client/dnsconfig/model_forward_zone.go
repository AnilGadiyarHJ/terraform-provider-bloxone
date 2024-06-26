/*
DNS Configuration API

The DNS application is a BloxOne DDI service that provides cloud-based DNS configuration with on-prem host serving DNS protocol. It is part of the full-featured BloxOne DDI solution that enables customers the ability to deploy large numbers of protocol servers in the delivery of DNS and DHCP throughout their enterprise network.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dnsconfig

import (
	"encoding/json"
	"time"
)

// checks if the ForwardZone type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardZone{}

// ForwardZone Forward zone
type ForwardZone struct {
	// Optional. Comment for zone configuration.
	Comment *string `json:"comment,omitempty"`
	// The timestamp when the object has been created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// Optional. _true_ to disable object. A disabled object is effectively non-existent when generating configuration.
	Disabled *bool `json:"disabled,omitempty"`
	// Optional. External DNS servers to forward to. Order is not significant.
	ExternalForwarders []Forwarder `json:"external_forwarders,omitempty"`
	// Optional. _true_ to only forward.
	ForwardOnly *bool `json:"forward_only,omitempty"`
	// Zone FQDN. The FQDN supplied at creation will be converted to canonical form.  Read-only after creation.
	Fqdn *string `json:"fqdn,omitempty"`
	// The resource identifier.
	Hosts []string `json:"hosts,omitempty"`
	// The resource identifier.
	Id *string `json:"id,omitempty"`
	// The resource identifier.
	InternalForwarders []string `json:"internal_forwarders,omitempty"`
	// Reverse zone network address in the following format: \"ip-address/cidr\". Defaults to empty.
	MappedSubnet *string `json:"mapped_subnet,omitempty"`
	// Read-only. Zone mapping type. Allowed values:  * _forward_,  * _ipv4_reverse_.  * _ipv6_reverse_.  Defaults to _forward_.
	Mapping *string `json:"mapping,omitempty"`
	// The resource identifier.
	Nsgs []string `json:"nsgs,omitempty"`
	// The resource identifier.
	Parent *string `json:"parent,omitempty"`
	// Zone FQDN in punycode.
	ProtocolFqdn *string `json:"protocol_fqdn,omitempty"`
	// Tagging specifics.
	Tags map[string]interface{} `json:"tags,omitempty"`
	// The timestamp when the object has been updated. Equals to _created_at_ if not updated after creation.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The resource identifier.
	View *string `json:"view,omitempty"`
	// The list of a forward zone warnings.
	Warnings             []Warning `json:"warnings,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ForwardZone ForwardZone

// NewForwardZone instantiates a new ForwardZone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardZone() *ForwardZone {
	this := ForwardZone{}
	return &this
}

// NewForwardZoneWithDefaults instantiates a new ForwardZone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardZoneWithDefaults() *ForwardZone {
	this := ForwardZone{}
	return &this
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *ForwardZone) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *ForwardZone) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *ForwardZone) SetComment(v string) {
	o.Comment = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ForwardZone) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ForwardZone) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ForwardZone) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *ForwardZone) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *ForwardZone) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *ForwardZone) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetExternalForwarders returns the ExternalForwarders field value if set, zero value otherwise.
func (o *ForwardZone) GetExternalForwarders() []Forwarder {
	if o == nil || IsNil(o.ExternalForwarders) {
		var ret []Forwarder
		return ret
	}
	return o.ExternalForwarders
}

// GetExternalForwardersOk returns a tuple with the ExternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetExternalForwardersOk() ([]Forwarder, bool) {
	if o == nil || IsNil(o.ExternalForwarders) {
		return nil, false
	}
	return o.ExternalForwarders, true
}

// HasExternalForwarders returns a boolean if a field has been set.
func (o *ForwardZone) HasExternalForwarders() bool {
	if o != nil && !IsNil(o.ExternalForwarders) {
		return true
	}

	return false
}

// SetExternalForwarders gets a reference to the given []Forwarder and assigns it to the ExternalForwarders field.
func (o *ForwardZone) SetExternalForwarders(v []Forwarder) {
	o.ExternalForwarders = v
}

// GetForwardOnly returns the ForwardOnly field value if set, zero value otherwise.
func (o *ForwardZone) GetForwardOnly() bool {
	if o == nil || IsNil(o.ForwardOnly) {
		var ret bool
		return ret
	}
	return *o.ForwardOnly
}

// GetForwardOnlyOk returns a tuple with the ForwardOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetForwardOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.ForwardOnly) {
		return nil, false
	}
	return o.ForwardOnly, true
}

// HasForwardOnly returns a boolean if a field has been set.
func (o *ForwardZone) HasForwardOnly() bool {
	if o != nil && !IsNil(o.ForwardOnly) {
		return true
	}

	return false
}

// SetForwardOnly gets a reference to the given bool and assigns it to the ForwardOnly field.
func (o *ForwardZone) SetForwardOnly(v bool) {
	o.ForwardOnly = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *ForwardZone) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *ForwardZone) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *ForwardZone) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *ForwardZone) GetHosts() []string {
	if o == nil || IsNil(o.Hosts) {
		var ret []string
		return ret
	}
	return o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetHostsOk() ([]string, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *ForwardZone) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *ForwardZone) SetHosts(v []string) {
	o.Hosts = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ForwardZone) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ForwardZone) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ForwardZone) SetId(v string) {
	o.Id = &v
}

// GetInternalForwarders returns the InternalForwarders field value if set, zero value otherwise.
func (o *ForwardZone) GetInternalForwarders() []string {
	if o == nil || IsNil(o.InternalForwarders) {
		var ret []string
		return ret
	}
	return o.InternalForwarders
}

// GetInternalForwardersOk returns a tuple with the InternalForwarders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetInternalForwardersOk() ([]string, bool) {
	if o == nil || IsNil(o.InternalForwarders) {
		return nil, false
	}
	return o.InternalForwarders, true
}

// HasInternalForwarders returns a boolean if a field has been set.
func (o *ForwardZone) HasInternalForwarders() bool {
	if o != nil && !IsNil(o.InternalForwarders) {
		return true
	}

	return false
}

// SetInternalForwarders gets a reference to the given []string and assigns it to the InternalForwarders field.
func (o *ForwardZone) SetInternalForwarders(v []string) {
	o.InternalForwarders = v
}

// GetMappedSubnet returns the MappedSubnet field value if set, zero value otherwise.
func (o *ForwardZone) GetMappedSubnet() string {
	if o == nil || IsNil(o.MappedSubnet) {
		var ret string
		return ret
	}
	return *o.MappedSubnet
}

// GetMappedSubnetOk returns a tuple with the MappedSubnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetMappedSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.MappedSubnet) {
		return nil, false
	}
	return o.MappedSubnet, true
}

// HasMappedSubnet returns a boolean if a field has been set.
func (o *ForwardZone) HasMappedSubnet() bool {
	if o != nil && !IsNil(o.MappedSubnet) {
		return true
	}

	return false
}

// SetMappedSubnet gets a reference to the given string and assigns it to the MappedSubnet field.
func (o *ForwardZone) SetMappedSubnet(v string) {
	o.MappedSubnet = &v
}

// GetMapping returns the Mapping field value if set, zero value otherwise.
func (o *ForwardZone) GetMapping() string {
	if o == nil || IsNil(o.Mapping) {
		var ret string
		return ret
	}
	return *o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetMappingOk() (*string, bool) {
	if o == nil || IsNil(o.Mapping) {
		return nil, false
	}
	return o.Mapping, true
}

// HasMapping returns a boolean if a field has been set.
func (o *ForwardZone) HasMapping() bool {
	if o != nil && !IsNil(o.Mapping) {
		return true
	}

	return false
}

// SetMapping gets a reference to the given string and assigns it to the Mapping field.
func (o *ForwardZone) SetMapping(v string) {
	o.Mapping = &v
}

// GetNsgs returns the Nsgs field value if set, zero value otherwise.
func (o *ForwardZone) GetNsgs() []string {
	if o == nil || IsNil(o.Nsgs) {
		var ret []string
		return ret
	}
	return o.Nsgs
}

// GetNsgsOk returns a tuple with the Nsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetNsgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Nsgs) {
		return nil, false
	}
	return o.Nsgs, true
}

// HasNsgs returns a boolean if a field has been set.
func (o *ForwardZone) HasNsgs() bool {
	if o != nil && !IsNil(o.Nsgs) {
		return true
	}

	return false
}

// SetNsgs gets a reference to the given []string and assigns it to the Nsgs field.
func (o *ForwardZone) SetNsgs(v []string) {
	o.Nsgs = v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *ForwardZone) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *ForwardZone) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *ForwardZone) SetParent(v string) {
	o.Parent = &v
}

// GetProtocolFqdn returns the ProtocolFqdn field value if set, zero value otherwise.
func (o *ForwardZone) GetProtocolFqdn() string {
	if o == nil || IsNil(o.ProtocolFqdn) {
		var ret string
		return ret
	}
	return *o.ProtocolFqdn
}

// GetProtocolFqdnOk returns a tuple with the ProtocolFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetProtocolFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ProtocolFqdn) {
		return nil, false
	}
	return o.ProtocolFqdn, true
}

// HasProtocolFqdn returns a boolean if a field has been set.
func (o *ForwardZone) HasProtocolFqdn() bool {
	if o != nil && !IsNil(o.ProtocolFqdn) {
		return true
	}

	return false
}

// SetProtocolFqdn gets a reference to the given string and assigns it to the ProtocolFqdn field.
func (o *ForwardZone) SetProtocolFqdn(v string) {
	o.ProtocolFqdn = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ForwardZone) GetTags() map[string]interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret map[string]interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetTagsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return map[string]interface{}{}, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ForwardZone) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *ForwardZone) SetTags(v map[string]interface{}) {
	o.Tags = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ForwardZone) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ForwardZone) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ForwardZone) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetView returns the View field value if set, zero value otherwise.
func (o *ForwardZone) GetView() string {
	if o == nil || IsNil(o.View) {
		var ret string
		return ret
	}
	return *o.View
}

// GetViewOk returns a tuple with the View field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetViewOk() (*string, bool) {
	if o == nil || IsNil(o.View) {
		return nil, false
	}
	return o.View, true
}

// HasView returns a boolean if a field has been set.
func (o *ForwardZone) HasView() bool {
	if o != nil && !IsNil(o.View) {
		return true
	}

	return false
}

// SetView gets a reference to the given string and assigns it to the View field.
func (o *ForwardZone) SetView(v string) {
	o.View = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *ForwardZone) GetWarnings() []Warning {
	if o == nil || IsNil(o.Warnings) {
		var ret []Warning
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardZone) GetWarningsOk() ([]Warning, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *ForwardZone) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []Warning and assigns it to the Warnings field.
func (o *ForwardZone) SetWarnings(v []Warning) {
	o.Warnings = v
}

func (o ForwardZone) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardZone) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.ExternalForwarders) {
		toSerialize["external_forwarders"] = o.ExternalForwarders
	}
	if !IsNil(o.ForwardOnly) {
		toSerialize["forward_only"] = o.ForwardOnly
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InternalForwarders) {
		toSerialize["internal_forwarders"] = o.InternalForwarders
	}
	if !IsNil(o.MappedSubnet) {
		toSerialize["mapped_subnet"] = o.MappedSubnet
	}
	if !IsNil(o.Mapping) {
		toSerialize["mapping"] = o.Mapping
	}
	if !IsNil(o.Nsgs) {
		toSerialize["nsgs"] = o.Nsgs
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.ProtocolFqdn) {
		toSerialize["protocol_fqdn"] = o.ProtocolFqdn
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.View) {
		toSerialize["view"] = o.View
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ForwardZone) UnmarshalJSON(data []byte) (err error) {
	varForwardZone := _ForwardZone{}

	err = json.Unmarshal(data, &varForwardZone)

	if err != nil {
		return err
	}

	*o = ForwardZone(varForwardZone)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "comment")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "disabled")
		delete(additionalProperties, "external_forwarders")
		delete(additionalProperties, "forward_only")
		delete(additionalProperties, "fqdn")
		delete(additionalProperties, "hosts")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal_forwarders")
		delete(additionalProperties, "mapped_subnet")
		delete(additionalProperties, "mapping")
		delete(additionalProperties, "nsgs")
		delete(additionalProperties, "parent")
		delete(additionalProperties, "protocol_fqdn")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "updated_at")
		delete(additionalProperties, "view")
		delete(additionalProperties, "warnings")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableForwardZone struct {
	value *ForwardZone
	isSet bool
}

func (v NullableForwardZone) Get() *ForwardZone {
	return v.value
}

func (v *NullableForwardZone) Set(val *ForwardZone) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardZone) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardZone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardZone(val *ForwardZone) *NullableForwardZone {
	return &NullableForwardZone{value: val, isSet: true}
}

func (v NullableForwardZone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardZone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
