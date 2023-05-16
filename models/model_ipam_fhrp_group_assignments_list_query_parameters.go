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
)

// IpamFhrpGroupAssignmentsListQueryParameters is an object.
type IpamFhrpGroupAssignmentsListQueryParameters struct {
	// Created:
	Created []time.Time `json:"created,omitempty" mapstructure:"created,omitempty"`
	// CreatedGt:
	CreatedGt []time.Time `json:"created__gt,omitempty" mapstructure:"created__gt,omitempty"`
	// CreatedGte:
	CreatedGte []time.Time `json:"created__gte,omitempty" mapstructure:"created__gte,omitempty"`
	// CreatedLt:
	CreatedLt []time.Time `json:"created__lt,omitempty" mapstructure:"created__lt,omitempty"`
	// CreatedLte:
	CreatedLte []time.Time `json:"created__lte,omitempty" mapstructure:"created__lte,omitempty"`
	// CreatedN:
	CreatedN []time.Time `json:"created__n,omitempty" mapstructure:"created__n,omitempty"`
	// CreatedByRequest:
	CreatedByRequest string `json:"created_by_request,omitempty" mapstructure:"created_by_request,omitempty"`
	// Device:
	Device []string `json:"device,omitempty" mapstructure:"device,omitempty"`
	// DeviceId:
	DeviceId []int32 `json:"device_id,omitempty" mapstructure:"device_id,omitempty"`
	// GroupId: Group (ID)
	GroupId []int32 `json:"group_id,omitempty" mapstructure:"group_id,omitempty"`
	// GroupIdN: Group (ID)
	GroupIdN []int32 `json:"group_id__n,omitempty" mapstructure:"group_id__n,omitempty"`
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
	// InterfaceId:
	InterfaceId []int32 `json:"interface_id,omitempty" mapstructure:"interface_id,omitempty"`
	// InterfaceIdGt:
	InterfaceIdGt []int32 `json:"interface_id__gt,omitempty" mapstructure:"interface_id__gt,omitempty"`
	// InterfaceIdGte:
	InterfaceIdGte []int32 `json:"interface_id__gte,omitempty" mapstructure:"interface_id__gte,omitempty"`
	// InterfaceIdLt:
	InterfaceIdLt []int32 `json:"interface_id__lt,omitempty" mapstructure:"interface_id__lt,omitempty"`
	// InterfaceIdLte:
	InterfaceIdLte []int32 `json:"interface_id__lte,omitempty" mapstructure:"interface_id__lte,omitempty"`
	// InterfaceIdN:
	InterfaceIdN []int32 `json:"interface_id__n,omitempty" mapstructure:"interface_id__n,omitempty"`
	// InterfaceType:
	InterfaceType string `json:"interface_type,omitempty" mapstructure:"interface_type,omitempty"`
	// InterfaceTypeN:
	InterfaceTypeN string `json:"interface_type__n,omitempty" mapstructure:"interface_type__n,omitempty"`
	// LastUpdated:
	LastUpdated []time.Time `json:"last_updated,omitempty" mapstructure:"last_updated,omitempty"`
	// LastUpdatedGt:
	LastUpdatedGt []time.Time `json:"last_updated__gt,omitempty" mapstructure:"last_updated__gt,omitempty"`
	// LastUpdatedGte:
	LastUpdatedGte []time.Time `json:"last_updated__gte,omitempty" mapstructure:"last_updated__gte,omitempty"`
	// LastUpdatedLt:
	LastUpdatedLt []time.Time `json:"last_updated__lt,omitempty" mapstructure:"last_updated__lt,omitempty"`
	// LastUpdatedLte:
	LastUpdatedLte []time.Time `json:"last_updated__lte,omitempty" mapstructure:"last_updated__lte,omitempty"`
	// LastUpdatedN:
	LastUpdatedN []time.Time `json:"last_updated__n,omitempty" mapstructure:"last_updated__n,omitempty"`
	// Limit: Number of results to return per page.
	Limit int32 `json:"limit,omitempty" mapstructure:"limit,omitempty"`
	// Offset: The initial index from which to return the results.
	Offset int32 `json:"offset,omitempty" mapstructure:"offset,omitempty"`
	// Ordering: Which field to use when ordering the results.
	Ordering string `json:"ordering,omitempty" mapstructure:"ordering,omitempty"`
	// Priority:
	Priority []int32 `json:"priority,omitempty" mapstructure:"priority,omitempty"`
	// PriorityGt:
	PriorityGt []int32 `json:"priority__gt,omitempty" mapstructure:"priority__gt,omitempty"`
	// PriorityGte:
	PriorityGte []int32 `json:"priority__gte,omitempty" mapstructure:"priority__gte,omitempty"`
	// PriorityLt:
	PriorityLt []int32 `json:"priority__lt,omitempty" mapstructure:"priority__lt,omitempty"`
	// PriorityLte:
	PriorityLte []int32 `json:"priority__lte,omitempty" mapstructure:"priority__lte,omitempty"`
	// PriorityN:
	PriorityN []int32 `json:"priority__n,omitempty" mapstructure:"priority__n,omitempty"`
	// UpdatedByRequest:
	UpdatedByRequest string `json:"updated_by_request,omitempty" mapstructure:"updated_by_request,omitempty"`
	// VirtualMachine:
	VirtualMachine []string `json:"virtual_machine,omitempty" mapstructure:"virtual_machine,omitempty"`
	// VirtualMachineId:
	VirtualMachineId []int32 `json:"virtual_machine_id,omitempty" mapstructure:"virtual_machine_id,omitempty"`
}

