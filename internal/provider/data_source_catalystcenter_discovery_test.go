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
func TestAccDataSourceCcDiscovery(t *testing.T) {
	var checks []resource.TestCheckFunc
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "cdp_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "discovery_type", "RANGE"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "enable_password_list.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "global_credential_id_list.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "http_read_credential", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "http_write_credential", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "ip_address_list", "192.168.0.1-192.168.0.99"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "ip_filter_list.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "lldp_level", "10"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "name", "disco42"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "netconf_port", "830"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "password_list.0", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "preferred_ip_method", "UseLoopBack"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "protocol_order", "ssh"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "retry", "3"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_auth_passphrase", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_auth_protocol", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_mode", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_priv_passphrase", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_priv_protocol", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_ro_community", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_ro_community_desc", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_rw_community", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_rw_community_desc", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_user_name", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "snmp_version", ""))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "timeout", "5"))
	checks = append(checks, resource.TestCheckResourceAttr("data.catalystcenter_discovery.test", "user_name_list.0", ""))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceCcDiscoveryConfig(),
				Check:  resource.ComposeTestCheckFunc(checks...),
			},
		},
	})
}

//template:end testAccDataSource

//template:begin testPrerequisites
//template:end testPrerequisites

//template:begin testAccDataSourceConfig
func testAccDataSourceCcDiscoveryConfig() string {
	config := `resource "catalystcenter_discovery" "test" {` + "\n"
	config += `	cdp_level = 10` + "\n"
	config += `	discovery_type = "RANGE"` + "\n"
	config += `	enable_password_list = [""]` + "\n"
	config += `	global_credential_id_list = [""]` + "\n"
	config += `	http_read_credential = ""` + "\n"
	config += `	http_write_credential = ""` + "\n"
	config += `	ip_address_list = "192.168.0.1-192.168.0.99"` + "\n"
	config += `	ip_filter_list = [""]` + "\n"
	config += `	lldp_level = 10` + "\n"
	config += `	name = "disco42"` + "\n"
	config += `	netconf_port = "830"` + "\n"
	config += `	password_list = [""]` + "\n"
	config += `	preferred_ip_method = "UseLoopBack"` + "\n"
	config += `	protocol_order = "ssh"` + "\n"
	config += `	retry = 3` + "\n"
	config += `	snmp_auth_passphrase = ""` + "\n"
	config += `	snmp_auth_protocol = ""` + "\n"
	config += `	snmp_mode = ""` + "\n"
	config += `	snmp_priv_passphrase = ""` + "\n"
	config += `	snmp_priv_protocol = ""` + "\n"
	config += `	snmp_ro_community = ""` + "\n"
	config += `	snmp_ro_community_desc = ""` + "\n"
	config += `	snmp_rw_community = ""` + "\n"
	config += `	snmp_rw_community_desc = ""` + "\n"
	config += `	snmp_user_name = ""` + "\n"
	config += `	snmp_version = ""` + "\n"
	config += `	timeout = 5` + "\n"
	config += `	user_name_list = [""]` + "\n"
	config += `}` + "\n"

	config += `
		data "catalystcenter_discovery" "test" {
			id = catalystcenter_discovery.test.id
		}
	`
	return config
}

//template:end testAccDataSourceConfig
