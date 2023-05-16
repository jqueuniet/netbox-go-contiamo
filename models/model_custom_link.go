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

// CustomLink is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type CustomLink struct {
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
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// GroupName: Links with the same group will appear as a dropdown menu
	GroupName string `json:"group_name,omitempty" mapstructure:"group_name,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// LinkText: Jinja2 template code for link text
	LinkText string `json:"link_text" mapstructure:"link_text"`
	// LinkUrl: Jinja2 template code for link URL
	LinkUrl string `json:"link_url" mapstructure:"link_url"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// NewWindow: Force link to open in a new window
	NewWindow bool `json:"new_window,omitempty" mapstructure:"new_window,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// Weight:
	Weight int32 `json:"weight,omitempty" mapstructure:"weight,omitempty"`
}

// Validate implements basic validation for this model
func (m CustomLink) Validate() error {
	return validation.Errors{
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"groupName": validation.Validate(
			m.GroupName, validation.Length(0, 50),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 100),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"weight": validation.Validate(
			m.Weight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
	}.Filter()
}

// GetButtonClass returns the ButtonClass property
func (m CustomLink) GetButtonClass() string {
	return m.ButtonClass
}

// SetButtonClass sets the ButtonClass property
func (m *CustomLink) SetButtonClass(val string) {
	m.ButtonClass = val
}

// GetContentTypes returns the ContentTypes property
func (m CustomLink) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *CustomLink) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetCreated returns the Created property
func (m CustomLink) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *CustomLink) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m CustomLink) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *CustomLink) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m CustomLink) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *CustomLink) SetEnabled(val bool) {
	m.Enabled = val
}

// GetGroupName returns the GroupName property
func (m CustomLink) GetGroupName() string {
	return m.GroupName
}

// SetGroupName sets the GroupName property
func (m *CustomLink) SetGroupName(val string) {
	m.GroupName = val
}

// GetId returns the Id property
func (m CustomLink) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *CustomLink) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m CustomLink) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *CustomLink) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetLinkText returns the LinkText property
func (m CustomLink) GetLinkText() string {
	return m.LinkText
}

// SetLinkText sets the LinkText property
func (m *CustomLink) SetLinkText(val string) {
	m.LinkText = val
}

// GetLinkUrl returns the LinkUrl property
func (m CustomLink) GetLinkUrl() string {
	return m.LinkUrl
}

// SetLinkUrl sets the LinkUrl property
func (m *CustomLink) SetLinkUrl(val string) {
	m.LinkUrl = val
}

// GetName returns the Name property
func (m CustomLink) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *CustomLink) SetName(val string) {
	m.Name = val
}

// GetNewWindow returns the NewWindow property
func (m CustomLink) GetNewWindow() bool {
	return m.NewWindow
}

// SetNewWindow sets the NewWindow property
func (m *CustomLink) SetNewWindow(val bool) {
	m.NewWindow = val
}

// GetUrl returns the Url property
func (m CustomLink) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *CustomLink) SetUrl(val string) {
	m.Url = val
}

// GetWeight returns the Weight property
func (m CustomLink) GetWeight() int32 {
	return m.Weight
}

// SetWeight sets the Weight property
func (m *CustomLink) SetWeight(val int32) {
	m.Weight = val
}
