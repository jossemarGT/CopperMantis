package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// DocumentType Document object
var DocumentType = Type("DocumentPayload", func() {
	Description("Document representation")
	Attribute("title", String, "Document title", func() {
		MinLength(5)
		MaxLength(140) // Like twitter does
		Example("The ultimate programming algorithm")
	})
	Attribute("status", String, "Document status", func() {
		MinLength(3)
		MaxLength(15)
		Example("Draft")
		Example("Published")
	})
	Attribute("excerpt", String, "Brief description of the document", func() {
		MinLength(5)
		MaxLength(140)
		Example("Lorem ipsum dolorem, foobar?")
	})
	Attribute("content", String, "Document content in markdown notation", func() {
		MinLength(5)
		MaxLength(2048)
		Example("Lorem ipsum dolorem, foobar?")
	})
})

var DocumentMediaType = MediaType("application/vnd.cms.document+json", func() {
	Description("Document representation")
	Reference(DocumentType)
	Attributes(func() {
		Attribute("id", Integer, "Document ID", func() {
			Example(1)
		})
		Attribute("href", String, "Document href", func() {
			Example("/document/1")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		Attribute("deleted_at", DateTime, "Date of soft delete")
		// References
		// Attribute("term", TermType, "Labels and tags")

		// Attributes below inherit from the base type
		Attribute("title")
		Attribute("status")
		Attribute("excerpt")
		Attribute("content")

		Required("id", "href", "title")
	})
	//
	// Links(func() {
	// 	Link("term")
	// })

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
		Attribute("excerpt")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
		Attribute("status")
		Attribute("excerpt")
		Attribute("content")
		// Attribute("term")
	})
})
