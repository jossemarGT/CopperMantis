package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("document", func() {
	BasePath("/document")

	Action("list", func() {
		Description("List documents with public data")
		Routing(GET(""))
		Params(func() {
			Param("page", Integer, "Page number", func() {
				Default(1)
			})
			Param("sort", String, func() {
				Enum("asc", "desc")
				Default("desc")
			})
			Param("kind", String, func() {
				Enum("page", "problem", "any")
				Default("any")
			})
		})
		// TODO: Extend CollectionOf to API Guidelines
		Response(OK, CollectionOf(DocumentMediaType))
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("show", func() {
		Description("Get specific document with public data")
		Routing(GET("/:documentID"))
		Params(func() {
			Param("documentID", Integer, "Document number ID")
		})
		Response(OK, DocumentMediaType)
		Response(BadRequest, ErrorMedia)
		Response(NotFound)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new document")
		Payload(DocumentType)
		Response(Created, "/document/[0-9]+")
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			// FIXME: PUT and PATCH have the same behavior
			PUT("/:documentID"),
			PATCH("/:documentID"),
		)
		Description("Overwrite document data")
		Params(func() {
			Param("documentID", Integer, "Document ID")
		})
		Payload(DocumentType)
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("delete", func() {
		Routing(
			DELETE("/:documentID"),
		)
		Description("Soft delete document data")
		Params(func() {
			Param("documentID", Integer, "Document ID")
		})
		Response(NoContent)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})
