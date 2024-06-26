/*
BloxOne Redirect API

You can configure BloxOne Threat Defense Cloud to redirect traffic to the Infoblox server that displays the default or customized redirect page. You can redirect traffic to a custom destination using custom redirects.

API version: v1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redirect

import (
	"encoding/json"
)

// checks if the CustomRedirectsReadCustomRedirect404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomRedirectsReadCustomRedirect404Response{}

// CustomRedirectsReadCustomRedirect404Response struct for CustomRedirectsReadCustomRedirect404Response
type CustomRedirectsReadCustomRedirect404Response struct {
	Error                *CustomRedirectsReadCustomRedirect404ResponseError `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CustomRedirectsReadCustomRedirect404Response CustomRedirectsReadCustomRedirect404Response

// NewCustomRedirectsReadCustomRedirect404Response instantiates a new CustomRedirectsReadCustomRedirect404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomRedirectsReadCustomRedirect404Response() *CustomRedirectsReadCustomRedirect404Response {
	this := CustomRedirectsReadCustomRedirect404Response{}
	return &this
}

// NewCustomRedirectsReadCustomRedirect404ResponseWithDefaults instantiates a new CustomRedirectsReadCustomRedirect404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomRedirectsReadCustomRedirect404ResponseWithDefaults() *CustomRedirectsReadCustomRedirect404Response {
	this := CustomRedirectsReadCustomRedirect404Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CustomRedirectsReadCustomRedirect404Response) GetError() CustomRedirectsReadCustomRedirect404ResponseError {
	if o == nil || IsNil(o.Error) {
		var ret CustomRedirectsReadCustomRedirect404ResponseError
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomRedirectsReadCustomRedirect404Response) GetErrorOk() (*CustomRedirectsReadCustomRedirect404ResponseError, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CustomRedirectsReadCustomRedirect404Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given CustomRedirectsReadCustomRedirect404ResponseError and assigns it to the Error field.
func (o *CustomRedirectsReadCustomRedirect404Response) SetError(v CustomRedirectsReadCustomRedirect404ResponseError) {
	o.Error = &v
}

func (o CustomRedirectsReadCustomRedirect404Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomRedirectsReadCustomRedirect404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CustomRedirectsReadCustomRedirect404Response) UnmarshalJSON(data []byte) (err error) {
	varCustomRedirectsReadCustomRedirect404Response := _CustomRedirectsReadCustomRedirect404Response{}

	err = json.Unmarshal(data, &varCustomRedirectsReadCustomRedirect404Response)

	if err != nil {
		return err
	}

	*o = CustomRedirectsReadCustomRedirect404Response(varCustomRedirectsReadCustomRedirect404Response)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCustomRedirectsReadCustomRedirect404Response struct {
	value *CustomRedirectsReadCustomRedirect404Response
	isSet bool
}

func (v NullableCustomRedirectsReadCustomRedirect404Response) Get() *CustomRedirectsReadCustomRedirect404Response {
	return v.value
}

func (v *NullableCustomRedirectsReadCustomRedirect404Response) Set(val *CustomRedirectsReadCustomRedirect404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomRedirectsReadCustomRedirect404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomRedirectsReadCustomRedirect404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomRedirectsReadCustomRedirect404Response(val *CustomRedirectsReadCustomRedirect404Response) *NullableCustomRedirectsReadCustomRedirect404Response {
	return &NullableCustomRedirectsReadCustomRedirect404Response{value: val, isSet: true}
}

func (v NullableCustomRedirectsReadCustomRedirect404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomRedirectsReadCustomRedirect404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
