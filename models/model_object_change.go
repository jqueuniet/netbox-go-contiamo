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

	"time"
)

// ObjectChange is an object.
type ObjectChange struct {
	// Action:
	Action struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"action" mapstructure:"action"`
	// ChangedObject:
	ChangedObject map[string]interface{} `json:"changed_object" mapstructure:"changed_object"`
	// ChangedObjectId:
	ChangedObjectId int64 `json:"changed_object_id" mapstructure:"changed_object_id"`
	// ChangedObjectType:
	ChangedObjectType string `json:"changed_object_type" mapstructure:"changed_object_type"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// PostchangeData:
	PostchangeData map[string]interface{} `json:"postchange_data" mapstructure:"postchange_data"`
	// PrechangeData:
	PrechangeData map[string]interface{} `json:"prechange_data" mapstructure:"prechange_data"`
	// RequestId:
	RequestId string `json:"request_id" mapstructure:"request_id"`
	// Time:
	Time time.Time `json:"time" mapstructure:"time"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUser `json:"user" mapstructure:"user"`
	// UserName:
	UserName string `json:"user_name" mapstructure:"user_name"`
}

// Validate implements basic validation for this model
func (m ObjectChange) Validate() error {
	return validation.Errors{
		"changedObject": validation.Validate(
			m.ChangedObject,
		),
		"changedObjectId": validation.Validate(
			m.ChangedObjectId, validation.Required, validation.Min(int64(0)), validation.Max(int64(9.223372036854776e+18)),
		),
		"postchangeData": validation.Validate(
			m.PostchangeData,
		),
		"prechangeData": validation.Validate(
			m.PrechangeData,
		),
		"requestId": validation.Validate(
			m.RequestId, validation.Required, is.UUID,
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetAction returns the Action property
func (m ObjectChange) GetAction() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Action
}

// SetAction sets the Action property
func (m *ObjectChange) SetAction(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Action = val
}

// GetChangedObject returns the ChangedObject property
func (m ObjectChange) GetChangedObject() map[string]interface{} {
	return m.ChangedObject
}

// SetChangedObject sets the ChangedObject property
func (m *ObjectChange) SetChangedObject(val map[string]interface{}) {
	m.ChangedObject = val
}

// GetChangedObjectId returns the ChangedObjectId property
func (m ObjectChange) GetChangedObjectId() int64 {
	return m.ChangedObjectId
}

// SetChangedObjectId sets the ChangedObjectId property
func (m *ObjectChange) SetChangedObjectId(val int64) {
	m.ChangedObjectId = val
}

// GetChangedObjectType returns the ChangedObjectType property
func (m ObjectChange) GetChangedObjectType() string {
	return m.ChangedObjectType
}

// SetChangedObjectType sets the ChangedObjectType property
func (m *ObjectChange) SetChangedObjectType(val string) {
	m.ChangedObjectType = val
}

// GetDisplay returns the Display property
func (m ObjectChange) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *ObjectChange) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m ObjectChange) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ObjectChange) SetId(val int32) {
	m.Id = val
}

// GetPostchangeData returns the PostchangeData property
func (m ObjectChange) GetPostchangeData() map[string]interface{} {
	return m.PostchangeData
}

// SetPostchangeData sets the PostchangeData property
func (m *ObjectChange) SetPostchangeData(val map[string]interface{}) {
	m.PostchangeData = val
}

// GetPrechangeData returns the PrechangeData property
func (m ObjectChange) GetPrechangeData() map[string]interface{} {
	return m.PrechangeData
}

// SetPrechangeData sets the PrechangeData property
func (m *ObjectChange) SetPrechangeData(val map[string]interface{}) {
	m.PrechangeData = val
}

// GetRequestId returns the RequestId property
func (m ObjectChange) GetRequestId() string {
	return m.RequestId
}

// SetRequestId sets the RequestId property
func (m *ObjectChange) SetRequestId(val string) {
	m.RequestId = val
}

// GetTime returns the Time property
func (m ObjectChange) GetTime() time.Time {
	return m.Time
}

// SetTime sets the Time property
func (m *ObjectChange) SetTime(val time.Time) {
	m.Time = val
}

// GetUrl returns the Url property
func (m ObjectChange) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *ObjectChange) SetUrl(val string) {
	m.Url = val
}

// GetUser returns the User property
func (m ObjectChange) GetUser() NestedUser {
	return m.User
}

// SetUser sets the User property
func (m *ObjectChange) SetUser(val NestedUser) {
	m.User = val
}

// GetUserName returns the UserName property
func (m ObjectChange) GetUserName() string {
	return m.UserName
}

// SetUserName sets the UserName property
func (m *ObjectChange) SetUserName(val string) {
	m.UserName = val
}
