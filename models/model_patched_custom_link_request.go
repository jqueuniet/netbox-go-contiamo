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

// PatchedCustomLinkRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedCustomLinkRequest struct {
	// ButtonClass: The class of the first link in a group will be used for the dropdown button
	//
	// * `outline-dark` - Default
	// * `blue` - Blue
	// * `indigo` - Indigo
	// * `purple` - Purple
	// * `pink` - Pink
	// * `red` - Red
	// * `orange` - Orange
	// * `yellow` - Yellow
	// * `green` - Green
	// * `teal` - Teal
	// * `cyan` - Cyan
	// * `gray` - Gray
	// * `black` - Black
	// * `white` - White
	// * `ghost-dark` - Link
	ButtonClass string `json:"button_class,omitempty" mapstructure:"button_class,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types,omitempty" mapstructure:"content_types,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// GroupName: Links with the same group will appear as a dropdown menu
	GroupName string `json:"group_name,omitempty" mapstructure:"group_name,omitempty"`
	// LinkText: Jinja2 template code for link text
	LinkText string `json:"link_text,omitempty" mapstructure:"link_text,omitempty"`
	// LinkUrl: Jinja2 template code for link URL
	LinkUrl string `json:"link_url,omitempty" mapstructure:"link_url,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// NewWindow: Force link to open in a new window
	NewWindow bool `json:"new_window,omitempty" mapstructure:"new_window,omitempty"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedCustomLinkRequest) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes,
		),
		"groupName": validation.Validate(
			m.GroupName, validation.Length(0, 50),
		),
		"linkText": validation.Validate(
			m.LinkText, validation.Length(1, 0),
		),
		"linkUrl": validation.Validate(
			m.LinkUrl, validation.Length(1, 0),
		),
		"name": validation.Validate(
			m.Name, validation.Length(1, 100),
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetButtonClass returns the ButtonClass property
func (m PatchedCustomLinkRequest) GetButtonClass() string {
	return m.ButtonClass
}

// SetButtonClass sets the ButtonClass property
func (m *PatchedCustomLinkRequest) SetButtonClass(val string) {
	m.ButtonClass = val
}

// GetContentTypes returns the ContentTypes property
func (m PatchedCustomLinkRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *PatchedCustomLinkRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetEnabled returns the Enabled property
func (m PatchedCustomLinkRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *PatchedCustomLinkRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroupName returns the GroupName property
func (m PatchedCustomLinkRequest) GetGroupName() string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *PatchedCustomLinkRequest) SetGroupName(val string) {
	m.GroupName = val
}

// GetLinkText returns the LinkText property
func (m PatchedCustomLinkRequest) GetLinkText() string {
	return m.LinkText
}

// SetLinkText sets the LinkText property
func (m *PatchedCustomLinkRequest) SetLinkText(val string) {
	m.LinkText = val
}

// GetLinkUrl returns the LinkUrl property
func (m PatchedCustomLinkRequest) GetLinkUrl() string {
	return m.LinkUrl
}

// SetLinkUrl sets the LinkUrl property
func (m *PatchedCustomLinkRequest) SetLinkUrl(val string) {
	m.LinkUrl = val
}

// GetName returns the Name property
func (m PatchedCustomLinkRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedCustomLinkRequest) SetName(val string) {
	m.Name = val
}

// GetNewWindow returns the NewWindow property
func (m PatchedCustomLinkRequest) GetNewWindow() bool {
	return m.NewWindow
}

// SetNewWindow sets the NewWindow property
func (m *PatchedCustomLinkRequest) SetNewWindow(val bool) {
	m.NewWindow = val
}

// GetWeight returns the Weight property
func (m PatchedCustomLinkRequest) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *PatchedCustomLinkRequest) SetWeight(val int32) {
	m.Weight = val
}
