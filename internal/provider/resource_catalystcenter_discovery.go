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
	"strings"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

//template:end imports

//template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &DiscoveryResource{}
var _ resource.ResourceWithImportState = &DiscoveryResource{}

func NewDiscoveryResource() resource.Resource {
	return &DiscoveryResource{}
}

type DiscoveryResource struct {
	client *cc.Client
}

func (r *DiscoveryResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_discovery"
}

func (r *DiscoveryResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("This resource can manage a Discovery.").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"cdp_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("CDP level is the number of hops between neighbor devices.").String,
				Optional:            true,
			},
			"discovery_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Type of Discovery.").AddStringEnumDescription("SINGLE", "RANGE", "MULTI RANGE", "CDP", "LLDP", "CIDR").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SINGLE", "RANGE", "MULTI RANGE", "CDP", "LLDP", "CIDR"),
				},
			},
			"enable_password_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"global_credential_id_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of IDs, which must include SNMP credential and CLI credential.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"http_read_credential": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"http_write_credential": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"ip_address_list": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A string of IP address ranges to discover.  E.g.: '172.30.0.1' for SINGLE, CDP and LLDP; '172.30.0.1-172.30.0.4' for RANGE; '72.30.0.1-172.30.0.4,172.31.0.1-172.31.0.4' for MULTI RANGE; '172.30.0.1/20' for CIDR.").String,
				Optional:            true,
			},
			"ip_filter_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A list of IP address ranges to exclude from the discovery.").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"lldp_level": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("LLDP level to which neighbor devices to be discovered.").String,
				Optional:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A name for the Discovery.").String,
				Required:            true,
			},
			"netconf_port": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Port number for netconf as a string. It requires valid SSH credentials to work.").String,
				Optional:            true,
			},
			"password_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
			"preferred_ip_method": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Preferred method for selecting management IP address.").AddStringEnumDescription("None", "UseLoopBack").AddDefaultValueDescription("None").String,
				Optional:            true,
				Computed:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("None", "UseLoopBack"),
				},
				Default: stringdefault.StaticString("None"),
			},
			"protocol_order": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("A string of comma-separated protocols (ssh/telnet), in the same order in which the connections to each device are attempted. E.g.: 'telnet': only telnet; 'ssh,telnet': ssh first, with telnet fallback.").String,
				Required:            true,
			},
			"retry": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of times to try establishing SSH connection to a device.").String,
				Optional:            true,
			},
			"snmp_auth_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Auth passphrase for SNMP.").String,
				Optional:            true,
			},
			"snmp_auth_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP auth protocol.").AddStringEnumDescription("SHA", "MD5").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("SHA", "MD5"),
				},
			},
			"snmp_mode": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("AUTHPRIV", "AUTHNOPRIV", "NOAUTHNOPRIV").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AUTHPRIV", "AUTHNOPRIV", "NOAUTHNOPRIV"),
				},
			},
			"snmp_priv_passphrase": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				Optional:            true,
			},
			"snmp_priv_protocol": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").AddStringEnumDescription("DES", "AES128").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DES", "AES128"),
				},
			},
			"snmp_ro_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP RO community of the devices to be discovered.").String,
				Optional:            true,
			},
			"snmp_ro_community_desc": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for snmp_ro_community.").String,
				Optional:            true,
			},
			"snmp_rw_community": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP RW community of the devices to be discovered.").String,
				Optional:            true,
			},
			"snmp_rw_community_desc": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Description for snmp_rw_community").String,
				Optional:            true,
			},
			"snmp_user_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP username of the devices to be discovered.").String,
				Optional:            true,
			},
			"snmp_version": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("SNMP version").AddStringEnumDescription("v2", "v3").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("v2", "v3"),
				},
			},
			"timeout": schema.Int64Attribute{
				MarkdownDescription: helpers.NewAttributeDescription("Number of seconds to wait for each SSH connection to a device.").String,
				Optional:            true,
			},
			"user_name_list": schema.ListAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("").String,
				ElementType:         types.StringType,
				Optional:            true,
			},
		},
	}
}

func (r *DiscoveryResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

//template:begin create
func (r *DiscoveryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Discovery

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, Discovery{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:end create

//template:begin read
func (r *DiscoveryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Discovery

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	params := ""
	params += "/" + state.Id.ValueString()
	res, err := r.client.Get("/dna/intent/api/v1/discovery" + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	state.updateFromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

//template:end read

//template:begin update
func (r *DiscoveryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Discovery

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// Read state
	diags = req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Update", plan.Id.ValueString()))

	body := plan.toBody(ctx, state)
	params := ""

	res, err := r.client.Put(plan.getPath()+"/"+plan.Id.ValueString()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (PUT), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:end update

//template:begin delete
func (r *DiscoveryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Discovery

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	res, err := r.client.Delete(state.getPath() + "/" + state.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

//template:end delete

//template:begin import
func (r *DiscoveryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

//template:end import
