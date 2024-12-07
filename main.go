package main

import (
	"context"

	"github.com/gionnid/terraform-provider-notion/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	pageRepo := notion.NewHTTPApiRepository()

	providerserver.Serve(context.Background(), provider.New(pageRepo), providerserver.ServeOpts{
		Address: "registry.terraform.io/gionnid/notion",
	})
}
