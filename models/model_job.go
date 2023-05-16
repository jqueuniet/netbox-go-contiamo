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

// Job is an object.
type Job struct {
	// Completed:
	Completed *time.Time `json:"completed,omitempty" mapstructure:"completed,omitempty"`
	// Created:
	Created time.Time `json:"created" mapstructure:"created"`
	// Data:
	Data map[string]interface{} `json:"data,omitempty" mapstructure:"data,omitempty"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// Interval: Recurrence interval (in minutes)
	Interval *int32 `json:"interval,omitempty" mapstructure:"interval,omitempty"`
	// JobId:
	JobId string `json:"job_id" mapstructure:"job_id"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// ObjectId:
	ObjectId *int64 `json:"object_id,omitempty" mapstructure:"object_id,omitempty"`
	// ObjectType:
	ObjectType string `json:"object_type" mapstructure:"object_type"`
	// Scheduled:
	Scheduled *time.Time `json:"scheduled,omitempty" mapstructure:"scheduled,omitempty"`
	// Started:
	Started *time.Time `json:"started,omitempty" mapstructure:"started,omitempty"`
	// Status:
	Status struct {
		Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
		Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
	} `json:"status" mapstructure:"status"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
	// User: Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a
	// dictionary of attributes which can be used to uniquely identify the related object. This class should be
	// subclassed to return a full representation of the related object on read.
	User NestedUser `json:"user" mapstructure:"user"`
}

// Validate implements basic validation for this model
func (m Job) Validate() error {
	return validation.Errors{
		"data": validation.Validate(
			m.Data,
		),
		"interval": validation.Validate(
			m.Interval, validation.Min(*int32(1)), validation.Max(*int32(2.147483647e+09)),
		),
		"jobId": validation.Validate(
			m.JobId, validation.Required, is.UUID,
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 200),
		),
		"objectId": validation.Validate(
			m.ObjectId, validation.Min(*int64(0)), validation.Max(*int64(9.223372036854776e+18)),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
		"user": validation.Validate(
			m.User, validation.NotNil,
		),
	}.Filter()
}

// GetCompleted returns the Completed property
func (m Job) GetCompleted() *time.Time {
	return m.Completed
}

// SetCompleted sets the Completed property
func (m *Job) SetCompleted(val *time.Time) {
	m.Completed = val
}

// GetCreated returns the Created property
func (m Job) GetCreated() time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Job) SetCreated(val time.Time) {
	m.Created = val
}

// GetData returns the Data property
func (m Job) GetData() map[string]interface{} {
	return m.Data
}

// SetData sets the Data property
func (m *Job) SetData(val map[string]interface{}) {
	m.Data = val
}

// GetDisplay returns the Display property
func (m Job) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Job) SetDisplay(val string) {
	m.Display = val
}

// GetId returns the Id property
func (m Job) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Job) SetId(val int32) {
	m.Id = val
}

// GetInterval returns the Interval property
func (m Job) GetInterval() *int32 {
	return m.Interval
}

// SetInterval sets the Interval property
func (m *Job) SetInterval(val *int32) {
	m.Interval = val
}

// GetJobId returns the JobId property
func (m Job) GetJobId() string {
	return m.JobId
}

// SetJobId sets the JobId property
func (m *Job) SetJobId(val string) {
	m.JobId = val
}

// GetName returns the Name property
func (m Job) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Job) SetName(val string) {
	m.Name = val
}

// GetObjectId returns the ObjectId property
func (m Job) GetObjectId() *int64 {
	return m.ObjectId
}

// SetObjectId sets the ObjectId property
func (m *Job) SetObjectId(val *int64) {
	m.ObjectId = val
}

// GetObjectType returns the ObjectType property
func (m Job) GetObjectType() string {
	return m.ObjectType
}

// SetObjectType sets the ObjectType property
func (m *Job) SetObjectType(val string) {
	m.ObjectType = val
}

// GetScheduled returns the Scheduled property
func (m Job) GetScheduled() *time.Time {
	return m.Scheduled
}

// SetScheduled sets the Scheduled property
func (m *Job) SetScheduled(val *time.Time) {
	m.Scheduled = val
}

// GetStarted returns the Started property
func (m Job) GetStarted() *time.Time {
	return m.Started
}

// SetStarted sets the Started property
func (m *Job) SetStarted(val *time.Time) {
	m.Started = val
}

// GetStatus returns the Status property
func (m Job) GetStatus() struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
} {
	return m.Status
}

// SetStatus sets the Status property
func (m *Job) SetStatus(val struct {
	Label string `json:"label,omitempty" mapstructure:"label,omitempty"`
	Value string `json:"value,omitempty" mapstructure:"value,omitempty"`
}) {
	m.Status = val
}

// GetUrl returns the Url property
func (m Job) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Job) SetUrl(val string) {
	m.Url = val
}

// GetUser returns the User property
func (m Job) GetUser() NestedUser {
	return m.User
}

// SetUser sets the User property
func (m *Job) SetUser(val NestedUser) {
	m.User = val
}
