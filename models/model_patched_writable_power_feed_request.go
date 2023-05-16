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

// PatchedWritablePowerFeedRequest is an object. Adds support for custom fields and tags.
type PatchedWritablePowerFeedRequest struct {
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
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Phase: * `single-phase` - Single phase
	// * `three-phase` - Three-phase
	Phase string `json:"phase,omitempty" mapstructure:"phase,omitempty"`
	// PowerPanel:
	PowerPanel int32 `json:"power_panel,omitempty" mapstructure:"power_panel,omitempty"`
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
func (m PatchedWritablePowerFeedRequest) Validate() error {
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
			m.Name, validation.Length(1, 100),
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
func (m PatchedWritablePowerFeedRequest) GetAmperage() int32 {
	return m.Amperage
}

// SetAmperage sets the Amperage property
func (m *PatchedWritablePowerFeedRequest) SetAmperage(val int32) {
	m.Amperage = val
}

// GetComments returns the Comments property
func (m PatchedWritablePowerFeedRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritablePowerFeedRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritablePowerFeedRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritablePowerFeedRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritablePowerFeedRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritablePowerFeedRequest) SetDescription(val string) {
	m.Description = val
}

// GetMarkConnected returns the MarkConnected property
func (m PatchedWritablePowerFeedRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *PatchedWritablePowerFeedRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetMaxUtilization returns the MaxUtilization property
func (m PatchedWritablePowerFeedRequest) GetMaxUtilization() int32 {
	return m.MaxUtilization
}

// SetMaxUtilization sets the MaxUtilization property
func (m *PatchedWritablePowerFeedRequest) SetMaxUtilization(val int32) {
	m.MaxUtilization = val
}

// GetName returns the Name property
func (m PatchedWritablePowerFeedRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritablePowerFeedRequest) SetName(val string) {
	m.Name = val
}

// GetPhase returns the Phase property
func (m PatchedWritablePowerFeedRequest) GetPhase() string {
	return m.Phase
}

// SetPhase sets the Phase property
func (m *PatchedWritablePowerFeedRequest) SetPhase(val string) {
	m.Phase = val
}

// GetPowerPanel returns the PowerPanel property
func (m PatchedWritablePowerFeedRequest) GetPowerPanel() int32 {
	return m.PowerPanel
}

// SetPowerPanel sets the PowerPanel property
func (m *PatchedWritablePowerFeedRequest) SetPowerPanel(val int32) {
	m.PowerPanel = val
}

// GetRack returns the Rack property
func (m PatchedWritablePowerFeedRequest) GetRack() *int32 {
	return m.Rack
}

// SetRack sets the Rack property
func (m *PatchedWritablePowerFeedRequest) SetRack(val *int32) {
	m.Rack = val
}

// GetStatus returns the Status property
func (m PatchedWritablePowerFeedRequest) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *PatchedWritablePowerFeedRequest) SetStatus(val string) {
	m.Status = val
}

// GetSupply returns the Supply property
func (m PatchedWritablePowerFeedRequest) GetSupply() string {
	return m.Supply
}

// SetSupply sets the Supply property
func (m *PatchedWritablePowerFeedRequest) SetSupply(val string) {
	m.Supply = val
}

// GetTags returns the Tags property
func (m PatchedWritablePowerFeedRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritablePowerFeedRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetType returns the Type property
func (m PatchedWritablePowerFeedRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *PatchedWritablePowerFeedRequest) SetType(val string) {
	m.Type = val
}

// GetVoltage returns the Voltage property
func (m PatchedWritablePowerFeedRequest) GetVoltage() int32 {
	return m.Voltage
}

// SetVoltage sets the Voltage property
func (m *PatchedWritablePowerFeedRequest) SetVoltage(val int32) {
	m.Voltage = val
}
