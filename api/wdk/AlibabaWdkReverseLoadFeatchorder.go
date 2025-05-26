package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkReverseLoadFeatchorder 取货详情
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
func AlibabaWdkReverseLoadFeatchorder(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseLoadFeatchorderAPIRequest, resp *wdk.AlibabaWdkReverseLoadFeatchorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
