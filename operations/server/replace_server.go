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

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReplaceServerHandlerFunc turns a function with the right signature into a replace server handler
type ReplaceServerHandlerFunc func(ReplaceServerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceServerHandlerFunc) Handle(params ReplaceServerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceServerHandler interface for that can handle valid replace server params
type ReplaceServerHandler interface {
	Handle(ReplaceServerParams, interface{}) middleware.Responder
}

// NewReplaceServer creates a new http.Handler for the replace server operation
func NewReplaceServer(ctx *middleware.Context, handler ReplaceServerHandler) *ReplaceServer {
	return &ReplaceServer{Context: ctx, Handler: handler}
}

/*ReplaceServer swagger:route PUT /services/haproxy/configuration/servers/{name} Server replaceServer

Replace a server

Replaces a server configuration by it's name in the specified backend.

*/
type ReplaceServer struct {
	Context *middleware.Context
	Handler ReplaceServerHandler
}

func (o *ReplaceServer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceServerParams()

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