// Validate implements basic validation for this model
func (m IpamFhrpGroupAssignmentsListQueryParameters) Validate() error {
	return validation.Errors{
		"created": validation.Validate(
			m.Created,
		),
		"createdGt": validation.Validate(
			m.CreatedGt,
		),
		"createdGte": validation.Validate(
			m.CreatedGte,
		),
		"createdLt": validation.Validate(
			m.CreatedLt,
		),
		"createdLte": validation.Validate(
			m.CreatedLte,
		),
		"createdN": validation.Validate(
			m.CreatedN,
		),
		"createdByRequest": validation.Validate(
			m.CreatedByRequest, is.UUID,
		),
		"device": validation.Validate(
			m.Device,
		),
		"deviceId": validation.Validate(
			m.DeviceId,
		),
		"groupId": validation.Validate(
			m.GroupId,
		),
		"groupIdN": validation.Validate(
			m.GroupIdN,
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
		"interfaceId": validation.Validate(
			m.InterfaceId,
		),
		"interfaceIdGt": validation.Validate(
			m.InterfaceIdGt,
		),
		"interfaceIdGte": validation.Validate(
			m.InterfaceIdGte,
		),
		"interfaceIdLt": validation.Validate(
			m.InterfaceIdLt,
		),
		"interfaceIdLte": validation.Validate(
			m.InterfaceIdLte,
		),
		"interfaceIdN": validation.Validate(
			m.InterfaceIdN,
		),
		"lastUpdated": validation.Validate(
			m.LastUpdated,
		),
		"lastUpdatedGt": validation.Validate(
			m.LastUpdatedGt,
		),
		"lastUpdatedGte": validation.Validate(
			m.LastUpdatedGte,
		),
		"lastUpdatedLt": validation.Validate(
			m.LastUpdatedLt,
		),
		"lastUpdatedLte": validation.Validate(
			m.LastUpdatedLte,
		),
		"lastUpdatedN": validation.Validate(
			m.LastUpdatedN,
		),
		"priority": validation.Validate(
			m.Priority,
		),
		"priorityGt": validation.Validate(
			m.PriorityGt,
		),
		"priorityGte": validation.Validate(
			m.PriorityGte,
		),
		"priorityLt": validation.Validate(
			m.PriorityLt,
		),
		"priorityLte": validation.Validate(
			m.PriorityLte,
		),
		"priorityN": validation.Validate(
			m.PriorityN,
		),
		"updatedByRequest": validation.Validate(
			m.UpdatedByRequest, is.UUID,
		),
		"virtualMachine": validation.Validate(
			m.VirtualMachine,
		),
		"virtualMachineId": validation.Validate(
			m.VirtualMachineId,
		),
	}.Filter()
}

// GetCreated returns the Created property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreated() []time.Time {
	return m.Created
}

// SetCreated sets the Created property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreated(val []time.Time) {
	m.Created = val
}

// GetCreatedGt returns the CreatedGt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedGt() []time.Time {
	return m.CreatedGt
}

// SetCreatedGt sets the CreatedGt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedGt(val []time.Time) {
	m.CreatedGt = val
}

// GetCreatedGte returns the CreatedGte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedGte() []time.Time {
	return m.CreatedGte
}

// SetCreatedGte sets the CreatedGte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedGte(val []time.Time) {
	m.CreatedGte = val
}

// GetCreatedLt returns the CreatedLt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedLt() []time.Time {
	return m.CreatedLt
}

// SetCreatedLt sets the CreatedLt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedLt(val []time.Time) {
	m.CreatedLt = val
}

