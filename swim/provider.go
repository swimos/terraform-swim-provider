package swim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SWIM_URL", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"swim_value_downlink": resourceValueDownlink(),
			"swim_map_downlink":   resourceMapDownlink(),
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
