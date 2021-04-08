package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/tree/main/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Hostname of the OpenAM server",
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OPENAM_HOST", nil),
			},
			"token_id": &schema.Schema{
				Type:        schema.TypeString,
				Description: "Token identifier with administrative privileges",
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("OPENAM_ADMIN_TOKEN", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"openam_realm": resourceOpenamRealm(),
		},
	}
}

