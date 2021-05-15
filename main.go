package main

import (
	"math/rand"
	"time"

	"provider/internal/core"

	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: core.Provider,
	})
}
