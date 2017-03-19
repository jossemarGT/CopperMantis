// Code generated by goagen v1.1.0, command line:
// $ goa_gen
//
// API "CopperMantis": Application User Types
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// Contest representation
type contestPayload struct {
	FinishTime *time.Time `form:"finish_time,omitempty" json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	StartTime  *time.Time `form:"start_time,omitempty" json:"start_time,omitempty" xml:"start_time,omitempty"`
	// Contest status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Contest title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the contestPayload type instance.
func (ut *contestPayload) Validate() (err error) {
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 3, true))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 15, false))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 3, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 140, false))
		}
	}
	return
}

// Publicize creates ContestPayload from contestPayload
func (ut *contestPayload) Publicize() *ContestPayload {
	var pub ContestPayload
	if ut.FinishTime != nil {
		pub.FinishTime = ut.FinishTime
	}
	if ut.StartTime != nil {
		pub.StartTime = ut.StartTime
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Contest representation
type ContestPayload struct {
	FinishTime *time.Time `form:"finish_time,omitempty" json:"finish_time,omitempty" xml:"finish_time,omitempty"`
	StartTime  *time.Time `form:"start_time,omitempty" json:"start_time,omitempty" xml:"start_time,omitempty"`
	// Contest status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Contest title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the ContestPayload type instance.
func (ut *ContestPayload) Validate() (err error) {
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 3, true))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 15, false))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 3, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 140, false))
		}
	}
	return
}

// Document representation
type documentPayload struct {
	// Document content in markdown notation
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Brief description of the document
	Excerpt *string `form:"excerpt,omitempty" json:"excerpt,omitempty" xml:"excerpt,omitempty"`
	// Document status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Document title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the documentPayload type instance.
func (ut *documentPayload) Validate() (err error) {
	if ut.Content != nil {
		if utf8.RuneCountInString(*ut.Content) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.content`, *ut.Content, utf8.RuneCountInString(*ut.Content), 5, true))
		}
	}
	if ut.Content != nil {
		if utf8.RuneCountInString(*ut.Content) > 2048 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.content`, *ut.Content, utf8.RuneCountInString(*ut.Content), 2048, false))
		}
	}
	if ut.Excerpt != nil {
		if utf8.RuneCountInString(*ut.Excerpt) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.excerpt`, *ut.Excerpt, utf8.RuneCountInString(*ut.Excerpt), 5, true))
		}
	}
	if ut.Excerpt != nil {
		if utf8.RuneCountInString(*ut.Excerpt) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.excerpt`, *ut.Excerpt, utf8.RuneCountInString(*ut.Excerpt), 140, false))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 3, true))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 15, false))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 5, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 140, false))
		}
	}
	return
}

// Publicize creates DocumentPayload from documentPayload
func (ut *documentPayload) Publicize() *DocumentPayload {
	var pub DocumentPayload
	if ut.Content != nil {
		pub.Content = ut.Content
	}
	if ut.Excerpt != nil {
		pub.Excerpt = ut.Excerpt
	}
	if ut.Status != nil {
		pub.Status = ut.Status
	}
	if ut.Title != nil {
		pub.Title = ut.Title
	}
	return &pub
}

// Document representation
type DocumentPayload struct {
	// Document content in markdown notation
	Content *string `form:"content,omitempty" json:"content,omitempty" xml:"content,omitempty"`
	// Brief description of the document
	Excerpt *string `form:"excerpt,omitempty" json:"excerpt,omitempty" xml:"excerpt,omitempty"`
	// Document status
	Status *string `form:"status,omitempty" json:"status,omitempty" xml:"status,omitempty"`
	// Document title
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the DocumentPayload type instance.
func (ut *DocumentPayload) Validate() (err error) {
	if ut.Content != nil {
		if utf8.RuneCountInString(*ut.Content) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.content`, *ut.Content, utf8.RuneCountInString(*ut.Content), 5, true))
		}
	}
	if ut.Content != nil {
		if utf8.RuneCountInString(*ut.Content) > 2048 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.content`, *ut.Content, utf8.RuneCountInString(*ut.Content), 2048, false))
		}
	}
	if ut.Excerpt != nil {
		if utf8.RuneCountInString(*ut.Excerpt) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.excerpt`, *ut.Excerpt, utf8.RuneCountInString(*ut.Excerpt), 5, true))
		}
	}
	if ut.Excerpt != nil {
		if utf8.RuneCountInString(*ut.Excerpt) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.excerpt`, *ut.Excerpt, utf8.RuneCountInString(*ut.Excerpt), 140, false))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 3, true))
		}
	}
	if ut.Status != nil {
		if utf8.RuneCountInString(*ut.Status) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.status`, *ut.Status, utf8.RuneCountInString(*ut.Status), 15, false))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 5, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 140 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 140, false))
		}
	}
	return
}

