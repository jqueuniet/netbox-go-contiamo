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

// WebhookRequest is an object. Extends the built-in ModelSerializer to enforce calling full_clean() on a copy of the associated instance during
// validation. (DRF does not do this by default; see https://github.com/encode/django-rest-framework/issues/3144)
type WebhookRequest struct {
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
}

// Validate implements basic validation for this model
func (m WebhookRequest) Validate() error {
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
			m.HttpContentType, validation.Length(1, 100),
		),
		"name": validation.Validate(
			m.Name, validation.Required, validation.Length(1, 150),
		),
		"payloadUrl": validation.Validate(
			m.PayloadUrl, validation.Required, validation.Length(1, 500),
		),
		"secret": validation.Validate(
			m.Secret, validation.Length(0, 255),
		),
	}.Filter()
}

// GetAdditionalHeaders returns the AdditionalHeaders property
func (m WebhookRequest) GetAdditionalHeaders() string {
	return m.AdditionalHeaders
}

// SetAdditionalHeaders sets the AdditionalHeaders property
func (m *WebhookRequest) SetAdditionalHeaders(val string) {
	m.AdditionalHeaders = val
}

// GetBodyTemplate returns the BodyTemplate property
func (m WebhookRequest) GetBodyTemplate() string {
	return m.BodyTemplate
}

// SetBodyTemplate sets the BodyTemplate property
func (m *WebhookRequest) SetBodyTemplate(val string) {
	m.BodyTemplate = val
}

// GetCaFilePath returns the CaFilePath property
func (m WebhookRequest) GetCaFilePath() *string {
	return m.CaFilePath
}

// SetCaFilePath sets the CaFilePath property
func (m *WebhookRequest) SetCaFilePath(val *string) {
	m.CaFilePath = val
}

// GetConditions returns the Conditions property
func (m WebhookRequest) GetConditions() map[string]interface{} {
	return m.Conditions
}

// SetConditions sets the Conditions property
func (m *WebhookRequest) SetConditions(val map[string]interface{}) {
	m.Conditions = val
}

// GetContentTypes returns the ContentTypes property
func (m WebhookRequest) GetContentTypes() []string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *WebhookRequest) SetContentTypes(val []string) {
	m.ContentTypes = val
}

// GetEnabled returns the Enabled property
func (m WebhookRequest) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *WebhookRequest) SetEnabled(val bool) {
	m.Enabled = val
}

// GetHttpContentType returns the HttpContentType property
func (m WebhookRequest) GetHttpContentType() string {
	return m.HttpContentType
}

// SetHttpContentType sets the HttpContentType property
func (m *WebhookRequest) SetHttpContentType(val string) {
	m.HttpContentType = val
}

// GetHttpMethod returns the HttpMethod property
func (m WebhookRequest) GetHttpMethod() string {
	return m.HttpMethod
}

// SetHttpMethod sets the HttpMethod property
func (m *WebhookRequest) SetHttpMethod(val string) {
	m.HttpMethod = val
}

// GetName returns the Name property
func (m WebhookRequest) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *WebhookRequest) SetName(val string) {
	m.Name = val
}

// GetPayloadUrl returns the PayloadUrl property
func (m WebhookRequest) GetPayloadUrl() string {
	return m.PayloadUrl
}

// SetPayloadUrl sets the PayloadUrl property
func (m *WebhookRequest) SetPayloadUrl(val string) {
	m.PayloadUrl = val
}

// GetSecret returns the Secret property
func (m WebhookRequest) GetSecret() string {
	return m.Secret
}

// SetSecret sets the Secret property
func (m *WebhookRequest) SetSecret(val string) {
	m.Secret = val
}

// GetSslVerification returns the SslVerification property
func (m WebhookRequest) GetSslVerification() bool {
	return m.SslVerification
}

// SetSslVerification sets the SslVerification property
func (m *WebhookRequest) SetSslVerification(val bool) {
	m.SslVerification = val
}

// GetTypeCreate returns the TypeCreate property
func (m WebhookRequest) GetTypeCreate() bool {
	return m.TypeCreate
}

// SetTypeCreate sets the TypeCreate property
func (m *WebhookRequest) SetTypeCreate(val bool) {
	m.TypeCreate = val
}

// GetTypeDelete returns the TypeDelete property
func (m WebhookRequest) GetTypeDelete() bool {
	return m.TypeDelete
}

// SetTypeDelete sets the TypeDelete property
func (m *WebhookRequest) SetTypeDelete(val bool) {
	m.TypeDelete = val
}

// GetTypeJobEnd returns the TypeJobEnd property
func (m WebhookRequest) GetTypeJobEnd() bool {
	return m.TypeJobEnd
}

// SetTypeJobEnd sets the TypeJobEnd property
func (m *WebhookRequest) SetTypeJobEnd(val bool) {
	m.TypeJobEnd = val
}

// GetTypeJobStart returns the TypeJobStart property
func (m WebhookRequest) GetTypeJobStart() bool {
	return m.TypeJobStart
}

// SetTypeJobStart sets the TypeJobStart property
func (m *WebhookRequest) SetTypeJobStart(val bool) {
	m.TypeJobStart = val
}

// GetTypeUpdate returns the TypeUpdate property
func (m WebhookRequest) GetTypeUpdate() bool {
	return m.TypeUpdate
}

// SetTypeUpdate sets the TypeUpdate property
func (m *WebhookRequest) SetTypeUpdate(val bool) {
	m.TypeUpdate = val
}
