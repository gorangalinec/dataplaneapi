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

package stick_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetStickRulesOKCode is the HTTP code returned for type GetStickRulesOK
const GetStickRulesOKCode int = 200

/*GetStickRulesOK Successful operation

swagger:response getStickRulesOK
*/
type GetStickRulesOK struct {

	/*
	  In: Body
	*/
	Payload *GetStickRulesOKBody `json:"body,omitempty"`
}

// NewGetStickRulesOK creates GetStickRulesOK with default headers values
func NewGetStickRulesOK() *GetStickRulesOK {

	return &GetStickRulesOK{}
}

// WithPayload adds the payload to the get stick rules o k response
func (o *GetStickRulesOK) WithPayload(payload *GetStickRulesOKBody) *GetStickRulesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick rules o k response
func (o *GetStickRulesOK) SetPayload(payload *GetStickRulesOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickRulesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetStickRulesDefault General Error

swagger:response getStickRulesDefault
*/
type GetStickRulesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetStickRulesDefault creates GetStickRulesDefault with default headers values
func NewGetStickRulesDefault(code int) *GetStickRulesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetStickRulesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get stick rules default response
func (o *GetStickRulesDefault) WithStatusCode(code int) *GetStickRulesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get stick rules default response
func (o *GetStickRulesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get stick rules default response
func (o *GetStickRulesDefault) WithPayload(payload *models.Error) *GetStickRulesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get stick rules default response
func (o *GetStickRulesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetStickRulesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
