package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("CopperMantis", func() {
	Title("Programming Contest Management System Rest API")
	Description("CopperMantis ecosystem public API")
	BasePath("/cms/v1")
	License(func() {
		Name("Apache-2.0")
		URL("https://github.com/CopperMantis/CopperMantis/blob/master/LICENSE")
	})
	Docs(func() {
		Description("CopperMantis Docs")
		URL("https://github.com/CopperMantis/CopperMantisDocs")
	})

	// Enhance Created response template with location header
	ResponseTemplate(Created, func(pattern string) {
		Description("Resource created")
		Status(201)
		Headers(func() {
			Header("Location", String, "href to created resource", func() {
				Pattern(pattern)
			})
		})
	})
})
