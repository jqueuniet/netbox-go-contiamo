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

// WritableCircuitTerminationRequest is an object. Adds support for custom fields and tags.
type WritableCircuitTerminationRequest struct {
	// Circuit:
	Circuit int32 `json:"circuit" mapstructure:"circuit"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// PortSpeed: Physical circuit speed
	PortSpeed *int32 `json:"port_speed,omitempty" mapstructure:"port_speed,omitempty"`
	// PpInfo: Patch panel ID and port number(s)
	PpInfo string `json:"pp_info,omitempty" mapstructure:"pp_info,omitempty"`
	// ProviderNetwork:
	ProviderNetwork *int32 `json:"provider_network,omitempty" mapstructure:"provider_network,omitempty"`
	// Site:
	Site *int32 `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TermSide: * `A` - A
	// * `Z` - Z
	TermSide string `json:"term_side" mapstructure:"term_side"`
	// UpstreamSpeed: Upstream speed, if different from port speed
	UpstreamSpeed *int32 `json:"upstream_speed,omitempty" mapstructure:"upstream_speed,omitempty"`
	// XconnectId: ID of the local cross-connect
	XconnectId string `json:"xconnect_id,omitempty" mapstructure:"xconnect_id,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableCircuitTerminationRequest) Validate() error {
	return validation.Errors{
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"portSpeed": validation.Validate(
			m.PortSpeed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"ppInfo": validation.Validate(
			m.PpInfo, validation.Length(0, 100),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"upstreamSpeed": validation.Validate(
			m.UpstreamSpeed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"xconnectId": validation.Validate(
			m.XconnectId, validation.Length(0, 50),
		),
	}.Filter()
}

// GetCircuit returns the Circuit property
func (m WritableCircuitTerminationRequest) GetCircuit() int32 {
	return m.Circuit
}

// SetCircuit sets the Circuit property
func (m *WritableCircuitTerminationRequest) SetCircuit(val int32) {
	m.Circuit = val
}

// GetCustomFields returns the CustomFields property
func (m WritableCircuitTerminationRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableCircuitTerminationRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableCircuitTerminationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableCircuitTerminationRequest) SetDescription(val string) {
	m.Description = val
}

// GetMarkConnected returns the MarkConnected property
func (m WritableCircuitTerminationRequest) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *WritableCircuitTerminationRequest) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetPortSpeed returns the PortSpeed property
func (m WritableCircuitTerminationRequest) GetPortSpeed() *int32 {
	return m.PortSpeed
}

// SetPortSpeed sets the PortSpeed property
func (m *WritableCircuitTerminationRequest) SetPortSpeed(val *int32) {
	m.PortSpeed = val
}

// GetPpInfo returns the PpInfo property
func (m WritableCircuitTerminationRequest) GetPpInfo() string {
	return m.PpInfo
}

// SetPpInfo sets the PpInfo property
func (m *WritableCircuitTerminationRequest) SetPpInfo(val string) {
	m.PpInfo = val
}

// GetProviderNetwork returns the ProviderNetwork property
func (m WritableCircuitTerminationRequest) GetProviderNetwork() *int32 {
	return m.ProviderNetwork
}

// SetProviderNetwork sets the ProviderNetwork property
func (m *WritableCircuitTerminationRequest) SetProviderNetwork(val *int32) {
	m.ProviderNetwork = val
}

// GetSite returns the Site property
func (m WritableCircuitTerminationRequest) GetSite() *int32 {
	return m.Site
}

// SetSite sets the Site property
func (m *WritableCircuitTerminationRequest) SetSite(val *int32) {
	m.Site = val
}

// GetTags returns the Tags property
func (m WritableCircuitTerminationRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableCircuitTerminationRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTermSide returns the TermSide property
func (m WritableCircuitTerminationRequest) GetTermSide() string {
	return m.TermSide
}

// SetTermSide sets the TermSide property
func (m *WritableCircuitTerminationRequest) SetTermSide(val string) {
	m.TermSide = val
}

// GetUpstreamSpeed returns the UpstreamSpeed property
func (m WritableCircuitTerminationRequest) GetUpstreamSpeed() *int32 {
	return m.UpstreamSpeed
}

// SetUpstreamSpeed sets the UpstreamSpeed property
func (m *WritableCircuitTerminationRequest) SetUpstreamSpeed(val *int32) {
	m.UpstreamSpeed = val
}

// GetXconnectId returns the XconnectId property
func (m WritableCircuitTerminationRequest) GetXconnectId() string {
	return m.XconnectId
}

// SetXconnectId sets the XconnectId property
func (m *WritableCircuitTerminationRequest) SetXconnectId(val string) {
	m.XconnectId = val
}
