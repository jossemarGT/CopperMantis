package controllers

import (
	"github.com/CopperMantis/CopperMantis/app"
	"github.com/goadesign/goa"
)

// ContestController implements the contest resource.
type ContestController struct {
	*goa.Controller
}

// NewContestController creates a contest controller.
func NewContestController(service *goa.Service) *ContestController {
	return &ContestController{Controller: service.NewController("ContestController")}
}

// Create runs the create action.
func (c *ContestController) Create(ctx *app.CreateContestContext) error {
	// ContestController_Create: start_implement

	// Put your logic here

	// ContestController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *ContestController) Delete(ctx *app.DeleteContestContext) error {
	// ContestController_Delete: start_implement

	// Put your logic here

	// ContestController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *ContestController) List(ctx *app.ListContestContext) error {
	// ContestController_List: start_implement

	// Put your logic here

	// ContestController_List: end_implement
	res := app.CmsContestCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *ContestController) Show(ctx *app.ShowContestContext) error {
	// ContestController_Show: start_implement

	// Put your logic here

	// ContestController_Show: end_implement
	res := &app.CmsContest{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *ContestController) Update(ctx *app.UpdateContestContext) error {
	// ContestController_Update: start_implement

	// Put your logic here

	// ContestController_Update: end_implement
	return nil
}
