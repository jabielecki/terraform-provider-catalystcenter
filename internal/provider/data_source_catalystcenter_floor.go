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
	_ datasource.DataSource              = &FloorDataSource{}
	_ datasource.DataSourceWithConfigure = &FloorDataSource{}
)

func NewFloorDataSource() datasource.DataSource {
	return &FloorDataSource{}
}

type FloorDataSource struct {
	client *cc.Client
}

func (d *FloorDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_floor"
}

func (d *FloorDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Floor.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the floor",
				Computed:            true,
			},
			"parent_name": schema.StringAttribute{
				MarkdownDescription: "The path of the parent building, e.g. `Global/Building1`",
				Computed:            true,
			},
			"floor_number": schema.Int64Attribute{
				MarkdownDescription: "Floor number",
				Computed:            true,
			},
			"rf_model": schema.StringAttribute{
				MarkdownDescription: "The RF model",
				Computed:            true,
			},
			"width": schema.Float64Attribute{
				MarkdownDescription: "Width",
				Computed:            true,
			},
			"length": schema.Float64Attribute{
				MarkdownDescription: "Length",
				Computed:            true,
			},
			"height": schema.Float64Attribute{
				MarkdownDescription: "Height",
				Computed:            true,
			},
		},
	}
}

func (d *FloorDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

//template:begin read
func (d *FloorDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Floor

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "/" + url.QueryEscape(config.Id.ValueString())
	res, err := d.client.Get("/api/v1/dna-maps-service/domains" + params)
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
