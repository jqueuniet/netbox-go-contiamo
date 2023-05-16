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

// NestedFHRPGroup is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedFHRPGroup struct {
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// GroupId:
	GroupId int32 `json:"group_id" mapstructure:"group_id"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Protocol: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	Protocol string `json:"protocol" mapstructure:"protocol"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedFHRPGroup) Validate() error {
	return validation.Errors{
		"groupId": validation.Validate(
			m.GroupId, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetDisplay returns the Display property
func (m NestedFHRPGroup) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedFHRPGroup) SetDisplay(val string) {
	m.Display = val
}

// GetGroupId returns the GroupId property
func (m NestedFHRPGroup) GetGroupId() int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *NestedFHRPGroup) SetGroupId(val int32) {
	m.GroupId = val
}

// GetId returns the Id property
func (m NestedFHRPGroup) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedFHRPGroup) SetId(val int32) {
	m.Id = val
}

// GetProtocol returns the Protocol property
func (m NestedFHRPGroup) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *NestedFHRPGroup) SetProtocol(val string) {
	m.Protocol = val
}

// GetUrl returns the Url property
func (m NestedFHRPGroup) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedFHRPGroup) SetUrl(val string) {
	m.Url = val
}
