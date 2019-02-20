// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/haproxytech/models"
)

// ReplaceTCPRequestRuleOKCode is the HTTP code returned for type ReplaceTCPRequestRuleOK
const ReplaceTCPRequestRuleOKCode int = 200

/*ReplaceTCPRequestRuleOK TCP Request Rule replaced

swagger:response replaceTcpRequestRuleOK
*/
type ReplaceTCPRequestRuleOK struct {

	/*
	  In: Body
	*/
	Payload *models.TCPRequestRule `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleOK creates ReplaceTCPRequestRuleOK with default headers values
func NewReplaceTCPRequestRuleOK() *ReplaceTCPRequestRuleOK {

	return &ReplaceTCPRequestRuleOK{}
}

// WithPayload adds the payload to the replace Tcp request rule o k response
func (o *ReplaceTCPRequestRuleOK) WithPayload(payload *models.TCPRequestRule) *ReplaceTCPRequestRuleOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule o k response
func (o *ReplaceTCPRequestRuleOK) SetPayload(payload *models.TCPRequestRule) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPRequestRuleBadRequestCode is the HTTP code returned for type ReplaceTCPRequestRuleBadRequest
const ReplaceTCPRequestRuleBadRequestCode int = 400

/*ReplaceTCPRequestRuleBadRequest Bad request

swagger:response replaceTcpRequestRuleBadRequest
*/
type ReplaceTCPRequestRuleBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleBadRequest creates ReplaceTCPRequestRuleBadRequest with default headers values
func NewReplaceTCPRequestRuleBadRequest() *ReplaceTCPRequestRuleBadRequest {

	return &ReplaceTCPRequestRuleBadRequest{}
}

// WithPayload adds the payload to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule bad request response
func (o *ReplaceTCPRequestRuleBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ReplaceTCPRequestRuleNotFoundCode is the HTTP code returned for type ReplaceTCPRequestRuleNotFound
const ReplaceTCPRequestRuleNotFoundCode int = 404

/*ReplaceTCPRequestRuleNotFound The specified resource was not found

swagger:response replaceTcpRequestRuleNotFound
*/
type ReplaceTCPRequestRuleNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleNotFound creates ReplaceTCPRequestRuleNotFound with default headers values
func NewReplaceTCPRequestRuleNotFound() *ReplaceTCPRequestRuleNotFound {

	return &ReplaceTCPRequestRuleNotFound{}
}

// WithPayload adds the payload to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace Tcp request rule not found response
func (o *ReplaceTCPRequestRuleNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ReplaceTCPRequestRuleDefault General Error

swagger:response replaceTcpRequestRuleDefault
*/
type ReplaceTCPRequestRuleDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewReplaceTCPRequestRuleDefault creates ReplaceTCPRequestRuleDefault with default headers values
func NewReplaceTCPRequestRuleDefault(code int) *ReplaceTCPRequestRuleDefault {
	if code <= 0 {
		code = 500
	}

	return &ReplaceTCPRequestRuleDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) WithStatusCode(code int) *ReplaceTCPRequestRuleDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) WithPayload(payload *models.Error) *ReplaceTCPRequestRuleDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the replace TCP request rule default response
func (o *ReplaceTCPRequestRuleDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ReplaceTCPRequestRuleDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
