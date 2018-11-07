package sumologic

import "github.com/hashicorp/terraform/helper/schema"

func timeZoneDiffSuppressFunc(k, old, new string, d *schema.ResourceData) bool {
	if old == "" && (new == "Etc/UTC" || new == "UTC") ||
	old == "Etc/UTC" && (new == "" || new == "UTC") ||
	old == "UTC" && (new == "Etc/UTC" || new == "") {
		return true
	}
	return false
}
