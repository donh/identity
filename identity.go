package main

import (
	"os"

	"github.com/donh/identity/api"
	"github.com/donh/identity/build"
	"github.com/donh/identity/test"
)

func main() {
	action := ""
	if len(os.Args) > 1 {
		action = os.Args[1]
	}
	switch action {
	case "build":
		build.Build()
	case "deploy":
		ledger := os.Args[2]
		build.Deploy(ledger)
	case "test":
		test.Test()
	default:
		api.SetAPIRoutes()
	}
}
