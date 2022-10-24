package swim

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMapDownlink() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceMapDownlinkRead,
		CreateContext: resourceMapDownlinkCreate,
		UpdateContext: resourceMapDownlinkUpdate,
		DeleteContext: resourceMapDownlinkDelete,
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
			"items": &schema.Schema{
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Computed: true,
				},
			},
		},
	}
}

func resourceMapDownlinkRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)

	items, diags := client.GetMapDownlink(node, lane)
	if diags != nil {
		return diags
	} else {
		data.Set("items", items)
		var diags diag.Diagnostics
		return diags
	}
}

func resourceMapDownlinkCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)
	items, setItems := data.GetOk("items")

	if setItems {
		diags := client.SetMapDownlink(node, lane, items.(map[string]interface{}))
		if diags == nil {
			data.SetId(node + "/" + lane)
		}
		return diags
	} else {
		items, diags := client.GetMapDownlink(node, lane)
		if diags == nil {
			data.SetId(node + "/" + lane)
			data.Set("items", items)
		}
		return diags

	}
}

func resourceMapDownlinkUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)
	items := data.Get("items").(map[string]interface{})

	diags := client.SetMapDownlink(node, lane, items)

	return diags
}

func resourceMapDownlinkDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(SwimClient)

	node := data.Get("node").(string)
	lane := data.Get("lane").(string)

	diags := client.ClearMapDownlink(node, lane)

	return diags
}
