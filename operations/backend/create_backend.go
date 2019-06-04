// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package backend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateBackendHandlerFunc turns a function with the right signature into a create backend handler
type CreateBackendHandlerFunc func(CreateBackendParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateBackendHandlerFunc) Handle(params CreateBackendParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CreateBackendHandler interface for that can handle valid create backend params
type CreateBackendHandler interface {
	Handle(CreateBackendParams, interface{}) middleware.Responder
}

// NewCreateBackend creates a new http.Handler for the create backend operation
func NewCreateBackend(ctx *middleware.Context, handler CreateBackendHandler) *CreateBackend {
	return &CreateBackend{Context: ctx, Handler: handler}
}

/*CreateBackend swagger:route POST /services/haproxy/configuration/backends Backend createBackend

Add a backend

Adds a new backend to the configuration file.

*/
type CreateBackend struct {
	Context *middleware.Context
	Handler CreateBackendHandler
}

func (o *CreateBackend) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateBackendParams()

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
