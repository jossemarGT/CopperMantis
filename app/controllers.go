// Code generated by goagen v1.1.0, command line:
// $ goa_gen
//
// API "CopperMantis": Application Controllers
//
// The content of this file is auto-generated, DO NOT MODIFY

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ContestController is the controller interface for the Contest actions.
type ContestController interface {
	goa.Muxer
	Create(*CreateContestContext) error
	Delete(*DeleteContestContext) error
	List(*ListContestContext) error
	Show(*ShowContestContext) error
	Update(*UpdateContestContext) error
}

// MountContestController "mounts" a Contest resource controller on the given service.
func MountContestController(service *goa.Service, ctrl ContestController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateContestContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ContestPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cms/v1/contest", ctrl.MuxHandler("Create", h, unmarshalCreateContestPayload))
	service.LogInfo("mount", "ctrl", "Contest", "action", "Create", "route", "POST /cms/v1/contest")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteContestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cms/v1/contest/:contestID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Contest", "action", "Delete", "route", "DELETE /cms/v1/contest/:contestID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListContestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/contest", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Contest", "action", "List", "route", "GET /cms/v1/contest")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowContestContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/contest/:contestID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Contest", "action", "Show", "route", "GET /cms/v1/contest/:contestID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateContestContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ContestPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/cms/v1/contest/:contestID", ctrl.MuxHandler("Update", h, unmarshalUpdateContestPayload))
	service.LogInfo("mount", "ctrl", "Contest", "action", "Update", "route", "PUT /cms/v1/contest/:contestID")
	service.Mux.Handle("PATCH", "/cms/v1/contest/:contestID", ctrl.MuxHandler("Update", h, unmarshalUpdateContestPayload))
	service.LogInfo("mount", "ctrl", "Contest", "action", "Update", "route", "PATCH /cms/v1/contest/:contestID")
}

// unmarshalCreateContestPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateContestPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &contestPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateContestPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateContestPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &contestPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// DocumentController is the controller interface for the Document actions.
type DocumentController interface {
	goa.Muxer
	Create(*CreateDocumentContext) error
	Delete(*DeleteDocumentContext) error
	List(*ListDocumentContext) error
	Show(*ShowDocumentContext) error
	Update(*UpdateDocumentContext) error
}

// MountDocumentController "mounts" a Document resource controller on the given service.
func MountDocumentController(service *goa.Service, ctrl DocumentController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateDocumentContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*DocumentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cms/v1/document", ctrl.MuxHandler("Create", h, unmarshalCreateDocumentPayload))
	service.LogInfo("mount", "ctrl", "Document", "action", "Create", "route", "POST /cms/v1/document")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteDocumentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cms/v1/document/:documentID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Document", "action", "Delete", "route", "DELETE /cms/v1/document/:documentID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListDocumentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/document", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Document", "action", "List", "route", "GET /cms/v1/document")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowDocumentContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/document/:documentID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Document", "action", "Show", "route", "GET /cms/v1/document/:documentID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateDocumentContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*DocumentPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/cms/v1/document/:documentID", ctrl.MuxHandler("Update", h, unmarshalUpdateDocumentPayload))
	service.LogInfo("mount", "ctrl", "Document", "action", "Update", "route", "PUT /cms/v1/document/:documentID")
	service.Mux.Handle("PATCH", "/cms/v1/document/:documentID", ctrl.MuxHandler("Update", h, unmarshalUpdateDocumentPayload))
	service.LogInfo("mount", "ctrl", "Document", "action", "Update", "route", "PATCH /cms/v1/document/:documentID")
}

// unmarshalCreateDocumentPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateDocumentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &documentPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateDocumentPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateDocumentPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &documentPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// MonitoringController is the controller interface for the Monitoring actions.
type MonitoringController interface {
	goa.Muxer
	Ping(*PingMonitoringContext) error
}

// MountMonitoringController "mounts" a Monitoring resource controller on the given service.
func MountMonitoringController(service *goa.Service, ctrl MonitoringController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPingMonitoringContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Ping(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/_ping", ctrl.MuxHandler("Ping", h, nil))
	service.LogInfo("mount", "ctrl", "Monitoring", "action", "Ping", "route", "GET /cms/v1/_ping")
}

// ProfileController is the controller interface for the Profile actions.
type ProfileController interface {
	goa.Muxer
	Create(*CreateProfileContext) error
	Delete(*DeleteProfileContext) error
	List(*ListProfileContext) error
	Show(*ShowProfileContext) error
	Update(*UpdateProfileContext) error
}

// MountProfileController "mounts" a Profile resource controller on the given service.
func MountProfileController(service *goa.Service, ctrl ProfileController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ProfilePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	service.Mux.Handle("POST", "/cms/v1/profile", ctrl.MuxHandler("Create", h, unmarshalCreateProfilePayload))
	service.LogInfo("mount", "ctrl", "Profile", "action", "Create", "route", "POST /cms/v1/profile")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	service.Mux.Handle("DELETE", "/cms/v1/profile/:profileID", ctrl.MuxHandler("Delete", h, nil))
	service.LogInfo("mount", "ctrl", "Profile", "action", "Delete", "route", "DELETE /cms/v1/profile/:profileID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/profile", ctrl.MuxHandler("List", h, nil))
	service.LogInfo("mount", "ctrl", "Profile", "action", "List", "route", "GET /cms/v1/profile")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewShowProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/cms/v1/profile/:profileID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Profile", "action", "Show", "route", "GET /cms/v1/profile/:profileID")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateProfileContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*ProfilePayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	service.Mux.Handle("PUT", "/cms/v1/profile/:profileID", ctrl.MuxHandler("Update", h, unmarshalUpdateProfilePayload))
	service.LogInfo("mount", "ctrl", "Profile", "action", "Update", "route", "PUT /cms/v1/profile/:profileID")
	service.Mux.Handle("PATCH", "/cms/v1/profile/:profileID", ctrl.MuxHandler("Update", h, unmarshalUpdateProfilePayload))
	service.LogInfo("mount", "ctrl", "Profile", "action", "Update", "route", "PATCH /cms/v1/profile/:profileID")
}

// unmarshalCreateProfilePayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateProfilePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &profilePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateProfilePayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateProfilePayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &profilePayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}
