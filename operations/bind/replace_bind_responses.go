// Code generated by go-swagger; DO NOT EDIT.

package bind

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceBindOKCode is the HTTP code returned for type ReplaceBindOK
const ReplaceBindOKCode int = 200

/*ReplaceBindOK Bind replaced

swagger:response replaceBindOK
*/
type ReplaceBindOK struct {

	/*
	  In: Body
	*/
	Payload *models.Bind `json:"body,omitempty"`
}

// NewReplaceBindOK creates ReplaceBindOK with default headers values
func NewReplaceBindOK() *ReplaceBindOK {

	return &ReplaceBindOK{}
}

// WithPayload adds the payload to the replace bind o k response
func (o *ReplaceBindOK) WithPayload(payload *models.Bind) *ReplaceBindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace bind o k response
func (o *ReplaceBindOK) SetPayload(payload *models.Bind) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBindAcceptedCode is the HTTP code returned for type ReplaceBindAccepted
const ReplaceBindAcceptedCode int = 202

/*ReplaceBindAccepted Configuration change accepted and reload requested

swagger:response replaceBindAccepted
*/
type ReplaceBindAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.Bind `json:"body,omitempty"`
}

// NewReplaceBindAccepted creates ReplaceBindAccepted with default headers values
func NewReplaceBindAccepted() *ReplaceBindAccepted {

	return &ReplaceBindAccepted{}
}

// WithReloadID adds the reloadId to the replace bind accepted response
func (o *ReplaceBindAccepted) WithReloadID(reloadID string) *ReplaceBindAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the replace bind accepted response
func (o *ReplaceBindAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the replace bind accepted response
func (o *ReplaceBindAccepted) WithPayload(payload *models.Bind) *ReplaceBindAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace bind accepted response
func (o *ReplaceBindAccepted) SetPayload(payload *models.Bind) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBindAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header Reload-ID

	reloadID := o.ReloadID
	if reloadID != "" {
		rw.Header().Set("Reload-ID", reloadID)
	}

	rw.WriteHeader(202)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBindBadRequestCode is the HTTP code returned for type ReplaceBindBadRequest
const ReplaceBindBadRequestCode int = 400

/*ReplaceBindBadRequest Bad request

swagger:response replaceBindBadRequest
*/
type ReplaceBindBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBindBadRequest creates ReplaceBindBadRequest with default headers values
func NewReplaceBindBadRequest() *ReplaceBindBadRequest {

	return &ReplaceBindBadRequest{}
}

// WithPayload adds the payload to the replace bind bad request response
func (o *ReplaceBindBadRequest) WithPayload(payload *models.Error) *ReplaceBindBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace bind bad request response
func (o *ReplaceBindBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBindBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceBindNotFoundCode is the HTTP code returned for type ReplaceBindNotFound
const ReplaceBindNotFoundCode int = 404

/*ReplaceBindNotFound The specified resource was not found

swagger:response replaceBindNotFound
*/
type ReplaceBindNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBindNotFound creates ReplaceBindNotFound with default headers values
func NewReplaceBindNotFound() *ReplaceBindNotFound {

	return &ReplaceBindNotFound{}
}

// WithPayload adds the payload to the replace bind not found response
func (o *ReplaceBindNotFound) WithPayload(payload *models.Error) *ReplaceBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace bind not found response
func (o *ReplaceBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceBindDefault General Error

swagger:response replaceBindDefault
*/
type ReplaceBindDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceBindDefault creates ReplaceBindDefault with default headers values
func NewReplaceBindDefault(code int) *ReplaceBindDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceBindDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace bind default response
func (o *ReplaceBindDefault) WithStatusCode(code int) *ReplaceBindDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace bind default response
func (o *ReplaceBindDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace bind default response
func (o *ReplaceBindDefault) WithPayload(payload *models.Error) *ReplaceBindDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace bind default response
func (o *ReplaceBindDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceBindDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}