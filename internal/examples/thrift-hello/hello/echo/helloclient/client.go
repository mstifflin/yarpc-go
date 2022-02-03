// Code generated by thriftrw-plugin-yarpc
// @generated

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package helloclient

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	echo "go.uber.org/yarpc/internal/examples/thrift-hello/hello/echo"
	reflect "reflect"
)

// Interface is a client for the Hello service.
type Interface interface {
	Echo(
		ctx context.Context,
		Echo *echo.EchoRequest,
		opts ...yarpc.CallOption,
	) (*echo.EchoResponse, error)
}

// New builds a new client for the Hello service.
//
// 	client := helloclient.New(dispatcher.ClientConfig("hello"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Hello",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "Hello",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) Echo(
	ctx context.Context,
	_Echo *echo.EchoRequest,
	opts ...yarpc.CallOption,
) (success *echo.EchoResponse, err error) {

	var result echo.Hello_Echo_Result
	args := echo.Hello_Echo_Helper.Args(_Echo)

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = echo.Hello_Echo_Helper.UnwrapResponse(&result)
	return
}
