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

// PatchedImageAttachmentRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type PatchedImageAttachmentRequest struct {
	// ContentType:
	ContentType string `json:"content_type,omitempty" mapstructure:"content_type,omitempty"`
	// Image:
	Image string `json:"image,omitempty" mapstructure:"image,omitempty"`
	// ImageHeight:
	ImageHeight int32 `json:"image_height,omitempty" mapstructure:"image_height,omitempty"`
	// ImageWidth:
	ImageWidth int32 `json:"image_width,omitempty" mapstructure:"image_width,omitempty"`
	// Name:
	Name string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// ObjectId:
	ObjectId int64 `json:"object_id,omitempty" mapstructure:"object_id,omitempty"`
}

// Validate implements basic validation for this model
func (m PatchedImageAttachmentRequest) Validate() error {
	return validation.Errors{
		"imageHeight": validation.Validate(
			m.ImageHeight, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"imageWidth": validation.Validate(
			m.ImageWidth, validation.Min(int32(0)), validation.Max(int32(32767)),
		),
		"name": validation.Validate(
			m.Name, validation.Length(0, 50),
		),
		"objectId": validation.Validate(
			m.ObjectId, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
	}.Filter()
}

// GetContentType returns the ContentType property
func (m PatchedImageAttachmentRequest) GetContentType() string {
	return m.ContentType
}

// SetContentType sets the ContentType property
func (m *PatchedImageAttachmentRequest) SetContentType(val string) {
	m.ContentType = val
}

// GetImage returns the Image property
func (m PatchedImageAttachmentRequest) GetImage() string {
	return m.Image
}

// SetImage sets the Image property
func (m *PatchedImageAttachmentRequest) SetImage(val string) {
	m.Image = val
}

// GetImageHeight returns the ImageHeight property
func (m PatchedImageAttachmentRequest) GetImageHeight() int32 {
	return m.ImageHeight
}

// SetImageHeight sets the ImageHeight property
func (m *PatchedImageAttachmentRequest) SetImageHeight(val int32) {
	m.ImageHeight = val
}

// GetImageWidth returns the ImageWidth property
func (m PatchedImageAttachmentRequest) GetImageWidth() int32 {
	return m.ImageWidth
}

// SetImageWidth sets the ImageWidth property
func (m *PatchedImageAttachmentRequest) SetImageWidth(val int32) {
	m.ImageWidth = val
}

// GetName returns the Name property
func (m PatchedImageAttachmentRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *PatchedImageAttachmentRequest) SetName(val string) {
	m.Name = val
}

// GetObjectId returns the ObjectId property
func (m PatchedImageAttachmentRequest) GetObjectId() int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *PatchedImageAttachmentRequest) SetObjectId(val int64) {
	m.ObjectId = val
}
