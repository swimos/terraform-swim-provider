package swim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceValueDownlink() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceValueDownlinkRead,
		CreateContext: resourceValueDownlinkCreate,
		UpdateContext: resourceValueDownlinkUpdate,
		DeleteContext: resourceValueDownlinkDelete,
		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"lane": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceValueDownlinkRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)

	value, diags := client.GetValueDownlink(node, lane)
	if diags != nil {
		return diags
	} else {
		data.Set("value", value)
		var diags diag.Diagnostics
		return diags
	}
}

func resourceValueDownlinkCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)
	value, setValue := data.GetOk("value")

	if setValue {
		diags := client.SetValueDownlink(node, lane, value.(string))
		if diags == nil {
			data.SetId(node + "/" + lane)
		}
		return diags
	} else {
		value, diags := client.GetValueDownlink(node, lane)
		if diags == nil {
			data.SetId(node + "/" + lane)
			data.Set("value", value)
		}
		return diags
	}
}

func resourceValueDownlinkUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)
	value := data.Get("value").(string)

	diags := client.SetValueDownlink(node, lane, value)

	return diags
}

func resourceValueDownlinkDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)

	diags := client.ClearValueDownlink(node, lane)

	return diags
}
