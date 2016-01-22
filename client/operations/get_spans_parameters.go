package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetSpansParams creates a new GetSpansParams object
// with the default values initialized.
func NewGetSpansParams() *GetSpansParams {
	var ()
	return &GetSpansParams{}
}

/*GetSpansParams contains all the parameters to send to the API endpoint
for the get spans operation typically these are written to a http.Request
*/
type GetSpansParams struct {

	/*ServiceName
	  Ex zipkin-web (required) - service that logged an annotation in the
	span.


	*/
	ServiceName string
}

// WithServiceName adds the serviceName to the get spans params
func (o *GetSpansParams) WithServiceName(serviceName string) *GetSpansParams {
	o.ServiceName = serviceName
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpansParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param serviceName
	qrServiceName := o.ServiceName
	qServiceName := qrServiceName
	if qServiceName != "" {
		if err := r.SetQueryParam("serviceName", qServiceName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
