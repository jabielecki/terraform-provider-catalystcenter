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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

//template:end imports

//template:begin types
type DeviceToSiteAssignment struct {
	Id types.String `tfsdk:"id"`
	Ip types.String `tfsdk:"ip"`
}

//template:end types

//template:begin getPath
func (data DeviceToSiteAssignment) getPath() string {
	return "/dna/intent/api/v1/assign-device-to-site//device"
}

//template:end getPath

//template:begin toBody
func (data DeviceToSiteAssignment) toBody(ctx context.Context, state DeviceToSiteAssignment) string {
	body := ""
	if !data.Ip.IsNull() {
		body, _ = sjson.Set(body, "device.0.ip", data.Ip.ValueString())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *DeviceToSiteAssignment) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("device.0.ip"); value.Exists() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *DeviceToSiteAssignment) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("device.0.ip"); value.Exists() && !data.Ip.IsNull() {
		data.Ip = types.StringValue(value.String())
	} else {
		data.Ip = types.StringNull()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *DeviceToSiteAssignment) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Ip.IsNull() {
		return false
	}
	return true
}

//template:end isNull
