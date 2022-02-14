/*
 * Ory Oathkeeper API
 *
 * Documentation for all of Ory Oathkeeper's APIs.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// FlushLoginConsentRequest FlushLoginConsentRequest flush login consent request
type FlushLoginConsentRequest struct {
	// NotAfter sets after which point tokens should not be flushed. This is useful when you want to keep a history of recent login and consent database entries for auditing. Format: date-time
	NotAfter *time.Time `json:"notAfter,omitempty"`
}

// NewFlushLoginConsentRequest instantiates a new FlushLoginConsentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlushLoginConsentRequest() *FlushLoginConsentRequest {
	this := FlushLoginConsentRequest{}
	return &this
}

// NewFlushLoginConsentRequestWithDefaults instantiates a new FlushLoginConsentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlushLoginConsentRequestWithDefaults() *FlushLoginConsentRequest {
	this := FlushLoginConsentRequest{}
	return &this
}

// GetNotAfter returns the NotAfter field value if set, zero value otherwise.
func (o *FlushLoginConsentRequest) GetNotAfter() time.Time {
	if o == nil || o.NotAfter == nil {
		var ret time.Time
		return ret
	}
	return *o.NotAfter
}

// GetNotAfterOk returns a tuple with the NotAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlushLoginConsentRequest) GetNotAfterOk() (*time.Time, bool) {
	if o == nil || o.NotAfter == nil {
		return nil, false
	}
	return o.NotAfter, true
}

// HasNotAfter returns a boolean if a field has been set.
func (o *FlushLoginConsentRequest) HasNotAfter() bool {
	if o != nil && o.NotAfter != nil {
		return true
	}

	return false
}

// SetNotAfter gets a reference to the given time.Time and assigns it to the NotAfter field.
func (o *FlushLoginConsentRequest) SetNotAfter(v time.Time) {
	o.NotAfter = &v
}

func (o FlushLoginConsentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NotAfter != nil {
		toSerialize["notAfter"] = o.NotAfter
	}
	return json.Marshal(toSerialize)
}

type NullableFlushLoginConsentRequest struct {
	value *FlushLoginConsentRequest
	isSet bool
}

func (v NullableFlushLoginConsentRequest) Get() *FlushLoginConsentRequest {
	return v.value
}

func (v *NullableFlushLoginConsentRequest) Set(val *FlushLoginConsentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFlushLoginConsentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFlushLoginConsentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlushLoginConsentRequest(val *FlushLoginConsentRequest) *NullableFlushLoginConsentRequest {
	return &NullableFlushLoginConsentRequest{value: val, isSet: true}
}

func (v NullableFlushLoginConsentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlushLoginConsentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
