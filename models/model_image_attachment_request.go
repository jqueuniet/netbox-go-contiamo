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

// ImageAttachmentRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type ImageAttachmentRequest struct {
	// ContentType:
	ContentType string `json:"content_type" mapstructure:"content_type"`
	// Image:
	Image string `json:"image" mapstructure:"image"`
	// ImageHeight:
	ImageHeight int32 `json:"image_height" mapstructure:"image_height"`
	// ImageWidth:
	ImageWidth int32 `json:"image_width" mapstructure:"image_width"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// ObjectId:
	ObjectId int64 `json:"object_id" mapstructure:"object_id"`
}

// Validate implements basic validation for this model
func (m ImageAttachmentRequest) Validate() error {
	return validation.Errors{
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
	}.Filter()
}

// GetContentType returns the ContentType property
func (m ImageAttachmentRequest) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *ImageAttachmentRequest) SetContentType(val string) {
	m.ContentType = val
}

// GetImage returns the Image property
func (m ImageAttachmentRequest) GetImage() string {
	return m.Image
}

// SetImage sets the Image property
func (m *ImageAttachmentRequest) SetImage(val string) {
	m.Image = val
}

// GetImageHeight returns the ImageHeight property
func (m ImageAttachmentRequest) GetImageHeight() int32 {
	return m.ImageHeight
}

// SetImageHeight sets the ImageHeight property
func (m *ImageAttachmentRequest) SetImageHeight(val int32) {
	m.ImageHeight = val
}

// GetImageWidth returns the ImageWidth property
func (m ImageAttachmentRequest) GetImageWidth() int32 {
	return m.ImageWidth
}

// SetImageWidth sets the ImageWidth property
func (m *ImageAttachmentRequest) SetImageWidth(val int32) {
	m.ImageWidth = val
}

// GetName returns the Name property
func (m ImageAttachmentRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *ImageAttachmentRequest) SetName(val string) {
	m.Name = val
}

// GetObjectId returns the ObjectId property
func (m ImageAttachmentRequest) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *ImageAttachmentRequest) SetObjectId(val int64) {
	m.ObjectId = val
}
