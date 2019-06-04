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

package reloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// GetReloadOKCode is the HTTP code returned for type GetReloadOK
const GetReloadOKCode int = 200

/*GetReloadOK Successful operation

swagger:response getReloadOK
*/
type GetReloadOK struct {

	/*
	  In: Body
	*/
	Payload *models.Reload `json:"body,omitempty"`
}

// NewGetReloadOK creates GetReloadOK with default headers values
func NewGetReloadOK() *GetReloadOK {

	return &GetReloadOK{}
}

// WithPayload adds the payload to the get reload o k response
func (o *GetReloadOK) WithPayload(payload *models.Reload) *GetReloadOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get reload o k response
func (o *GetReloadOK) SetPayload(payload *models.Reload) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReloadOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReloadNotFoundCode is the HTTP code returned for type GetReloadNotFound
const GetReloadNotFoundCode int = 404

/*GetReloadNotFound The specified resource was not found

swagger:response getReloadNotFound
*/
type GetReloadNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReloadNotFound creates GetReloadNotFound with default headers values
func NewGetReloadNotFound() *GetReloadNotFound {

	return &GetReloadNotFound{}
}

// WithPayload adds the payload to the get reload not found response
func (o *GetReloadNotFound) WithPayload(payload *models.Error) *GetReloadNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get reload not found response
func (o *GetReloadNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReloadNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*GetReloadDefault General Error

swagger:response getReloadDefault
*/
type GetReloadDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReloadDefault creates GetReloadDefault with default headers values
func NewGetReloadDefault(code int) *GetReloadDefault {
	if code <= 0 {
		code = 500
	}

	return &GetReloadDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get reload default response
func (o *GetReloadDefault) WithStatusCode(code int) *GetReloadDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get reload default response
func (o *GetReloadDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get reload default response
func (o *GetReloadDefault) WithPayload(payload *models.Error) *GetReloadDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get reload default response
func (o *GetReloadDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReloadDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
