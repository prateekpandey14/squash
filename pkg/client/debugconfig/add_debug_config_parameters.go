package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// NewAddDebugConfigParams creates a new AddDebugConfigParams object
// with the default values initialized.
func NewAddDebugConfigParams() *AddDebugConfigParams {
	var ()
	return &AddDebugConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddDebugConfigParamsWithTimeout creates a new AddDebugConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddDebugConfigParamsWithTimeout(timeout time.Duration) *AddDebugConfigParams {
	var ()
	return &AddDebugConfigParams{

		timeout: timeout,
	}
}

// NewAddDebugConfigParamsWithContext creates a new AddDebugConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddDebugConfigParamsWithContext(ctx context.Context) *AddDebugConfigParams {
	var ()
	return &AddDebugConfigParams{

		Context: ctx,
	}
}

// NewAddDebugConfigParamsWithHTTPClient creates a new AddDebugConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddDebugConfigParamsWithHTTPClient(client *http.Client) *AddDebugConfigParams {
	var ()
	return &AddDebugConfigParams{
		HTTPClient: client,
	}
}

/*AddDebugConfigParams contains all the parameters to send to the API endpoint
for the add debug config operation typically these are written to a http.Request
*/
type AddDebugConfigParams struct {

	/*Body
	  DebugConfig object that needs to be added

	*/
	Body *models.DebugConfig

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add debug config params
func (o *AddDebugConfigParams) WithTimeout(timeout time.Duration) *AddDebugConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add debug config params
func (o *AddDebugConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add debug config params
func (o *AddDebugConfigParams) WithContext(ctx context.Context) *AddDebugConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add debug config params
func (o *AddDebugConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add debug config params
func (o *AddDebugConfigParams) WithHTTPClient(client *http.Client) *AddDebugConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add debug config params
func (o *AddDebugConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add debug config params
func (o *AddDebugConfigParams) WithBody(body *models.DebugConfig) *AddDebugConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add debug config params
func (o *AddDebugConfigParams) SetBody(body *models.DebugConfig) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddDebugConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.DebugConfig)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
