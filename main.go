package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/marionettist/sumologic-terraform-provider/sumologic"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: sumologic.Provider,
	})
}
