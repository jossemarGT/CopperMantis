package main

import (
	_ "github.com/CopperMantis/CopperMantis/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genapp "github.com/goadesign/goa/goagen/gen_app"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	genmain "github.com/goadesign/goa/goagen/gen_main"
	"github.com/goadesign/goa/goagen/gen_schema"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
)

func main() {
	codegen.ParseDSL()
	codegen.Run(
		// Generate main.go and controllers
		&genmain.Generator{
			API:    design.Design,
			Target: "app",
			Force:  false,
		},
		// Generate all the wiring for the http server
		&genapp.Generator{
			API:    design.Design,
			OutDir: "app",
			Target: "app",
			NoTest: false,
		},
		// Generate CLI tool
		&genclient.Generator{
			API: design.Design,
		},
		// Generate API Schema or "pact"
		// Read more at: https://blog.heroku.com/json_schema_for_heroku_platform_api
		&genschema.Generator{
			API: design.Design,
		},
		// Generate swagger (another kind of "pact")
		&genswagger.Generator{
			API: design.Design,
		},
	)
}
