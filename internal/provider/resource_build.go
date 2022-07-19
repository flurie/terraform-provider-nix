package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBuild() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample resource in the Terraform provider scaffolding.",

		CreateContext: resourceBuildCreate,
		ReadContext:   resourceBuildRead,
		UpdateContext: resourceBuildUpdate,
		DeleteContext: resourceBuildDelete,

		Schema: map[string]*schema.Schema{
			"flake": {
				Description: "Name of flake to build. Defaults to '.'",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     ".",
			},
			"target": {
				Description: "Name of target to build. Defaults to flake's default target.",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"ouputSelection": {
				Description: "Subset of outputs to build.",
				Type:        schema.TypeSet,
				Optional:    true,
				Default:     []string{},
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"impure": {
				Description: "Name of target to build. Defaults to flake's default target.",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
			},
			"extra_attributes": {
				Description: "Extra attributes to pass to the nix builder.",
				Type:        schema.TypeString,
				Optional:    true,
			},
		},
	}
}

func resourceBuildCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*NixClient)

	idFromAPI := "my-id"
	d.SetId(idFromAPI)

	// write logs using the tflog package
	// see https://pkg.go.dev/github.com/hashicorp/terraform-plugin-log/tflog
	// for more information
	tflog.Trace(ctx, "created a resource")

	return diag.Errorf("not implemented")
}

func resourceBuildRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceBuildUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceBuildDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}
