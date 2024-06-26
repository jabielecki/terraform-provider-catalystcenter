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

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccCcImageActivationConfig_minimum() string {
	config := `resource "catalystcenter_image_activation" "test" {` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccCcImageActivationConfig_all() string {
	config := `resource "catalystcenter_image_activation" "test" {` + "\n"
	config += `	device_uuid = "138b3181-f9c5-4271-9292-cf3152ab4d3e"` + "\n"
	config += `	image_uuid_list = [""]` + "\n"
	config += `	activate_lower_image_version = true` + "\n"
	config += `	device_upgrade_mode = ""` + "\n"
	config += `	distribute_if_needed = true` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
