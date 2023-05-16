// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// InterfaceTemplate is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type InterfaceTemplate struct {
	// Bridge: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Bridge *NestedInterfaceTemplate `json:"bridge,omitempty" mapstructure:"bridge,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DeviceType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	DeviceType *NestedDeviceType `json:"device_type,omitempty" mapstructure:"device_type,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Label: Physical label
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// MgmtOnly:
	MgmtOnly bool `json:"mgmt_only,omitempty" mapstructure:"mgmt_only,omitempty"`
	// ModuleType: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ModuleType *NestedModuleType `json:"module_type,omitempty" mapstructure:"module_type,omitempty"`
	// Name:
	//         {module} is accepted as a substitution for the module bay position when attached to a module type.
	Name string `json:"name" mapstructure:"name"`
	// PoeMode:
	PoeMode *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"poe_mode,omitempty" mapstructure:"poe_mode,omitempty"`
	// PoeType:
	PoeType *struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"poe_type,omitempty" mapstructure:"poe_type,omitempty"`
	// Type:
	Type struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m InterfaceTemplate) Validate() error {
	return validation.Errors{
		"bridge": validation.Validate(
			m.Bridge,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"deviceType": validation.Validate(
			m.DeviceType,
		),
		"label": validation.Validate(
			m.Label, validation.Length(0, 64),
		),
		"moduleType": validation.Validate(
			m.ModuleType,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 64),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetBridge returns the Bridge property
func (m InterfaceTemplate) GetBridge() *NestedInterfaceTemplate {
	return m.Bridge
}

// SetBridge sets the Bridge property
func (m *InterfaceTemplate) SetBridge(val *NestedInterfaceTemplate) {
	m.Bridge = val
}

// GetCreated returns the Created property
func (m InterfaceTemplate) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *InterfaceTemplate) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDescription returns the Description property
func (m InterfaceTemplate) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *InterfaceTemplate) SetDescription(val string) {
	m.Description = val
}

// GetDeviceType returns the DeviceType property
func (m InterfaceTemplate) GetDeviceType() *NestedDeviceType {
	return m.DeviceType
}

// SetDeviceType sets the DeviceType property
func (m *InterfaceTemplate) SetDeviceType(val *NestedDeviceType) {
	m.DeviceType = val
}

// GetDisplay returns the Display property
func (m InterfaceTemplate) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *InterfaceTemplate) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m InterfaceTemplate) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *InterfaceTemplate) SetEnabled(val bool) {
	m.Enabled = val
}

// GetId returns the Id property
func (m InterfaceTemplate) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *InterfaceTemplate) SetId(val int32) {
	m.Id = val
}

// GetLabel returns the Label property
func (m InterfaceTemplate) GetLabel() string {
	return m.Label
}

// SetLabel sets the Label property
func (m *InterfaceTemplate) SetLabel(val string) {
	m.Label = val
}

// GetLastUpdated returns the LastUpdated property
func (m InterfaceTemplate) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *InterfaceTemplate) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetMgmtOnly returns the MgmtOnly property
func (m InterfaceTemplate) GetMgmtOnly() bool {
	return m.MgmtOnly
}

// SetMgmtOnly sets the MgmtOnly property
func (m *InterfaceTemplate) SetMgmtOnly(val bool) {
	m.MgmtOnly = val
}

// GetModuleType returns the ModuleType property
func (m InterfaceTemplate) GetModuleType() *NestedModuleType {
	return m.ModuleType
}

// SetModuleType sets the ModuleType property
func (m *InterfaceTemplate) SetModuleType(val *NestedModuleType) {
	m.ModuleType = val
}

// GetName returns the Name property
func (m InterfaceTemplate) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *InterfaceTemplate) SetName(val string) {
	m.Name = val
}

// GetPoeMode returns the PoeMode property
func (m InterfaceTemplate) GetPoeMode() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.PoeMode
}

// SetPoeMode sets the PoeMode property
func (m *InterfaceTemplate) SetPoeMode(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.PoeMode = val
}

// GetPoeType returns the PoeType property
func (m InterfaceTemplate) GetPoeType() *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.PoeType
}

// SetPoeType sets the PoeType property
func (m *InterfaceTemplate) SetPoeType(val *struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.PoeType = val
}

// GetType returns the Type property
func (m InterfaceTemplate) GetType() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Type
}

// SetType sets the Type property
func (m *InterfaceTemplate) SetType(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Type = val
}

// GetUrl returns the Url property
func (m InterfaceTemplate) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *InterfaceTemplate) SetUrl(val string) {
	m.Url = val
}
