package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("monitoring", func() {
	// Using API BasePath
	Action("ping", func() {
		Description("Perform health check.")
		Routing(GET("/_ping"))
		Response(OK, "text/plain")
	})
})
