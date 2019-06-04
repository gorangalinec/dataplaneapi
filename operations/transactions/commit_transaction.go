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

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CommitTransactionHandlerFunc turns a function with the right signature into a commit transaction handler
type CommitTransactionHandlerFunc func(CommitTransactionParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn CommitTransactionHandlerFunc) Handle(params CommitTransactionParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// CommitTransactionHandler interface for that can handle valid commit transaction params
type CommitTransactionHandler interface {
	Handle(CommitTransactionParams, interface{}) middleware.Responder
}

// NewCommitTransaction creates a new http.Handler for the commit transaction operation
func NewCommitTransaction(ctx *middleware.Context, handler CommitTransactionHandler) *CommitTransaction {
	return &CommitTransaction{Context: ctx, Handler: handler}
}

/*CommitTransaction swagger:route PUT /services/haproxy/transactions/{id} Transactions commitTransaction

Commit transaction

Commit transaction, execute all operations in transaction and return msg

*/
type CommitTransaction struct {
	Context *middleware.Context
	Handler CommitTransactionHandler
}

func (o *CommitTransaction) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCommitTransactionParams()

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
