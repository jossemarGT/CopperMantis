package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("profile", func() {
	BasePath("/profile")

	Action("list", func() {
		Description("List user profiles with public data")
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
		Response(OK, CollectionOf(ProfileMediaType))
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Get specific profile with public data")
		Routing(GET("/:profileID"))
		Params(func() {
			Param("profileID", Integer, "Profile number ID")
		})
		Response(OK, ProfileMediaType)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new user profile")
		Payload(ProfileType)
		Response(Created, "/profile/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			// FIXME: PUT and PATCH have the same behavior
			PUT("/:profileID"),
			PATCH("/:profileID"),
		)
		Description("Overwrite profile data")
		Params(func() {
			Param("profileID", Integer, "Profile ID")
		})
		Payload(ProfileType)
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:profileID"),
		)
		Description("Soft delete profile data")
		Params(func() {
			Param("profileID", Integer, "Profile ID")
		})
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})
