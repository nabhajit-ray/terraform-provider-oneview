provider "oneview" {
	ov_username = "<ov_username>"
	ov_password = "<ov_password"
	ov_endpoint = "<ov_endpoint>"
	ov_sslverify = false
	ov_apiversion = <ov_apiversion>
	ov_ifmatch = "*"
}

resource "oneview_appliancetime_and_local" "Appliance Time and Local" {
	Locale = "en_US.UTF-8"
	DateTime = "2014-09-11T12:10:33"
	Timezone = "UTC"
    NtpServers = [ "16.110.135.123" ]
}

/* Testing data source
data "oneview_appliancetime_and_local" "Appliance Time and Local" {
        name = "SYN03_Frame1"
}

output "oneview_appliancetime_and_local_value" {
        value = "${data.oneview_appliancetime_and_local.Appliance Time and Local.time_zone}"
}
*/