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

// ExtrasConfigTemplatesListQueryParameters is an object.
type ExtrasConfigTemplatesListQueryParameters struct {
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
	// Tag:
	Tag []string `json:"tag,omitempty" mapstructure:"tag,omitempty"`
	// TagN:
	TagN []string `json:"tag__n,omitempty" mapstructure:"tag__n,omitempty"`
}

// Validate implements basic validation for this model
func (m ExtrasConfigTemplatesListQueryParameters) Validate() error {
	return validation.Errors{
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
		"tag": validation.Validate(
			m.Tag,
		),
		"tagN": validation.Validate(
			m.TagN,
		),
	}.Filter()
}

// GetDataFileId returns the DataFileId property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataFileId() []*int32 {
	return m.DataFileId
}

// SetDataFileId sets the DataFileId property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataFileId(val []*int32) {
	m.DataFileId = val
}

// GetDataFileIdN returns the DataFileIdN property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataFileIdN() []*int32 {
	return m.DataFileIdN
}

// SetDataFileIdN sets the DataFileIdN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataFileIdN(val []*int32) {
	m.DataFileIdN = val
}

// GetDataSourceId returns the DataSourceId property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSourceId() []*int32 {
	return m.DataSourceId
}

// SetDataSourceId sets the DataSourceId property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSourceId(val []*int32) {
	m.DataSourceId = val
}

// GetDataSourceIdN returns the DataSourceIdN property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSourceIdN() []*int32 {
	return m.DataSourceIdN
}

// SetDataSourceIdN sets the DataSourceIdN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSourceIdN(val []*int32) {
	m.DataSourceIdN = val
}

// GetDataSynced returns the DataSynced property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSynced() []time.Time {
	return m.DataSynced
}

// SetDataSynced sets the DataSynced property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSynced(val []time.Time) {
	m.DataSynced = val
}

// GetDataSyncedGt returns the DataSyncedGt property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSyncedGt() []time.Time {
	return m.DataSyncedGt
}

// SetDataSyncedGt sets the DataSyncedGt property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSyncedGt(val []time.Time) {
	m.DataSyncedGt = val
}

// GetDataSyncedGte returns the DataSyncedGte property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSyncedGte() []time.Time {
	return m.DataSyncedGte
}

// SetDataSyncedGte sets the DataSyncedGte property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSyncedGte(val []time.Time) {
	m.DataSyncedGte = val
}

// GetDataSyncedLt returns the DataSyncedLt property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSyncedLt() []time.Time {
	return m.DataSyncedLt
}

// SetDataSyncedLt sets the DataSyncedLt property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSyncedLt(val []time.Time) {
	m.DataSyncedLt = val
}

// GetDataSyncedLte returns the DataSyncedLte property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSyncedLte() []time.Time {
	return m.DataSyncedLte
}

// SetDataSyncedLte sets the DataSyncedLte property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSyncedLte(val []time.Time) {
	m.DataSyncedLte = val
}

// GetDataSyncedN returns the DataSyncedN property
func (m ExtrasConfigTemplatesListQueryParameters) GetDataSyncedN() []time.Time {
	return m.DataSyncedN
}

// SetDataSyncedN sets the DataSyncedN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDataSyncedN(val []time.Time) {
	m.DataSyncedN = val
}

// GetDescription returns the Description property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescription() []string {
	return m.Description
}

// SetDescription sets the Description property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescription(val []string) {
	m.Description = val
}

// GetDescriptionEmpty returns the DescriptionEmpty property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionEmpty() []string {
	return m.DescriptionEmpty
}

// SetDescriptionEmpty sets the DescriptionEmpty property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionEmpty(val []string) {
	m.DescriptionEmpty = val
}

// GetDescriptionIc returns the DescriptionIc property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionIc() []string {
	return m.DescriptionIc
}

// SetDescriptionIc sets the DescriptionIc property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionIc(val []string) {
	m.DescriptionIc = val
}

// GetDescriptionIe returns the DescriptionIe property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionIe() []string {
	return m.DescriptionIe
}

// SetDescriptionIe sets the DescriptionIe property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionIe(val []string) {
	m.DescriptionIe = val
}

// GetDescriptionIew returns the DescriptionIew property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionIew() []string {
	return m.DescriptionIew
}

// SetDescriptionIew sets the DescriptionIew property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionIew(val []string) {
	m.DescriptionIew = val
}

// GetDescriptionIsw returns the DescriptionIsw property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionIsw() []string {
	return m.DescriptionIsw
}

// SetDescriptionIsw sets the DescriptionIsw property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionIsw(val []string) {
	m.DescriptionIsw = val
}

// GetDescriptionN returns the DescriptionN property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionN() []string {
	return m.DescriptionN
}

// SetDescriptionN sets the DescriptionN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionN(val []string) {
	m.DescriptionN = val
}

// GetDescriptionNic returns the DescriptionNic property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionNic() []string {
	return m.DescriptionNic
}

