/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.13
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// Volume Volume volume
type Volume struct {
	// Date/Time the volume was created.
	CreatedAt *string `json:"CreatedAt,omitempty"`
	// Name of the volume driver used by the volume.
	Driver string `json:"Driver"`
	// User-defined key/value metadata.
	Labels map[string]string `json:"Labels"`
	// Mount path of the volume on the host.
	Mountpoint string `json:"Mountpoint"`
	// Name of the volume.
	Name string `json:"Name"`
	// The driver specific options used when creating the volume.
	Options map[string]string `json:"Options"`
	// The level at which the volume exists. Either `global` for cluster-wide, or `local` for machine level.
	Scope string `json:"Scope"`
	// Low-level details about the volume, provided by the volume driver. Details are returned as a map with key/value pairs: `{\"key\":\"value\",\"key2\":\"value2\"}`.  The `Status` field is optional, and is omitted if the volume driver does not support this feature.
	Status map[string]map[string]interface{} `json:"Status,omitempty"`
	UsageData *VolumeUsageData `json:"UsageData,omitempty"`
}

// NewVolume instantiates a new Volume object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVolume(driver string, labels map[string]string, mountpoint string, name string, options map[string]string, scope string) *Volume {
	this := Volume{}
	this.Driver = driver
	this.Labels = labels
	this.Mountpoint = mountpoint
	this.Name = name
	this.Options = options
	this.Scope = scope
	return &this
}

// NewVolumeWithDefaults instantiates a new Volume object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeWithDefaults() *Volume {
	this := Volume{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Volume) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Volume) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Volume) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDriver returns the Driver field value
func (o *Volume) GetDriver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Driver
}

// GetDriverOk returns a tuple with the Driver field value
// and a boolean to check if the value has been set.
func (o *Volume) GetDriverOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Driver, true
}

// SetDriver sets field value
func (o *Volume) SetDriver(v string) {
	o.Driver = v
}

// GetLabels returns the Labels field value
func (o *Volume) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *Volume) GetLabelsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *Volume) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetMountpoint returns the Mountpoint field value
func (o *Volume) GetMountpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mountpoint
}

// GetMountpointOk returns a tuple with the Mountpoint field value
// and a boolean to check if the value has been set.
func (o *Volume) GetMountpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mountpoint, true
}

// SetMountpoint sets field value
func (o *Volume) SetMountpoint(v string) {
	o.Mountpoint = v
}

// GetName returns the Name field value
func (o *Volume) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Volume) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Volume) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value
func (o *Volume) GetOptions() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *Volume) GetOptionsOk() (*map[string]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *Volume) SetOptions(v map[string]string) {
	o.Options = v
}

// GetScope returns the Scope field value
func (o *Volume) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *Volume) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *Volume) SetScope(v string) {
	o.Scope = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Volume) GetStatus() map[string]map[string]interface{} {
	if o == nil || o.Status == nil {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetStatusOk() (map[string]map[string]interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Volume) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given map[string]map[string]interface{} and assigns it to the Status field.
func (o *Volume) SetStatus(v map[string]map[string]interface{}) {
	o.Status = v
}

// GetUsageData returns the UsageData field value if set, zero value otherwise.
func (o *Volume) GetUsageData() VolumeUsageData {
	if o == nil || o.UsageData == nil {
		var ret VolumeUsageData
		return ret
	}
	return *o.UsageData
}

// GetUsageDataOk returns a tuple with the UsageData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Volume) GetUsageDataOk() (*VolumeUsageData, bool) {
	if o == nil || o.UsageData == nil {
		return nil, false
	}
	return o.UsageData, true
}

// HasUsageData returns a boolean if a field has been set.
func (o *Volume) HasUsageData() bool {
	if o != nil && o.UsageData != nil {
		return true
	}

	return false
}

// SetUsageData gets a reference to the given VolumeUsageData and assigns it to the UsageData field.
func (o *Volume) SetUsageData(v VolumeUsageData) {
	o.UsageData = &v
}

func (o Volume) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["CreatedAt"] = o.CreatedAt
	}
	if true {
		toSerialize["Driver"] = o.Driver
	}
	if true {
		toSerialize["Labels"] = o.Labels
	}
	if true {
		toSerialize["Mountpoint"] = o.Mountpoint
	}
	if true {
		toSerialize["Name"] = o.Name
	}
	if true {
		toSerialize["Options"] = o.Options
	}
	if true {
		toSerialize["Scope"] = o.Scope
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.UsageData != nil {
		toSerialize["UsageData"] = o.UsageData
	}
	return json.Marshal(toSerialize)
}

type NullableVolume struct {
	value *Volume
	isSet bool
}

func (v NullableVolume) Get() *Volume {
	return v.value
}

func (v *NullableVolume) Set(val *Volume) {
	v.value = val
	v.isSet = true
}

func (v NullableVolume) IsSet() bool {
	return v.isSet
}

func (v *NullableVolume) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolume(val *Volume) *NullableVolume {
	return &NullableVolume{value: val, isSet: true}
}

func (v NullableVolume) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolume) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


