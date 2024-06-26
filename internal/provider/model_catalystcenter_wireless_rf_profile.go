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
type WirelessRFProfile struct {
	Id                           types.String `tfsdk:"id"`
	Name                         types.String `tfsdk:"name"`
	DefaultRfProfile             types.Bool   `tfsdk:"default_rf_profile"`
	EnableRadioTypeA             types.Bool   `tfsdk:"enable_radio_type_a"`
	EnableRadioTypeB             types.Bool   `tfsdk:"enable_radio_type_b"`
	EnableRadioTypeC             types.Bool   `tfsdk:"enable_radio_type_c"`
	ChannelWidth                 types.String `tfsdk:"channel_width"`
	EnableCustom                 types.Bool   `tfsdk:"enable_custom"`
	EnableBrownField             types.Bool   `tfsdk:"enable_brown_field"`
	RadioTypeAParentProfile      types.String `tfsdk:"radio_type_a_parent_profile"`
	RadioTypeARadioChannels      types.String `tfsdk:"radio_type_a_radio_channels"`
	RadioTypeADataRates          types.String `tfsdk:"radio_type_a_data_rates"`
	RadioTypeAMandatoryDataRates types.String `tfsdk:"radio_type_a_mandatory_data_rates"`
	RadioTypeAPowerTresholdV1    types.Int64  `tfsdk:"radio_type_a_power_treshold_v1"`
	RadioTypeARxSopThreshold     types.String `tfsdk:"radio_type_a_rx_sop_threshold"`
	RadioTypeAMinPowerLevel      types.Int64  `tfsdk:"radio_type_a_min_power_level"`
	RadioTypeAMaxPowerLevel      types.Int64  `tfsdk:"radio_type_a_max_power_level"`
	RadioTypeBParentProfile      types.String `tfsdk:"radio_type_b_parent_profile"`
	RadioTypeBRadioChannels      types.String `tfsdk:"radio_type_b_radio_channels"`
	RadioTypeBDataRates          types.String `tfsdk:"radio_type_b_data_rates"`
	RadioTypeBMandatoryDataRates types.String `tfsdk:"radio_type_b_mandatory_data_rates"`
	RadioTypeBPowerTresholdV1    types.Int64  `tfsdk:"radio_type_b_power_treshold_v1"`
	RadioTypeBRxSopThreshold     types.String `tfsdk:"radio_type_b_rx_sop_threshold"`
	RadioTypeBMinPowerLevel      types.Int64  `tfsdk:"radio_type_b_min_power_level"`
	RadioTypeBMaxPowerLevel      types.Int64  `tfsdk:"radio_type_b_max_power_level"`
	RadioTypeCParentProfile      types.String `tfsdk:"radio_type_c_parent_profile"`
	RadioTypeCRadioChannels      types.String `tfsdk:"radio_type_c_radio_channels"`
	RadioTypeCDataRates          types.String `tfsdk:"radio_type_c_data_rates"`
	RadioTypeCMandatoryDataRates types.String `tfsdk:"radio_type_c_mandatory_data_rates"`
	RadioTypeCPowerTresholdV1    types.Int64  `tfsdk:"radio_type_c_power_treshold_v1"`
	RadioTypeCRxSopThreshold     types.String `tfsdk:"radio_type_c_rx_sop_threshold"`
	RadioTypeCMinPowerLevel      types.Int64  `tfsdk:"radio_type_c_min_power_level"`
	RadioTypeCMaxPowerLevel      types.Int64  `tfsdk:"radio_type_c_max_power_level"`
}

//template:end types

//template:begin getPath
func (data WirelessRFProfile) getPath() string {
	return "/dna/intent/api/v1/wireless/rf-profile"
}

//template:end getPath

//template:begin getPathDelete

//template:end getPathDelete

