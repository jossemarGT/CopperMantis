package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// ProfileType User profile object
var ProfileType = Type("ProfilePayload", func() {
	Description("User profile representation")
	Attribute("display_name", String, "Unique nickname", func() {
		MinLength(5)
		MaxLength(15) // Like twitter does
		Example("thunderBetaTester")
	})
	Attribute("full_name", String, "Full real name", func() {
		MinLength(3)
		MaxLength(60)
		Example("John Beta")
	})
	Attribute("country", String, "Country where is registering", func() {
		MinLength(2)
		MaxLength(60)
		Example("USA")
	})
	Attribute("biography", String, "Brief user self description", func() {
		MinLength(3)
		MaxLength(350)
		Example("Number one tester")
	})
	Attribute("birth_date", DateTime)
})

var ProfileMediaType = MediaType("application/vnd.cms.profile+json", func() {
	Description("User profile representation")
	Reference(ProfileType)
	Attributes(func() {
		Attribute("id", Integer, "User Profile ID", func() {
			Example(1)
		})
		Attribute("href", String, "User profile href", func() {
			Example("/profile/1")
		})
		Attribute("created_at", DateTime, "Date of creation")
		Attribute("updated_at", DateTime, "Date of last update")
		Attribute("deleted_at", DateTime, "Date of soft delete")
		// References
		// Attribute("achievement", Rank, "Achievement related to user_profile") ?

		// Attributes below inherit from the base type
		Attribute("display_name")
		Attribute("full_name")
		Attribute("country")
		Attribute("biography")
		Attribute("birth_date")

		Required("id", "href", "display_name", "full_name", "birth_date", "created_at")
	})
	//
	// Links(func() {
	// 	Link("contest") ?
	// })

	View("default", func() {
		Attribute("id")
		Attribute("href")
		Attribute("display_name")
		Attribute("created_at")
	})

	View("tiny", func() {
		Attribute("id")
		Attribute("href")
		Attribute("display_name")
	})

	View("full", func() {
		Attribute("id")
		Attribute("href")
		Attribute("display_name")
		Attribute("full_name")
		Attribute("country")
		Attribute("biography")
		Attribute("birth_date")
		// Attribute("contest") ?
		// Rank ?
	})
})
