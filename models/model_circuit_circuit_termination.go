// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// CircuitCircuitTermination is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type CircuitCircuitTermination struct {
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// PortSpeed: Physical circuit speed
	PortSpeed *int32 `json:"port_speed,omitempty" mapstructure:"port_speed,omitempty"`
	// ProviderNetwork: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ProviderNetwork *NestedProviderNetwork `json:"provider_network" mapstructure:"provider_network"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSite `json:"site" mapstructure:"site"`
	// UpstreamSpeed: Upstream speed, if different from port speed
	UpstreamSpeed *int32 `json:"upstream_speed,omitempty" mapstructure:"upstream_speed,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// XconnectId: ID of the local cross-connect
	XconnectId string `json:"xconnect_id,omitempty" mapstructure:"xconnect_id,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitCircuitTermination) Validate() error {
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"xconnectId": validation.Validate(
			m.XconnectId, validation.Length(0, 50),
		),
	}.Filter()
}

// GetDescription returns the Description property
func (m CircuitCircuitTermination) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitCircuitTermination) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m CircuitCircuitTermination) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CircuitCircuitTermination) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m CircuitCircuitTermination) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitCircuitTermination) SetId(val int32) {
	m.Id = val
}

// GetPortSpeed returns the PortSpeed property
func (m CircuitCircuitTermination) GetPortSpeed() *int32 {
	return m.PortSpeed
}

// SetPortSpeed sets the PortSpeed property
func (m *CircuitCircuitTermination) SetPortSpeed(val *int32) {
	m.PortSpeed = val
}

// GetProviderNetwork returns the ProviderNetwork property
func (m CircuitCircuitTermination) GetProviderNetwork() *NestedProviderNetwork {
	return m.ProviderNetwork
}

// SetProviderNetwork sets the ProviderNetwork property
func (m *CircuitCircuitTermination) SetProviderNetwork(val *NestedProviderNetwork) {
	m.ProviderNetwork = val
}

// GetSite returns the Site property
func (m CircuitCircuitTermination) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitCircuitTermination) SetSite(val *NestedSite) {
	m.Site = val
}

// GetUpstreamSpeed returns the UpstreamSpeed property
func (m CircuitCircuitTermination) GetUpstreamSpeed() *int32 {
	return m.UpstreamSpeed
}

// SetUpstreamSpeed sets the UpstreamSpeed property
func (m *CircuitCircuitTermination) SetUpstreamSpeed(val *int32) {
	m.UpstreamSpeed = val
}

// GetUrl returns the Url property
func (m CircuitCircuitTermination) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CircuitCircuitTermination) SetUrl(val string) {
	m.Url = val
}

// GetXconnectId returns the XconnectId property
func (m CircuitCircuitTermination) GetXconnectId() string {
	return m.XconnectId
}

// SetXconnectId sets the XconnectId property
func (m *CircuitCircuitTermination) SetXconnectId(val string) {
	m.XconnectId = val
}