// GetCreatedLte returns the CreatedLte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedLte() []time.Time {
	return m.CreatedLte
}

// SetCreatedLte sets the CreatedLte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedLte(val []time.Time) {
	m.CreatedLte = val
}

// GetCreatedN returns the CreatedN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedN() []time.Time {
	return m.CreatedN
}

// SetCreatedN sets the CreatedN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedN(val []time.Time) {
	m.CreatedN = val
}

// GetCreatedByRequest returns the CreatedByRequest property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetCreatedByRequest() string {
	return m.CreatedByRequest
}

// SetCreatedByRequest sets the CreatedByRequest property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetCreatedByRequest(val string) {
	m.CreatedByRequest = val
}

// GetDevice returns the Device property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetDevice() []string {
	return m.Device
}

// SetDevice sets the Device property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetDevice(val []string) {
	m.Device = val
}

// GetDeviceId returns the DeviceId property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetDeviceId() []int32 {
	return m.DeviceId
}

// SetDeviceId sets the DeviceId property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetDeviceId(val []int32) {
	m.DeviceId = val
}

// GetGroupId returns the GroupId property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetGroupId() []int32 {
	return m.GroupId
}

// SetGroupId sets the GroupId property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetGroupId(val []int32) {
	m.GroupId = val
}

// GetGroupIdN returns the GroupIdN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetGroupIdN() []int32 {
	return m.GroupIdN
}

// SetGroupIdN sets the GroupIdN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetGroupIdN(val []int32) {
	m.GroupIdN = val
}

// GetId returns the Id property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetId() []int32 {
	return m.Id
}

// SetId sets the Id property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetId(val []int32) {
	m.Id = val
}

// GetIdGt returns the IdGt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetIdGt() []int32 {
	return m.IdGt
}

// SetIdGt sets the IdGt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetIdGt(val []int32) {
	m.IdGt = val
}

// GetIdGte returns the IdGte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetIdGte() []int32 {
	return m.IdGte
}

// SetIdGte sets the IdGte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetIdGte(val []int32) {
	m.IdGte = val
}

// GetIdLt returns the IdLt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetIdLt() []int32 {
	return m.IdLt
}

// SetIdLt sets the IdLt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetIdLt(val []int32) {
	m.IdLt = val
}

// GetIdLte returns the IdLte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetIdLte() []int32 {
	return m.IdLte
}

// SetIdLte sets the IdLte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetIdLte(val []int32) {
	m.IdLte = val
}

// GetIdN returns the IdN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetIdN() []int32 {
	return m.IdN
}

// SetIdN sets the IdN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetIdN(val []int32) {
	m.IdN = val
}

// GetInterfaceId returns the InterfaceId property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceId() []int32 {
	return m.InterfaceId
}

// SetInterfaceId sets the InterfaceId property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceId(val []int32) {
	m.InterfaceId = val
}

// GetInterfaceIdGt returns the InterfaceIdGt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceIdGt() []int32 {
	return m.InterfaceIdGt
}

// SetInterfaceIdGt sets the InterfaceIdGt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceIdGt(val []int32) {
	m.InterfaceIdGt = val
}

// GetInterfaceIdGte returns the InterfaceIdGte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceIdGte() []int32 {
	return m.InterfaceIdGte
}

// SetInterfaceIdGte sets the InterfaceIdGte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceIdGte(val []int32) {
	m.InterfaceIdGte = val
}

// GetInterfaceIdLt returns the InterfaceIdLt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceIdLt() []int32 {
	return m.InterfaceIdLt
}

// SetInterfaceIdLt sets the InterfaceIdLt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceIdLt(val []int32) {
	m.InterfaceIdLt = val
}

// GetInterfaceIdLte returns the InterfaceIdLte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceIdLte() []int32 {
	return m.InterfaceIdLte
}

// SetInterfaceIdLte sets the InterfaceIdLte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceIdLte(val []int32) {
	m.InterfaceIdLte = val
}

// GetInterfaceIdN returns the InterfaceIdN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceIdN() []int32 {
	return m.InterfaceIdN
}

// SetInterfaceIdN sets the InterfaceIdN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceIdN(val []int32) {
	m.InterfaceIdN = val
}

// GetInterfaceType returns the InterfaceType property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceType() string {
	return m.InterfaceType
}

// SetInterfaceType sets the InterfaceType property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceType(val string) {
	m.InterfaceType = val
}

// GetInterfaceTypeN returns the InterfaceTypeN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetInterfaceTypeN() string {
	return m.InterfaceTypeN
}

