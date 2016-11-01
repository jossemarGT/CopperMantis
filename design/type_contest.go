package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ContestType Contest object
var ContestType = Type("ContestPayload", func() {
	Description("Contest representation")
	Attribute("title", String, "Contest title", func() {
		MinLength(3)
		MaxLength(140)
		Example("Extreme programming contest 2016")
	})
	Attribute("status", String, "Contest status", func() {
		MinLength(3)
		MaxLength(15)
		Example("scheduled")
		Example("finished")
		Example("ongoing")
	})
	Attribute("start_time", DateTime)
	Attribute("finish_time", DateTime)
})

var ContestMediaType = MediaType("application/vnd.cms.contest+json", func() {
	Description("Contest representation")
	Reference(ContestType)
	Attributes(func() {
		Attribute("id", Integer, "Contest ID", func() {
			Example(1)
		})
		Attribute("href", String, "Contest href", func() {
			Example("/contest/1")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		Attribute("deleted_at", DateTime, "Date of soft delete")
		// References
		// Attribute("user", ProfileType, "Users related to this contest")
		// Attribute("document", DocumentType, "Documents related to this contest")

		// Attributes below inherit from the base type
		Attribute("title")
		Attribute("status")
		Attribute("start_time")
		Attribute("finish_time")

		Required("id", "href", "title", "status", "start_time", "finish_time", "created_at")
	})
	//
	// Links(func() {
	// 	Link("user")
	// 	Link("document")
	// })

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("title")
		Attribute("status")
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
		Attribute("start_time")
		Attribute("finish_time")
		// Attribute("document")
	})
})
