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

// PatchedWritableContactRequest is an object. Adds support for custom fields and tags.
type PatchedWritableContactRequest struct {
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
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// Phone:
	Phone string `json:"phone,omitempty" mapstructure:"phone,omitempty"`
	// Tags:
	Tags []NestedTagRequest `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Title:
	Title string `json:"title,omitempty" mapstructure:"title,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedWritableContactRequest) Validate() error {
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
			m.Name, validation.Length(1, 100),
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
func (m PatchedWritableContactRequest) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *PatchedWritableContactRequest) SetAddress(val string) {
	m.Address = val
}

// GetComments returns the Comments property
func (m PatchedWritableContactRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *PatchedWritableContactRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m PatchedWritableContactRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *PatchedWritableContactRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m PatchedWritableContactRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *PatchedWritableContactRequest) SetDescription(val string) {
	m.Description = val
}

// GetEmail returns the Email property
func (m PatchedWritableContactRequest) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *PatchedWritableContactRequest) SetEmail(val string) {
	m.Email = val
}

// GetGroup returns the Group property
func (m PatchedWritableContactRequest) GetGroup() *int32 {
	return m.Group
}

// SetGroup sets the Group property
func (m *PatchedWritableContactRequest) SetGroup(val *int32) {
	m.Group = val
}

// GetLink returns the Link property
func (m PatchedWritableContactRequest) GetLink() string {
	return m.Link
}

// SetLink sets the Link property
func (m *PatchedWritableContactRequest) SetLink(val string) {
	m.Link = val
}

// GetName returns the Name property
func (m PatchedWritableContactRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedWritableContactRequest) SetName(val string) {
	m.Name = val
}

// GetPhone returns the Phone property
func (m PatchedWritableContactRequest) GetPhone() string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *PatchedWritableContactRequest) SetPhone(val string) {
	m.Phone = val
}

// GetTags returns the Tags property
func (m PatchedWritableContactRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *PatchedWritableContactRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTitle returns the Title property
func (m PatchedWritableContactRequest) GetTitle() string {
	return m.Title
}

// SetTitle sets the Title property
func (m *PatchedWritableContactRequest) SetTitle(val string) {
	m.Title = val
}
