package main

import (
	_ "github.com/CopperMantis/CopperMantis/design"

	"github.com/goadesign/goa/design"
	"github.com/goadesign/goa/goagen/codegen"
	genapp "github.com/goadesign/goa/goagen/gen_app"
	genclient "github.com/goadesign/goa/goagen/gen_client"
	genschema "github.com/goadesign/goa/goagen/gen_schema"
	genswagger "github.com/goadesign/goa/goagen/gen_swagger"
	meta "github.com/goadesign/goa/goagen/meta"
)

func main() {
	// Get gorma generator
	genmeta, err := meta.NewGenerator(
		"gorma.Generate",
		[]*codegen.ImportSpec{codegen.SimpleImport("github.com/goadesign/gorma")},
		map[string]string{"design": "github.com/CopperMantis/CopperMantis/design", "out": ".", "pkg": "models"},
		[]string{},
	)

	if err != nil {
		panic(err)
	}

	codegen.ParseDSL()
	codegen.Run(
		// TODO: Generate main.go and controllers
		//  If needed please perform a  `goagen boostrap` instead
		// Generate all the wiring for the http server
		genapp.NewGenerator(
			genapp.API(design.Design),
			genapp.OutDir("app"),
			genapp.Target("app"),
			genapp.NoTest(false),
		),
		// Generate ER database gorm mapping
		genmeta,
		// Generate CLI tool
		genclient.NewGenerator(
			genclient.API(design.Design),
		),
		// Generate API Schema or "pact"
		// Read more at: https://blog.heroku.com/json_schema_for_heroku_platform_api
		genswagger.NewGenerator(
			genswagger.API(design.Design),
		),
		// Generate swagger (another more popular "pact")
		genschema.NewGenerator(
			genschema.API(design.Design),
		),
	)
}
