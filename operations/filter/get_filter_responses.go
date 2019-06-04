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

package filter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetFilterOKCode is the HTTP code returned for type GetFilterOK
const GetFilterOKCode int = 200

/*GetFilterOK Successful operation

swagger:response getFilterOK
*/
type GetFilterOK struct {

	/*
	  In: Body
	*/
	Payload *GetFilterOKBody `json:"body,omitempty"`
}

// NewGetFilterOK creates GetFilterOK with default headers values
func NewGetFilterOK() *GetFilterOK {

	return &GetFilterOK{}
}

// WithPayload adds the payload to the get filter o k response
func (o *GetFilterOK) WithPayload(payload *GetFilterOKBody) *GetFilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filter o k response
func (o *GetFilterOK) SetPayload(payload *GetFilterOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetFilterNotFoundCode is the HTTP code returned for type GetFilterNotFound
const GetFilterNotFoundCode int = 404

/*GetFilterNotFound The specified resource was not found

swagger:response getFilterNotFound
*/
type GetFilterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFilterNotFound creates GetFilterNotFound with default headers values
func NewGetFilterNotFound() *GetFilterNotFound {

	return &GetFilterNotFound{}
}

// WithPayload adds the payload to the get filter not found response
func (o *GetFilterNotFound) WithPayload(payload *models.Error) *GetFilterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filter not found response
func (o *GetFilterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFilterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetFilterDefault General Error

swagger:response getFilterDefault
*/
type GetFilterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFilterDefault creates GetFilterDefault with default headers values
func NewGetFilterDefault(code int) *GetFilterDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFilterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get filter default response
func (o *GetFilterDefault) WithStatusCode(code int) *GetFilterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get filter default response
func (o *GetFilterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get filter default response
func (o *GetFilterDefault) WithPayload(payload *models.Error) *GetFilterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get filter default response
func (o *GetFilterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFilterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