// SetDescriptionNic sets the DescriptionNic property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionNic(val []string) {
	m.DescriptionNic = val
}

// GetDescriptionNie returns the DescriptionNie property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionNie() []string {
	return m.DescriptionNie
}

// SetDescriptionNie sets the DescriptionNie property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionNie(val []string) {
	m.DescriptionNie = val
}

// GetDescriptionNiew returns the DescriptionNiew property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionNiew() []string {
	return m.DescriptionNiew
}

// SetDescriptionNiew sets the DescriptionNiew property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionNiew(val []string) {
	m.DescriptionNiew = val
}

// GetDescriptionNisw returns the DescriptionNisw property
func (m ExtrasConfigTemplatesListQueryParameters) GetDescriptionNisw() []string {
	return m.DescriptionNisw
}

// SetDescriptionNisw sets the DescriptionNisw property
func (m *ExtrasConfigTemplatesListQueryParameters) SetDescriptionNisw(val []string) {
	m.DescriptionNisw = val
}

// GetId returns the Id property
func (m ExtrasConfigTemplatesListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *ExtrasConfigTemplatesListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m ExtrasConfigTemplatesListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *ExtrasConfigTemplatesListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m ExtrasConfigTemplatesListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *ExtrasConfigTemplatesListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m ExtrasConfigTemplatesListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *ExtrasConfigTemplatesListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m ExtrasConfigTemplatesListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *ExtrasConfigTemplatesListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m ExtrasConfigTemplatesListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetLimit returns the Limit property
func (m ExtrasConfigTemplatesListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *ExtrasConfigTemplatesListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetName returns the Name property
func (m ExtrasConfigTemplatesListQueryParameters) GetName() []string {
	return m.Name
}

// SetName sets the Name property
func (m *ExtrasConfigTemplatesListQueryParameters) SetName(val []string) {
	m.Name = val
}

// GetNameEmpty returns the NameEmpty property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameEmpty() []string {
	return m.NameEmpty
}

// SetNameEmpty sets the NameEmpty property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameEmpty(val []string) {
	m.NameEmpty = val
}

// GetNameIc returns the NameIc property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameIc() []string {
	return m.NameIc
}

// SetNameIc sets the NameIc property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameIc(val []string) {
	m.NameIc = val
}

// GetNameIe returns the NameIe property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameIe() []string {
	return m.NameIe
}

// SetNameIe sets the NameIe property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameIe(val []string) {
	m.NameIe = val
}

// GetNameIew returns the NameIew property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameIew() []string {
	return m.NameIew
}

// SetNameIew sets the NameIew property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameIew(val []string) {
	m.NameIew = val
}

// GetNameIsw returns the NameIsw property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameIsw() []string {
	return m.NameIsw
}

// SetNameIsw sets the NameIsw property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameIsw(val []string) {
	m.NameIsw = val
}

// GetNameN returns the NameN property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameN() []string {
	return m.NameN
}

// SetNameN sets the NameN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameN(val []string) {
	m.NameN = val
}

// GetNameNic returns the NameNic property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameNic() []string {
	return m.NameNic
}

// SetNameNic sets the NameNic property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameNic(val []string) {
	m.NameNic = val
}

// GetNameNie returns the NameNie property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameNie() []string {
	return m.NameNie
}

// SetNameNie sets the NameNie property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameNie(val []string) {
	m.NameNie = val
}

// GetNameNiew returns the NameNiew property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameNiew() []string {
	return m.NameNiew
}

// SetNameNiew sets the NameNiew property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameNiew(val []string) {
	m.NameNiew = val
}

// GetNameNisw returns the NameNisw property
func (m ExtrasConfigTemplatesListQueryParameters) GetNameNisw() []string {
	return m.NameNisw
}

// SetNameNisw sets the NameNisw property
func (m *ExtrasConfigTemplatesListQueryParameters) SetNameNisw(val []string) {
	m.NameNisw = val
}

// GetOffset returns the Offset property
func (m ExtrasConfigTemplatesListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *ExtrasConfigTemplatesListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m ExtrasConfigTemplatesListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *ExtrasConfigTemplatesListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetQ returns the Q property
func (m ExtrasConfigTemplatesListQueryParameters) GetQ() string {
	return m.Q
}

// SetQ sets the Q property
func (m *ExtrasConfigTemplatesListQueryParameters) SetQ(val string) {
	m.Q = val
}

// GetTag returns the Tag property
func (m ExtrasConfigTemplatesListQueryParameters) GetTag() []string {
	return m.Tag
}

// SetTag sets the Tag property
func (m *ExtrasConfigTemplatesListQueryParameters) SetTag(val []string) {
	m.Tag = val
}

// GetTagN returns the TagN property
func (m ExtrasConfigTemplatesListQueryParameters) GetTagN() []string {
	return m.TagN
}

// SetTagN sets the TagN property
func (m *ExtrasConfigTemplatesListQueryParameters) SetTagN(val []string) {
	m.TagN = val
}
