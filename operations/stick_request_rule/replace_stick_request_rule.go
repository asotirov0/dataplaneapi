// Code generated by go-swagger; DO NOT EDIT.

package stick_request_rule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// ReplaceStickRequestRuleHandlerFunc turns a function with the right signature into a replace stick request rule handler
type ReplaceStickRequestRuleHandlerFunc func(ReplaceStickRequestRuleParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn ReplaceStickRequestRuleHandlerFunc) Handle(params ReplaceStickRequestRuleParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// ReplaceStickRequestRuleHandler interface for that can handle valid replace stick request rule params
type ReplaceStickRequestRuleHandler interface {
	Handle(ReplaceStickRequestRuleParams, interface{}) middleware.Responder
}

// NewReplaceStickRequestRule creates a new http.Handler for the replace stick request rule operation
func NewReplaceStickRequestRule(ctx *middleware.Context, handler ReplaceStickRequestRuleHandler) *ReplaceStickRequestRule {
	return &ReplaceStickRequestRule{Context: ctx, Handler: handler}
}

/*ReplaceStickRequestRule swagger:route PUT /services/haproxy/configuration/stick_request_rules/{id} StickRequestRule replaceStickRequestRule

Replace a Stick Request Rule

Replaces a Stick Request Rule configuration by it's ID in the specified backend.

*/
type ReplaceStickRequestRule struct {
	Context *middleware.Context
	Handler ReplaceStickRequestRuleHandler
}

func (o *ReplaceStickRequestRule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewReplaceStickRequestRuleParams()

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