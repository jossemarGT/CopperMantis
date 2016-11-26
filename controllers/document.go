package controllers

import (
	"github.com/CopperMantis/CopperMantis/app"
	"github.com/goadesign/goa"
)

// DocumentController implements the document resource.
type DocumentController struct {
	*goa.Controller
}

// NewDocumentController creates a document controller.
func NewDocumentController(service *goa.Service) *DocumentController {
	return &DocumentController{Controller: service.NewController("DocumentController")}
}

// Create runs the create action.
func (c *DocumentController) Create(ctx *app.CreateDocumentContext) error {
	// DocumentController_Create: start_implement

	// Put your logic here

	// DocumentController_Create: end_implement
	return nil
}

// Delete runs the delete action.
func (c *DocumentController) Delete(ctx *app.DeleteDocumentContext) error {
	// DocumentController_Delete: start_implement

	// Put your logic here

	// DocumentController_Delete: end_implement
	return nil
}

// List runs the list action.
func (c *DocumentController) List(ctx *app.ListDocumentContext) error {
	// DocumentController_List: start_implement

	// Put your logic here

	// DocumentController_List: end_implement
	res := app.CmsDocumentCollection{}
	return ctx.OK(res)
}

// Show runs the show action.
func (c *DocumentController) Show(ctx *app.ShowDocumentContext) error {
	// DocumentController_Show: start_implement

	// Put your logic here

	// DocumentController_Show: end_implement
	res := &app.CmsDocument{}
	return ctx.OK(res)
}

// Update runs the update action.
func (c *DocumentController) Update(ctx *app.UpdateDocumentContext) error {
	// DocumentController_Update: start_implement

	// Put your logic here

	// DocumentController_Update: end_implement
	return nil
}