//template:begin toBody
func (data WirelessRFProfile) toBody(ctx context.Context, state WirelessRFProfile) string {
	body := ""
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.DefaultRfProfile.IsNull() {
		body, _ = sjson.Set(body, "defaultRfProfile", data.DefaultRfProfile.ValueBool())
	}
	if !data.EnableRadioTypeA.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeA", data.EnableRadioTypeA.ValueBool())
	}
	if !data.EnableRadioTypeB.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeB", data.EnableRadioTypeB.ValueBool())
	}
	if !data.EnableRadioTypeC.IsNull() {
		body, _ = sjson.Set(body, "enableRadioTypeC", data.EnableRadioTypeC.ValueBool())
	}
	if !data.ChannelWidth.IsNull() {
		body, _ = sjson.Set(body, "channelWidth", data.ChannelWidth.ValueString())
	}
	if !data.EnableCustom.IsNull() {
		body, _ = sjson.Set(body, "enableCustom", data.EnableCustom.ValueBool())
	}
	if !data.EnableBrownField.IsNull() {
		body, _ = sjson.Set(body, "enableBrownField", data.EnableBrownField.ValueBool())
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.parentProfile", data.RadioTypeAParentProfile.ValueString())
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.radioChannels", data.RadioTypeARadioChannels.ValueString())
	}
	if !data.RadioTypeADataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.dataRates", data.RadioTypeADataRates.ValueString())
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.mandatoryDataRates", data.RadioTypeAMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeAPowerTresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.powerThresholdV1", data.RadioTypeAPowerTresholdV1.ValueInt64())
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.rxSopThreshold", data.RadioTypeARxSopThreshold.ValueString())
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.minPowerLevel", data.RadioTypeAMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeAProperties.maxPowerLevel", data.RadioTypeAMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.parentProfile", data.RadioTypeBParentProfile.ValueString())
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.radioChannels", data.RadioTypeBRadioChannels.ValueString())
	}
	if !data.RadioTypeBDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.dataRates", data.RadioTypeBDataRates.ValueString())
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.mandatoryDataRates", data.RadioTypeBMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeBPowerTresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.powerThresholdV1", data.RadioTypeBPowerTresholdV1.ValueInt64())
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.rxSopThreshold", data.RadioTypeBRxSopThreshold.ValueString())
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.minPowerLevel", data.RadioTypeBMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeBProperties.maxPowerLevel", data.RadioTypeBMaxPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.parentProfile", data.RadioTypeCParentProfile.ValueString())
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.radioChannels", data.RadioTypeCRadioChannels.ValueString())
	}
	if !data.RadioTypeCDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.dataRates", data.RadioTypeCDataRates.ValueString())
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.mandatoryDataRates", data.RadioTypeCMandatoryDataRates.ValueString())
	}
	if !data.RadioTypeCPowerTresholdV1.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.powerThresholdV1", data.RadioTypeCPowerTresholdV1.ValueInt64())
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.rxSopThreshold", data.RadioTypeCRxSopThreshold.ValueString())
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.minPowerLevel", data.RadioTypeCMinPowerLevel.ValueInt64())
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		body, _ = sjson.Set(body, "radioTypeCProperties.maxPowerLevel", data.RadioTypeCMaxPowerLevel.ValueInt64())
	}
	return body
}

//template:end toBody

