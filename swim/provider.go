package swim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	ws "github.com/sacOO7/gowebsocket"
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
			"swim_endpoint": resourceEndpoint(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	url := d.Get("url").(string)

	//Todo add error check
	client := ws.New(url)
	client.Connect()

	var diags diag.Diagnostics

	return client, diags
}
