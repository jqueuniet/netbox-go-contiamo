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

// Webhook is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type Webhook struct {
	// AdditionalHeaders: User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).
	AdditionalHeaders string `json:"additional_headers,omitempty" mapstructure:"additional_headers,omitempty"`
	// BodyTemplate: Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate string `json:"body_template,omitempty" mapstructure:"body_template,omitempty"`
	// CaFilePath: The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	CaFilePath *string `json:"ca_file_path,omitempty" mapstructure:"ca_file_path,omitempty"`
	// Conditions: A set of conditions which determine whether the webhook will be generated.
	Conditions map[string]interface{} `json:"conditions,omitempty" mapstructure:"conditions,omitempty"`
	// ContentTypes:
	ContentTypes []string `json:"content_types" mapstructure:"content_types"`
	// Created:
	Created *time.Time `json:"created" mapstructure:"created"`
	// Display:
	Display string `json:"display" mapstructure:"display"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// HttpContentType: The complete list of official content types is available <a href="https://www.iana.org/assignments/media-types/media-types.xhtml">here</a>.
	HttpContentType string `json:"http_content_type,omitempty" mapstructure:"http_content_type,omitempty"`
	// HttpMethod: * `GET` - GET
	// * `POST` - POST
	// * `PUT` - PUT
	// * `PATCH` - PATCH
	// * `DELETE` - DELETE
	HttpMethod string `json:"http_method,omitempty" mapstructure:"http_method,omitempty"`
	// Id:
	Id int32 `json:"id" mapstructure:"id"`
	// LastUpdated:
	LastUpdated *time.Time `json:"last_updated" mapstructure:"last_updated"`
	// Name:
	Name string `json:"name" mapstructure:"name"`
	// PayloadUrl: This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.
	PayloadUrl string `json:"payload_url" mapstructure:"payload_url"`
	// Secret: When provided, the request will include a 'X-Hook-Signature' header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	Secret string `json:"secret,omitempty" mapstructure:"secret,omitempty"`
	// SslVerification: Enable SSL certificate verification. Disable with caution!
	SslVerification bool `json:"ssl_verification,omitempty" mapstructure:"ssl_verification,omitempty"`
	// TypeCreate: Triggers when a matching object is created.
	TypeCreate bool `json:"type_create,omitempty" mapstructure:"type_create,omitempty"`
	// TypeDelete: Triggers when a matching object is deleted.
	TypeDelete bool `json:"type_delete,omitempty" mapstructure:"type_delete,omitempty"`
	// TypeJobEnd: Triggers when a job for a matching object terminates.
	TypeJobEnd bool `json:"type_job_end,omitempty" mapstructure:"type_job_end,omitempty"`
	// TypeJobStart: Triggers when a job for a matching object is started.
	TypeJobStart bool `json:"type_job_start,omitempty" mapstructure:"type_job_start,omitempty"`
	// TypeUpdate: Triggers when a matching object is updated.
	TypeUpdate bool `json:"type_update,omitempty" mapstructure:"type_update,omitempty"`
	// Url:
	Url string `json:"url" mapstructure:"url"`
}

// Validate implements basic validation for this model
func (m Webhook) Validate() error {
	return validation.Errors{
		"caFilePath": validation.Validate(
			m.CaFilePath, validation.Length(0, 4096),
		),
		"conditions": validation.Validate(
			m.Conditions,
		),
		"contentTypes": validation.Validate(
			m.ContentTypes, validation.NotNil,
		),
		"httpContentType": validation.Validate(
			m.HttpContentType, validation.Length(0, 100),
		),
		"name": validation.Validate(
			m.Name, validation.NotNil, validation.Length(0, 150),
		),
		"payloadUrl": validation.Validate(
			m.PayloadUrl, validation.NotNil, validation.Length(0, 500),
		),
		"secret": validation.Validate(
			m.Secret, validation.Length(0, 255),
		),
		"url": validation.Validate(
			m.Url, validation.Required, is.RequestURI,
		),
	}.Filter()
}

// GetAdditionalHeaders returns the AdditionalHeaders property
func (m Webhook) GetAdditionalHeaders() string {
	return m.AdditionalHeaders
}

// SetAdditionalHeaders sets the AdditionalHeaders property
func (m *Webhook) SetAdditionalHeaders(val string) {
	m.AdditionalHeaders = val
}

// GetBodyTemplate returns the BodyTemplate property
func (m Webhook) GetBodyTemplate() string {
	return m.BodyTemplate
}

// SetBodyTemplate sets the BodyTemplate property
func (m *Webhook) SetBodyTemplate(val string) {
	m.BodyTemplate = val
}

// GetCaFilePath returns the CaFilePath property
func (m Webhook) GetCaFilePath() *string {
	return m.CaFilePath
}

// SetCaFilePath sets the CaFilePath property
func (m *Webhook) SetCaFilePath(val *string) {
	m.CaFilePath = val
}

// GetConditions returns the Conditions property
func (m Webhook) GetConditions() map[string]interface{} {
	return m.Conditions
}

// SetConditions sets the Conditions property
func (m *Webhook) SetConditions(val map[string]interface{}) {
	m.Conditions = val
}

// GetContentTypes returns the ContentTypes property
func (m Webhook) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *Webhook) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetCreated returns the Created property
func (m Webhook) GetCreated() *time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *Webhook) SetCreated(val *time.Time) {
	m.Created = val
}

// GetDisplay returns the Display property
func (m Webhook) GetDisplay() string {
	return m.Display
}

// SetDisplay sets the Display property
func (m *Webhook) SetDisplay(val string) {
	m.Display = val
}

// GetEnabled returns the Enabled property
func (m Webhook) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *Webhook) SetEnabled(val bool) {
	m.Enabled = val
}

// GetHttpContentType returns the HttpContentType property
func (m Webhook) GetHttpContentType() string {
	return m.HttpContentType
}

// SetHttpContentType sets the HttpContentType property
func (m *Webhook) SetHttpContentType(val string) {
	m.HttpContentType = val
}

// GetHttpMethod returns the HttpMethod property
func (m Webhook) GetHttpMethod() string {
	return m.HttpMethod
}

// SetHttpMethod sets the HttpMethod property
func (m *Webhook) SetHttpMethod(val string) {
	m.HttpMethod = val
}

// GetId returns the Id property
func (m Webhook) GetId() int32 {
	return m.Id
}

// SetId sets the Id property
func (m *Webhook) SetId(val int32) {
	m.Id = val
}

// GetLastUpdated returns the LastUpdated property
func (m Webhook) GetLastUpdated() *time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *Webhook) SetLastUpdated(val *time.Time) {
	m.LastUpdated = val
}

// GetName returns the Name property
func (m Webhook) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Webhook) SetName(val string) {
	m.Name = val
}

// GetPayloadUrl returns the PayloadUrl property
func (m Webhook) GetPayloadUrl() string {
	return m.PayloadUrl
}

// SetPayloadUrl sets the PayloadUrl property
func (m *Webhook) SetPayloadUrl(val string) {
	m.PayloadUrl = val
}

// GetSecret returns the Secret property
func (m Webhook) GetSecret() string {
	return m.Secret
}

// SetSecret sets the Secret property
func (m *Webhook) SetSecret(val string) {
	m.Secret = val
}

// GetSslVerification returns the SslVerification property
func (m Webhook) GetSslVerification() bool {
	return m.SslVerification
}

// SetSslVerification sets the SslVerification property
func (m *Webhook) SetSslVerification(val bool) {
	m.SslVerification = val
}

// GetTypeCreate returns the TypeCreate property
func (m Webhook) GetTypeCreate() bool {
	return m.TypeCreate
}

// SetTypeCreate sets the TypeCreate property
func (m *Webhook) SetTypeCreate(val bool) {
	m.TypeCreate = val
}

// GetTypeDelete returns the TypeDelete property
func (m Webhook) GetTypeDelete() bool {
	return m.TypeDelete
}

// SetTypeDelete sets the TypeDelete property
func (m *Webhook) SetTypeDelete(val bool) {
	m.TypeDelete = val
}

// GetTypeJobEnd returns the TypeJobEnd property
func (m Webhook) GetTypeJobEnd() bool {
	return m.TypeJobEnd
}

// SetTypeJobEnd sets the TypeJobEnd property
func (m *Webhook) SetTypeJobEnd(val bool) {
	m.TypeJobEnd = val
}

// GetTypeJobStart returns the TypeJobStart property
func (m Webhook) GetTypeJobStart() bool {
	return m.TypeJobStart
}

// SetTypeJobStart sets the TypeJobStart property
func (m *Webhook) SetTypeJobStart(val bool) {
	m.TypeJobStart = val
}

// GetTypeUpdate returns the TypeUpdate property
func (m Webhook) GetTypeUpdate() bool {
	return m.TypeUpdate
}

// SetTypeUpdate sets the TypeUpdate property
func (m *Webhook) SetTypeUpdate(val bool) {
	m.TypeUpdate = val
}

// GetUrl returns the Url property
func (m Webhook) GetUrl() string {
	return m.Url
}

// SetUrl sets the Url property
func (m *Webhook) SetUrl(val string) {
	m.Url = val
}