//template:begin fromBody
func (data *WirelessRFProfile) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.defaultRfProfile"); value.Exists() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeA"); value.Exists() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeB"); value.Exists() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeC"); value.Exists() {
		data.EnableRadioTypeC = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeC = types.BoolNull()
	}
	if value := res.Get("response.0.channelWidth"); value.Exists() {
		data.ChannelWidth = types.StringValue(value.String())
	} else {
		data.ChannelWidth = types.StringNull()
	}
	if value := res.Get("response.0.enableCustom"); value.Exists() {
		data.EnableCustom = types.BoolValue(value.Bool())
	} else {
		data.EnableCustom = types.BoolNull()
	}
	if value := res.Get("response.0.enableBrownField"); value.Exists() {
		data.EnableBrownField = types.BoolValue(value.Bool())
	} else {
		data.EnableBrownField = types.BoolNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.parentProfile"); value.Exists() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.radioChannels"); value.Exists() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.dataRates"); value.Exists() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeAPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeAProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeAProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.parentProfile"); value.Exists() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.radioChannels"); value.Exists() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.dataRates"); value.Exists() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeBPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.parentProfile"); value.Exists() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.radioChannels"); value.Exists() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.dataRates"); value.Exists() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.mandatoryDataRates"); value.Exists() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.powerThresholdV1"); value.Exists() {
		data.RadioTypeCPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.rxSopThreshold"); value.Exists() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.minPowerLevel"); value.Exists() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.maxPowerLevel"); value.Exists() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
}

//template:end fromBody

