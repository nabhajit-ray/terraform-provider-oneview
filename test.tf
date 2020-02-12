provider "oneview" { 

  ov_username = "Administrator" 

  ov_password = "madhav123" 

  ov_endpoint = "https://10.20.7.10" 

  ov_sslverify = false 

  ov_apiversion = 1200 

  ov_ifmatch = "*" 

} 

  

// Get Server Hardware 

data "oneview_server_hardware" "server_hardware" { 

  name = "0000A66102, bay 5" 

} 

  

// Get Server Profile Template 

data "oneview_server_profile_template" "server_profile_template" { 

  name = " Test_spt " 

} 

  

// Create Server Profile from Server Profile Template 

resource "oneview_server_profile" "server_profile" { 

  name = "OV-Terraform-test" 

  template = "${data.oneview_server_profile_template.server_profile_template.name}" 

  hardware_name = "${data.oneview_server_hardware.server_hardware.name}" 

  type = "ServerProfileV10" 

} 
