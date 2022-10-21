package swim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Swim Provider
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SWIM_URL", nil),
			},
		},
		// Todo change this
		ResourcesMap: map[string]*schema.Resource{
			"swim_value_downlink": resourceValueDownlink(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := data.Get("url").(string)
	client := SwimClient{url}
	var diags diag.Diagnostics

	return client, diags
}
