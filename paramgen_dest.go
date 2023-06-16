// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/conduitio/conduit-connector-sdk/cmd/paramgen

package grpcclient

import (
	sdk "github.com/conduitio/conduit-connector-sdk"
)

func (Config) Parameters() map[string]sdk.Parameter {
	return map[string]sdk.Parameter{
		"maxDowntime": {
			Default:     "10m",
			Description: "max downtime accepted for the server to be off.",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"rateLimit": {
			Default:     "0",
			Description: "the bandwidth limit in bytes/second, use \"0\" to disable rate limiting.",
			Type:        sdk.ParameterTypeInt,
			Validations: []sdk.Validation{
				sdk.ValidationGreaterThan{Value: -1},
			},
		},
		"reconnectDelay": {
			Default:     "5s",
			Description: "delay between each gRPC request retry.",
			Type:        sdk.ParameterTypeDuration,
			Validations: []sdk.Validation{},
		},
		"url": {
			Default:     "",
			Description: "url to gRPC server",
			Type:        sdk.ParameterTypeString,
			Validations: []sdk.Validation{
				sdk.ValidationRequired{},
			},
		},
	}
}
