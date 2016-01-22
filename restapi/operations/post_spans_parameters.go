package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/middleware"

	"github.com/t-yuki/zipkin-go/models"
)

// NewPostSpansParams creates a new PostSpansParams object
// with the default values initialized.
func NewPostSpansParams() PostSpansParams {
	var ()
	return PostSpansParams{}
}

// PostSpansParams contains all the bound params for the post spans operation
// typically these are obtained from a http.Request
//
// swagger:parameters PostSpans
type PostSpansParams struct {
	/*A list of spans
	  Required: true
	  In: body
	*/
	Span models.ListOfSpans
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *PostSpansParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	var body models.ListOfSpans
	if err := route.Consumer.Consume(r.Body, &body); err != nil {
		res = append(res, errors.NewParseError("span", "body", "", err))
	} else {
		for _, io := range o.Span {
			if err := io.Validate(route.Formats); err != nil {
				res = append(res, err)
				break
			}
		}

		if len(res) == 0 {
			o.Span = body
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