// User profile representation
type profilePayload struct {
	// Brief user self description
	Biography *string    `form:"biography,omitempty" json:"biography,omitempty" xml:"biography,omitempty"`
	BirthDate *time.Time `form:"birth_date,omitempty" json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// Country where is registering
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Unique nickname
	DisplayName *string `form:"display_name,omitempty" json:"display_name,omitempty" xml:"display_name,omitempty"`
	// Full real name
	FullName *string `form:"full_name,omitempty" json:"full_name,omitempty" xml:"full_name,omitempty"`
}

// Validate validates the profilePayload type instance.
func (ut *profilePayload) Validate() (err error) {
	if ut.Biography != nil {
		if utf8.RuneCountInString(*ut.Biography) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.biography`, *ut.Biography, utf8.RuneCountInString(*ut.Biography), 3, true))
		}
	}
	if ut.Biography != nil {
		if utf8.RuneCountInString(*ut.Biography) > 350 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.biography`, *ut.Biography, utf8.RuneCountInString(*ut.Biography), 350, false))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 2, true))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) > 60 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 60, false))
		}
	}
	if ut.DisplayName != nil {
		if utf8.RuneCountInString(*ut.DisplayName) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.display_name`, *ut.DisplayName, utf8.RuneCountInString(*ut.DisplayName), 5, true))
		}
	}
	if ut.DisplayName != nil {
		if utf8.RuneCountInString(*ut.DisplayName) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.display_name`, *ut.DisplayName, utf8.RuneCountInString(*ut.DisplayName), 15, false))
		}
	}
	if ut.FullName != nil {
		if utf8.RuneCountInString(*ut.FullName) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.full_name`, *ut.FullName, utf8.RuneCountInString(*ut.FullName), 3, true))
		}
	}
	if ut.FullName != nil {
		if utf8.RuneCountInString(*ut.FullName) > 60 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.full_name`, *ut.FullName, utf8.RuneCountInString(*ut.FullName), 60, false))
		}
	}
	return
}

// Publicize creates ProfilePayload from profilePayload
func (ut *profilePayload) Publicize() *ProfilePayload {
	var pub ProfilePayload
	if ut.Biography != nil {
		pub.Biography = ut.Biography
	}
	if ut.BirthDate != nil {
		pub.BirthDate = ut.BirthDate
	}
	if ut.Country != nil {
		pub.Country = ut.Country
	}
	if ut.DisplayName != nil {
		pub.DisplayName = ut.DisplayName
	}
	if ut.FullName != nil {
		pub.FullName = ut.FullName
	}
	return &pub
}

// User profile representation
type ProfilePayload struct {
	// Brief user self description
	Biography *string    `form:"biography,omitempty" json:"biography,omitempty" xml:"biography,omitempty"`
	BirthDate *time.Time `form:"birth_date,omitempty" json:"birth_date,omitempty" xml:"birth_date,omitempty"`
	// Country where is registering
	Country *string `form:"country,omitempty" json:"country,omitempty" xml:"country,omitempty"`
	// Unique nickname
	DisplayName *string `form:"display_name,omitempty" json:"display_name,omitempty" xml:"display_name,omitempty"`
	// Full real name
	FullName *string `form:"full_name,omitempty" json:"full_name,omitempty" xml:"full_name,omitempty"`
}

// Validate validates the ProfilePayload type instance.
func (ut *ProfilePayload) Validate() (err error) {
	if ut.Biography != nil {
		if utf8.RuneCountInString(*ut.Biography) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.biography`, *ut.Biography, utf8.RuneCountInString(*ut.Biography), 3, true))
		}
	}
	if ut.Biography != nil {
		if utf8.RuneCountInString(*ut.Biography) > 350 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.biography`, *ut.Biography, utf8.RuneCountInString(*ut.Biography), 350, false))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) < 2 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 2, true))
		}
	}
	if ut.Country != nil {
		if utf8.RuneCountInString(*ut.Country) > 60 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.country`, *ut.Country, utf8.RuneCountInString(*ut.Country), 60, false))
		}
	}
	if ut.DisplayName != nil {
		if utf8.RuneCountInString(*ut.DisplayName) < 5 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.display_name`, *ut.DisplayName, utf8.RuneCountInString(*ut.DisplayName), 5, true))
		}
	}
	if ut.DisplayName != nil {
		if utf8.RuneCountInString(*ut.DisplayName) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.display_name`, *ut.DisplayName, utf8.RuneCountInString(*ut.DisplayName), 15, false))
		}
	}
	if ut.FullName != nil {
		if utf8.RuneCountInString(*ut.FullName) < 3 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.full_name`, *ut.FullName, utf8.RuneCountInString(*ut.FullName), 3, true))
		}
	}
	if ut.FullName != nil {
		if utf8.RuneCountInString(*ut.FullName) > 60 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`response.full_name`, *ut.FullName, utf8.RuneCountInString(*ut.FullName), 60, false))
		}
	}
	return
}
