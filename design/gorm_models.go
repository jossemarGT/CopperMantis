package design

import (
	"github.com/goadesign/gorma"
	. "github.com/goadesign/gorma/dsl"
)

var _ = StorageGroup("CopperMantis", func() {
	Store("postgres", gorma.Postgres, func() {
		Description("CopperMantis Postgres relational store")
		// user_profile model
		Model("Profile", func() {
			Alias("profile")
			RendersTo(ProfileMediaType)
			BuildsFrom(func() {
				Payload("profile", "create")
				Payload("profile", "update")
			})
			Description("User Profile")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			// Relationships
			// HasMany("contest", "contest")
		})

		// Contest Model
		Model("Contest", func() {
			Alias("contest")
			RendersTo(ContestMediaType)
			BuildsFrom(func() {
				Payload("contest", "create")
				Payload("contest", "update")
			})
			Description("Programming Contest")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			// Relationships
			// HasMany("document", "document")
		})

		// Document Model
		Model("Document", func() {
			Alias("document")
			RendersTo(DocumentMediaType)
			BuildsFrom(func() {
				Payload("document", "create")
				Payload("document", "update")
			})
			Description("Document")
			Field("id", gorma.Integer, func() {
				PrimaryKey()
			})
			// Relationships
			// HasMany("term", "term")
			// HasMany("contest", "contest")
		})
	})
})
