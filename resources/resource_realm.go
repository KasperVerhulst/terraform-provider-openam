package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceRealm() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRealmCreate,
		ReadContext:   resourceRealmRead,
		UpdateContext: resourceRealmUpdate,
		DeleteContext: resourceRealmDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"parent": {
				Type:     schema.TypeString,
				Required: true,
				Default:  "/",
			},
		},
	}
}

// API call to create openam realm - POST
func resourceRealmCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

// API call to read openam realm - GET
func resourceRealmRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

// API call to update openam realm - PUT
func resourceRealmUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}

// API call to delete openam realm - DELETE
func resourceRealmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diag diag.Diagnostics
	return diag
}
