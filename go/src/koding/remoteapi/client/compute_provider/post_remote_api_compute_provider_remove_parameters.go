package compute_provider

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIComputeProviderRemoveParams creates a new PostRemoteAPIComputeProviderRemoveParams object
// with the default values initialized.
func NewPostRemoteAPIComputeProviderRemoveParams() *PostRemoteAPIComputeProviderRemoveParams {
	var ()
	return &PostRemoteAPIComputeProviderRemoveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIComputeProviderRemoveParamsWithTimeout creates a new PostRemoteAPIComputeProviderRemoveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIComputeProviderRemoveParamsWithTimeout(timeout time.Duration) *PostRemoteAPIComputeProviderRemoveParams {
	var ()
	return &PostRemoteAPIComputeProviderRemoveParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIComputeProviderRemoveParamsWithContext creates a new PostRemoteAPIComputeProviderRemoveParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIComputeProviderRemoveParamsWithContext(ctx context.Context) *PostRemoteAPIComputeProviderRemoveParams {
	var ()
	return &PostRemoteAPIComputeProviderRemoveParams{

		Context: ctx,
	}
}

/*PostRemoteAPIComputeProviderRemoveParams contains all the parameters to send to the API endpoint
for the post remote API compute provider remove operation typically these are written to a http.Request
*/
type PostRemoteAPIComputeProviderRemoveParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) WithTimeout(timeout time.Duration) *PostRemoteAPIComputeProviderRemoveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) WithContext(ctx context.Context) *PostRemoteAPIComputeProviderRemoveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIComputeProviderRemoveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API compute provider remove params
func (o *PostRemoteAPIComputeProviderRemoveParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIComputeProviderRemoveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
