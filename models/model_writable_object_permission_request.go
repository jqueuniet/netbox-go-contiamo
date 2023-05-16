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

// WritableObjectPermissionRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WritableObjectPermissionRequest struct {
	// Actions: The list of actions granted by this permission
	Actions []string `json:"actions" mapstructure:"actions"`
	// Constraints: Queryset filter matching the applicable objects of the selected type(s)
	Constraints map[string]interface{} `json:"constraints,omitempty" mapstructure:"constraints,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Groups:
	Groups []int32 `json:"groups,omitempty" mapstructure:"groups,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// ObjectTypes:
	ObjectTypes []string `json:"object_types" mapstructure:"object_types"`
	// Users:
	Users []int32 `json:"users,omitempty" mapstructure:"users,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableObjectPermissionRequest) Validate() error {
	return validation.Errors{
		"actions": validation.Validate(
			m.Actions, validation.NotNil,
		),
		"constraints": validation.Validate(
			m.Constraints,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"groups": validation.Validate(
			m.Groups,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"objectTypes": validation.Validate(
			m.ObjectTypes, validation.NotNil,
		),
		"users": validation.Validate(
			m.Users,
		),
	}.Filter()
}

// GetActions returns the Actions property
func (m WritableObjectPermissionRequest) GetActions() []string {
	return m.Actions
}

// SetActions sets the Actions property
func (m *WritableObjectPermissionRequest) SetActions(val []string) {
	m.Actions = val
}

// GetConstraints returns the Constraints property
func (m WritableObjectPermissionRequest) GetConstraints() map[string]interface{} {
	return m.Constraints
}

// SetConstraints sets the Constraints property
func (m *WritableObjectPermissionRequest) SetConstraints(val map[string]interface{}) {
	m.Constraints = val
}

// GetDescription returns the Description property
func (m WritableObjectPermissionRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableObjectPermissionRequest) SetDescription(val string) {
	m.Description = val
}

// GetEnabled returns the Enabled property
func (m WritableObjectPermissionRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *WritableObjectPermissionRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroups returns the Groups property
func (m WritableObjectPermissionRequest) GetGroups() []int32 {
	return m.Groups
}

// SetGroups sets the Groups property
func (m *WritableObjectPermissionRequest) SetGroups(val []int32) {
	m.Groups = val
}

// GetName returns the Name property
func (m WritableObjectPermissionRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableObjectPermissionRequest) SetName(val string) {
	m.Name = val
}

// GetObjectTypes returns the ObjectTypes property
func (m WritableObjectPermissionRequest) GetObjectTypes() []string {
	return m.ObjectTypes
}

// SetObjectTypes sets the ObjectTypes property
func (m *WritableObjectPermissionRequest) SetObjectTypes(val []string) {
	m.ObjectTypes = val
}

// GetUsers returns the Users property
func (m WritableObjectPermissionRequest) GetUsers() []int32 {
	return m.Users
}

// SetUsers sets the Users property
func (m *WritableObjectPermissionRequest) SetUsers(val []int32) {
	m.Users = val
}
