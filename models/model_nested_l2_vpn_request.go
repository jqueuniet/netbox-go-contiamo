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
)

// nestedL2VPNRequestSlugPattern is the validation pattern for NestedL2VPNRequest.Slug
var nestedL2VPNRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedL2VPNRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedL2VPNRequest struct {
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
}

// Validate implements basic validation for this model
func (m NestedL2VPNRequest) Validate() error {
	return validation.Errors{
		"identifier": validation.Validate(
			m.Identifier, validation.Min(*int64(-9.223372036854776e+18)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(nestedL2VPNRequestSlugPattern),
		),
	}.Filter()
}

// GetIdentifier returns the Identifier property
func (m NestedL2VPNRequest) GetIdentifier() *int64 {
	return m.Identifier
}

// SetIdentifier sets the Identifier property
func (m *NestedL2VPNRequest) SetIdentifier(val *int64) {
	m.Identifier = val
}

// GetName returns the Name property
func (m NestedL2VPNRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedL2VPNRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedL2VPNRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedL2VPNRequest) SetSlug(val string) {
	m.Slug = val
}

// GetType returns the Type property
func (m NestedL2VPNRequest) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *NestedL2VPNRequest) SetType(val string) {
	m.Type = val
}
