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

//template:begin testAcc
func TestAccCcVirtualNetworkToFabricSite(t *testing.T) {
	var checks []resource.TestCheckFunc

	var steps []resource.TestStep
	steps = append(steps, resource.TestStep{
		Config: testAccCcVirtualNetworkToFabricSitePrerequisitesConfig + testAccCcVirtualNetworkToFabricSiteConfig_all(),
		Check:  resource.ComposeTestCheckFunc(checks...),
	})

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps:                    steps,
	})
}

//template:end testAcc

//template:begin testPrerequisites
const testAccCcVirtualNetworkToFabricSitePrerequisitesConfig = `
resource "catalystcenter_area" "test" {
  name        = "Area1"
  parent_name = "Global"
}
resource "catalystcenter_fabric_site" "test" {
  site_name_hierarchy = "Global/Area1"
  fabric_type         = "FABRIC_SITE"

  depends_on = [catalystcenter_area.test]
}
resource "catalystcenter_fabric_virtual_network" "test" {
  virtual_network_name = "SDA_VN1"
  is_guest             = false
  sg_names             = ["Employees"]
}

`

//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccCcVirtualNetworkToFabricSiteConfig_minimum() string {
	config := `resource "catalystcenter_virtual_network_to_fabric_site" "test" {` + "\n"
	config += `	virtual_network_name = catalystcenter_fabric_virtual_network.test.virtual_network_name` + "\n"
	config += `	site_name_hierarchy = catalystcenter_fabric_site.test.site_name_hierarchy` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccCcVirtualNetworkToFabricSiteConfig_all() string {
	config := `resource "catalystcenter_virtual_network_to_fabric_site" "test" {` + "\n"
	config += `	virtual_network_name = catalystcenter_fabric_virtual_network.test.virtual_network_name` + "\n"
	config += `	site_name_hierarchy = catalystcenter_fabric_site.test.site_name_hierarchy` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
