// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"

	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// nestedL2VPNSlugPattern is the validation pattern for NestedL2VPN.Slug
var nestedL2VPNSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedL2VPN is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedL2VPN struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Identifier:
	Identifier *int64 `json:"identifier,omitempty" mapstructure:"identifier,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
	// Type: * `vpws` - VPWS
	// * `vpls` - VPLS
	// * `vxlan` - VXLAN
	// * `vxlan-evpn` - VXLAN-EVPN
	// * `mpls-evpn` - MPLS EVPN
	// * `pbb-evpn` - PBB EVPN
	// * `epl` - EPL
	// * `evpl` - EVPL
	// * `ep-lan` - Ethernet Private LAN
	// * `evp-lan` - Ethernet Virtual Private LAN
	// * `ep-tree` - Ethernet Private Tree
	// * `evp-tree` - Ethernet Virtual Private Tree
	Type string `json:"type" mapstructure:"type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedL2VPN) Validate() error {
	return validation.Errors{
		"identifier": validation.Validate(
			m.Identifier, validation.Min(*int64(-9.223372036854776e+18)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.NotNil, validation.Length(0, 100), validation.Match(nestedL2VPNSlugPattern),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedL2VPN) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedL2VPN) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedL2VPN) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedL2VPN) SetId(val int32) {
	m.Id = val
}

// GetIdentifier returns the Identifier property
func (m NestedL2VPN) GetIdentifier() *int64 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *NestedL2VPN) SetIdentifier(val *int64) {
	m.Identifier = val
}

// GetName returns the Name property
func (m NestedL2VPN) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedL2VPN) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedL2VPN) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedL2VPN) SetSlug(val string) {
	m.Slug = val
}

// GetType returns the Type property
func (m NestedL2VPN) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *NestedL2VPN) SetType(val string) {
	m.Type = val
}

// GetUrl returns the Url property
func (m NestedL2VPN) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedL2VPN) SetUrl(val string) {
	m.Url = val
}
