// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

//template:begin imports
import (
	"context"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type FabricVirtualNetwork struct {
	Id                 types.String `tfsdk:"id"`
	VirtualNetworkName types.String `tfsdk:"virtual_network_name"`
	IsGuest            types.Bool   `tfsdk:"is_guest"`
	SgNames            types.List   `tfsdk:"sg_names"`
	VmanageVpnId       types.String `tfsdk:"vmanage_vpn_id"`
}

//template:end types

//template:begin getPath
func (data FabricVirtualNetwork) getPath() string {
	return "/dna/intent/api/v1/virtual-network"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data FabricVirtualNetwork) toBody(ctx context.Context, state FabricVirtualNetwork) string {
	body := ""
	if !data.VirtualNetworkName.IsNull() {
		body, _ = sjson.Set(body, "virtualNetworkName", data.VirtualNetworkName.ValueString())
	}
	if !data.IsGuest.IsNull() {
		body, _ = sjson.Set(body, "isGuestVirtualNetwork", data.IsGuest.ValueBool())
	}
	if !data.SgNames.IsNull() {
		var values []string
		data.SgNames.ElementsAs(ctx, &values, false)
		body, _ = sjson.Set(body, "scalableGroupNames", values)
	}
	if !data.VmanageVpnId.IsNull() {
		body, _ = sjson.Set(body, "vManageVpnId", data.VmanageVpnId.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *FabricVirtualNetwork) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("isGuestVirtualNetwork"); value.Exists() {
		data.IsGuest = types.BoolValue(value.Bool())
	} else {
		data.IsGuest = types.BoolValue(false)
	}
	if value := res.Get("scalableGroupNames"); value.Exists() {
		data.SgNames = helpers.GetStringList(value.Array())
	} else {
		data.SgNames = types.ListNull(types.StringType)
	}
	if value := res.Get("vManageVpnId"); value.Exists() {
		data.VmanageVpnId = types.StringValue(value.String())
	} else {
		data.VmanageVpnId = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *FabricVirtualNetwork) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("virtualNetworkName"); value.Exists() && !data.VirtualNetworkName.IsNull() {
		data.VirtualNetworkName = types.StringValue(value.String())
	} else {
		data.VirtualNetworkName = types.StringNull()
	}
	if value := res.Get("isGuestVirtualNetwork"); value.Exists() && !data.IsGuest.IsNull() {
		data.IsGuest = types.BoolValue(value.Bool())
	} else if data.IsGuest.ValueBool() != false {
		data.IsGuest = types.BoolNull()
	}
	if value := res.Get("scalableGroupNames"); value.Exists() && !data.SgNames.IsNull() {
		data.SgNames = helpers.GetStringList(value.Array())
	} else {
		data.SgNames = types.ListNull(types.StringType)
	}
	if value := res.Get("vManageVpnId"); value.Exists() && !data.VmanageVpnId.IsNull() {
		data.VmanageVpnId = types.StringValue(value.String())
	} else {
		data.VmanageVpnId = types.StringNull()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *FabricVirtualNetwork) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.VirtualNetworkName.IsNull() {
		return false
	}
	if !data.IsGuest.IsNull() {
		return false
	}
	if !data.SgNames.IsNull() {
		return false
	}
	if !data.VmanageVpnId.IsNull() {
		return false
	}
	return true
}

//template:end isNull