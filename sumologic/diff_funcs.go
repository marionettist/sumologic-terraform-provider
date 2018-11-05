package sumologic

import "github.com/hashicorp/terraform/helper/schema"

func ignoreDiffs(k, old, new string, d *schema.ResourceData) bool {return true}
