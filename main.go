package main

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"terraform-provider-transfer.sh/transfer"
)

func main() {
	//goland:noinspection GoUnhandledErrorResult
	providerserver.Serve(context.Background(), transfer.New, providerserver.ServeOpts{
		Address: "hashicorp.com/edu/transfer.sh",
	})
}
