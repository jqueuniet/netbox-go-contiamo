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

// WritableContactRequest is an object. Adds support for custom fields and tags.
type WritableContactRequest struct {
	// Address:
	Address string `json:"address,omitempty" mapstructure:"address,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Email:
	Email string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// Group:
	Group *int32 `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Link:
	Link string `json:"link,omitempty" mapstructure:"link,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Phone:
	Phone string `json:"phone,omitempty" mapstructure:"phone,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Title:
	Title string `json:"title,omitempty" mapstructure:"title,omitempty"`
}

// Validate implements basic validation for this model
func (m WritableContactRequest) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address, validation.Length(0, 200),
		),
		"customFields": validation.Validate(
			m.CustomFields,
		),
		"description": validation.Validate(
			m.Description, validation.Length(0, 200),
		),
		"email": validation.Validate(
			m.Email, validation.Length(0, 254), is.EmailFormat,
		),
		"link": validation.Validate(
			m.Link, validation.Length(0, 200), is.RequestURI,
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"phone": validation.Validate(
			m.Phone, validation.Length(0, 50),
		),
		"tags": validation.Validate(
			m.Tags,
		),
		"title": validation.Validate(
			m.Title, validation.Length(0, 100),
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m WritableContactRequest) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *WritableContactRequest) SetAddress(val string) {
	m.Address = val
}

// GetComments returns the Comments property
func (m WritableContactRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *WritableContactRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m WritableContactRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *WritableContactRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m WritableContactRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *WritableContactRequest) SetDescription(val string) {
	m.Description = val
}

// GetEmail returns the Email property
func (m WritableContactRequest) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *WritableContactRequest) SetEmail(val string) {
	m.Email = val
}

// GetGroup returns the Group property
func (m WritableContactRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *WritableContactRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetLink returns the Link property
func (m WritableContactRequest) GetLink() string {
	return m.Link
}

// SetLink sets the Link property
func (m *WritableContactRequest) SetLink(val string) {
	m.Link = val
}

// GetName returns the Name property
func (m WritableContactRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WritableContactRequest) SetName(val string) {
	m.Name = val
}

// GetPhone returns the Phone property
func (m WritableContactRequest) GetPhone() string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *WritableContactRequest) SetPhone(val string) {
	m.Phone = val
}

// GetTags returns the Tags property
func (m WritableContactRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *WritableContactRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTitle returns the Title property
func (m WritableContactRequest) GetTitle() string {
	return m.Title
}

// SetTitle sets the Title property
func (m *WritableContactRequest) SetTitle(val string) {
	m.Title = val
}
