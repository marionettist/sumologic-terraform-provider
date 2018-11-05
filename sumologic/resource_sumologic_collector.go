package sumologic

import (
	"log"
	"strconv"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSumologicCollector() *schema.Resource {
	//noDiffFunc := func(k, old, new string, d *schema.ResourceData) bool {return true}
	return &schema.Resource{
		Create: resourceSumologicCollectorCreate,
		Read:   resourceSumologicCollectorRead,
		Delete: resourceSumologicCollectorDelete,
		Update: resourceSumologicCollectorUpdate,

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
				Default:  "Etc/UTC",
			},
			"lookup_by_name": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				//Computed: true,
				Default:  false,
				DiffSuppressFunc: ignoreDiffs,
			},
			"foo": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				Default:  true,
				DiffSuppressFunc: func (k, old, new string, d *schema.ResourceData) bool {return true},
			},
			"destroy": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: false,
				//Computed: true,
				Default:  true,
				DiffSuppressFunc: ignoreDiffs,
			},
			//"foo": {
			//	Type:     schema.TypeBool,
			//	Optional: true,
			//	ForceNew: false,
			//	Default:  true,
			//},
			//"foo2": {
			//	Type:     schema.TypeBool,
			//	Optional: true,
			//	ForceNew: false,
			//	Computed: true,
			//	//DiffSuppressFunc: ignoreDiffs,
			//},
			//"foo3": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	ForceNew: false,
			//	Default:  "foo3",
			//},
			//"foo4": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	ForceNew: false,
			//	Default:  "foo4",
			//	DiffSuppressFunc: ignoreDiffs,
			//},
		},
	}
}

func resourceSumologicCollectorRead(d *schema.ResourceData, meta interface{}) error {
	log.Printf("*********** destroy at start of resourceSumologicCollectorRead: %v", d.Get("destroy").(bool))
	c := meta.(*Client)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	category := d.Get("category").(string)
	timeZone := d.Get("timezone").(string)
	lookUpByName := d.Get("lookup_by_name")
	destroy := d.Get("destroy")
	//foo := d.Get("foo")
	//foo2 := d.Get("foo2")
	//foo3 := d.Get("foo3")
	//foo4 := d.Get("foo4")
	log.Println(name)
	log.Println(description)
	log.Println(category)
	log.Println(timeZone)
	log.Println(lookUpByName)
	log.Println(destroy)
	//log.Println(foo)
	//log.Println(foo2)
	//log.Println(foo3)
	//log.Println(foo4)

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
	//d.Set("destroy", collector.Destroy)
	//d.Set("lookup_by_name", collector.LookupByName)

	log.Printf("*********** destroy at end of resourceSumologicCollectorRead: %v", d.Get("destroy").(bool))
	return nil
}

func resourceSumologicCollectorDelete(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*Client)

	if d.Get("destroy").(bool) {
		id, _ := strconv.Atoi(d.Id())
		log.Println("[DEBUG] sumologic deleting collector")
		return c.DeleteCollector(id)
	}

	log.Println("[DEBUG] sumologic not deleting collector (destroy set to false)")
	return nil
}

func resourceSumologicCollectorCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	description := d.Get("description").(string)
	category := d.Get("category").(string)
	timeZone := d.Get("timezone").(string)
	lookUpByName := d.Get("lookup_by_name")
	destroy := d.Get("destroy")
	log.Printf("*********** destroy at start of resourceSumologicCollectorCreate: %v", d.Get("destroy").(bool))
	//foo := d.Get("foo")
	//foo2 := d.Get("foo2")
	//foo3 := d.Get("foo3")
	//foo4 := d.Get("foo4")
	log.Println(name)
	log.Println(description)
	log.Println(category)
	log.Println(timeZone)
	log.Println(lookUpByName)
	log.Println(destroy)
	//log.Println(foo)
	//log.Println(foo2)
	//log.Println(foo3)
	//log.Println(foo4)

	c := meta.(*Client)

	if d.Get("lookup_by_name").(bool) {
		collector, err := c.GetCollectorName(d.Get("name").(string))

		if err != nil {
			return err
		}

		if collector != nil {
			d.SetId(strconv.Itoa(collector.ID))
		}
	}

	if d.Id() == "" {
		id, err := c.CreateCollector(Collector{
			CollectorType: "Hosted",
			Name:          d.Get("name").(string),
		})

		if err != nil {
			return err
		}

		d.SetId(strconv.Itoa(id))
	}

	retVal := resourceSumologicCollectorUpdate(d, meta)
	log.Printf("*********** destroy at end of resourceSumologicCollectorCreate: %v", d.Get("destroy").(bool))
	return retVal
}

func resourceSumologicCollectorUpdate(d *schema.ResourceData, meta interface{}) error {
	log.Printf("*********** destroy at start of resourceSumologicCollectorUpdate: %v", d.Get("destroy").(bool))

	collector := resourceToCollector(d)

	c := meta.(*Client)
	err := c.UpdateCollector(collector)

	if err != nil {
		return err
	}

	retVal := resourceSumologicCollectorRead(d, meta)
	log.Printf("*********** destroy at end of resourceSumologicCollectorUpdate: %v", d.Get("destroy").(bool))
	return retVal
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
		//Destroy:       d.Get("destroy").(bool),
		//LookupByName:  d.Get("lookup_by_name").(bool),
	}
}
