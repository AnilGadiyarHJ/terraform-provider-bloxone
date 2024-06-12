/*
Schedule Software/Config Updates

Infoblox by default does automatic software updates when they become available. Updates are applied to all on-prem hosts, physical or virtual. However, you can override and schedule the software updates. You can also defer the updates to a later date and time. You can configure up to a total of 50 deferrals (scheduled and deferred software updates), which means you have the flexibility to create up to 50 update groups across different on-prem hosts by mapping with appropriate tags. Tags are be used to associate deferrals (scheduled or deferred) with a specific or group of onprem-hosts. Apart from software update deferrals, config update deferrals also can be configured using these overrides.

API version: v2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package upgradepolicy

import (
	"encoding/json"
)

// checks if the CreateMaintenanceWindowResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMaintenanceWindowResponse{}

// CreateMaintenanceWindowResponse struct for CreateMaintenanceWindowResponse
type CreateMaintenanceWindowResponse struct {
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateMaintenanceWindowResponse CreateMaintenanceWindowResponse

// NewCreateMaintenanceWindowResponse instantiates a new CreateMaintenanceWindowResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMaintenanceWindowResponse() *CreateMaintenanceWindowResponse {
	this := CreateMaintenanceWindowResponse{}
	return &this
}

// NewCreateMaintenanceWindowResponseWithDefaults instantiates a new CreateMaintenanceWindowResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMaintenanceWindowResponseWithDefaults() *CreateMaintenanceWindowResponse {
	this := CreateMaintenanceWindowResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateMaintenanceWindowResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMaintenanceWindowResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateMaintenanceWindowResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateMaintenanceWindowResponse) SetId(v string) {
	o.Id = &v
}

func (o CreateMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMaintenanceWindowResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CreateMaintenanceWindowResponse) UnmarshalJSON(data []byte) (err error) {
	varCreateMaintenanceWindowResponse := _CreateMaintenanceWindowResponse{}

	err = json.Unmarshal(data, &varCreateMaintenanceWindowResponse)

	if err != nil {
		return err
	}

	*o = CreateMaintenanceWindowResponse(varCreateMaintenanceWindowResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateMaintenanceWindowResponse struct {
	value *CreateMaintenanceWindowResponse
	isSet bool
}

func (v NullableCreateMaintenanceWindowResponse) Get() *CreateMaintenanceWindowResponse {
	return v.value
}

func (v *NullableCreateMaintenanceWindowResponse) Set(val *CreateMaintenanceWindowResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMaintenanceWindowResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMaintenanceWindowResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMaintenanceWindowResponse(val *CreateMaintenanceWindowResponse) *NullableCreateMaintenanceWindowResponse {
	return &NullableCreateMaintenanceWindowResponse{value: val, isSet: true}
}

func (v NullableCreateMaintenanceWindowResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMaintenanceWindowResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
