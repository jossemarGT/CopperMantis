package controllers

import (
	"github.com/CopperMantis/CopperMantis/app"
	"github.com/goadesign/goa"
)

// MonitoringController implements the monitoring resource.
type MonitoringController struct {
	*goa.Controller
}

// NewMonitoringController creates a monitoring controller.
func NewMonitoringController(service *goa.Service) *MonitoringController {
	return &MonitoringController{Controller: service.NewController("MonitoringController")}
}

// Ping runs the ping action.
func (c *MonitoringController) Ping(ctx *app.PingMonitoringContext) error {
	// MonitoringController_Ping: start_implement

	// Put your logic here

	// MonitoringController_Ping: end_implement
	return nil
}
