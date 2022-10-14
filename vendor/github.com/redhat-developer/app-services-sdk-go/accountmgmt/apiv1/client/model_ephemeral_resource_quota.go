/*
 * Account Management Service API
 *
 * Manage user subscriptions and clusters
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accountmgmtclient

import (
	"encoding/json"
)

// EphemeralResourceQuota struct for EphemeralResourceQuota
type EphemeralResourceQuota struct {
	AvailabilityZoneType *string `json:"availability_zone_type,omitempty"`
	Byoc *bool `json:"byoc,omitempty"`
	ResourceName *string `json:"resource_name,omitempty"`
	ResourceType *string `json:"resource_type,omitempty"`
	Sku *string `json:"sku,omitempty"`
	SkuCount *int32 `json:"sku_count,omitempty"`
}

// NewEphemeralResourceQuota instantiates a new EphemeralResourceQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEphemeralResourceQuota() *EphemeralResourceQuota {
	this := EphemeralResourceQuota{}
	return &this
}

// NewEphemeralResourceQuotaWithDefaults instantiates a new EphemeralResourceQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEphemeralResourceQuotaWithDefaults() *EphemeralResourceQuota {
	this := EphemeralResourceQuota{}
	return &this
}

// GetAvailabilityZoneType returns the AvailabilityZoneType field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetAvailabilityZoneType() string {
	if o == nil || o.AvailabilityZoneType == nil {
		var ret string
		return ret
	}
	return *o.AvailabilityZoneType
}

// GetAvailabilityZoneTypeOk returns a tuple with the AvailabilityZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetAvailabilityZoneTypeOk() (*string, bool) {
	if o == nil || o.AvailabilityZoneType == nil {
		return nil, false
	}
	return o.AvailabilityZoneType, true
}

// HasAvailabilityZoneType returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasAvailabilityZoneType() bool {
	if o != nil && o.AvailabilityZoneType != nil {
		return true
	}

	return false
}

// SetAvailabilityZoneType gets a reference to the given string and assigns it to the AvailabilityZoneType field.
func (o *EphemeralResourceQuota) SetAvailabilityZoneType(v string) {
	o.AvailabilityZoneType = &v
}

// GetByoc returns the Byoc field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetByoc() bool {
	if o == nil || o.Byoc == nil {
		var ret bool
		return ret
	}
	return *o.Byoc
}

// GetByocOk returns a tuple with the Byoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetByocOk() (*bool, bool) {
	if o == nil || o.Byoc == nil {
		return nil, false
	}
	return o.Byoc, true
}

// HasByoc returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasByoc() bool {
	if o != nil && o.Byoc != nil {
		return true
	}

	return false
}

// SetByoc gets a reference to the given bool and assigns it to the Byoc field.
func (o *EphemeralResourceQuota) SetByoc(v bool) {
	o.Byoc = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetResourceName() string {
	if o == nil || o.ResourceName == nil {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetResourceNameOk() (*string, bool) {
	if o == nil || o.ResourceName == nil {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasResourceName() bool {
	if o != nil && o.ResourceName != nil {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *EphemeralResourceQuota) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetResourceType() string {
	if o == nil || o.ResourceType == nil {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetResourceTypeOk() (*string, bool) {
	if o == nil || o.ResourceType == nil {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasResourceType() bool {
	if o != nil && o.ResourceType != nil {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *EphemeralResourceQuota) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetSku() string {
	if o == nil || o.Sku == nil {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetSkuOk() (*string, bool) {
	if o == nil || o.Sku == nil {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasSku() bool {
	if o != nil && o.Sku != nil {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *EphemeralResourceQuota) SetSku(v string) {
	o.Sku = &v
}

// GetSkuCount returns the SkuCount field value if set, zero value otherwise.
func (o *EphemeralResourceQuota) GetSkuCount() int32 {
	if o == nil || o.SkuCount == nil {
		var ret int32
		return ret
	}
	return *o.SkuCount
}

// GetSkuCountOk returns a tuple with the SkuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EphemeralResourceQuota) GetSkuCountOk() (*int32, bool) {
	if o == nil || o.SkuCount == nil {
		return nil, false
	}
	return o.SkuCount, true
}

// HasSkuCount returns a boolean if a field has been set.
func (o *EphemeralResourceQuota) HasSkuCount() bool {
	if o != nil && o.SkuCount != nil {
		return true
	}

	return false
}

// SetSkuCount gets a reference to the given int32 and assigns it to the SkuCount field.
func (o *EphemeralResourceQuota) SetSkuCount(v int32) {
	o.SkuCount = &v
}

func (o EphemeralResourceQuota) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AvailabilityZoneType != nil {
		toSerialize["availability_zone_type"] = o.AvailabilityZoneType
	}
	if o.Byoc != nil {
		toSerialize["byoc"] = o.Byoc
	}
	if o.ResourceName != nil {
		toSerialize["resource_name"] = o.ResourceName
	}
	if o.ResourceType != nil {
		toSerialize["resource_type"] = o.ResourceType
	}
	if o.Sku != nil {
		toSerialize["sku"] = o.Sku
	}
	if o.SkuCount != nil {
		toSerialize["sku_count"] = o.SkuCount
	}
	return json.Marshal(toSerialize)
}

type NullableEphemeralResourceQuota struct {
	value *EphemeralResourceQuota
	isSet bool
}

func (v NullableEphemeralResourceQuota) Get() *EphemeralResourceQuota {
	return v.value
}

func (v *NullableEphemeralResourceQuota) Set(val *EphemeralResourceQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableEphemeralResourceQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableEphemeralResourceQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEphemeralResourceQuota(val *EphemeralResourceQuota) *NullableEphemeralResourceQuota {
	return &NullableEphemeralResourceQuota{value: val, isSet: true}
}

func (v NullableEphemeralResourceQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEphemeralResourceQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

