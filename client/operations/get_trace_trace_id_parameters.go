package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetTraceTraceIDParams creates a new GetTraceTraceIDParams object
// with the default values initialized.
func NewGetTraceTraceIDParams() *GetTraceTraceIDParams {
	var ()
	return &GetTraceTraceIDParams{}
}

/*GetTraceTraceIDParams contains all the parameters to send to the API endpoint
for the get trace trace ID operation typically these are written to a http.Request
*/
type GetTraceTraceIDParams struct {

	/*TraceID
	  the 64-bit hex-encoded id of the trace as a path parameter.

	*/
	TraceID string
}

// WithTraceID adds the traceId to the get trace trace ID params
func (o *GetTraceTraceIDParams) WithTraceID(traceId string) *GetTraceTraceIDParams {
	o.TraceID = traceId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetTraceTraceIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param traceId
	if err := r.SetPathParam("traceId", o.TraceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
