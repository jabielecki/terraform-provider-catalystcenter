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
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

//template:end imports

//template:begin testAcc
func TestAccCcImageFromFile(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_from_file.test", "third_party_application_type", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_from_file.test", "family", ""))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_from_file.test", "name", "basename(\"../software.bin\")"))
	checks = append(checks, resource.TestCheckResourceAttr("catalystcenter_image_from_file.test", "third_party_vendor", ""))

	var steps []resource.TestStep
	if os.Getenv("SKIP_MINIMUM_TEST") == "" {
		steps = append(steps, resource.TestStep{
			Config: testAccCcImageFromFileConfig_minimum(),
		})
	}
	steps = append(steps, resource.TestStep{
		Config: testAccCcImageFromFileConfig_all(),
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
//template:end testPrerequisites

//template:begin testAccConfigMinimal
func testAccCcImageFromFileConfig_minimum() string {
	config := `resource "catalystcenter_image_from_file" "test" {` + "\n"
	config += `	source_path = "../software.bin"` + "\n"
	config += `	name = "basename(\"../software.bin\")"` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigMinimal

//template:begin testAccConfigAll
func testAccCcImageFromFileConfig_all() string {
	config := `resource "catalystcenter_image_from_file" "test" {` + "\n"
	config += `	third_party_application_type = ""` + "\n"
	config += `	family = ""` + "\n"
	config += `	source_path = "../software.bin"` + "\n"
	config += `	name = "basename(\"../software.bin\")"` + "\n"
	config += `	third_party_vendor = ""` + "\n"
	config += `	is_third_party = ` + "\n"
	config += `}` + "\n"
	return config
}

//template:end testAccConfigAll
