package main
import (
	"context"
	"code.byted.org/gopkg/logs"
	"example/thrift_gen/webarch/kite/example"
	"example/thrift_gen/base"
)
// Implement ExampleMethod Method.
func ExampleMethod(ctx context.Context, r *example.ExampleReq) (*example.ExampleResp, error) {
	logs.Infof("ExampleMethod request: %+v", r)
	resp := example.NewExampleResp()
	resp.BaseResp = base.NewBaseResp()
	resp.Msg = "Hi! " + r.GetMsg()
	return resp, nil
}