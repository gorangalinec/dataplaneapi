// Code generated by go-swagger; DO NOT EDIT.

package log_target

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetLogTargetHandlerFunc turns a function with the right signature into a get log target handler
type GetLogTargetHandlerFunc func(GetLogTargetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetLogTargetHandlerFunc) Handle(params GetLogTargetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetLogTargetHandler interface for that can handle valid get log target params
type GetLogTargetHandler interface {
	Handle(GetLogTargetParams, interface{}) middleware.Responder
}

// NewGetLogTarget creates a new http.Handler for the get log target operation
func NewGetLogTarget(ctx *middleware.Context, handler GetLogTargetHandler) *GetLogTarget {
	return &GetLogTarget{Context: ctx, Handler: handler}
}

/*GetLogTarget swagger:route GET /services/haproxy/configuration/log_targets/{id} LogTarget getLogTarget

Return one Log Target

Returns one Log Target configuration by it's ID in the specified parent.

*/
type GetLogTarget struct {
	Context *middleware.Context
	Handler GetLogTargetHandler
}

func (o *GetLogTarget) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetLogTargetParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
