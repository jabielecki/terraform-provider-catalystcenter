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
	"strings"
	"time"

	"github.com/CiscoDevNet/terraform-provider-catalystcenter/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	cc "github.com/netascode/go-catalystcenter"
)

//template:end imports

//template:begin model

// Ensure provider defined types fully satisfy framework interfaces
var _ resource.Resource = &VirtualNetworkIPPoolResource{}

func NewVirtualNetworkIPPoolResource() resource.Resource {
	return &VirtualNetworkIPPoolResource{}
}

type VirtualNetworkIPPoolResource struct {
	client *cc.Client
}

func (r *VirtualNetworkIPPoolResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_network_ip_pool"
}

func (r *VirtualNetworkIPPoolResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: helpers.NewAttributeDescription("Manages IP Pool in SDA Virtual Network").String,

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"virtual_network_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Virtual Network Name, that is associated to Fabric Site").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"site_name_hierarchy": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Path of sda Fabric Site").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"layer2_only": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Layer2 Only enablement flag and default value is False").AddDefaultValueDescription("false").String,
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"ip_pool_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ip Pool Name, that is reserved to Fabric Site (Required for L3 and INFRA_VN)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vlan_id": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("vlan Id(applicable for L3 , L2 and INFRA_VN)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"vlan_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Vlan name represent the segment name, if empty, vlanName would be auto generated by API").String,
				Required:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"auto_generate_vlan_name": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("It will auto generate vlanName, if vlanName is empty(applicable for L3 and INFRA_VN)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"traffic_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Traffic type(applicable for L3 and L2)").AddStringEnumDescription("DATA", "VOICE").String,
				Required:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("DATA", "VOICE"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"scalable_group_name": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Scalable Group Name(applicable for L3)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"l2_flooding_enabled": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Layer2 flooding enablement flag(applicable for L3 , L2 and always true for L2 and default value is False)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"critical_pool": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Critical pool enablement flag(applicable for L3 and default value is False )").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"wireless_pool": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Wireless Pool enablement flag(applicable for L3 and L2 and default value is False)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"ip_directed_broadcast": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Ip Directed Broadcast enablement flag(applicable for L3 and default value is False)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"common_pool": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Common Pool enablement flag(applicable for L3 and L2 and default value is False)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"bridge_mode_vm": schema.BoolAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Bridge Mode Vm enablement flag (applicable for L3 and L2 and default value is False)").String,
				Optional:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.RequiresReplace(),
				},
			},
			"pool_type": schema.StringAttribute{
				MarkdownDescription: helpers.NewAttributeDescription("Pool Type (applicable for INFRA_VN)").AddStringEnumDescription("AP", "Extended").String,
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("AP", "Extended"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
		},
	}
}

func (r *VirtualNetworkIPPoolResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*CcProviderData).Client
}

//template:end model

func (r *VirtualNetworkIPPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan VirtualNetworkIPPool

	// Read plan
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Create", plan.Id.ValueString()))

	// Create object
	body := plan.toBody(ctx, VirtualNetworkIPPool{})

	params := ""
	res, err := r.client.Post(plan.getPath()+params, body)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to configure object (POST), got error: %s, %s", err, res.String()))
		return
	}

	plan.Id = basetypes.NewStringValue(plan.IpPoolName.ValueString() + "-" + plan.VirtualNetworkName.ValueString() + "-" + plan.SiteNameHierarchy.ValueString())
	tflog.Debug(ctx, fmt.Sprintf("%s: Create finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

func (r *VirtualNetworkIPPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state VirtualNetworkIPPool

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", state.Id.String()))

	time.Sleep(15 * time.Second)

	params := ""
	params += "?ipPoolName=" + url.QueryEscape(state.IpPoolName.ValueString())
	params += "&siteNameHierarchy=" + url.QueryEscape(state.SiteNameHierarchy.ValueString())
	params += "&virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
	res, err := r.client.Get(state.getPath() + params)
	if err != nil && strings.Contains(err.Error(), "StatusCode 404") {
		resp.State.RemoveResource(ctx)
		return
	} else if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object (GET), got error: %s, %s", err, res.String()))
		return
	}

	// If every attribute is set to null we are dealing with an import operation and therefore reading all attributes
	if state.isNull(ctx, res) {
		state.fromBody(ctx, res)
	} else {
		state.updateFromBody(ctx, res)
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", state.Id.ValueString()))

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

//template:begin update
func (r *VirtualNetworkIPPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VirtualNetworkIPPool

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

	tflog.Debug(ctx, fmt.Sprintf("%s: Update finished successfully", plan.Id.ValueString()))

	diags = resp.State.Set(ctx, &plan)
	resp.Diagnostics.Append(diags...)
}

//template:end update

func (r *VirtualNetworkIPPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state VirtualNetworkIPPool

	// Read state
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Delete", state.Id.ValueString()))
	time.Sleep(15 * time.Second)

	params := ""
	params += "?ipPoolName=" + url.QueryEscape(state.IpPoolName.ValueString())
	params += "&siteNameHierarchy=" + url.QueryEscape(state.SiteNameHierarchy.ValueString())
	params += "&virtualNetworkName=" + url.QueryEscape(state.VirtualNetworkName.ValueString())
	res, err := r.client.Delete(state.getPath() + params)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to delete object (DELETE), got error: %s, %s", err, res.String()))
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Delete finished successfully", state.Id.ValueString()))

	resp.State.RemoveResource(ctx)
}

//template:begin import
//template:end import
