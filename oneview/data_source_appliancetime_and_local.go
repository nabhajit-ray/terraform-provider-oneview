// (C) Copyright 2020 Hewlett Packard Enterprise Development LP
//
// Licensed under the Apache License, Version 2.0 (the "License");
// You may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed
// under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
// CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package oneview

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceApplianceTimeandLocal() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplianceTimeandLocalRead,

		Schema: map[string]*schema.Schema{
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"eTag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"modified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locale_display_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"polling_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"locale": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_servers": {
				Computed: true,
				Type:     schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
		},
	}
}

func dataSourceApplianceTimeandLocalRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	applianceTimeandLocals, err := config.ovClient.GetApplianceTimeandLocals("", "", "", "")
	if err != nil || applianceTimeandLocals.IsNil() {
		return nil
	}
	d.Set("type", applianceTimeandLocals.Type)
	d.Set("category", applianceTimeandLocals.Category)
	d.Set("uri", applianceTimeandLocals.URI.String())
	d.Set("eTag", applianceTimeandLocals.ETAG)
	d.Set("created", applianceTimeandLocals.Created)
	d.Set("modified", applianceTimeandLocals.Modified)
	d.Set("locale_display_time", applianceTimeandLocals.LocaleDisplayName.String())
	d.Set("polling_interval", applianceTimeandLocals.PollingInterval)
	d.Set("locale", applianceTimeandLocals.Locale)
	d.Set("date_time", applianceTimeandLocals.DateTime.String())
	d.Set("time_zone", applianceTimeandLocals.TimeZone.String())
	d.Set("ntp_servers", applianceTimeandLocals.NtpServers)
	return nil
}