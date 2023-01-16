package cause

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCause() *schema.Resource {
	return &schema.Resource{
		CreateContext: causeCreate,
		ReadContext:   causeRead,
		UpdateContext: causeUpdate,
		DeleteContext: causeDelete,

		Schema: map[string]*schema.Schema{
			"number": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func causeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	count := d.Get("number").(string)
	d.SetId(count)
	// https://www.uuidtools.com/api/generate/v1/count/uuid_count
	resp, err := http.Get("https://www.uuidtools.com/api/generate/v1/count/" + count)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("[ERROR] %s: causeCreate filed", err)
	}
	defer resp.Body.Close()
	log.Printf("[ERROR] %s: causeCreate successfully", err)
	fmt.Println(string(body))
	return diags
}

func causeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	return nil
}

func causeUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return causeRead(ctx, d, m)
}

func causeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}

