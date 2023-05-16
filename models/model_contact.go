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

// Contact is an object. Adds support for custom fields and tags.
type Contact struct {
	// Address:
	Address string `json:"address,omitempty" mapstructure:"address,omitempty"`
	// Comments:
	Comments string `json:"comments,omitempty" mapstructure:"comments,omitempty"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// CustomFields:
	CustomFields map[string]interface{} `json:"custom_fields,omitempty" mapstructure:"custom_fields,omitempty"`
	// Description:
	Description string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Email:
	Email string `json:"email,omitempty" mapstructure:"email,omitempty"`
	// Group: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	Group *NestedContactGroup `json:"group,omitempty" mapstructure:"group,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Link:
	Link string `json:"link,omitempty" mapstructure:"link,omitempty"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Phone:
	Phone string `json:"phone,omitempty" mapstructure:"phone,omitempty"`
	// Tags:
	Tags []NestedTag `json:"tags,omitempty" mapstructure:"tags,omitempty"`
	// Title:
	Title string `json:"title,omitempty" mapstructure:"title,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Contact) Validate() error {
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
			m.Name, validation.NotNil, validation.Length(0, 100),
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
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m Contact) GetAddress() string {
	return m.Address
}

// SetAddress sets the Address property
func (m *Contact) SetAddress(val string) {
	m.Address = val
}

// GetComments returns the Comments property
func (m Contact) GetComments() string {
	return m.Comments
}

// SetComments sets the Comments property
func (m *Contact) SetComments(val string) {
	m.Comments = val
}

// GetCreated returns the Created property
func (m Contact) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Contact) SetCreated(val *time.Time) {
	m.Created = val
}

// GetCustomFields returns the CustomFields property
func (m Contact) GetCustomFields() map[string]interface{} {
	return m.CustomFields
}

// SetCustomFields sets the CustomFields property
func (m *Contact) SetCustomFields(val map[string]interface{}) {
	m.CustomFields = val
}

// GetDescription returns the Description property
func (m Contact) GetDescription() string {
	return m.Description
}

// SetDescription sets the Description property
func (m *Contact) SetDescription(val string) {
	m.Description = val
}

// GetDisplay returns the Display property
func (m Contact) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Contact) SetDisplay(val string) {
	m.Display = val
}

// GetEmail returns the Email property
func (m Contact) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *Contact) SetEmail(val string) {
	m.Email = val
}

// GetGroup returns the Group property
func (m Contact) GetGroup() *NestedContactGroup {
	return m.Group
}

// SetGroup sets the Group property
func (m *Contact) SetGroup(val *NestedContactGroup) {
	m.Group = val
}

// GetId returns the Id property
func (m Contact) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Contact) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Contact) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Contact) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLink returns the Link property
func (m Contact) GetLink() string {
	return m.Link
}

// SetLink sets the Link property
func (m *Contact) SetLink(val string) {
	m.Link = val
}

// GetName returns the Name property
func (m Contact) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Contact) SetName(val string) {
	m.Name = val
}

// GetPhone returns the Phone property
func (m Contact) GetPhone() string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *Contact) SetPhone(val string) {
	m.Phone = val
}

// GetTags returns the Tags property
func (m Contact) GetTags() []NestedTag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Contact) SetTags(val []NestedTag) {
	m.Tags = val
}

// GetTitle returns the Title property
func (m Contact) GetTitle() string {
	return m.Title
}

// SetTitle sets the Title property
func (m *Contact) SetTitle(val string) {
	m.Title = val
}

// GetUrl returns the Url property
func (m Contact) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Contact) SetUrl(val string) {
	m.Url = val
}
