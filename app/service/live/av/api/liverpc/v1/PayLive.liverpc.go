// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/PayLive.proto

/*
Package v1 is a generated liverpc stub package.
This code was generated with go-common/app/tool/liverpc/protoc-gen-liverpc v0.1.

It is generated from these files:
	v1/PayLive.proto
	v1/Pk.proto
*/
package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports
// Imports only used by utility functions:

// =================
// PayLive Interface
// =================

type PayLive interface {
	// * 付费直播鉴权
	//
	LiveValidate(context.Context, *PayLiveLiveValidateReq) (*PayLiveLiveValidateResp, error)
}

// =======================
// PayLive Live Rpc Client
// =======================

type payLiveRpcClient struct {
	client *liverpc.Client
}

// NewPayLiveRpcClient creates a Rpc client that implements the PayLive interface.
// It communicates using Rpc and can be configured with a custom HTTPClient.
func NewPayLiveRpcClient(client *liverpc.Client) PayLive {
	return &payLiveRpcClient{
		client: client,
	}
}

func (c *payLiveRpcClient) LiveValidate(ctx context.Context, in *PayLiveLiveValidateReq) (*PayLiveLiveValidateResp, error) {
	out := new(PayLiveLiveValidateResp)
	err := doRpcRequest(ctx, c.client, 1, "PayLive.liveValidate", in, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// =====
// Utils
// =====

func doRpcRequest(ctx context.Context, client *liverpc.Client, version int, method string, in, out proto.Message) (err error) {
	err = client.Call(ctx, version, method, in, out)
	return
}
