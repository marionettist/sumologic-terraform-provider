package sumologic

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccSumologicCollector_simple(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicCollectorConfigSimple,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectorExists("sumologic_collector.test", t),
					resource.TestCheckResourceAttrSet("sumologic_collector.test", "id"),
				),
			},
		},
	})
}

func TestAccSumologicCollector_all(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicCollectorConfigAll,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectorExists("sumologic_collector.test", t),
					resource.TestCheckResourceAttrSet("sumologic_collector.test", "id"),
				),
			},
		},
	})
}

func TestAccSumologicCollector_change(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccSumologicCollectorConfigSimple,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectorExists("sumologic_collector.test", t),
					resource.TestCheckResourceAttr("sumologic_collector.test", "name", "MyCollector"),
					resource.TestCheckResourceAttr("sumologic_collector.test", "description", ""),
					resource.TestCheckResourceAttr("sumologic_collector.test", "category", ""),
					resource.TestCheckResourceAttr("sumologic_collector.test", "timezone", ""),
				),
			},
			{
				Config: testAccSumologicCollectorConfigAll,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckCollectorExists("sumologic_collector.test", t),
					resource.TestCheckResourceAttr("sumologic_collector.test", "name", "MyCollectorName"),
					resource.TestCheckResourceAttr("sumologic_collector.test", "description", "CollectorDesc"),
					resource.TestCheckResourceAttr("sumologic_collector.test", "category", "Category"),
					resource.TestCheckResourceAttr("sumologic_collector.test", "timezone", "Europe/Berlin"),
				),
			},
		},
	})
}

func testAccCheckCollectorExists(n string, t *testing.T) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		return nil
	}
}

var testAccSumologicCollectorConfigSimple = `

resource "sumologic_collector" "test" {
  name = "MyCollector"
  # description="CollectorDesc"
  category="Category"
  timezone = "Europe/Berlin"
}
`

var testAccSumologicCollectorConfigAll = `
resource "sumologic_collector" "test" {
  name="MyCollectorName"
  description="CollectorDesc"
  category="Category"
  timezone="Europe/Berlin"
}
`
