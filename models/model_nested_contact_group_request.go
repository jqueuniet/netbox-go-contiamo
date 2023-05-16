// This file is auto-generated, DO NOT EDIT.
//
// Source:
//
//	Title: NetBox REST API
//	Version: 3.5.1 (3.5)
package models

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
)

// nestedContactGroupRequestSlugPattern is the validation pattern for NestedContactGroupRequest.Slug
var nestedContactGroupRequestSlugPattern = regexp.MustCompile(`^[-a-zA-Z0-9_]+$`)

// NestedContactGroupRequest is an object. Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
// dictionary of attributes which can be used to uniquely identify the related object. This class should be
// subclassed to return a full representation of the related object on read.
type NestedContactGroupRequest struct {
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// Slug:
	Slug string `json:"slug" mapstructure:"slug"`
}

// Validate implements basic validation for this model
func (m NestedContactGroupRequest) Validate() error {
	return validation.Errors{
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 100),
		),
		"slug": validation.Validate(
			m.Slug, validation.Required, validation.Length(1, 100), validation.Match(nestedContactGroupRequestSlugPattern),
		),
	}.Filter()
}

// GetName returns the Name property
func (m NestedContactGroupRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *NestedContactGroupRequest) SetName(val string) {
	m.Name = val
}

// GetSlug returns the Slug property
func (m NestedContactGroupRequest) GetSlug() string {
	return m.Slug
}

// SetSlug sets the Slug property
func (m *NestedContactGroupRequest) SetSlug(val string) {
	m.Slug = val
}
