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

// checks if the ItemStructs type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemStructs{}

// ItemStructs The Items Structure which contains the item and its description
type ItemStructs struct {
	// The description of the item
	Description *string `json:"description,omitempty"`
	// The data of the Item
	Item                 *string `json:"item,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ItemStructs ItemStructs

// NewItemStructs instantiates a new ItemStructs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemStructs() *ItemStructs {
	this := ItemStructs{}
	return &this
}

// NewItemStructsWithDefaults instantiates a new ItemStructs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemStructsWithDefaults() *ItemStructs {
	this := ItemStructs{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ItemStructs) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemStructs) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ItemStructs) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ItemStructs) SetDescription(v string) {
	o.Description = &v
}

// GetItem returns the Item field value if set, zero value otherwise.
func (o *ItemStructs) GetItem() string {
	if o == nil || IsNil(o.Item) {
		var ret string
		return ret
	}
	return *o.Item
}

// GetItemOk returns a tuple with the Item field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemStructs) GetItemOk() (*string, bool) {
	if o == nil || IsNil(o.Item) {
		return nil, false
	}
	return o.Item, true
}

// HasItem returns a boolean if a field has been set.
func (o *ItemStructs) HasItem() bool {
	if o != nil && !IsNil(o.Item) {
		return true
	}

	return false
}

// SetItem gets a reference to the given string and assigns it to the Item field.
func (o *ItemStructs) SetItem(v string) {
	o.Item = &v
}

func (o ItemStructs) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ItemStructs) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Item) {
		toSerialize["item"] = o.Item
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ItemStructs) UnmarshalJSON(data []byte) (err error) {
	varItemStructs := _ItemStructs{}

	err = json.Unmarshal(data, &varItemStructs)

	if err != nil {
		return err
	}

	*o = ItemStructs(varItemStructs)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "description")
		delete(additionalProperties, "item")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemStructs struct {
	value *ItemStructs
	isSet bool
}

func (v NullableItemStructs) Get() *ItemStructs {
	return v.value
}

func (v *NullableItemStructs) Set(val *ItemStructs) {
	v.value = val
	v.isSet = true
}

func (v NullableItemStructs) IsSet() bool {
	return v.isSet
}

func (v *NullableItemStructs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemStructs(val *ItemStructs) *NullableItemStructs {
	return &NullableItemStructs{value: val, isSet: true}
}

func (v NullableItemStructs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemStructs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
