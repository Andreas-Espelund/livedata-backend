// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// IndividualsDeleteV1HandlerFunc turns a function with the right signature into a individuals delete v1 handler
type IndividualsDeleteV1HandlerFunc func(IndividualsDeleteV1Params) middleware.Responder

// Handle executing the request and returning a response
func (fn IndividualsDeleteV1HandlerFunc) Handle(params IndividualsDeleteV1Params) middleware.Responder {
	return fn(params)
}

// IndividualsDeleteV1Handler interface for that can handle valid individuals delete v1 params
type IndividualsDeleteV1Handler interface {
	Handle(IndividualsDeleteV1Params) middleware.Responder
}

// NewIndividualsDeleteV1 creates a new http.Handler for the individuals delete v1 operation
func NewIndividualsDeleteV1(ctx *middleware.Context, handler IndividualsDeleteV1Handler) *IndividualsDeleteV1 {
	return &IndividualsDeleteV1{Context: ctx, Handler: handler}
}

/*
	IndividualsDeleteV1 swagger:route DELETE /v1/individuals/{id} individualsDeleteV1

# Delete individual

Delete an individual by id
*/
type IndividualsDeleteV1 struct {
	Context *middleware.Context
	Handler IndividualsDeleteV1Handler
}

func (o *IndividualsDeleteV1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewIndividualsDeleteV1Params()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
