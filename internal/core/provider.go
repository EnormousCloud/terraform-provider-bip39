package core

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// providerConfig embeds internal terraform provider configuration
type providerConfig struct {
}

// Provider creates the Docker provider
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{}, // no parameters for provider
		ResourcesMap: map[string]*schema.Resource{
			"bip39_entropy": resourceBip39Entropy(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &providerConfig{}, nil
}
