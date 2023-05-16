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

// CircuitCircuitTerminationRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type CircuitCircuitTerminationRequest struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// PortSpeed: Physical circuit speed
	PortSpeed *int32 `json:"port_speed,omitempty" mapstructure:"port_speed,omitempty"`
	// ProviderNetwork: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ProviderNetwork *NestedProviderNetworkRequest `json:"provider_network" mapstructure:"provider_network"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSiteRequest `json:"site" mapstructure:"site"`
	// UpstreamSpeed: Upstream speed, if different from port speed
	UpstreamSpeed *int32 `json:"upstream_speed,omitempty" mapstructure:"upstream_speed,omitempty"`
	// XconnectId: ID of the local cross-connect
	XconnectId string `json:"xconnect_id,omitempty" mapstructure:"xconnect_id,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitCircuitTerminationRequest) Validate() error {
	return validation.Errors{
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"portSpeed": validation.Validate(
			m.PortSpeed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"providerNetwork": validation.Validate(
			m.ProviderNetwork,
		),
		"site": validation.Validate(
			m.Site,
		),
		"upstreamSpeed": validation.Validate(
			m.UpstreamSpeed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"xconnectId": validation.Validate(
			m.XconnectId, validation.Length(0, 50),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m CircuitCircuitTerminationRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitCircuitTerminationRequest) SetDescription(val string) {
	m.Description = val
}

// GetPortSpeed returns the PortSpeed property
func (m CircuitCircuitTerminationRequest) GetPortSpeed() *int32 {
	return m.PortSpeed
}

// SetPortSpeed sets the PortSpeed property
func (m *CircuitCircuitTerminationRequest) SetPortSpeed(val *int32) {
	m.PortSpeed = val
}

// GetProviderNetwork returns the ProviderNetwork property
func (m CircuitCircuitTerminationRequest) GetProviderNetwork() *NestedProviderNetworkRequest {
	return m.ProviderNetwork
}

// SetProviderNetwork sets the ProviderNetwork property
func (m *CircuitCircuitTerminationRequest) SetProviderNetwork(val *NestedProviderNetworkRequest) {
	m.ProviderNetwork = val
}

// GetSite returns the Site property
func (m CircuitCircuitTerminationRequest) GetSite() *NestedSiteRequest {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitCircuitTerminationRequest) SetSite(val *NestedSiteRequest) {
	m.Site = val
}

// GetUpstreamSpeed returns the UpstreamSpeed property
func (m CircuitCircuitTerminationRequest) GetUpstreamSpeed() *int32 {
	return m.UpstreamSpeed
}

// SetUpstreamSpeed sets the UpstreamSpeed property
func (m *CircuitCircuitTerminationRequest) SetUpstreamSpeed(val *int32) {
	m.UpstreamSpeed = val
}

// GetXconnectId returns the XconnectId property
func (m CircuitCircuitTerminationRequest) GetXconnectId() string {
	return m.XconnectId
}

// SetXconnectId sets the XconnectId property
func (m *CircuitCircuitTerminationRequest) SetXconnectId(val string) {
	m.XconnectId = val
}
