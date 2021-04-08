package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/tree/main/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/tree/main/plugin"
  "terraform-provider-openam/provider"
)

func main() {
  plugin.Serve(&plugin.ServeOpts{
    ProviderFunc: func() *schema.Provider {
      return provider.Provider()
    },
  })
}