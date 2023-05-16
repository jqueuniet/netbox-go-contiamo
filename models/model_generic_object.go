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

// GenericObject is an object. Minimal representation of some generic object identified by ContentType and PK.
type GenericObject struct {
	// Object:
	Object map[string]interface{} `json:"object" mapstructure:"object"`
	// ObjectId:
	ObjectId int32 `json:"object_id" mapstructure:"object_id"`
	// ObjectType:
	ObjectType string `json:"object_type" mapstructure:"object_type"`
}

// Validate implements basic validation for this model
func (m GenericObject) Validate() error {
	return validation.Errors{
		"object": validation.Validate(
			m.Object,
		),
	}.Filter()
}

// GetObject returns the Object property
func (m GenericObject) GetObject() map[string]interface{} {
	return m.Object
}

// SetObject sets the Object property
func (m *GenericObject) SetObject(val map[string]interface{}) {
	m.Object = val
}

// GetObjectId returns the ObjectId property
func (m GenericObject) GetObjectId() int32 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *GenericObject) SetObjectId(val int32) {
	m.ObjectId = val
}

// GetObjectType returns the ObjectType property
func (m GenericObject) GetObjectType() string {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *GenericObject) SetObjectType(val string) {
	m.ObjectType = val
}
