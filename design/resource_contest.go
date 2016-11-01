package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("contest", func() {
	BasePath("/contest")

	Action("list", func() {
		Description("List contests with public data")
		Routing(GET(""))
		Params(func() {
			Param("page", Integer, "Page number", func() {
				Default(1)
			})
			Param("sort", String, func() {
				Enum("asc", "desc")
				Default("desc")
			})
		})
		// TODO: Extend CollectionOf to API Guidelines
		Response(OK, CollectionOf(ContestMediaType))
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Get specific contest with public data")
		Routing(GET("/:contestID"))
		Params(func() {
			Param("contestID", Integer, "Profile number ID")
		})
		Response(OK, ContestMediaType)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new contest")
		Payload(ContestType)
		Response(Created, "/contest/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			// FIXME: PUT and PATCH have the same behavior
			PUT("/:contestID"),
			PATCH("/:contestID"),
		)
		Description("Overwrite contest data")
		Params(func() {
			Param("contestID", Integer, "Profile ID")
		})
		Payload(ContestType)
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:contestID"),
		)
		Description("Soft delete contest data")
		Params(func() {
			Param("contestID", Integer, "Profile ID")
		})
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})
