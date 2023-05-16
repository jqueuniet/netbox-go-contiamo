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

// GenericObjectRequest is an object. Minimal representation of some generic object identified by ContentType and PK.
type GenericObjectRequest struct {
	// ObjectId:
	ObjectId int32 `json:"object_id" mapstructure:"object_id"`
	// ObjectType:
	ObjectType string `json:"object_type" mapstructure:"object_type"`
}

// Validate implements basic validation for this model
func (m GenericObjectRequest) Validate() error {
	return validation.Errors{}.Filter()
}

// GetObjectId returns the ObjectId property
func (m GenericObjectRequest) GetObjectId() int32 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *GenericObjectRequest) SetObjectId(val int32) {
	m.ObjectId = val
}

// GetObjectType returns the ObjectType property
func (m GenericObjectRequest) GetObjectType() string {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *GenericObjectRequest) SetObjectType(val string) {
	m.ObjectType = val
}
