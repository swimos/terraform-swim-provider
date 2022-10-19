package swim

import (
	"context"
	"strconv"
	"time"

	ws "github.com/sacOO7/gowebsocket"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEndpoint() *schema.Resource {
	return &schema.Resource{
		//Todo implement all
		ReadContext:   resourceEndpointRead,
		CreateContext: resourceEndpointUpdate,
		UpdateContext: resourceEndpointUpdate,
		DeleteContext: resourceEndpointDelete,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"lane": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceEndpointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	var diags diag.Diagnostics
	return diags
}

func resourceEndpointUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(ws.Socket)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	node := d.Get("node").(string)
	lane := d.Get("lane").(string)
	value := d.Get("value").(string)

	client.SendText("@command(node: \"" + node + "\", lane: \"" + lane + "\")\"" + value + "\"")

	var diags diag.Diagnostics
	return diags
}

func resourceEndpointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(ws.Socket)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	node := d.Get("node").(string)
	lane := d.Get("lane").(string)

	client.SendText("@command(node: \"" + node + "\", lane: \"" + lane + "\")")

	var diags diag.Diagnostics
	return diags
}
