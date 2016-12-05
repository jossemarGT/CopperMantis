package controllers

import (
	"github.com/CopperMantis/CopperMantis/app"
	"github.com/CopperMantis/CopperMantis/models"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
)

// ContestController implements the contest resource.
type ContestController struct {
	*goa.Controller
	contestStorage models.ContestStorage
}

// NewContestController creates a contest controller.
func NewContestController(service *goa.Service, db *gorm.DB) *ContestController {
	m := models.NewContestDB(db)
	return &ContestController{Controller: service.NewController("ContestController"), contestStorage: m}
}

// Create runs the create action.
func (c *ContestController) Create(ctx *app.CreateContestContext) error {
	// ContestController_Create: start_implement
	m := models.ContestFromContestPayload(ctx.Payload)
	err := c.contestStorage.Add(ctx.Context, m)
	if err != nil {
		return ctx.InternalServerError(err)
	}
	ctx.ResponseData.Header().Set("Location", app.ContestHref(m.ID))
	return ctx.Created()
	// ContestController_Create: end_implement
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
	m, err := c.contestStorage.OneCmsContest(ctx.Context, ctx.ContestID)
	if err == gorm.ErrRecordNotFound {
		return ctx.NotFound()
	} else if err != nil {
		return ctx.InternalServerError(err)
	}
	m.Href = app.ContestHref(m.ID)
	return ctx.OK(m)
	// ContestController_Show: end_implement
}

// Update runs the update action.
func (c *ContestController) Update(ctx *app.UpdateContestContext) error {
	// ContestController_Update: start_implement

	// Put your logic here

	// ContestController_Update: end_implement
	return nil
}
