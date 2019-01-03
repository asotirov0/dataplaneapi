// Code generated by go-swagger; DO NOT EDIT.

package frontend

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetFrontendsHandlerFunc turns a function with the right signature into a get frontends handler
type GetFrontendsHandlerFunc func(GetFrontendsParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetFrontendsHandlerFunc) Handle(params GetFrontendsParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetFrontendsHandler interface for that can handle valid get frontends params
type GetFrontendsHandler interface {
	Handle(GetFrontendsParams, interface{}) middleware.Responder
}

// NewGetFrontends creates a new http.Handler for the get frontends operation
func NewGetFrontends(ctx *middleware.Context, handler GetFrontendsHandler) *GetFrontends {
	return &GetFrontends{Context: ctx, Handler: handler}
}

/*GetFrontends swagger:route GET /services/haproxy/configuration/frontends Frontend getFrontends

Return an array of frontends

Returns an array of all configured frontends.

*/
type GetFrontends struct {
	Context *middleware.Context
	Handler GetFrontendsHandler
}

func (o *GetFrontends) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetFrontendsParams()

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