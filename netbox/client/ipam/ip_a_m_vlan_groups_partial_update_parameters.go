// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/hosting-de-labs/go-netbox/netbox/models"
)

// NewIPAMVlanGroupsPartialUpdateParams creates a new IPAMVlanGroupsPartialUpdateParams object
// with the default values initialized.
func NewIPAMVlanGroupsPartialUpdateParams() *IPAMVlanGroupsPartialUpdateParams {
	var ()
	return &IPAMVlanGroupsPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVlanGroupsPartialUpdateParamsWithTimeout creates a new IPAMVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVlanGroupsPartialUpdateParamsWithTimeout(timeout time.Duration) *IPAMVlanGroupsPartialUpdateParams {
	var ()
	return &IPAMVlanGroupsPartialUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMVlanGroupsPartialUpdateParamsWithContext creates a new IPAMVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVlanGroupsPartialUpdateParamsWithContext(ctx context.Context) *IPAMVlanGroupsPartialUpdateParams {
	var ()
	return &IPAMVlanGroupsPartialUpdateParams{

		Context: ctx,
	}
}

// NewIPAMVlanGroupsPartialUpdateParamsWithHTTPClient creates a new IPAMVlanGroupsPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVlanGroupsPartialUpdateParamsWithHTTPClient(client *http.Client) *IPAMVlanGroupsPartialUpdateParams {
	var ()
	return &IPAMVlanGroupsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMVlanGroupsPartialUpdateParams contains all the parameters to send to the API endpoint
for the ipam vlan groups partial update operation typically these are written to a http.Request
*/
type IPAMVlanGroupsPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableVLANGroup
	/*ID
	  A unique integer value identifying this VLAN group.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) WithTimeout(timeout time.Duration) *IPAMVlanGroupsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) WithContext(ctx context.Context) *IPAMVlanGroupsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) WithHTTPClient(client *http.Client) *IPAMVlanGroupsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) WithData(data *models.WritableVLANGroup) *IPAMVlanGroupsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) SetData(data *models.WritableVLANGroup) {
	o.Data = data
}

// WithID adds the id to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) WithID(id int64) *IPAMVlanGroupsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam vlan groups partial update params
func (o *IPAMVlanGroupsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVlanGroupsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
