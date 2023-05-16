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

// NestedFHRPGroupRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedFHRPGroupRequest struct {
	// GroupId:
	GroupId int32 `json:"group_id" mapstructure:"group_id"`
	// Protocol: * `vrrp2` - VRRPv2
	// * `vrrp3` - VRRPv3
	// * `carp` - CARP
	// * `clusterxl` - ClusterXL
	// * `hsrp` - HSRP
	// * `glbp` - GLBP
	// * `other` - Other
	Protocol string `json:"protocol" mapstructure:"protocol"`
}

// Validate implements basic validation for this model
func (m NestedFHRPGroupRequest) Validate() error {
	return validation.Errors{
		"groupId": validation.Validate(
			m.GroupId, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetGroupId returns the GroupId property
func (m NestedFHRPGroupRequest) GetGroupId() int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *NestedFHRPGroupRequest) SetGroupId(val int32) {
	m.GroupId = val
}

// GetProtocol returns the Protocol property
func (m NestedFHRPGroupRequest) GetProtocol() string {
	return m.Protocol
}

// SetProtocol sets the Protocol property
func (m *NestedFHRPGroupRequest) SetProtocol(val string) {
	m.Protocol = val
}
