package sumologic

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/customdiff"
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSumologicCollector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSumologicCollectorCreate,
		Read:   resourceSumologicCollectorRead,
		Delete: resourceSumologicCollectorDelete,
		Update: resourceSumologicCollectorUpdate,
		CustomizeDiff: resourceSumologicCollectorCustomizeDiff(),

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default:  "",
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: false,
				Default: "UTC",
				//DiffSuppressFunc: timeZoneDiffSuppressFunc,
			},
		},
	}
}

func resourceSumologicCollectorRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id, err := strconv.Atoi(d.Id())

	if err != nil {
		return err
	}

	collector, err := c.GetCollector(id)

	if err != nil {
		log.Printf("[WARN] Collector not found, removing from state: %v - %v", id, err)
		d.SetId("")

		return nil
	}

	d.Set("name", collector.Name)
	d.Set("description", collector.Description)
	d.Set("category", collector.Category)
	d.Set("timezone", collector.TimeZone)

	return nil
}

func resourceSumologicCollectorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return fmt.Errorf("expected collector id to be int; got %s; %s", d.Id(), err)
	}
	return c.DeleteCollector(id)
}

func resourceSumologicCollectorCreate(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	id, err := c.CreateCollector(Collector{
		CollectorType: "Hosted",
		Name:          d.Get("name").(string),
	})

	if err != nil {
		return err
	}

	d.SetId(strconv.Itoa(id))

	return resourceSumologicCollectorUpdate(d, meta)
}

func resourceSumologicCollectorUpdate(d *schema.ResourceData, meta interface{}) error {

	collector := resourceToCollector(d)

	c := meta.(*Client)
	err := c.UpdateCollector(collector)

	if err != nil {
		return err
	}

	return resourceSumologicCollectorRead(d, meta)
}

//func resourceSumologicCollectorCustomizeDiff(diff *schema.ResourceDiff, meta interface{}) error {
//	return nil
//}

func resourceSumologicCollectorCustomizeDiff() schema.CustomizeDiffFunc {
	return customdiff.All(
		customdiff.IfValueChange("description", condFunc, diffFunc),
		)
}

func condFunc(old, new, meta interface{}) bool {
	return false
}
func diffFunc(*schema.ResourceDiff, interface{}) error {
	return nil
}
func resourceToCollector(d *schema.ResourceData) Collector {
	id, _ := strconv.Atoi(d.Id())

	return Collector{
		ID:            id,
		CollectorType: "Hosted",
		Name:          d.Get("name").(string),
		Description:   d.Get("description").(string),
		Category:      d.Get("category").(string),
		TimeZone:      d.Get("timezone").(string),
	}
}
