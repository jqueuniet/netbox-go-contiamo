// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// WritablePowerFeedRequest is an object. Adds support for custom fields and tags.
type WritablePowerFeedRequest struct {
	// Amperage:
	Amperage int32 `json:"amperage,omitempty" mapstructure:"amperage,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// MaxUtilization: Maximum permissible draw (percentage)
	MaxUtilization int32 `json:"max_utilization,omitempty" mapstructure:"max_utilization,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Phase: * `single-phase` - Single phase
	// * `three-phase` - Three-phase
	Phase string `json:"phase,omitempty" mapstructure:"phase,omitempty"`
	// PowerPanel:
	PowerPanel int32 `json:"power_panel" mapstructure:"power_panel"`
	// Rack:
	Rack *int32 `json:"rack,omitempty" mapstructure:"rack,omitempty"`
	// Status: * `offline` - Offline
	// * `active` - Active
	// * `planned` - Planned
	// * `failed` - Failed
	Status string `json:"status,omitempty" mapstructure:"status,omitempty"`
	// Supply: * `ac` - AC
	// * `dc` - DC
	Supply string `json:"supply,omitempty" mapstructure:"supply,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Type: * `primary` - Primary
	// * `redundant` - Redundant
	Type string `json:"type,omitempty" mapstructure:"type,omitempty"`
	// Voltage:
	Voltage int32 `json:"voltage,omitempty" mapstructure:"voltage,omitempty"`
}

// Validate implements basic validation for this model
func (m WritablePowerFeedRequest) Validate() error {
	return validation.Errors{
		"amperage": validation.Validate(
			m.Amperage, validation.Min(int32(1)), validation.Max(int32(32767)),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"maxUtilization": validation.Validate(
			m.MaxUtilization, validation.Min(int32(1)), validation.Max(int32(100)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"voltage": validation.Validate(
			m.Voltage, validation.Min(int32(-32768)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetAmperage returns the Amperage property
func (m WritablePowerFeedRequest) GetAmperage() int32 {
	return m.Amperage
}

// SetAmperage sets the Amperage property
func (m *WritablePowerFeedRequest) SetAmperage(val int32) {
	m.Amperage = val
}

// GetComments returns the Comments property
func (m WritablePowerFeedRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritablePowerFeedRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritablePowerFeedRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritablePowerFeedRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritablePowerFeedRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritablePowerFeedRequest) SetDescription(val string) {
	m.Description = val
}

// GetMarkConnected returns the MarkConnected property
func (m WritablePowerFeedRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *WritablePowerFeedRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMaxUtilization returns the MaxUtilization property
func (m WritablePowerFeedRequest) GetMaxUtilization() int32 {
	return m.MaxUtilization
}

// SetMaxUtilization sets the MaxUtilization property
func (m *WritablePowerFeedRequest) SetMaxUtilization(val int32) {
	m.MaxUtilization = val
}

// GetName returns the Name property
func (m WritablePowerFeedRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritablePowerFeedRequest) SetName(val string) {
	m.Name = val
}

// GetPhase returns the Phase property
func (m WritablePowerFeedRequest) GetPhase() string {
	return m.Phase
}

// SetPhase sets the Phase property
func (m *WritablePowerFeedRequest) SetPhase(val string) {
	m.Phase = val
}

// GetPowerPanel returns the PowerPanel property
func (m WritablePowerFeedRequest) GetPowerPanel() int32 {
	return m.PowerPanel
}

// SetPowerPanel sets the PowerPanel property
func (m *WritablePowerFeedRequest) SetPowerPanel(val int32) {
	m.PowerPanel = val
}

// GetRack returns the Rack property
func (m WritablePowerFeedRequest) GetRack() *int32 {
	return m.Rack
}

// SetRack sets the Rack property
func (m *WritablePowerFeedRequest) SetRack(val *int32) {
	m.Rack = val
}

// GetStatus returns the Status property
func (m WritablePowerFeedRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *WritablePowerFeedRequest) SetStatus(val string) {
	m.Status = val
}

// GetSupply returns the Supply property
func (m WritablePowerFeedRequest) GetSupply() string {
	return m.Supply
}

// SetSupply sets the Supply property
func (m *WritablePowerFeedRequest) SetSupply(val string) {
	m.Supply = val
}

// GetTags returns the Tags property
func (m WritablePowerFeedRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritablePowerFeedRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m WritablePowerFeedRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *WritablePowerFeedRequest) SetType(val string) {
	m.Type = val
}

// GetVoltage returns the Voltage property
func (m WritablePowerFeedRequest) GetVoltage() int32 {
	return m.Voltage
}

// SetVoltage sets the Voltage property
func (m *WritablePowerFeedRequest) SetVoltage(val int32) {
	m.Voltage = val
}
