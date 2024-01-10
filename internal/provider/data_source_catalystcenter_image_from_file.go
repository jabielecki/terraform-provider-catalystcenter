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

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

//template:end imports

//template:begin model

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &ImageFromFileDataSource{}
	_ datasource.DataSourceWithConfigure = &ImageFromFileDataSource{}
)

func NewImageFromFileDataSource() datasource.DataSource {
	return &ImageFromFileDataSource{}
}

type ImageFromFileDataSource struct {
	client *cc.Client
}

func (d *ImageFromFileDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_image_from_file"
}

func (d *ImageFromFileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Remove me.",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"third_party_application_type": schema.StringAttribute{
				MarkdownDescription: "Third party application type",
				Computed:            true,
			},
			"family": schema.StringAttribute{
				MarkdownDescription: "Third party image family",
				Computed:            true,
			},
			"source_path": schema.StringAttribute{
				MarkdownDescription: "Local path where the software image file resides. Supported file extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz, qcow2.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "File name that uniquely identifies the software image. It should not contain any path. Usually this can be specified as `basename(source_path)`",
				Required:            true,
			},
			"third_party_vendor": schema.StringAttribute{
				MarkdownDescription: "Third party Vendor",
				Computed:            true,
			},
			"is_third_party": schema.BoolAttribute{
				MarkdownDescription: "Whether the software image is from a third party.",
				Computed:            true,
			},
		},
	}
}

func (d *ImageFromFileDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

//template:begin read
func (d *ImageFromFileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config ImageFromFile

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	params := ""
	params += "?imageUuid=" + config.Id.ValueString()
	res, err := d.client.Get("/dna/intent/api/v1/image/importation" + params)
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
