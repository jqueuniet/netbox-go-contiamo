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

// ObjectPermission is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ObjectPermission struct {
	// Actions: The list of actions granted by this permission
	Actions []string `json:"actions" mapstructure:"actions"`
	// Constraints: Queryset filter matching the applicable objects of the selected type(s)
	Constraints map[string]interface{} `json:"constraints,omitempty" mapstructure:"constraints,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// Groups:
	Groups []int32 `json:"groups,omitempty" mapstructure:"groups,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// ObjectTypes:
	ObjectTypes []string `json:"object_types" mapstructure:"object_types"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Users:
	Users []int32 `json:"users,omitempty" mapstructure:"users,omitempty"`
}

// Validate implements basic validation for this model
func (m ObjectPermission) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"objectTypes": validation.Validate(
			m.ObjectTypes, validation.NotNil,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"users": validation.Validate(
			m.Users,
		),
	}.Filter()
}

// GetActions returns the Actions property
func (m ObjectPermission) GetActions() []string {
	return m.Actions
}

// SetActions sets the Actions property
func (m *ObjectPermission) SetActions(val []string) {
	m.Actions = val
}

// GetConstraints returns the Constraints property
func (m ObjectPermission) GetConstraints() map[string]interface{} {
	return m.Constraints
}

// SetConstraints sets the Constraints property
func (m *ObjectPermission) SetConstraints(val map[string]interface{}) {
	m.Constraints = val
}

// GetDescription returns the Description property
func (m ObjectPermission) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ObjectPermission) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m ObjectPermission) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ObjectPermission) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m ObjectPermission) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *ObjectPermission) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroups returns the Groups property
func (m ObjectPermission) GetGroups() []int32 {
	return m.Groups
}

// SetGroups sets the Groups property
func (m *ObjectPermission) SetGroups(val []int32) {
	m.Groups = val
}

// GetId returns the Id property
func (m ObjectPermission) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ObjectPermission) SetId(val int32) {
	m.Id = val
}

// GetName returns the Name property
func (m ObjectPermission) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ObjectPermission) SetName(val string) {
	m.Name = val
}

// GetObjectTypes returns the ObjectTypes property
func (m ObjectPermission) GetObjectTypes() []string {
	return m.ObjectTypes
}

// SetObjectTypes sets the ObjectTypes property
func (m *ObjectPermission) SetObjectTypes(val []string) {
	m.ObjectTypes = val
}

// GetUrl returns the Url property
func (m ObjectPermission) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ObjectPermission) SetUrl(val string) {
	m.Url = val
}

// GetUsers returns the Users property
func (m ObjectPermission) GetUsers() []int32 {
	return m.Users
}

// SetUsers sets the Users property
func (m *ObjectPermission) SetUsers(val []int32) {
	m.Users = val
}