//template:begin updateFromBody
func (data *WirelessRFProfile) updateFromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("response.0.name"); value.Exists() && !data.Name.IsNull() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("response.0.defaultRfProfile"); value.Exists() && !data.DefaultRfProfile.IsNull() {
		data.DefaultRfProfile = types.BoolValue(value.Bool())
	} else {
		data.DefaultRfProfile = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeA"); value.Exists() && !data.EnableRadioTypeA.IsNull() {
		data.EnableRadioTypeA = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeA = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeB"); value.Exists() && !data.EnableRadioTypeB.IsNull() {
		data.EnableRadioTypeB = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeB = types.BoolNull()
	}
	if value := res.Get("response.0.enableRadioTypeC"); value.Exists() && !data.EnableRadioTypeC.IsNull() {
		data.EnableRadioTypeC = types.BoolValue(value.Bool())
	} else {
		data.EnableRadioTypeC = types.BoolNull()
	}
	if value := res.Get("response.0.channelWidth"); value.Exists() && !data.ChannelWidth.IsNull() {
		data.ChannelWidth = types.StringValue(value.String())
	} else {
		data.ChannelWidth = types.StringNull()
	}
	if value := res.Get("response.0.enableCustom"); value.Exists() && !data.EnableCustom.IsNull() {
		data.EnableCustom = types.BoolValue(value.Bool())
	} else {
		data.EnableCustom = types.BoolNull()
	}
	if value := res.Get("response.0.enableBrownField"); value.Exists() && !data.EnableBrownField.IsNull() {
		data.EnableBrownField = types.BoolValue(value.Bool())
	} else {
		data.EnableBrownField = types.BoolNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.parentProfile"); value.Exists() && !data.RadioTypeAParentProfile.IsNull() {
		data.RadioTypeAParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeAParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.radioChannels"); value.Exists() && !data.RadioTypeARadioChannels.IsNull() {
		data.RadioTypeARadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeARadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.dataRates"); value.Exists() && !data.RadioTypeADataRates.IsNull() {
		data.RadioTypeADataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeADataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeAMandatoryDataRates.IsNull() {
		data.RadioTypeAMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeAMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeAPowerTresholdV1.IsNull() {
		data.RadioTypeAPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeAProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeARxSopThreshold.IsNull() {
		data.RadioTypeARxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeARxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeAProperties.minPowerLevel"); value.Exists() && !data.RadioTypeAMinPowerLevel.IsNull() {
		data.RadioTypeAMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeAProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeAMaxPowerLevel.IsNull() {
		data.RadioTypeAMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeAMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.parentProfile"); value.Exists() && !data.RadioTypeBParentProfile.IsNull() {
		data.RadioTypeBParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeBParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.radioChannels"); value.Exists() && !data.RadioTypeBRadioChannels.IsNull() {
		data.RadioTypeBRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeBRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.dataRates"); value.Exists() && !data.RadioTypeBDataRates.IsNull() {
		data.RadioTypeBDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeBMandatoryDataRates.IsNull() {
		data.RadioTypeBMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeBMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeBPowerTresholdV1.IsNull() {
		data.RadioTypeBPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeBRxSopThreshold.IsNull() {
		data.RadioTypeBRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeBRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeBProperties.minPowerLevel"); value.Exists() && !data.RadioTypeBMinPowerLevel.IsNull() {
		data.RadioTypeBMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeBProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeBMaxPowerLevel.IsNull() {
		data.RadioTypeBMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeBMaxPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.parentProfile"); value.Exists() && !data.RadioTypeCParentProfile.IsNull() {
		data.RadioTypeCParentProfile = types.StringValue(value.String())
	} else {
		data.RadioTypeCParentProfile = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.radioChannels"); value.Exists() && !data.RadioTypeCRadioChannels.IsNull() {
		data.RadioTypeCRadioChannels = types.StringValue(value.String())
	} else {
		data.RadioTypeCRadioChannels = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.dataRates"); value.Exists() && !data.RadioTypeCDataRates.IsNull() {
		data.RadioTypeCDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.mandatoryDataRates"); value.Exists() && !data.RadioTypeCMandatoryDataRates.IsNull() {
		data.RadioTypeCMandatoryDataRates = types.StringValue(value.String())
	} else {
		data.RadioTypeCMandatoryDataRates = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.powerThresholdV1"); value.Exists() && !data.RadioTypeCPowerTresholdV1.IsNull() {
		data.RadioTypeCPowerTresholdV1 = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCPowerTresholdV1 = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.rxSopThreshold"); value.Exists() && !data.RadioTypeCRxSopThreshold.IsNull() {
		data.RadioTypeCRxSopThreshold = types.StringValue(value.String())
	} else {
		data.RadioTypeCRxSopThreshold = types.StringNull()
	}
	if value := res.Get("response.0.radioTypeCProperties.minPowerLevel"); value.Exists() && !data.RadioTypeCMinPowerLevel.IsNull() {
		data.RadioTypeCMinPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMinPowerLevel = types.Int64Null()
	}
	if value := res.Get("response.0.radioTypeCProperties.maxPowerLevel"); value.Exists() && !data.RadioTypeCMaxPowerLevel.IsNull() {
		data.RadioTypeCMaxPowerLevel = types.Int64Value(value.Int())
	} else {
		data.RadioTypeCMaxPowerLevel = types.Int64Null()
	}
}

//template:end updateFromBody

//template:begin isNull
func (data *WirelessRFProfile) isNull(ctx context.Context, res gjson.Result) bool {
	if !data.Name.IsNull() {
		return false
	}
	if !data.DefaultRfProfile.IsNull() {
		return false
	}
	if !data.EnableRadioTypeA.IsNull() {
		return false
	}
	if !data.EnableRadioTypeB.IsNull() {
		return false
	}
	if !data.EnableRadioTypeC.IsNull() {
		return false
	}
	if !data.ChannelWidth.IsNull() {
		return false
	}
	if !data.EnableCustom.IsNull() {
		return false
	}
	if !data.EnableBrownField.IsNull() {
		return false
	}
	if !data.RadioTypeAParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeARadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeADataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeAPowerTresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeARxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeAMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeAMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeBRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeBDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeBPowerTresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeBRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeBMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeBMaxPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCParentProfile.IsNull() {
		return false
	}
	if !data.RadioTypeCRadioChannels.IsNull() {
		return false
	}
	if !data.RadioTypeCDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCMandatoryDataRates.IsNull() {
		return false
	}
	if !data.RadioTypeCPowerTresholdV1.IsNull() {
		return false
	}
	if !data.RadioTypeCRxSopThreshold.IsNull() {
		return false
	}
	if !data.RadioTypeCMinPowerLevel.IsNull() {
		return false
	}
	if !data.RadioTypeCMaxPowerLevel.IsNull() {
		return false
	}
	return true
}

//template:end isNull
