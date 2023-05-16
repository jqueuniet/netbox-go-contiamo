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

// CableTermination is an object. Adds support for custom fields and tags.
type CableTermination struct {
	// Cable:
	Cable int32 `json:"cable" mapstructure:"cable"`
	// CableEnd: * `A` - A
	// * `B` - B
	CableEnd string `json:"cable_end" mapstructure:"cable_end"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Termination:
	Termination map[string]interface{} `json:"termination" mapstructure:"termination"`
	// TerminationId:
	TerminationId int64 `json:"termination_id" mapstructure:"termination_id"`
	// TerminationType:
	TerminationType string `json:"termination_type" mapstructure:"termination_type"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m CableTermination) Validate() error {
	return validation.Errors{
		"termination": validation.Validate(
			m.Termination,
		),
		"terminationId": validation.Validate(
			m.TerminationId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetCable returns the Cable property
func (m CableTermination) GetCable() int32 {
	return m.Cable
}

// SetCable sets the Cable property
func (m *CableTermination) SetCable(val int32) {
	m.Cable = val
}

// GetCableEnd returns the CableEnd property
func (m CableTermination) GetCableEnd() string {
	return m.CableEnd
}

// SetCableEnd sets the CableEnd property
func (m *CableTermination) SetCableEnd(val string) {
	m.CableEnd = val
}

// GetCreated returns the Created property
func (m CableTermination) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CableTermination) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m CableTermination) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CableTermination) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m CableTermination) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CableTermination) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m CableTermination) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CableTermination) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetTermination returns the Termination property
func (m CableTermination) GetTermination() map[string]interface{} {
	return m.Termination
}

// SetTermination sets the Termination property
func (m *CableTermination) SetTermination(val map[string]interface{}) {
	m.Termination = val
}

// GetTerminationId returns the TerminationId property
func (m CableTermination) GetTerminationId() int64 {
	return m.TerminationId
}

// SetTerminationId sets the TerminationId property
func (m *CableTermination) SetTerminationId(val int64) {
	m.TerminationId = val
}

// GetTerminationType returns the TerminationType property
func (m CableTermination) GetTerminationType() string {
	return m.TerminationType
}

// SetTerminationType sets the TerminationType property
func (m *CableTermination) SetTerminationType(val string) {
	m.TerminationType = val
}

// GetUrl returns the Url property
func (m CableTermination) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CableTermination) SetUrl(val string) {
	m.Url = val
}
