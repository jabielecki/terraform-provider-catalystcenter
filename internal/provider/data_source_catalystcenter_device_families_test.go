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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAccDataSource
func TestAccDataSourceCcDeviceFamilies(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_families.test", "response.0.device_family", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_device_families.test", "response.0.device_family_identifier", ""))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcDeviceFamiliesConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceCcDeviceFamiliesConfig() string {
	config := `resource "catalystcenter_device_families" "test" {` + "\n"
	config += `	response = [{` + "\n"
	config += `	  device_family = ""` + "\n"
	config += `	  device_family_identifier = ""` + "\n"
	config += `	}]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_device_families" "test" {
			id = catalystcenter_device_families.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig