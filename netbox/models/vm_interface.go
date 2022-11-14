// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VMInterface VM interface
//
// swagger:model VMInterface
type VMInterface struct {

	// bridge
	Bridge *NestedVMInterface `json:"bridge,omitempty"`

	// Count fhrp groups
	// Read Only: true
	CountFhrpGroups int64 `json:"count_fhrp_groups,omitempty"`

	// Count ipaddresses
	// Read Only: true
	CountIpaddresses int64 `json:"count_ipaddresses,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Enabled
	Enabled bool `json:"enabled,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// l2vpn termination
	L2vpnTermination *NestedL2VPNTermination `json:"l2vpn_termination,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// MAC Address
	MacAddress *string `json:"mac_address,omitempty"`

	// mode
	Mode *VMInterfaceMode `json:"mode,omitempty"`

	// MTU
	// Maximum: 65536
	// Minimum: 1
	Mtu *int64 `json:"mtu,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// parent
	Parent *NestedVMInterface `json:"parent,omitempty"`

	// tagged vlans
	// Unique: true
	TaggedVlans []*NestedVLAN `json:"tagged_vlans"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// untagged vlan
	UntaggedVlan *NestedVLAN `json:"untagged_vlan,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// virtual machine
	// Required: true
	VirtualMachine *NestedVirtualMachine `json:"virtual_machine"`

	// vrf
	Vrf *NestedVRF `json:"vrf,omitempty"`
}

// Validate validates this VM interface
func (m *VMInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBridge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateL2vpnTermination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMtu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTaggedVlans(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntaggedVlan(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirtualMachine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVrf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMInterface) validateBridge(formats strfmt.Registry) error {
	if swag.IsZero(m.Bridge) { // not required
		return nil
	}

	if m.Bridge != nil {
		if err := m.Bridge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bridge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bridge")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateL2vpnTermination(formats strfmt.Registry) error {
	if swag.IsZero(m.L2vpnTermination) { // not required
		return nil
	}

	if m.L2vpnTermination != nil {
		if err := m.L2vpnTermination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l2vpn_termination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l2vpn_termination")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	if m.Mode != nil {
		if err := m.Mode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateMtu(formats strfmt.Registry) error {
	if swag.IsZero(m.Mtu) { // not required
		return nil
	}

	if err := validate.MinimumInt("mtu", "body", *m.Mtu, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("mtu", "body", *m.Mtu, 65536, false); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateParent(formats strfmt.Registry) error {
	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateTaggedVlans(formats strfmt.Registry) error {
	if swag.IsZero(m.TaggedVlans) { // not required
		return nil
	}

	if err := validate.UniqueItems("tagged_vlans", "body", m.TaggedVlans); err != nil {
		return err
	}

	for i := 0; i < len(m.TaggedVlans); i++ {
		if swag.IsZero(m.TaggedVlans[i]) { // not required
			continue
		}

		if m.TaggedVlans[i] != nil {
			if err := m.TaggedVlans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMInterface) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMInterface) validateUntaggedVlan(formats strfmt.Registry) error {
	if swag.IsZero(m.UntaggedVlan) { // not required
		return nil
	}

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) validateVirtualMachine(formats strfmt.Registry) error {

	if err := validate.Required("virtual_machine", "body", m.VirtualMachine); err != nil {
		return err
	}

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) validateVrf(formats strfmt.Registry) error {
	if swag.IsZero(m.Vrf) { // not required
		return nil
	}

	if m.Vrf != nil {
		if err := m.Vrf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this VM interface based on the context it is used
func (m *VMInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBridge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountFhrpGroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCountIpaddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateL2vpnTermination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTaggedVlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUntaggedVlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVirtualMachine(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVrf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VMInterface) contextValidateBridge(ctx context.Context, formats strfmt.Registry) error {

	if m.Bridge != nil {
		if err := m.Bridge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bridge")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bridge")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateCountFhrpGroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count_fhrp_groups", "body", int64(m.CountFhrpGroups)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateCountIpaddresses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "count_ipaddresses", "body", int64(m.CountIpaddresses)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateL2vpnTermination(ctx context.Context, formats strfmt.Registry) error {

	if m.L2vpnTermination != nil {
		if err := m.L2vpnTermination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l2vpn_termination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("l2vpn_termination")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateMode(ctx context.Context, formats strfmt.Registry) error {

	if m.Mode != nil {
		if err := m.Mode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateParent(ctx context.Context, formats strfmt.Registry) error {

	if m.Parent != nil {
		if err := m.Parent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateTaggedVlans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TaggedVlans); i++ {

		if m.TaggedVlans[i] != nil {
			if err := m.TaggedVlans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tagged_vlans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMInterface) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VMInterface) contextValidateUntaggedVlan(ctx context.Context, formats strfmt.Registry) error {

	if m.UntaggedVlan != nil {
		if err := m.UntaggedVlan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("untagged_vlan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("untagged_vlan")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *VMInterface) contextValidateVirtualMachine(ctx context.Context, formats strfmt.Registry) error {

	if m.VirtualMachine != nil {
		if err := m.VirtualMachine.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("virtual_machine")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("virtual_machine")
			}
			return err
		}
	}

	return nil
}

func (m *VMInterface) contextValidateVrf(ctx context.Context, formats strfmt.Registry) error {

	if m.Vrf != nil {
		if err := m.Vrf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vrf")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vrf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VMInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMInterface) UnmarshalBinary(b []byte) error {
	var res VMInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// VMInterfaceMode Mode
//
// swagger:model VMInterfaceMode
type VMInterfaceMode struct {

	// label
	// Required: true
	// Enum: [Access Tagged Tagged (All)]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [access tagged tagged-all]
	Value *string `json:"value"`
}

// Validate validates this VM interface mode
func (m *VMInterfaceMode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var vmInterfaceModeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Access","Tagged","Tagged (All)"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmInterfaceModeTypeLabelPropEnum = append(vmInterfaceModeTypeLabelPropEnum, v)
	}
}

const (

	// VMInterfaceModeLabelAccess captures enum value "Access"
	VMInterfaceModeLabelAccess string = "Access"

	// VMInterfaceModeLabelTagged captures enum value "Tagged"
	VMInterfaceModeLabelTagged string = "Tagged"

	// VMInterfaceModeLabelTaggedAll captures enum value "Tagged (All)"
	VMInterfaceModeLabelTaggedAll string = "Tagged (All)"
)

// prop value enum
func (m *VMInterfaceMode) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmInterfaceModeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VMInterfaceMode) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("mode"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var vmInterfaceModeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","tagged","tagged-all"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmInterfaceModeTypeValuePropEnum = append(vmInterfaceModeTypeValuePropEnum, v)
	}
}

const (

	// VMInterfaceModeValueAccess captures enum value "access"
	VMInterfaceModeValueAccess string = "access"

	// VMInterfaceModeValueTagged captures enum value "tagged"
	VMInterfaceModeValueTagged string = "tagged"

	// VMInterfaceModeValueTaggedDashAll captures enum value "tagged-all"
	VMInterfaceModeValueTaggedDashAll string = "tagged-all"
)

// prop value enum
func (m *VMInterfaceMode) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, vmInterfaceModeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VMInterfaceMode) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("mode"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("mode"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this VM interface mode based on context it is used
func (m *VMInterfaceMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *VMInterfaceMode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VMInterfaceMode) UnmarshalBinary(b []byte) error {
	var res VMInterfaceMode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
