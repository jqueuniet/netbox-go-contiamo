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

// ExtrasWebhooksListQueryParameters is an object.
type ExtrasWebhooksListQueryParameters struct {
	// CaFilePath:
	CaFilePath []string `json:"ca_file_path,omitempty" mapstructure:"ca_file_path,omitempty"`
	// CaFilePathEmpty:
	CaFilePathEmpty []string `json:"ca_file_path__empty,omitempty" mapstructure:"ca_file_path__empty,omitempty"`
	// CaFilePathIc:
	CaFilePathIc []string `json:"ca_file_path__ic,omitempty" mapstructure:"ca_file_path__ic,omitempty"`
	// CaFilePathIe:
	CaFilePathIe []string `json:"ca_file_path__ie,omitempty" mapstructure:"ca_file_path__ie,omitempty"`
	// CaFilePathIew:
	CaFilePathIew []string `json:"ca_file_path__iew,omitempty" mapstructure:"ca_file_path__iew,omitempty"`
	// CaFilePathIsw:
	CaFilePathIsw []string `json:"ca_file_path__isw,omitempty" mapstructure:"ca_file_path__isw,omitempty"`
	// CaFilePathN:
	CaFilePathN []string `json:"ca_file_path__n,omitempty" mapstructure:"ca_file_path__n,omitempty"`
	// CaFilePathNic:
	CaFilePathNic []string `json:"ca_file_path__nic,omitempty" mapstructure:"ca_file_path__nic,omitempty"`
	// CaFilePathNie:
	CaFilePathNie []string `json:"ca_file_path__nie,omitempty" mapstructure:"ca_file_path__nie,omitempty"`
	// CaFilePathNiew:
	CaFilePathNiew []string `json:"ca_file_path__niew,omitempty" mapstructure:"ca_file_path__niew,omitempty"`
	// CaFilePathNisw:
	CaFilePathNisw []string `json:"ca_file_path__nisw,omitempty" mapstructure:"ca_file_path__nisw,omitempty"`
	// ContentTypeId:
	ContentTypeId []int32 `json:"content_type_id,omitempty" mapstructure:"content_type_id,omitempty"`
	// ContentTypeIdGt:
	ContentTypeIdGt []int32 `json:"content_type_id__gt,omitempty" mapstructure:"content_type_id__gt,omitempty"`
	// ContentTypeIdGte:
	ContentTypeIdGte []int32 `json:"content_type_id__gte,omitempty" mapstructure:"content_type_id__gte,omitempty"`
	// ContentTypeIdLt:
	ContentTypeIdLt []int32 `json:"content_type_id__lt,omitempty" mapstructure:"content_type_id__lt,omitempty"`
	// ContentTypeIdLte:
	ContentTypeIdLte []int32 `json:"content_type_id__lte,omitempty" mapstructure:"content_type_id__lte,omitempty"`
	// ContentTypeIdN:
	ContentTypeIdN []int32 `json:"content_type_id__n,omitempty" mapstructure:"content_type_id__n,omitempty"`
	// ContentTypes:
	ContentTypes string `json:"content_types,omitempty" mapstructure:"content_types,omitempty"`
	// ContentTypesIc:
	ContentTypesIc string `json:"content_types__ic,omitempty" mapstructure:"content_types__ic,omitempty"`
	// ContentTypesIe:
	ContentTypesIe string `json:"content_types__ie,omitempty" mapstructure:"content_types__ie,omitempty"`
	// ContentTypesIew:
	ContentTypesIew string `json:"content_types__iew,omitempty" mapstructure:"content_types__iew,omitempty"`
	// ContentTypesIsw:
	ContentTypesIsw string `json:"content_types__isw,omitempty" mapstructure:"content_types__isw,omitempty"`
	// ContentTypesN:
	ContentTypesN string `json:"content_types__n,omitempty" mapstructure:"content_types__n,omitempty"`
	// ContentTypesNic:
	ContentTypesNic string `json:"content_types__nic,omitempty" mapstructure:"content_types__nic,omitempty"`
	// ContentTypesNie:
	ContentTypesNie string `json:"content_types__nie,omitempty" mapstructure:"content_types__nie,omitempty"`
	// ContentTypesNiew:
	ContentTypesNiew string `json:"content_types__niew,omitempty" mapstructure:"content_types__niew,omitempty"`
	// ContentTypesNisw:
	ContentTypesNisw string `json:"content_types__nisw,omitempty" mapstructure:"content_types__nisw,omitempty"`
	// Enabled:
	Enabled bool `json:"enabled,omitempty" mapstructure:"enabled,omitempty"`
	// HttpContentType:
	HttpContentType []string `json:"http_content_type,omitempty" mapstructure:"http_content_type,omitempty"`
	// HttpContentTypeEmpty:
	HttpContentTypeEmpty []string `json:"http_content_type__empty,omitempty" mapstructure:"http_content_type__empty,omitempty"`
	// HttpContentTypeIc:
	HttpContentTypeIc []string `json:"http_content_type__ic,omitempty" mapstructure:"http_content_type__ic,omitempty"`
	// HttpContentTypeIe:
	HttpContentTypeIe []string `json:"http_content_type__ie,omitempty" mapstructure:"http_content_type__ie,omitempty"`
	// HttpContentTypeIew:
	HttpContentTypeIew []string `json:"http_content_type__iew,omitempty" mapstructure:"http_content_type__iew,omitempty"`
	// HttpContentTypeIsw:
	HttpContentTypeIsw []string `json:"http_content_type__isw,omitempty" mapstructure:"http_content_type__isw,omitempty"`
	// HttpContentTypeN:
	HttpContentTypeN []string `json:"http_content_type__n,omitempty" mapstructure:"http_content_type__n,omitempty"`
	// HttpContentTypeNic:
	HttpContentTypeNic []string `json:"http_content_type__nic,omitempty" mapstructure:"http_content_type__nic,omitempty"`
	// HttpContentTypeNie:
	HttpContentTypeNie []string `json:"http_content_type__nie,omitempty" mapstructure:"http_content_type__nie,omitempty"`
	// HttpContentTypeNiew:
	HttpContentTypeNiew []string `json:"http_content_type__niew,omitempty" mapstructure:"http_content_type__niew,omitempty"`
	// HttpContentTypeNisw:
	HttpContentTypeNisw []string `json:"http_content_type__nisw,omitempty" mapstructure:"http_content_type__nisw,omitempty"`
	// HttpMethod: * `GET` - GET
	// * `POST` - POST
	// * `PUT` - PUT
	// * `PATCH` - PATCH
	// * `DELETE` - DELETE
	HttpMethod []string `json:"http_method,omitempty" mapstructure:"http_method,omitempty"`
	// HttpMethodN: * `GET` - GET
	// * `POST` - POST
	// * `PUT` - PUT
	// * `PATCH` - PATCH
	// * `DELETE` - DELETE
	HttpMethodN []string `json:"http_method__n,omitempty" mapstructure:"http_method__n,omitempty"`
	// Id:
	Id []int32 `json:"id,omitempty" mapstructure:"id,omitempty"`
	// IdGt:
	IdGt []int32 `json:"id__gt,omitempty" mapstructure:"id__gt,omitempty"`
	// IdGte:
	IdGte []int32 `json:"id__gte,omitempty" mapstructure:"id__gte,omitempty"`
	// IdLt:
	IdLt []int32 `json:"id__lt,omitempty" mapstructure:"id__lt,omitempty"`
	// IdLte:
	IdLte []int32 `json:"id__lte,omitempty" mapstructure:"id__lte,omitempty"`
	// IdN:
	IdN []int32 `json:"id__n,omitempty" mapstructure:"id__n,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Name:
	Name []string `json:"name,omitempty" mapstructure:"name,omitempty"`
	// NameEmpty:
	NameEmpty []string `json:"name__empty,omitempty" mapstructure:"name__empty,omitempty"`
	// NameIc:
	NameIc []string `json:"name__ic,omitempty" mapstructure:"name__ic,omitempty"`
	// NameIe:
	NameIe []string `json:"name__ie,omitempty" mapstructure:"name__ie,omitempty"`
	// NameIew:
	NameIew []string `json:"name__iew,omitempty" mapstructure:"name__iew,omitempty"`
	// NameIsw:
	NameIsw []string `json:"name__isw,omitempty" mapstructure:"name__isw,omitempty"`
	// NameN:
	NameN []string `json:"name__n,omitempty" mapstructure:"name__n,omitempty"`
	// NameNic:
	NameNic []string `json:"name__nic,omitempty" mapstructure:"name__nic,omitempty"`
	// NameNie:
	NameNie []string `json:"name__nie,omitempty" mapstructure:"name__nie,omitempty"`
	// NameNiew:
	NameNiew []string `json:"name__niew,omitempty" mapstructure:"name__niew,omitempty"`
	// NameNisw:
	NameNisw []string `json:"name__nisw,omitempty" mapstructure:"name__nisw,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// PayloadUrl:
	PayloadUrl []string `json:"payload_url,omitempty" mapstructure:"payload_url,omitempty"`
	// PayloadUrlEmpty:
	PayloadUrlEmpty []string `json:"payload_url__empty,omitempty" mapstructure:"payload_url__empty,omitempty"`
	// PayloadUrlIc:
	PayloadUrlIc []string `json:"payload_url__ic,omitempty" mapstructure:"payload_url__ic,omitempty"`
	// PayloadUrlIe:
	PayloadUrlIe []string `json:"payload_url__ie,omitempty" mapstructure:"payload_url__ie,omitempty"`
	// PayloadUrlIew:
	PayloadUrlIew []string `json:"payload_url__iew,omitempty" mapstructure:"payload_url__iew,omitempty"`
	// PayloadUrlIsw:
	PayloadUrlIsw []string `json:"payload_url__isw,omitempty" mapstructure:"payload_url__isw,omitempty"`
	// PayloadUrlN:
	PayloadUrlN []string `json:"payload_url__n,omitempty" mapstructure:"payload_url__n,omitempty"`
	// PayloadUrlNic:
	PayloadUrlNic []string `json:"payload_url__nic,omitempty" mapstructure:"payload_url__nic,omitempty"`
	// PayloadUrlNie:
	PayloadUrlNie []string `json:"payload_url__nie,omitempty" mapstructure:"payload_url__nie,omitempty"`
	// PayloadUrlNiew:
	PayloadUrlNiew []string `json:"payload_url__niew,omitempty" mapstructure:"payload_url__niew,omitempty"`
	// PayloadUrlNisw:
	PayloadUrlNisw []string `json:"payload_url__nisw,omitempty" mapstructure:"payload_url__nisw,omitempty"`
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
	// Secret:
	Secret []string `json:"secret,omitempty" mapstructure:"secret,omitempty"`
	// SecretEmpty:
	SecretEmpty []string `json:"secret__empty,omitempty" mapstructure:"secret__empty,omitempty"`
	// SecretIc:
	SecretIc []string `json:"secret__ic,omitempty" mapstructure:"secret__ic,omitempty"`
	// SecretIe:
	SecretIe []string `json:"secret__ie,omitempty" mapstructure:"secret__ie,omitempty"`
	// SecretIew:
	SecretIew []string `json:"secret__iew,omitempty" mapstructure:"secret__iew,omitempty"`
	// SecretIsw:
	SecretIsw []string `json:"secret__isw,omitempty" mapstructure:"secret__isw,omitempty"`
	// SecretN:
	SecretN []string `json:"secret__n,omitempty" mapstructure:"secret__n,omitempty"`
	// SecretNic:
	SecretNic []string `json:"secret__nic,omitempty" mapstructure:"secret__nic,omitempty"`
	// SecretNie:
	SecretNie []string `json:"secret__nie,omitempty" mapstructure:"secret__nie,omitempty"`
	// SecretNiew:
	SecretNiew []string `json:"secret__niew,omitempty" mapstructure:"secret__niew,omitempty"`
	// SecretNisw:
	SecretNisw []string `json:"secret__nisw,omitempty" mapstructure:"secret__nisw,omitempty"`
	// SslVerification:
	SslVerification bool `json:"ssl_verification,omitempty" mapstructure:"ssl_verification,omitempty"`
	// TypeCreate:
	TypeCreate bool `json:"type_create,omitempty" mapstructure:"type_create,omitempty"`
	// TypeDelete:
	TypeDelete bool `json:"type_delete,omitempty" mapstructure:"type_delete,omitempty"`
	// TypeJobEnd:
	TypeJobEnd bool `json:"type_job_end,omitempty" mapstructure:"type_job_end,omitempty"`
	// TypeJobStart:
	TypeJobStart bool `json:"type_job_start,omitempty" mapstructure:"type_job_start,omitempty"`
	// TypeUpdate:
	TypeUpdate bool `json:"type_update,omitempty" mapstructure:"type_update,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasWebhooksListQueryParameters) Validate() error {
	return validation.Errors{
		"caFilePath": validation.Validate(
			m.CaFilePath,
		),
		"caFilePathEmpty": validation.Validate(
			m.CaFilePathEmpty,
		),
		"caFilePathIc": validation.Validate(
			m.CaFilePathIc,
		),
		"caFilePathIe": validation.Validate(
			m.CaFilePathIe,
		),
		"caFilePathIew": validation.Validate(
			m.CaFilePathIew,
		),
		"caFilePathIsw": validation.Validate(
			m.CaFilePathIsw,
		),
		"caFilePathN": validation.Validate(
			m.CaFilePathN,
		),
		"caFilePathNic": validation.Validate(
			m.CaFilePathNic,
		),
		"caFilePathNie": validation.Validate(
			m.CaFilePathNie,
		),
		"caFilePathNiew": validation.Validate(
			m.CaFilePathNiew,
		),
		"caFilePathNisw": validation.Validate(
			m.CaFilePathNisw,
		),
		"contentTypeId": validation.Validate(
			m.ContentTypeId,
		),
		"contentTypeIdGt": validation.Validate(
			m.ContentTypeIdGt,
		),
		"contentTypeIdGte": validation.Validate(
			m.ContentTypeIdGte,
		),
		"contentTypeIdLt": validation.Validate(
			m.ContentTypeIdLt,
		),
		"contentTypeIdLte": validation.Validate(
			m.ContentTypeIdLte,
		),
		"contentTypeIdN": validation.Validate(
			m.ContentTypeIdN,
		),
		"httpContentType": validation.Validate(
			m.HttpContentType,
		),
		"httpContentTypeEmpty": validation.Validate(
			m.HttpContentTypeEmpty,
		),
		"httpContentTypeIc": validation.Validate(
			m.HttpContentTypeIc,
		),
		"httpContentTypeIe": validation.Validate(
			m.HttpContentTypeIe,
		),
		"httpContentTypeIew": validation.Validate(
			m.HttpContentTypeIew,
		),
		"httpContentTypeIsw": validation.Validate(
			m.HttpContentTypeIsw,
		),
		"httpContentTypeN": validation.Validate(
			m.HttpContentTypeN,
		),
		"httpContentTypeNic": validation.Validate(
			m.HttpContentTypeNic,
		),
		"httpContentTypeNie": validation.Validate(
			m.HttpContentTypeNie,
		),
		"httpContentTypeNiew": validation.Validate(
			m.HttpContentTypeNiew,
		),
		"httpContentTypeNisw": validation.Validate(
			m.HttpContentTypeNisw,
		),
		"httpMethod": validation.Validate(
			m.HttpMethod,
		),
		"httpMethodN": validation.Validate(
			m.HttpMethodN,
		),
		"id": validation.Validate(
			m.Id,
		),
		"idGt": validation.Validate(
			m.IdGt,
		),
		"idGte": validation.Validate(
			m.IdGte,
		),
		"idLt": validation.Validate(
			m.IdLt,
		),
		"idLte": validation.Validate(
			m.IdLte,
		),
		"idN": validation.Validate(
			m.IdN,
		),
		"name": validation.Validate(
			m.Name,
		),
		"nameEmpty": validation.Validate(
			m.NameEmpty,
		),
		"nameIc": validation.Validate(
			m.NameIc,
		),
		"nameIe": validation.Validate(
			m.NameIe,
		),
		"nameIew": validation.Validate(
			m.NameIew,
		),
		"nameIsw": validation.Validate(
			m.NameIsw,
		),
		"nameN": validation.Validate(
			m.NameN,
		),
		"nameNic": validation.Validate(
			m.NameNic,
		),
		"nameNie": validation.Validate(
			m.NameNie,
		),
		"nameNiew": validation.Validate(
			m.NameNiew,
		),
		"nameNisw": validation.Validate(
			m.NameNisw,
		),
		"payloadUrl": validation.Validate(
			m.PayloadUrl,
		),
		"payloadUrlEmpty": validation.Validate(
			m.PayloadUrlEmpty,
		),
		"payloadUrlIc": validation.Validate(
			m.PayloadUrlIc,
		),
		"payloadUrlIe": validation.Validate(
			m.PayloadUrlIe,
		),
		"payloadUrlIew": validation.Validate(
			m.PayloadUrlIew,
		),
		"payloadUrlIsw": validation.Validate(
			m.PayloadUrlIsw,
		),
		"payloadUrlN": validation.Validate(
			m.PayloadUrlN,
		),
		"payloadUrlNic": validation.Validate(
			m.PayloadUrlNic,
		),
		"payloadUrlNie": validation.Validate(
			m.PayloadUrlNie,
		),
		"payloadUrlNiew": validation.Validate(
			m.PayloadUrlNiew,
		),
		"payloadUrlNisw": validation.Validate(
			m.PayloadUrlNisw,
		),
		"secret": validation.Validate(
			m.Secret,
		),
		"secretEmpty": validation.Validate(
			m.SecretEmpty,
		),
		"secretIc": validation.Validate(
			m.SecretIc,
		),
		"secretIe": validation.Validate(
			m.SecretIe,
		),
		"secretIew": validation.Validate(
			m.SecretIew,
		),
		"secretIsw": validation.Validate(
			m.SecretIsw,
		),
		"secretN": validation.Validate(
			m.SecretN,
		),
		"secretNic": validation.Validate(
			m.SecretNic,
		),
		"secretNie": validation.Validate(
			m.SecretNie,
		),
		"secretNiew": validation.Validate(
			m.SecretNiew,
		),
		"secretNisw": validation.Validate(
			m.SecretNisw,
		),
	}.Filter()
}

// GetCaFilePath returns the CaFilePath property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePath() []string {
	return m.CaFilePath
}

// SetCaFilePath sets the CaFilePath property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePath(val []string) {
	m.CaFilePath = val
}

// GetCaFilePathEmpty returns the CaFilePathEmpty property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathEmpty() []string {
	return m.CaFilePathEmpty
}

// SetCaFilePathEmpty sets the CaFilePathEmpty property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathEmpty(val []string) {
	m.CaFilePathEmpty = val
}

// GetCaFilePathIc returns the CaFilePathIc property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathIc() []string {
	return m.CaFilePathIc
}

// SetCaFilePathIc sets the CaFilePathIc property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathIc(val []string) {
	m.CaFilePathIc = val
}

// GetCaFilePathIe returns the CaFilePathIe property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathIe() []string {
	return m.CaFilePathIe
}

// SetCaFilePathIe sets the CaFilePathIe property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathIe(val []string) {
	m.CaFilePathIe = val
}

// GetCaFilePathIew returns the CaFilePathIew property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathIew() []string {
	return m.CaFilePathIew
}

// SetCaFilePathIew sets the CaFilePathIew property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathIew(val []string) {
	m.CaFilePathIew = val
}

// GetCaFilePathIsw returns the CaFilePathIsw property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathIsw() []string {
	return m.CaFilePathIsw
}

// SetCaFilePathIsw sets the CaFilePathIsw property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathIsw(val []string) {
	m.CaFilePathIsw = val
}

// GetCaFilePathN returns the CaFilePathN property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathN() []string {
	return m.CaFilePathN
}

// SetCaFilePathN sets the CaFilePathN property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathN(val []string) {
	m.CaFilePathN = val
}

// GetCaFilePathNic returns the CaFilePathNic property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathNic() []string {
	return m.CaFilePathNic
}

// SetCaFilePathNic sets the CaFilePathNic property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathNic(val []string) {
	m.CaFilePathNic = val
}

// GetCaFilePathNie returns the CaFilePathNie property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathNie() []string {
	return m.CaFilePathNie
}

// SetCaFilePathNie sets the CaFilePathNie property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathNie(val []string) {
	m.CaFilePathNie = val
}

// GetCaFilePathNiew returns the CaFilePathNiew property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathNiew() []string {
	return m.CaFilePathNiew
}

// SetCaFilePathNiew sets the CaFilePathNiew property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathNiew(val []string) {
	m.CaFilePathNiew = val
}

// GetCaFilePathNisw returns the CaFilePathNisw property
func (m ExtrasWebhooksListQueryParameters) GetCaFilePathNisw() []string {
	return m.CaFilePathNisw
}

// SetCaFilePathNisw sets the CaFilePathNisw property
func (m *ExtrasWebhooksListQueryParameters) SetCaFilePathNisw(val []string) {
	m.CaFilePathNisw = val
}

// GetContentTypeId returns the ContentTypeId property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeId() []int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeId(val []int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdGt returns the ContentTypeIdGt property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeIdGt() []int32 {
	return m.ContentTypeIdGt
}

// SetContentTypeIdGt sets the ContentTypeIdGt property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeIdGt(val []int32) {
	m.ContentTypeIdGt = val
}

// GetContentTypeIdGte returns the ContentTypeIdGte property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeIdGte() []int32 {
	return m.ContentTypeIdGte
}

// SetContentTypeIdGte sets the ContentTypeIdGte property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeIdGte(val []int32) {
	m.ContentTypeIdGte = val
}

// GetContentTypeIdLt returns the ContentTypeIdLt property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeIdLt() []int32 {
	return m.ContentTypeIdLt
}

// SetContentTypeIdLt sets the ContentTypeIdLt property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeIdLt(val []int32) {
	m.ContentTypeIdLt = val
}

// GetContentTypeIdLte returns the ContentTypeIdLte property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeIdLte() []int32 {
	return m.ContentTypeIdLte
}

// SetContentTypeIdLte sets the ContentTypeIdLte property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeIdLte(val []int32) {
	m.ContentTypeIdLte = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasWebhooksListQueryParameters) GetContentTypeIdN() []int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypeIdN(val []int32) {
	m.ContentTypeIdN = val
}

// GetContentTypes returns the ContentTypes property
func (m ExtrasWebhooksListQueryParameters) GetContentTypes() string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypes(val string) {
	m.ContentTypes = val
}

// GetContentTypesIc returns the ContentTypesIc property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesIc() string {
	return m.ContentTypesIc
}

// SetContentTypesIc sets the ContentTypesIc property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesIc(val string) {
	m.ContentTypesIc = val
}

// GetContentTypesIe returns the ContentTypesIe property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesIe() string {
	return m.ContentTypesIe
}

// SetContentTypesIe sets the ContentTypesIe property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesIe(val string) {
	m.ContentTypesIe = val
}

// GetContentTypesIew returns the ContentTypesIew property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesIew() string {
	return m.ContentTypesIew
}

// SetContentTypesIew sets the ContentTypesIew property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesIew(val string) {
	m.ContentTypesIew = val
}

// GetContentTypesIsw returns the ContentTypesIsw property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesIsw() string {
	return m.ContentTypesIsw
}

// SetContentTypesIsw sets the ContentTypesIsw property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesIsw(val string) {
	m.ContentTypesIsw = val
}

// GetContentTypesN returns the ContentTypesN property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesN() string {
	return m.ContentTypesN
}

// SetContentTypesN sets the ContentTypesN property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesN(val string) {
	m.ContentTypesN = val
}

// GetContentTypesNic returns the ContentTypesNic property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesNic() string {
	return m.ContentTypesNic
}

// SetContentTypesNic sets the ContentTypesNic property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesNic(val string) {
	m.ContentTypesNic = val
}

// GetContentTypesNie returns the ContentTypesNie property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesNie() string {
	return m.ContentTypesNie
}

// SetContentTypesNie sets the ContentTypesNie property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesNie(val string) {
	m.ContentTypesNie = val
}

// GetContentTypesNiew returns the ContentTypesNiew property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesNiew() string {
	return m.ContentTypesNiew
}

// SetContentTypesNiew sets the ContentTypesNiew property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesNiew(val string) {
	m.ContentTypesNiew = val
}

// GetContentTypesNisw returns the ContentTypesNisw property
func (m ExtrasWebhooksListQueryParameters) GetContentTypesNisw() string {
	return m.ContentTypesNisw
}

// SetContentTypesNisw sets the ContentTypesNisw property
func (m *ExtrasWebhooksListQueryParameters) SetContentTypesNisw(val string) {
	m.ContentTypesNisw = val
}

// GetEnabled returns the Enabled property
func (m ExtrasWebhooksListQueryParameters) GetEnabled() bool {
	return m.Enabled
}

// SetEnabled sets the Enabled property
func (m *ExtrasWebhooksListQueryParameters) SetEnabled(val bool) {
	m.Enabled = val
}

// GetHttpContentType returns the HttpContentType property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentType() []string {
	return m.HttpContentType
}

// SetHttpContentType sets the HttpContentType property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentType(val []string) {
	m.HttpContentType = val
}

// GetHttpContentTypeEmpty returns the HttpContentTypeEmpty property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeEmpty() []string {
	return m.HttpContentTypeEmpty
}

// SetHttpContentTypeEmpty sets the HttpContentTypeEmpty property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeEmpty(val []string) {
	m.HttpContentTypeEmpty = val
}

// GetHttpContentTypeIc returns the HttpContentTypeIc property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeIc() []string {
	return m.HttpContentTypeIc
}

// SetHttpContentTypeIc sets the HttpContentTypeIc property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeIc(val []string) {
	m.HttpContentTypeIc = val
}

// GetHttpContentTypeIe returns the HttpContentTypeIe property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeIe() []string {
	return m.HttpContentTypeIe
}

// SetHttpContentTypeIe sets the HttpContentTypeIe property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeIe(val []string) {
	m.HttpContentTypeIe = val
}

// GetHttpContentTypeIew returns the HttpContentTypeIew property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeIew() []string {
	return m.HttpContentTypeIew
}

// SetHttpContentTypeIew sets the HttpContentTypeIew property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeIew(val []string) {
	m.HttpContentTypeIew = val
}

// GetHttpContentTypeIsw returns the HttpContentTypeIsw property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeIsw() []string {
	return m.HttpContentTypeIsw
}

// SetHttpContentTypeIsw sets the HttpContentTypeIsw property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeIsw(val []string) {
	m.HttpContentTypeIsw = val
}

// GetHttpContentTypeN returns the HttpContentTypeN property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeN() []string {
	return m.HttpContentTypeN
}

// SetHttpContentTypeN sets the HttpContentTypeN property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeN(val []string) {
	m.HttpContentTypeN = val
}

// GetHttpContentTypeNic returns the HttpContentTypeNic property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeNic() []string {
	return m.HttpContentTypeNic
}

// SetHttpContentTypeNic sets the HttpContentTypeNic property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeNic(val []string) {
	m.HttpContentTypeNic = val
}

// GetHttpContentTypeNie returns the HttpContentTypeNie property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeNie() []string {
	return m.HttpContentTypeNie
}

// SetHttpContentTypeNie sets the HttpContentTypeNie property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeNie(val []string) {
	m.HttpContentTypeNie = val
}

// GetHttpContentTypeNiew returns the HttpContentTypeNiew property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeNiew() []string {
	return m.HttpContentTypeNiew
}

// SetHttpContentTypeNiew sets the HttpContentTypeNiew property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeNiew(val []string) {
	m.HttpContentTypeNiew = val
}

// GetHttpContentTypeNisw returns the HttpContentTypeNisw property
func (m ExtrasWebhooksListQueryParameters) GetHttpContentTypeNisw() []string {
	return m.HttpContentTypeNisw
}

// SetHttpContentTypeNisw sets the HttpContentTypeNisw property
func (m *ExtrasWebhooksListQueryParameters) SetHttpContentTypeNisw(val []string) {
	m.HttpContentTypeNisw = val
}

// GetHttpMethod returns the HttpMethod property
func (m ExtrasWebhooksListQueryParameters) GetHttpMethod() []string {
	return m.HttpMethod
}

// SetHttpMethod sets the HttpMethod property
func (m *ExtrasWebhooksListQueryParameters) SetHttpMethod(val []string) {
	m.HttpMethod = val
}

// GetHttpMethodN returns the HttpMethodN property
func (m ExtrasWebhooksListQueryParameters) GetHttpMethodN() []string {
	return m.HttpMethodN
}

// SetHttpMethodN sets the HttpMethodN property
func (m *ExtrasWebhooksListQueryParameters) SetHttpMethodN(val []string) {
	m.HttpMethodN = val
}

// GetId returns the Id property
func (m ExtrasWebhooksListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasWebhooksListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasWebhooksListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasWebhooksListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasWebhooksListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasWebhooksListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasWebhooksListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasWebhooksListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasWebhooksListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasWebhooksListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasWebhooksListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasWebhooksListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasWebhooksListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasWebhooksListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasWebhooksListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasWebhooksListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasWebhooksListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasWebhooksListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasWebhooksListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasWebhooksListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasWebhooksListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasWebhooksListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasWebhooksListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasWebhooksListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasWebhooksListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasWebhooksListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasWebhooksListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasWebhooksListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasWebhooksListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasWebhooksListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasWebhooksListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasWebhooksListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasWebhooksListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasWebhooksListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasWebhooksListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasWebhooksListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasWebhooksListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasWebhooksListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasWebhooksListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasWebhooksListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPayloadUrl returns the PayloadUrl property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrl() []string {
	return m.PayloadUrl
}

// SetPayloadUrl sets the PayloadUrl property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrl(val []string) {
	m.PayloadUrl = val
}

// GetPayloadUrlEmpty returns the PayloadUrlEmpty property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlEmpty() []string {
	return m.PayloadUrlEmpty
}

// SetPayloadUrlEmpty sets the PayloadUrlEmpty property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlEmpty(val []string) {
	m.PayloadUrlEmpty = val
}

// GetPayloadUrlIc returns the PayloadUrlIc property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlIc() []string {
	return m.PayloadUrlIc
}

// SetPayloadUrlIc sets the PayloadUrlIc property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlIc(val []string) {
	m.PayloadUrlIc = val
}

// GetPayloadUrlIe returns the PayloadUrlIe property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlIe() []string {
	return m.PayloadUrlIe
}

// SetPayloadUrlIe sets the PayloadUrlIe property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlIe(val []string) {
	m.PayloadUrlIe = val
}

// GetPayloadUrlIew returns the PayloadUrlIew property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlIew() []string {
	return m.PayloadUrlIew
}

// SetPayloadUrlIew sets the PayloadUrlIew property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlIew(val []string) {
	m.PayloadUrlIew = val
}

// GetPayloadUrlIsw returns the PayloadUrlIsw property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlIsw() []string {
	return m.PayloadUrlIsw
}

// SetPayloadUrlIsw sets the PayloadUrlIsw property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlIsw(val []string) {
	m.PayloadUrlIsw = val
}

// GetPayloadUrlN returns the PayloadUrlN property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlN() []string {
	return m.PayloadUrlN
}

// SetPayloadUrlN sets the PayloadUrlN property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlN(val []string) {
	m.PayloadUrlN = val
}

// GetPayloadUrlNic returns the PayloadUrlNic property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlNic() []string {
	return m.PayloadUrlNic
}

// SetPayloadUrlNic sets the PayloadUrlNic property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlNic(val []string) {
	m.PayloadUrlNic = val
}

// GetPayloadUrlNie returns the PayloadUrlNie property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlNie() []string {
	return m.PayloadUrlNie
}

// SetPayloadUrlNie sets the PayloadUrlNie property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlNie(val []string) {
	m.PayloadUrlNie = val
}

// GetPayloadUrlNiew returns the PayloadUrlNiew property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlNiew() []string {
	return m.PayloadUrlNiew
}

// SetPayloadUrlNiew sets the PayloadUrlNiew property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlNiew(val []string) {
	m.PayloadUrlNiew = val
}

// GetPayloadUrlNisw returns the PayloadUrlNisw property
func (m ExtrasWebhooksListQueryParameters) GetPayloadUrlNisw() []string {
	return m.PayloadUrlNisw
}

// SetPayloadUrlNisw sets the PayloadUrlNisw property
func (m *ExtrasWebhooksListQueryParameters) SetPayloadUrlNisw(val []string) {
	m.PayloadUrlNisw = val
}

// GetQ returns the Q property
func (m ExtrasWebhooksListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasWebhooksListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetSecret returns the Secret property
func (m ExtrasWebhooksListQueryParameters) GetSecret() []string {
	return m.Secret
}

// SetSecret sets the Secret property
func (m *ExtrasWebhooksListQueryParameters) SetSecret(val []string) {
	m.Secret = val
}

// GetSecretEmpty returns the SecretEmpty property
func (m ExtrasWebhooksListQueryParameters) GetSecretEmpty() []string {
	return m.SecretEmpty
}

// SetSecretEmpty sets the SecretEmpty property
func (m *ExtrasWebhooksListQueryParameters) SetSecretEmpty(val []string) {
	m.SecretEmpty = val
}

// GetSecretIc returns the SecretIc property
func (m ExtrasWebhooksListQueryParameters) GetSecretIc() []string {
	return m.SecretIc
}

// SetSecretIc sets the SecretIc property
func (m *ExtrasWebhooksListQueryParameters) SetSecretIc(val []string) {
	m.SecretIc = val
}

// GetSecretIe returns the SecretIe property
func (m ExtrasWebhooksListQueryParameters) GetSecretIe() []string {
	return m.SecretIe
}

// SetSecretIe sets the SecretIe property
func (m *ExtrasWebhooksListQueryParameters) SetSecretIe(val []string) {
	m.SecretIe = val
}

// GetSecretIew returns the SecretIew property
func (m ExtrasWebhooksListQueryParameters) GetSecretIew() []string {
	return m.SecretIew
}

// SetSecretIew sets the SecretIew property
func (m *ExtrasWebhooksListQueryParameters) SetSecretIew(val []string) {
	m.SecretIew = val
}

// GetSecretIsw returns the SecretIsw property
func (m ExtrasWebhooksListQueryParameters) GetSecretIsw() []string {
	return m.SecretIsw
}

// SetSecretIsw sets the SecretIsw property
func (m *ExtrasWebhooksListQueryParameters) SetSecretIsw(val []string) {
	m.SecretIsw = val
}

// GetSecretN returns the SecretN property
func (m ExtrasWebhooksListQueryParameters) GetSecretN() []string {
	return m.SecretN
}

// SetSecretN sets the SecretN property
func (m *ExtrasWebhooksListQueryParameters) SetSecretN(val []string) {
	m.SecretN = val
}

// GetSecretNic returns the SecretNic property
func (m ExtrasWebhooksListQueryParameters) GetSecretNic() []string {
	return m.SecretNic
}

// SetSecretNic sets the SecretNic property
func (m *ExtrasWebhooksListQueryParameters) SetSecretNic(val []string) {
	m.SecretNic = val
}

// GetSecretNie returns the SecretNie property
func (m ExtrasWebhooksListQueryParameters) GetSecretNie() []string {
	return m.SecretNie
}

// SetSecretNie sets the SecretNie property
func (m *ExtrasWebhooksListQueryParameters) SetSecretNie(val []string) {
	m.SecretNie = val
}

// GetSecretNiew returns the SecretNiew property
func (m ExtrasWebhooksListQueryParameters) GetSecretNiew() []string {
	return m.SecretNiew
}

// SetSecretNiew sets the SecretNiew property
func (m *ExtrasWebhooksListQueryParameters) SetSecretNiew(val []string) {
	m.SecretNiew = val
}

// GetSecretNisw returns the SecretNisw property
func (m ExtrasWebhooksListQueryParameters) GetSecretNisw() []string {
	return m.SecretNisw
}

// SetSecretNisw sets the SecretNisw property
func (m *ExtrasWebhooksListQueryParameters) SetSecretNisw(val []string) {
	m.SecretNisw = val
}

// GetSslVerification returns the SslVerification property
func (m ExtrasWebhooksListQueryParameters) GetSslVerification() bool {
	return m.SslVerification
}

// SetSslVerification sets the SslVerification property
func (m *ExtrasWebhooksListQueryParameters) SetSslVerification(val bool) {
	m.SslVerification = val
}

// GetTypeCreate returns the TypeCreate property
func (m ExtrasWebhooksListQueryParameters) GetTypeCreate() bool {
	return m.TypeCreate
}

// SetTypeCreate sets the TypeCreate property
func (m *ExtrasWebhooksListQueryParameters) SetTypeCreate(val bool) {
	m.TypeCreate = val
}

// GetTypeDelete returns the TypeDelete property
func (m ExtrasWebhooksListQueryParameters) GetTypeDelete() bool {
	return m.TypeDelete
}

// SetTypeDelete sets the TypeDelete property
func (m *ExtrasWebhooksListQueryParameters) SetTypeDelete(val bool) {
	m.TypeDelete = val
}

// GetTypeJobEnd returns the TypeJobEnd property
func (m ExtrasWebhooksListQueryParameters) GetTypeJobEnd() bool {
	return m.TypeJobEnd
}

// SetTypeJobEnd sets the TypeJobEnd property
func (m *ExtrasWebhooksListQueryParameters) SetTypeJobEnd(val bool) {
	m.TypeJobEnd = val
}

// GetTypeJobStart returns the TypeJobStart property
func (m ExtrasWebhooksListQueryParameters) GetTypeJobStart() bool {
	return m.TypeJobStart
}

// SetTypeJobStart sets the TypeJobStart property
func (m *ExtrasWebhooksListQueryParameters) SetTypeJobStart(val bool) {
	m.TypeJobStart = val
}

// GetTypeUpdate returns the TypeUpdate property
func (m ExtrasWebhooksListQueryParameters) GetTypeUpdate() bool {
	return m.TypeUpdate
}

// SetTypeUpdate sets the TypeUpdate property
func (m *ExtrasWebhooksListQueryParameters) SetTypeUpdate(val bool) {
	m.TypeUpdate = val
}
