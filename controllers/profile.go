package controllers

import (
	"github.com/CopperMantis/CopperMantis/app"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ProfileController implements the profile resource.
type ProfileController struct {
	*goa.Controller
	db *gorm.DB
}

// NewProfileController creates a profile controller.
func NewProfileController(service *goa.Service, db *gorm.DB) *ProfileController {
	return &ProfileController{Controller: service.NewController("ProfileController"), db: db}
}

// Create runs the create action.
func (c *ProfileController) Create(ctx *app.CreateProfileContext) error {
	// ProfileController_Create: start_implement

	// Put your logic here

	// ProfileController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *ProfileController) Delete(ctx *app.DeleteProfileContext) error {
	// ProfileController_Delete: start_implement

	// Put your logic here

	// ProfileController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *ProfileController) List(ctx *app.ListProfileContext) error {
	// ProfileController_List: start_implement

	// Put your logic here

	// ProfileController_List: end_implement
	res := app.CmsProfileCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ProfileController) Show(ctx *app.ShowProfileContext) error {
	// ProfileController_Show: start_implement

	// Put your logic here

	// ProfileController_Show: end_implement
	res := &app.CmsProfile{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ProfileController) Update(ctx *app.UpdateProfileContext) error {
	// ProfileController_Update: start_implement

	// Put your logic here

	// ProfileController_Update: end_implement
	return nil
}
