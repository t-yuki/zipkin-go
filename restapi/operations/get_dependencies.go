package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// GetDependenciesHandlerFunc turns a function with the right signature into a get dependencies handler
type GetDependenciesHandlerFunc func(GetDependenciesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetDependenciesHandlerFunc) Handle(params GetDependenciesParams) middleware.Responder {
	return fn(params)
}

// GetDependenciesHandler interface for that can handle valid get dependencies params
type GetDependenciesHandler interface {
	Handle(GetDependenciesParams) middleware.Responder
}

// NewGetDependencies creates a new http.Handler for the get dependencies operation
func NewGetDependencies(ctx *middleware.Context, handler GetDependenciesHandler) *GetDependencies {
	return &GetDependencies{Context: ctx, Handler: handler}
}

/*GetDependencies swagger:route GET /dependencies getDependencies

Returns dependency links derived from spans.

Span names are in lowercase, rpc method for example. Conventionally,
when the span name isn't known, name = "unknown".


*/
type GetDependencies struct {
	Context *middleware.Context
	Params  GetDependenciesParams
	Handler GetDependenciesHandler
}

func (o *GetDependencies) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewGetDependenciesParams()

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
