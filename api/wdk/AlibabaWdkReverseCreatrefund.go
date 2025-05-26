package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkReverseCreatrefund 逆向提交
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
func AlibabaWdkReverseCreatrefund(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkReverseCreatrefundAPIRequest, resp *wdk.AlibabaWdkReverseCreatrefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
