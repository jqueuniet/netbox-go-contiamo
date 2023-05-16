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

// ImageAttachment is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ImageAttachment struct {
	// ContentType:
	ContentType string `json:"content_type" mapstructure:"content_type"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Image:
	Image string `json:"image" mapstructure:"image"`
	// ImageHeight:
	ImageHeight int32 `json:"image_height" mapstructure:"image_height"`
	// ImageWidth:
	ImageWidth int32 `json:"image_width" mapstructure:"image_width"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// ObjectId:
	ObjectId int64 `json:"object_id" mapstructure:"object_id"`
	// Parent:
	Parent map[string]interface{} `json:"parent" mapstructure:"parent"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m ImageAttachment) Validate() error {
	return validation.Errors{
		"image": validation.Validate(
			m.Image, validation.Required, is.RequestURI,
		),
		"imageHeight": validation.Validate(
			m.ImageHeight, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"imageWidth": validation.Validate(
			m.ImageWidth, validation.NotNil, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 50),
		),
		"objectId": validation.Validate(
			m.ObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"parent": validation.Validate(
			m.Parent,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetContentType returns the ContentType property
func (m ImageAttachment) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *ImageAttachment) SetContentType(val string) {
	m.ContentType = val
}

// GetCreated returns the Created property
func (m ImageAttachment) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *ImageAttachment) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m ImageAttachment) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ImageAttachment) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ImageAttachment) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ImageAttachment) SetId(val int32) {
	m.Id = val
}

// GetImage returns the Image property
func (m ImageAttachment) GetImage() string {
	return m.Image
}

// SetImage sets the Image property
func (m *ImageAttachment) SetImage(val string) {
	m.Image = val
}

// GetImageHeight returns the ImageHeight property
func (m ImageAttachment) GetImageHeight() int32 {
	return m.ImageHeight
}

// SetImageHeight sets the ImageHeight property
func (m *ImageAttachment) SetImageHeight(val int32) {
	m.ImageHeight = val
}

// GetImageWidth returns the ImageWidth property
func (m ImageAttachment) GetImageWidth() int32 {
	return m.ImageWidth
}

// SetImageWidth sets the ImageWidth property
func (m *ImageAttachment) SetImageWidth(val int32) {
	m.ImageWidth = val
}

// GetLastUpdated returns the LastUpdated property
func (m ImageAttachment) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *ImageAttachment) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m ImageAttachment) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ImageAttachment) SetName(val string) {
	m.Name = val
}

// GetObjectId returns the ObjectId property
func (m ImageAttachment) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *ImageAttachment) SetObjectId(val int64) {
	m.ObjectId = val
}

// GetParent returns the Parent property
func (m ImageAttachment) GetParent() map[string]interface{} {
	return m.Parent
}

// SetParent sets the Parent property
func (m *ImageAttachment) SetParent(val map[string]interface{}) {
	m.Parent = val
}

// GetUrl returns the Url property
func (m ImageAttachment) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ImageAttachment) SetUrl(val string) {
	m.Url = val
}
