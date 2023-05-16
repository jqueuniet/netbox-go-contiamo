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

// CircuitTermination is an object. Adds support for custom fields and tags.
type CircuitTermination struct {
	// Occupied:
	Occupied bool `json:"_occupied" mapstructure:"_occupied"`
	// Cable: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Cable *NestedCable `json:"cable" mapstructure:"cable"`
	// CableEnd:
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// Circuit: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Circuit NestedCircuit `json:"circuit" mapstructure:"circuit"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LinkPeers:
	LinkPeers []interface{} `json:"link_peers" mapstructure:"link_peers"`
	// LinkPeersType: Return the type of the peer link terminations, or None.
	LinkPeersType string `json:"link_peers_type" mapstructure:"link_peers_type"`
	// MarkConnected: Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty" mapstructure:"mark_connected,omitempty"`
	// PortSpeed: Physical circuit speed
	PortSpeed *int32 `json:"port_speed,omitempty" mapstructure:"port_speed,omitempty"`
	// PpInfo: Patch panel ID and port number(s)
	PpInfo string `json:"pp_info,omitempty" mapstructure:"pp_info,omitempty"`
	// ProviderNetwork: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	ProviderNetwork *NestedProviderNetwork `json:"provider_network,omitempty" mapstructure:"provider_network,omitempty"`
	// Site: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Site *NestedSite `json:"site,omitempty" mapstructure:"site,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// TermSide: * `A` - A
	// * `Z` - Z
	TermSide string `json:"term_side" mapstructure:"term_side"`
	// UpstreamSpeed: Upstream speed, if different from port speed
	UpstreamSpeed *int32 `json:"upstream_speed,omitempty" mapstructure:"upstream_speed,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// XconnectId: ID of the local cross-connect
	XconnectId string `json:"xconnect_id,omitempty" mapstructure:"xconnect_id,omitempty"`
}

// Validate implements basic validation for this model
func (m CircuitTermination) Validate() error {
	return validation.Errors{
		"cable": validation.Validate(
			m.Cable,
		),
		"circuit": validation.Validate(
			m.Circuit, validation.NotNil,
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"linkPeers": validation.Validate(
			m.LinkPeers, validation.NotNil,
		),
		"portSpeed": validation.Validate(
			m.PortSpeed, validation.Min(*int32(0)), validation.Max(*int32(2.147483647e+09)),
		),
		"ppInfo": validation.Validate(
			m.PpInfo, validation.Length(0, 100),
		),
		"providerNetwork": validation.Validate(
			m.ProviderNetwork,
		),
		"site": validation.Validate(
			m.Site,
		),
		"tags": validation.Validate(
			m.Tags,
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

// GetOccupied returns the Occupied property
func (m CircuitTermination) GetOccupied() bool {
	return m.Occupied
}

// SetOccupied sets the Occupied property
func (m *CircuitTermination) SetOccupied(val bool) {
	m.Occupied = val
}

// GetCable returns the Cable property
func (m CircuitTermination) GetCable() *NestedCable {
	return m.Cable
}

// SetCable sets the Cable property
func (m *CircuitTermination) SetCable(val *NestedCable) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m CircuitTermination) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *CircuitTermination) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCircuit returns the Circuit property
func (m CircuitTermination) GetCircuit() NestedCircuit {
	return m.Circuit
}

// SetCircuit sets the Circuit property
func (m *CircuitTermination) SetCircuit(val NestedCircuit) {
	m.Circuit = val
}

// GetCreated returns the Created property
func (m CircuitTermination) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CircuitTermination) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m CircuitTermination) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *CircuitTermination) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m CircuitTermination) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *CircuitTermination) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m CircuitTermination) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CircuitTermination) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m CircuitTermination) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CircuitTermination) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m CircuitTermination) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CircuitTermination) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkPeers returns the LinkPeers property
func (m CircuitTermination) GetLinkPeers() []interface{} {
	return m.LinkPeers
}

// SetLinkPeers sets the LinkPeers property
func (m *CircuitTermination) SetLinkPeers(val []interface{}) {
	m.LinkPeers = val
}

// GetLinkPeersType returns the LinkPeersType property
func (m CircuitTermination) GetLinkPeersType() string {
	return m.LinkPeersType
}

// SetLinkPeersType sets the LinkPeersType property
func (m *CircuitTermination) SetLinkPeersType(val string) {
	m.LinkPeersType = val
}

// GetMarkConnected returns the MarkConnected property
func (m CircuitTermination) GetMarkConnected() bool {
	return m.MarkConnected
}

// SetMarkConnected sets the MarkConnected property
func (m *CircuitTermination) SetMarkConnected(val bool) {
	m.MarkConnected = val
}

// GetPortSpeed returns the PortSpeed property
func (m CircuitTermination) GetPortSpeed() *int32 {
	return m.PortSpeed
}

// SetPortSpeed sets the PortSpeed property
func (m *CircuitTermination) SetPortSpeed(val *int32) {
	m.PortSpeed = val
}

// GetPpInfo returns the PpInfo property
func (m CircuitTermination) GetPpInfo() string {
	return m.PpInfo
}

// SetPpInfo sets the PpInfo property
func (m *CircuitTermination) SetPpInfo(val string) {
	m.PpInfo = val
}

// GetProviderNetwork returns the ProviderNetwork property
func (m CircuitTermination) GetProviderNetwork() *NestedProviderNetwork {
	return m.ProviderNetwork
}

// SetProviderNetwork sets the ProviderNetwork property
func (m *CircuitTermination) SetProviderNetwork(val *NestedProviderNetwork) {
	m.ProviderNetwork = val
}

// GetSite returns the Site property
func (m CircuitTermination) GetSite() *NestedSite {
	return m.Site
}

// SetSite sets the Site property
func (m *CircuitTermination) SetSite(val *NestedSite) {
	m.Site = val
}

// GetTags returns the Tags property
func (m CircuitTermination) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *CircuitTermination) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTermSide returns the TermSide property
func (m CircuitTermination) GetTermSide() string {
	return m.TermSide
}

// SetTermSide sets the TermSide property
func (m *CircuitTermination) SetTermSide(val string) {
	m.TermSide = val
}

// GetUpstreamSpeed returns the UpstreamSpeed property
func (m CircuitTermination) GetUpstreamSpeed() *int32 {
	return m.UpstreamSpeed
}

// SetUpstreamSpeed sets the UpstreamSpeed property
func (m *CircuitTermination) SetUpstreamSpeed(val *int32) {
	m.UpstreamSpeed = val
}

// GetUrl returns the Url property
func (m CircuitTermination) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CircuitTermination) SetUrl(val string) {
	m.Url = val
}

// GetXconnectId returns the XconnectId property
func (m CircuitTermination) GetXconnectId() string {
	return m.XconnectId
}

// SetXconnectId sets the XconnectId property
func (m *CircuitTermination) SetXconnectId(val string) {
	m.XconnectId = val
}
