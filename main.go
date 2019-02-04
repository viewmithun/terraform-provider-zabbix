package main

import (
	"github.com/viewmithun/terraform-provider-zabbix/provider"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	p := plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	}

	plugin.Serve(&p)
}
