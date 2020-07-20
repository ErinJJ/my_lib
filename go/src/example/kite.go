/*
Autogenerated by kitool v3.7.14
DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
*/

package main

import (
	"context"

	"example/thrift_gen/webarch/kite/example"

	"code.byted.org/kite/endpoint"
	"code.byted.org/kite/kite"
	"code.byted.org/kite/kitutil"
)

func init() {
	kite.Processor = example.NewExampleServiceProcessor(&ExampleServiceHandler{})
}

type ExampleServiceHandler struct{}

func (h *ExampleServiceHandler) ExampleMethod(ctx context.Context, r *example.ExampleReq) (*example.ExampleResp, error) {
	resp, err := kite.KiteMW(mkExampleMethod())(kite.InitMethodContext(ctx, "ExampleMethod"), &KiteExampleReq{r})

	if resp == nil {
		return nil, err
	}
	return resp.(endpoint.KiteResponse).RealResponse().(*example.ExampleResp), err

}

type KiteExampleReq struct {
	*example.ExampleReq
}

func (kr *KiteExampleReq) IsSetBase() bool {
	if kr.ExampleReq == nil {
		return false
	}
	return kr.ExampleReq.IsSetBase()
}

func (kr *KiteExampleReq) GetBase() endpoint.KiteBase {
	if kr.ExampleReq == nil {
		return nil
	}
	return kitutil.ConvertRequestBase(kr.ExampleReq.GetBase())
}

func (kr *KiteExampleReq) RealRequest() interface{} {
	return kr.ExampleReq
}

type KiteExampleResp struct {
	*example.ExampleResp
}

func (kr *KiteExampleResp) GetBaseResp() endpoint.KiteBaseResp {
	if kr.ExampleResp != nil {
		if ret := kr.ExampleResp.GetBaseResp(); ret != nil {
			return ret
		}
	}
	return nil
}

func (kr *KiteExampleResp) RealResponse() interface{} {
	return kr.ExampleResp
}

func mkExampleMethod() endpoint.EndPoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoint.KiteRequest).RealRequest()
		r := req.(*example.ExampleReq)
		resp, err := ExampleMethod(ctx, r)
		if resp == nil {
			return nil, err
		}
		return &KiteExampleResp{resp}, err
	}
}