// Code generated by go-swagger; DO NOT EDIT.

package acl

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// CreateACLCreatedCode is the HTTP code returned for type CreateACLCreated
const CreateACLCreatedCode int = 201

/*CreateACLCreated ACL Rule created

swagger:response createAclCreated
*/
type CreateACLCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ACL `json:"body,omitempty"`
}

// NewCreateACLCreated creates CreateACLCreated with default headers values
func NewCreateACLCreated() *CreateACLCreated {

	return &CreateACLCreated{}
}

// WithPayload adds the payload to the create Acl created response
func (o *CreateACLCreated) WithPayload(payload *models.ACL) *CreateACLCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Acl created response
func (o *CreateACLCreated) SetPayload(payload *models.ACL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateACLCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateACLAcceptedCode is the HTTP code returned for type CreateACLAccepted
const CreateACLAcceptedCode int = 202

/*CreateACLAccepted Configuration change accepted and reload requested

swagger:response createAclAccepted
*/
type CreateACLAccepted struct {
	/*ID of the requested reload

	 */
	ReloadID string `json:"Reload-ID"`

	/*
	  In: Body
	*/
	Payload *models.ACL `json:"body,omitempty"`
}

// NewCreateACLAccepted creates CreateACLAccepted with default headers values
func NewCreateACLAccepted() *CreateACLAccepted {

	return &CreateACLAccepted{}
}

// WithReloadID adds the reloadId to the create Acl accepted response
func (o *CreateACLAccepted) WithReloadID(reloadID string) *CreateACLAccepted {
	o.ReloadID = reloadID
	return o
}

// SetReloadID sets the reloadId to the create Acl accepted response
func (o *CreateACLAccepted) SetReloadID(reloadID string) {
	o.ReloadID = reloadID
}

// WithPayload adds the payload to the create Acl accepted response
func (o *CreateACLAccepted) WithPayload(payload *models.ACL) *CreateACLAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Acl accepted response
func (o *CreateACLAccepted) SetPayload(payload *models.ACL) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateACLAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

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

// CreateACLBadRequestCode is the HTTP code returned for type CreateACLBadRequest
const CreateACLBadRequestCode int = 400

/*CreateACLBadRequest Bad request

swagger:response createAclBadRequest
*/
type CreateACLBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateACLBadRequest creates CreateACLBadRequest with default headers values
func NewCreateACLBadRequest() *CreateACLBadRequest {

	return &CreateACLBadRequest{}
}

// WithPayload adds the payload to the create Acl bad request response
func (o *CreateACLBadRequest) WithPayload(payload *models.Error) *CreateACLBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Acl bad request response
func (o *CreateACLBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateACLBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateACLConflictCode is the HTTP code returned for type CreateACLConflict
const CreateACLConflictCode int = 409

/*CreateACLConflict The specified resource already exists

swagger:response createAclConflict
*/
type CreateACLConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateACLConflict creates CreateACLConflict with default headers values
func NewCreateACLConflict() *CreateACLConflict {

	return &CreateACLConflict{}
}

// WithPayload adds the payload to the create Acl conflict response
func (o *CreateACLConflict) WithPayload(payload *models.Error) *CreateACLConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Acl conflict response
func (o *CreateACLConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateACLConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CreateACLDefault General Error

swagger:response createAclDefault
*/
type CreateACLDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateACLDefault creates CreateACLDefault with default headers values
func NewCreateACLDefault(code int) *CreateACLDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateACLDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create Acl default response
func (o *CreateACLDefault) WithStatusCode(code int) *CreateACLDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create Acl default response
func (o *CreateACLDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create Acl default response
func (o *CreateACLDefault) WithPayload(payload *models.Error) *CreateACLDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create Acl default response
func (o *CreateACLDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateACLDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
