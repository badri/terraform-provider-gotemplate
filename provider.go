package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"gotemplate": dataSourceFile(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"gotemplate": schema.DataSourceResourceShim(
				"gotemplate",
				dataSourceFile(),
			),
		},
	}
}
