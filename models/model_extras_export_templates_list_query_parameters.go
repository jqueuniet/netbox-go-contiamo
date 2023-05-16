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

// ExtrasExportTemplatesListQueryParameters is an object.
type ExtrasExportTemplatesListQueryParameters struct {
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
	// DataFileId: Data file (ID)
	DataFileId []*int32 `json:"data_file_id,omitempty" mapstructure:"data_file_id,omitempty"`
	// DataFileIdN: Data file (ID)
	DataFileIdN []*int32 `json:"data_file_id__n,omitempty" mapstructure:"data_file_id__n,omitempty"`
	// DataSourceId: Data source (ID)
	DataSourceId []*int32 `json:"data_source_id,omitempty" mapstructure:"data_source_id,omitempty"`
	// DataSourceIdN: Data source (ID)
	DataSourceIdN []*int32 `json:"data_source_id__n,omitempty" mapstructure:"data_source_id__n,omitempty"`
	// DataSynced:
	DataSynced []time.Time `json:"data_synced,omitempty" mapstructure:"data_synced,omitempty"`
	// DataSyncedGt:
	DataSyncedGt []time.Time `json:"data_synced__gt,omitempty" mapstructure:"data_synced__gt,omitempty"`
	// DataSyncedGte:
	DataSyncedGte []time.Time `json:"data_synced__gte,omitempty" mapstructure:"data_synced__gte,omitempty"`
	// DataSyncedLt:
	DataSyncedLt []time.Time `json:"data_synced__lt,omitempty" mapstructure:"data_synced__lt,omitempty"`
	// DataSyncedLte:
	DataSyncedLte []time.Time `json:"data_synced__lte,omitempty" mapstructure:"data_synced__lte,omitempty"`
	// DataSyncedN:
	DataSyncedN []time.Time `json:"data_synced__n,omitempty" mapstructure:"data_synced__n,omitempty"`
	// Description:
	Description []string `json:"description,omitempty" mapstructure:"description,omitempty"`
	// DescriptionEmpty:
	DescriptionEmpty []string `json:"description__empty,omitempty" mapstructure:"description__empty,omitempty"`
	// DescriptionIc:
	DescriptionIc []string `json:"description__ic,omitempty" mapstructure:"description__ic,omitempty"`
	// DescriptionIe:
	DescriptionIe []string `json:"description__ie,omitempty" mapstructure:"description__ie,omitempty"`
	// DescriptionIew:
	DescriptionIew []string `json:"description__iew,omitempty" mapstructure:"description__iew,omitempty"`
	// DescriptionIsw:
	DescriptionIsw []string `json:"description__isw,omitempty" mapstructure:"description__isw,omitempty"`
	// DescriptionN:
	DescriptionN []string `json:"description__n,omitempty" mapstructure:"description__n,omitempty"`
	// DescriptionNic:
	DescriptionNic []string `json:"description__nic,omitempty" mapstructure:"description__nic,omitempty"`
	// DescriptionNie:
	DescriptionNie []string `json:"description__nie,omitempty" mapstructure:"description__nie,omitempty"`
	// DescriptionNiew:
	DescriptionNiew []string `json:"description__niew,omitempty" mapstructure:"description__niew,omitempty"`
	// DescriptionNisw:
	DescriptionNisw []string `json:"description__nisw,omitempty" mapstructure:"description__nisw,omitempty"`
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
	// Q: Search
	Q string `json:"q,omitempty" mapstructure:"q,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasExportTemplatesListQueryParameters) Validate() error {
	return validation.Errors{
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
		"dataFileId": validation.Validate(
			m.DataFileId,
		),
		"dataFileIdN": validation.Validate(
			m.DataFileIdN,
		),
		"dataSourceId": validation.Validate(
			m.DataSourceId,
		),
		"dataSourceIdN": validation.Validate(
			m.DataSourceIdN,
		),
		"dataSynced": validation.Validate(
			m.DataSynced,
		),
		"dataSyncedGt": validation.Validate(
			m.DataSyncedGt,
		),
		"dataSyncedGte": validation.Validate(
			m.DataSyncedGte,
		),
		"dataSyncedLt": validation.Validate(
			m.DataSyncedLt,
		),
		"dataSyncedLte": validation.Validate(
			m.DataSyncedLte,
		),
		"dataSyncedN": validation.Validate(
			m.DataSyncedN,
		),
		"description": validation.Validate(
			m.Description,
		),
		"descriptionEmpty": validation.Validate(
			m.DescriptionEmpty,
		),
		"descriptionIc": validation.Validate(
			m.DescriptionIc,
		),
		"descriptionIe": validation.Validate(
			m.DescriptionIe,
		),
		"descriptionIew": validation.Validate(
			m.DescriptionIew,
		),
		"descriptionIsw": validation.Validate(
			m.DescriptionIsw,
		),
		"descriptionN": validation.Validate(
			m.DescriptionN,
		),
		"descriptionNic": validation.Validate(
			m.DescriptionNic,
		),
		"descriptionNie": validation.Validate(
			m.DescriptionNie,
		),
		"descriptionNiew": validation.Validate(
			m.DescriptionNiew,
		),
		"descriptionNisw": validation.Validate(
			m.DescriptionNisw,
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
	}.Filter()
}

// GetContentTypeId returns the ContentTypeId property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeId() []int32 {
	return m.ContentTypeId
}

// SetContentTypeId sets the ContentTypeId property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeId(val []int32) {
	m.ContentTypeId = val
}

// GetContentTypeIdGt returns the ContentTypeIdGt property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeIdGt() []int32 {
	return m.ContentTypeIdGt
}

// SetContentTypeIdGt sets the ContentTypeIdGt property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeIdGt(val []int32) {
	m.ContentTypeIdGt = val
}

// GetContentTypeIdGte returns the ContentTypeIdGte property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeIdGte() []int32 {
	return m.ContentTypeIdGte
}

// SetContentTypeIdGte sets the ContentTypeIdGte property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeIdGte(val []int32) {
	m.ContentTypeIdGte = val
}

// GetContentTypeIdLt returns the ContentTypeIdLt property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeIdLt() []int32 {
	return m.ContentTypeIdLt
}

// SetContentTypeIdLt sets the ContentTypeIdLt property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeIdLt(val []int32) {
	m.ContentTypeIdLt = val
}

// GetContentTypeIdLte returns the ContentTypeIdLte property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeIdLte() []int32 {
	return m.ContentTypeIdLte
}

// SetContentTypeIdLte sets the ContentTypeIdLte property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeIdLte(val []int32) {
	m.ContentTypeIdLte = val
}

// GetContentTypeIdN returns the ContentTypeIdN property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypeIdN() []int32 {
	return m.ContentTypeIdN
}

// SetContentTypeIdN sets the ContentTypeIdN property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypeIdN(val []int32) {
	m.ContentTypeIdN = val
}

// GetContentTypes returns the ContentTypes property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypes() string {
	return m.ContentTypes
}

// SetContentTypes sets the ContentTypes property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypes(val string) {
	m.ContentTypes = val
}

// GetContentTypesIc returns the ContentTypesIc property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesIc() string {
	return m.ContentTypesIc
}

// SetContentTypesIc sets the ContentTypesIc property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesIc(val string) {
	m.ContentTypesIc = val
}

// GetContentTypesIe returns the ContentTypesIe property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesIe() string {
	return m.ContentTypesIe
}

// SetContentTypesIe sets the ContentTypesIe property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesIe(val string) {
	m.ContentTypesIe = val
}

// GetContentTypesIew returns the ContentTypesIew property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesIew() string {
	return m.ContentTypesIew
}

// SetContentTypesIew sets the ContentTypesIew property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesIew(val string) {
	m.ContentTypesIew = val
}

// GetContentTypesIsw returns the ContentTypesIsw property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesIsw() string {
	return m.ContentTypesIsw
}

// SetContentTypesIsw sets the ContentTypesIsw property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesIsw(val string) {
	m.ContentTypesIsw = val
}

// GetContentTypesN returns the ContentTypesN property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesN() string {
	return m.ContentTypesN
}

// SetContentTypesN sets the ContentTypesN property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesN(val string) {
	m.ContentTypesN = val
}

// GetContentTypesNic returns the ContentTypesNic property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesNic() string {
	return m.ContentTypesNic
}

// SetContentTypesNic sets the ContentTypesNic property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesNic(val string) {
	m.ContentTypesNic = val
}

// GetContentTypesNie returns the ContentTypesNie property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesNie() string {
	return m.ContentTypesNie
}

// SetContentTypesNie sets the ContentTypesNie property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesNie(val string) {
	m.ContentTypesNie = val
}

// GetContentTypesNiew returns the ContentTypesNiew property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesNiew() string {
	return m.ContentTypesNiew
}

// SetContentTypesNiew sets the ContentTypesNiew property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesNiew(val string) {
	m.ContentTypesNiew = val
}

// GetContentTypesNisw returns the ContentTypesNisw property
func (m ExtrasExportTemplatesListQueryParameters) GetContentTypesNisw() string {
	return m.ContentTypesNisw
}

// SetContentTypesNisw sets the ContentTypesNisw property
func (m *ExtrasExportTemplatesListQueryParameters) SetContentTypesNisw(val string) {
	m.ContentTypesNisw = val
}

// GetDataFileId returns the DataFileId property
func (m ExtrasExportTemplatesListQueryParameters) GetDataFileId() []*int32 {
	return m.DataFileId
}

// SetDataFileId sets the DataFileId property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataFileId(val []*int32) {
	m.DataFileId = val
}

// GetDataFileIdN returns the DataFileIdN property
func (m ExtrasExportTemplatesListQueryParameters) GetDataFileIdN() []*int32 {
	return m.DataFileIdN
}

// SetDataFileIdN sets the DataFileIdN property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataFileIdN(val []*int32) {
	m.DataFileIdN = val
}

// GetDataSourceId returns the DataSourceId property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSourceId() []*int32 {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSourceId(val []*int32) {
	m.DataSourceId = val
}

// GetDataSourceIdN returns the DataSourceIdN property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSourceIdN() []*int32 {
	return m.DataSourceIdN
}

// SetDataSourceIdN sets the DataSourceIdN property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSourceIdN(val []*int32) {
	m.DataSourceIdN = val
}

// GetDataSynced returns the DataSynced property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSynced() []time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSynced(val []time.Time) {
	m.DataSynced = val
}

// GetDataSyncedGt returns the DataSyncedGt property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSyncedGt() []time.Time {
	return m.DataSyncedGt
}

// SetDataSyncedGt sets the DataSyncedGt property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSyncedGt(val []time.Time) {
	m.DataSyncedGt = val
}

// GetDataSyncedGte returns the DataSyncedGte property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSyncedGte() []time.Time {
	return m.DataSyncedGte
}

// SetDataSyncedGte sets the DataSyncedGte property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSyncedGte(val []time.Time) {
	m.DataSyncedGte = val
}

// GetDataSyncedLt returns the DataSyncedLt property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSyncedLt() []time.Time {
	return m.DataSyncedLt
}

// SetDataSyncedLt sets the DataSyncedLt property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSyncedLt(val []time.Time) {
	m.DataSyncedLt = val
}

// GetDataSyncedLte returns the DataSyncedLte property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSyncedLte() []time.Time {
	return m.DataSyncedLte
}

// SetDataSyncedLte sets the DataSyncedLte property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSyncedLte(val []time.Time) {
	m.DataSyncedLte = val
}

// GetDataSyncedN returns the DataSyncedN property
func (m ExtrasExportTemplatesListQueryParameters) GetDataSyncedN() []time.Time {
	return m.DataSyncedN
}

// SetDataSyncedN sets the DataSyncedN property
func (m *ExtrasExportTemplatesListQueryParameters) SetDataSyncedN(val []time.Time) {
	m.DataSyncedN = val
}

// GetDescription returns the Description property
func (m ExtrasExportTemplatesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m ExtrasExportTemplatesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *ExtrasExportTemplatesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m ExtrasExportTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasExportTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasExportTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasExportTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasExportTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasExportTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasExportTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasExportTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasExportTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasExportTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasExportTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasExportTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasExportTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasExportTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasExportTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasExportTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasExportTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasExportTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasExportTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasExportTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasExportTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasExportTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasExportTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasExportTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasExportTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasExportTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasExportTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasExportTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasExportTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasExportTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasExportTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasExportTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasExportTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}
