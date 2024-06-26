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
	"fmt"
	"net/url"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &WirelessRFProfileDataSource{}
	_ datasource.DataSourceWithConfigure = &WirelessRFProfileDataSource{}
)

func NewWirelessRFProfileDataSource() datasource.DataSource {
	return &WirelessRFProfileDataSource{}
}

type WirelessRFProfileDataSource struct {
	client *cc.Client
}

func (d *WirelessRFProfileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_rf_profile"
}

func (d *WirelessRFProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Wireless RF Profile.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "RF Profile Name",
				Computed:            true,
			},
			"default_rf_profile": schema.BoolAttribute{
				MarkdownDescription: "is Default Rf Profile",
				Computed:            true,
			},
			"enable_radio_type_a": schema.BoolAttribute{
				MarkdownDescription: "Enable Radio Type A",
				Computed:            true,
			},
			"enable_radio_type_b": schema.BoolAttribute{
				MarkdownDescription: "Enable Radio Type B",
				Computed:            true,
			},
			"enable_radio_type_c": schema.BoolAttribute{
				MarkdownDescription: "Enable Radio Type C (6GHz)",
				Computed:            true,
			},
			"channel_width": schema.StringAttribute{
				MarkdownDescription: "Channel Width",
				Computed:            true,
			},
			"enable_custom": schema.BoolAttribute{
				MarkdownDescription: "Enable Custom",
				Computed:            true,
			},
			"enable_brown_field": schema.BoolAttribute{
				MarkdownDescription: "Enable Brown Field",
				Computed:            true,
			},
			"radio_type_a_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Parent Profile",
				Computed:            true,
			},
			"radio_type_a_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Radio Channels",
				Computed:            true,
			},
			"radio_type_a_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Data Rates",
				Computed:            true,
			},
			"radio_type_a_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Mandatory Data Rates",
				Computed:            true,
			},
			"radio_type_a_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Power Threshold V1",
				Computed:            true,
			},
			"radio_type_a_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeA Properties - Rx Sop Threshold",
				Computed:            true,
			},
			"radio_type_a_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Min Power Level",
				Computed:            true,
			},
			"radio_type_a_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeA Properties - Max Power Level",
				Computed:            true,
			},
			"radio_type_b_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Parent Profile",
				Computed:            true,
			},
			"radio_type_b_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Radio Channels",
				Computed:            true,
			},
			"radio_type_b_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Data Rates",
				Computed:            true,
			},
			"radio_type_b_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Mandatory Data Rates",
				Computed:            true,
			},
			"radio_type_b_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Power Threshold V1",
				Computed:            true,
			},
			"radio_type_b_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeB Properties - Rx Sop Threshold",
				Computed:            true,
			},
			"radio_type_b_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Min Power Level",
				Computed:            true,
			},
			"radio_type_b_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeB Properties - Max Power Level",
				Computed:            true,
			},
			"radio_type_c_parent_profile": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Parent Profile",
				Computed:            true,
			},
			"radio_type_c_radio_channels": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Radio Channels",
				Computed:            true,
			},
			"radio_type_c_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Data Rates",
				Computed:            true,
			},
			"radio_type_c_mandatory_data_rates": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Mandatory Data Rates",
				Computed:            true,
			},
			"radio_type_c_power_treshold_v1": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Power Threshold V1",
				Computed:            true,
			},
			"radio_type_c_rx_sop_threshold": schema.StringAttribute{
				MarkdownDescription: "Radio TypeC Properties - Rx Sop Threshold",
				Computed:            true,
			},
			"radio_type_c_min_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Min Power Level",
				Computed:            true,
			},
			"radio_type_c_max_power_level": schema.Int64Attribute{
				MarkdownDescription: "Radio TypeC Properties - Max Power Level",
				Computed:            true,
			},
		},
	}
}

func (d *WirelessRFProfileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

//template:begin read
func (d *WirelessRFProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config WirelessRFProfile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?rf-profile-name=" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get(config.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}

//template:end read
