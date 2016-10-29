package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("CopperMantis", func() {
	Title("Programming Contest Management System Rest API")
	Description("Public API for CopperMantis ecosystem")
	BasePath("/cms/v1")
})
