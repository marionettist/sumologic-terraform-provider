package sumologic

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccDataSourcSumologicCollector(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testDataSourceAccSumologicCollectorConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccDataSourceCollectorCheck("data.sumologic_collector.by_name"),
					testAccDataSourceCollectorCheck("data.sumologic_collector.by_id"),
					resource.TestCheckResourceAttrSet("data.sumologic_collector.by_name", "id"),
					resource.TestCheckResourceAttrSet("data.sumologic_collector.by_id", "id"),
				),
			},
		},
	})
}

func testAccDataSourceCollectorCheck(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]
		if !ok {
			return fmt.Errorf("root module has no resource called %s", name)
		}

		collectorRs, ok := s.RootModule().Resources["sumologic_collector.test"]
		if !ok {
			return fmt.Errorf("can't find sumologic_collector.test in state")
		}

		attr := rs.Primary.Attributes

		actualName := attr["name"]
		expectedName := collectorRs.Primary.Attributes["name"]
		if actualName != expectedName {
			return fmt.Errorf(
				"name is %s; want %s",
				actualName,
				expectedName,
			)
		}

		actualId := attr["id"]
		expectedId := collectorRs.Primary.Attributes["id"]
		if actualId != expectedId {
			return fmt.Errorf(
				"id is %s; want %s",
				actualId,
				expectedId,
			)
		}

		actualDescription := attr["description"]
		expectedDescription := collectorRs.Primary.Attributes["description"]
		if actualDescription != expectedDescription {
			return fmt.Errorf(
				"description is %s; want %s",
				actualDescription,
				expectedDescription,
			)
		}

		actualCategory := attr["category"]
		expectedCategory := collectorRs.Primary.Attributes["category"]
		if actualCategory != expectedCategory {
			return fmt.Errorf(
				"category is %s; want %s",
				actualCategory,
				expectedCategory,
			)
		}

		actualTimezone := attr["timezone"]
		expectedTimezone := collectorRs.Primary.Attributes["timezone"]
		if actualTimezone != expectedTimezone {
			return fmt.Errorf(
				"timezone is %s; want %s",
				actualTimezone,
				expectedTimezone,
			)
		}

		return nil
	}
}

var testDataSourceAccSumologicCollectorConfig = `
resource "sumologic_collector" "test" {
  name = "MyCollector"
  description = "MyCollectorDesc"
  category = "Cat"
  timezone = "Europe/Berlin"
}

data "sumologic_collector" "by_name" {
  name = "${sumologic_collector.test.name}"
}

data "sumologic_collector" "by_id" {
  id = "${sumologic_collector.test.id}"
}
`