// SetInterfaceTypeN sets the InterfaceTypeN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetInterfaceTypeN(val string) {
	m.InterfaceTypeN = val
}

// GetLastUpdated returns the LastUpdated property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdated() []time.Time {
	return m.LastUpdated
}

// SetLastUpdated sets the LastUpdated property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdated(val []time.Time) {
	m.LastUpdated = val
}

// GetLastUpdatedGt returns the LastUpdatedGt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdatedGt() []time.Time {
	return m.LastUpdatedGt
}

// SetLastUpdatedGt sets the LastUpdatedGt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdatedGt(val []time.Time) {
	m.LastUpdatedGt = val
}

// GetLastUpdatedGte returns the LastUpdatedGte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdatedGte() []time.Time {
	return m.LastUpdatedGte
}

// SetLastUpdatedGte sets the LastUpdatedGte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdatedGte(val []time.Time) {
	m.LastUpdatedGte = val
}

// GetLastUpdatedLt returns the LastUpdatedLt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdatedLt() []time.Time {
	return m.LastUpdatedLt
}

// SetLastUpdatedLt sets the LastUpdatedLt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdatedLt(val []time.Time) {
	m.LastUpdatedLt = val
}

// GetLastUpdatedLte returns the LastUpdatedLte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdatedLte() []time.Time {
	return m.LastUpdatedLte
}

// SetLastUpdatedLte sets the LastUpdatedLte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdatedLte(val []time.Time) {
	m.LastUpdatedLte = val
}

// GetLastUpdatedN returns the LastUpdatedN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLastUpdatedN() []time.Time {
	return m.LastUpdatedN
}

// SetLastUpdatedN sets the LastUpdatedN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLastUpdatedN(val []time.Time) {
	m.LastUpdatedN = val
}

// GetLimit returns the Limit property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetLimit() int32 {
	return m.Limit
}

// SetLimit sets the Limit property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetLimit(val int32) {
	m.Limit = val
}

// GetOffset returns the Offset property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetOffset() int32 {
	return m.Offset
}

// SetOffset sets the Offset property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetOffset(val int32) {
	m.Offset = val
}

// GetOrdering returns the Ordering property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetOrdering() string {
	return m.Ordering
}

// SetOrdering sets the Ordering property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetOrdering(val string) {
	m.Ordering = val
}

// GetPriority returns the Priority property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriority() []int32 {
	return m.Priority
}

// SetPriority sets the Priority property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriority(val []int32) {
	m.Priority = val
}

// GetPriorityGt returns the PriorityGt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriorityGt() []int32 {
	return m.PriorityGt
}

// SetPriorityGt sets the PriorityGt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriorityGt(val []int32) {
	m.PriorityGt = val
}

// GetPriorityGte returns the PriorityGte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriorityGte() []int32 {
	return m.PriorityGte
}

// SetPriorityGte sets the PriorityGte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriorityGte(val []int32) {
	m.PriorityGte = val
}

// GetPriorityLt returns the PriorityLt property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriorityLt() []int32 {
	return m.PriorityLt
}

// SetPriorityLt sets the PriorityLt property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriorityLt(val []int32) {
	m.PriorityLt = val
}

// GetPriorityLte returns the PriorityLte property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriorityLte() []int32 {
	return m.PriorityLte
}

// SetPriorityLte sets the PriorityLte property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriorityLte(val []int32) {
	m.PriorityLte = val
}

// GetPriorityN returns the PriorityN property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetPriorityN() []int32 {
	return m.PriorityN
}

// SetPriorityN sets the PriorityN property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetPriorityN(val []int32) {
	m.PriorityN = val
}

// GetUpdatedByRequest returns the UpdatedByRequest property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetUpdatedByRequest() string {
	return m.UpdatedByRequest
}

// SetUpdatedByRequest sets the UpdatedByRequest property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetUpdatedByRequest(val string) {
	m.UpdatedByRequest = val
}

// GetVirtualMachine returns the VirtualMachine property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetVirtualMachine() []string {
	return m.VirtualMachine
}

// SetVirtualMachine sets the VirtualMachine property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetVirtualMachine(val []string) {
	m.VirtualMachine = val
}

// GetVirtualMachineId returns the VirtualMachineId property
func (m IpamFhrpGroupAssignmentsListQueryParameters) GetVirtualMachineId() []int32 {
	return m.VirtualMachineId
}

// SetVirtualMachineId sets the VirtualMachineId property
func (m *IpamFhrpGroupAssignmentsListQueryParameters) SetVirtualMachineId(val []int32) {
	m.VirtualMachineId = val
}
