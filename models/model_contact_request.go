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

// ContactRequest is an object. Adds support for custom fields and tags.
type ContactRequest struct {
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
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedContactGroupRequest `json:"group,omitempty" mapstructure:"group,omitempty"`
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
func (m ContactRequest) Validate() error {
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
		"group": validation.Validate(
			m.Group,
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
func (m ContactRequest) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *ContactRequest) SetAddress(val string) {
	m.Address = val
}

// GetComments returns the Comments property
func (m ContactRequest) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *ContactRequest) SetComments(val string) {
	m.Comments = val
}

// GetCustomFields returns the CustomFields property
func (m ContactRequest) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *ContactRequest) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m ContactRequest) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ContactRequest) SetDescription(val string) {
	m.Description = val
}

// GetEmail returns the Email property
func (m ContactRequest) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *ContactRequest) SetEmail(val string) {
	m.Email = val
}

// GetGroup returns the Group property
func (m ContactRequest) GetGroup() *NestedContactGroupRequest {
	return m.Group
}

// SetGroup sets the Group property
func (m *ContactRequest) SetGroup(val *NestedContactGroupRequest) {
	m.Group = val
}

// GetLink returns the Link property
func (m ContactRequest) GetLink() string {
	return m.Link
}

// SetLink sets the Link property
func (m *ContactRequest) SetLink(val string) {
	m.Link = val
}

// GetName returns the Name property
func (m ContactRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ContactRequest) SetName(val string) {
	m.Name = val
}

// GetPhone returns the Phone property
func (m ContactRequest) GetPhone() string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *ContactRequest) SetPhone(val string) {
	m.Phone = val
}

// GetTags returns the Tags property
func (m ContactRequest) GetTags() []NestedTagRequest {
	return m.Tags
}

// SetTags sets the Tags property
func (m *ContactRequest) SetTags(val []NestedTagRequest) {
	m.Tags = val
}

// GetTitle returns the Title property
func (m ContactRequest) GetTitle() string {
	return m.Title
}

// SetTitle sets the Title property
func (m *ContactRequest) SetTitle(val string) {
	m.Title = val
}
