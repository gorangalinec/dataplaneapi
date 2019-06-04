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

package defaults

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetDefaultsOKCode is the HTTP code returned for type GetDefaultsOK
const GetDefaultsOKCode int = 200

/*GetDefaultsOK Successful operation

swagger:response getDefaultsOK
*/
type GetDefaultsOK struct {

	/*
	  In: Body
	*/
	Payload *GetDefaultsOKBody `json:"body,omitempty"`
}

// NewGetDefaultsOK creates GetDefaultsOK with default headers values
func NewGetDefaultsOK() *GetDefaultsOK {

	return &GetDefaultsOK{}
}

// WithPayload adds the payload to the get defaults o k response
func (o *GetDefaultsOK) WithPayload(payload *GetDefaultsOKBody) *GetDefaultsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get defaults o k response
func (o *GetDefaultsOK) SetPayload(payload *GetDefaultsOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefaultsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetDefaultsDefault General Error

swagger:response getDefaultsDefault
*/
type GetDefaultsDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetDefaultsDefault creates GetDefaultsDefault with default headers values
func NewGetDefaultsDefault(code int) *GetDefaultsDefault {
	if code <= 0 {
		code = 500
	}

	return &GetDefaultsDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get defaults default response
func (o *GetDefaultsDefault) WithStatusCode(code int) *GetDefaultsDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get defaults default response
func (o *GetDefaultsDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get defaults default response
func (o *GetDefaultsDefault) WithPayload(payload *models.Error) *GetDefaultsDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get defaults default response
func (o *GetDefaultsDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetDefaultsDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
