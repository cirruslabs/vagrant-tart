package main

import (
	"github.com/cirruslabs/vagrant-tart/internal/provider"
	"os"

	sdk "github.com/hashicorp/vagrant-plugin-sdk"
)

var ComponentOptions = []sdk.Option{
	sdk.WithComponents(
		&provider.TartProvider{},
	),
	sdk.WithName("tart"),
}

func main() {
	sdk.Main(ComponentOptions...)
	os.Exit(0)
}
