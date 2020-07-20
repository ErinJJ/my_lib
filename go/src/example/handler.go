package main

import (
	"code.byted.org/gopkg/logs"
	"context"
	"example/thrift_gen/base"
	"example/thrift_gen/webarch/kite/example"
)

// Implement ExampleMethod Method.
func ExampleMethod(ctx context.Context, r *example.ExampleReq) (*example.ExampleResp, error) {
	logs.Infof("ExampleMethod request: %+v", r)
	resp := example.NewExampleResp()
	resp.BaseResp = base.NewBaseResp()
	resp.Msg = "Hi! " + r.GetMsg()
	return resp, nil
}
