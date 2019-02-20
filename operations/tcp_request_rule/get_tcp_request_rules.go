// Code generated by go-swagger; DO NOT EDIT.

package tcp_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTCPRequestRulesHandlerFunc turns a function with the right signature into a get TCP request rules handler
type GetTCPRequestRulesHandlerFunc func(GetTCPRequestRulesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTCPRequestRulesHandlerFunc) Handle(params GetTCPRequestRulesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTCPRequestRulesHandler interface for that can handle valid get TCP request rules params
type GetTCPRequestRulesHandler interface {
	Handle(GetTCPRequestRulesParams, interface{}) middleware.Responder
}

// NewGetTCPRequestRules creates a new http.Handler for the get TCP request rules operation
func NewGetTCPRequestRules(ctx *middleware.Context, handler GetTCPRequestRulesHandler) *GetTCPRequestRules {
	return &GetTCPRequestRules{Context: ctx, Handler: handler}
}

/*GetTCPRequestRules swagger:route GET /services/haproxy/configuration/tcp_request_rules TCPRequestRule getTcpRequestRules

Return an array of all TCP Request Rules

Returns all TCP Request Rules that are configured in specified parent and parent type.

*/
type GetTCPRequestRules struct {
	Context *middleware.Context
	Handler GetTCPRequestRulesHandler
}

func (o *GetTCPRequestRules) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTCPRequestRulesParams()

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
