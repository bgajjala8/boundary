// Code generated by "make api"; DO NOT EDIT.
// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hosts

import (
	"strconv"
	"strings"

	"github.com/hashicorp/boundary/api"
)

// Option is a func that sets optional attributes for a call. This does not need
// to be used directly, but instead option arguments are built from the
// functions in this package. WithX options set a value to that given in the
// argument; DefaultX options indicate that the value should be set to its
// default. When an API call is made options are processed in the order they
// appear in the function call, so for a given argument X, a succession of WithX
// or DefaultX calls will result in the last call taking effect.
type Option func(*options)

type options struct {
	postMap                      map[string]any
	queryMap                     map[string]string
	withAutomaticVersioning      bool
	withSkipCurlOutput           bool
	withFilter                   string
	withListToken                string
	withClientDirectedPagination bool
	withPageSize                 uint32
	withResourcePathOverride     string
}

func getDefaultOptions() options {
	return options{
		postMap:  make(map[string]any),
		queryMap: make(map[string]string),
	}
}

func getOpts(opt ...Option) (options, []api.Option) {
	opts := getDefaultOptions()
	for _, o := range opt {
		if o != nil {
			o(&opts)
		}
	}
	var apiOpts []api.Option
	if opts.withSkipCurlOutput {
		apiOpts = append(apiOpts, api.WithSkipCurlOutput(true))
	}
	if opts.withFilter != "" {
		opts.queryMap["filter"] = opts.withFilter
	}
	if opts.withListToken != "" {
		opts.queryMap["list_token"] = opts.withListToken
	}
	if opts.withPageSize != 0 {
		opts.queryMap["page_size"] = strconv.FormatUint(uint64(opts.withPageSize), 10)
	}
	return opts, apiOpts
}

// If set, and if the version is zero during an update, the API will perform a
// fetch to get the current version of the resource and populate it during the
// update call. This is convenient but opens up the possibility for subtle
// order-of-modification issues, so use carefully.
func WithAutomaticVersioning(enable bool) Option {
	return func(o *options) {
		o.withAutomaticVersioning = enable
	}
}

// WithSkipCurlOutput tells the API to not use the current call for cURL output.
// Useful for when we need to look up versions.
func WithSkipCurlOutput(skip bool) Option {
	return func(o *options) {
		o.withSkipCurlOutput = skip
	}
}

// WithListToken tells the API to use the provided list token
// for listing operations on this resource.
func WithListToken(listToken string) Option {
	return func(o *options) {
		o.withListToken = listToken
	}
}

// WithFilter tells the API to filter the items returned using the provided
// filter term.  The filter should be in a format supported by
// hashicorp/go-bexpr.
func WithFilter(filter string) Option {
	return func(o *options) {
		o.withFilter = strings.TrimSpace(filter)
	}
}

// WithClientDirectedPagination tells the List function to return only the first
// page, if more pages are available
func WithClientDirectedPagination(with bool) Option {
	return func(o *options) {
		o.withClientDirectedPagination = with
	}
}

// WithPageSize controls the size of pages used during List
func WithPageSize(with uint32) Option {
	return func(o *options) {
		o.withPageSize = with
	}
}

// WithResourcePathOverride tells the API to use the provided resource path
func WithResourcePathOverride(path string) Option {
	return func(o *options) {
		o.withResourcePathOverride = path
	}
}

func WithStaticHostAddress(inAddress string) Option {
	return func(o *options) {
		raw, ok := o.postMap["attributes"]
		if !ok {
			raw = any(map[string]any{})
		}
		val := raw.(map[string]any)
		val["address"] = inAddress
		o.postMap["attributes"] = val
	}
}

func DefaultStaticHostAddress() Option {
	return func(o *options) {
		raw, ok := o.postMap["attributes"]
		if !ok {
			raw = any(map[string]any{})
		}
		val := raw.(map[string]any)
		val["address"] = nil
		o.postMap["attributes"] = val
	}
}

func WithAttributes(inAttributes map[string]interface{}) Option {
	return func(o *options) {
		o.postMap["attributes"] = inAttributes
	}
}

func DefaultAttributes() Option {
	return func(o *options) {
		o.postMap["attributes"] = nil
	}
}

func WithDescription(inDescription string) Option {
	return func(o *options) {
		o.postMap["description"] = inDescription
	}
}

func DefaultDescription() Option {
	return func(o *options) {
		o.postMap["description"] = nil
	}
}

func WithName(inName string) Option {
	return func(o *options) {
		o.postMap["name"] = inName
	}
}

func DefaultName() Option {
	return func(o *options) {
		o.postMap["name"] = nil
	}
}
