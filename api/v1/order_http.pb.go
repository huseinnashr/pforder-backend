// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.1
// - protoc             (unknown)
// source: v1/order.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAccountServiceListOrder = "/v1.AccountService/ListOrder"

type AccountServiceHTTPServer interface {
	ListOrder(context.Context, *ListOrderRequest) (*ListOrderResponse, error)
}

func RegisterAccountServiceHTTPServer(s *http.Server, srv AccountServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/orders:list", _AccountService_ListOrder0_HTTP_Handler(srv))
}

func _AccountService_ListOrder0_HTTP_Handler(srv AccountServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListOrderRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAccountServiceListOrder)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListOrder(ctx, req.(*ListOrderRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListOrderResponse)
		return ctx.Result(200, reply)
	}
}

type AccountServiceHTTPClient interface {
	ListOrder(ctx context.Context, req *ListOrderRequest, opts ...http.CallOption) (rsp *ListOrderResponse, err error)
}

type AccountServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewAccountServiceHTTPClient(client *http.Client) AccountServiceHTTPClient {
	return &AccountServiceHTTPClientImpl{client}
}

func (c *AccountServiceHTTPClientImpl) ListOrder(ctx context.Context, in *ListOrderRequest, opts ...http.CallOption) (*ListOrderResponse, error) {
	var out ListOrderResponse
	pattern := "/orders:list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAccountServiceListOrder))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
