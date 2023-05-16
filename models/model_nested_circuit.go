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

// NestedCircuit is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedCircuit struct {
	// Cid: Unique circuit ID
	Cid string `json:"cid" mapstructure:"cid"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m NestedCircuit) Validate() error {
	return validation.Errors{
		"cid": validation.Validate(
			m.Cid, validation.NotNil, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCid returns the Cid property
func (m NestedCircuit) GetCid() string {
	return m.Cid
}

// SetCid sets the Cid property
func (m *NestedCircuit) SetCid(val string) {
	m.Cid = val
}

// GetDisplay returns the Display property
func (m NestedCircuit) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *NestedCircuit) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m NestedCircuit) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *NestedCircuit) SetId(val int32) {
	m.Id = val
}

// GetUrl returns the Url property
func (m NestedCircuit) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *NestedCircuit) SetUrl(val string) {
	m.Url = val
}
