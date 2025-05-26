package aliqin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aliqin"
)

// AlibabaAliqinFcIvrNumCall ivr呼叫
// alibaba.aliqin.fc.ivr.num.call
//
// ivr呼叫
func AlibabaAliqinFcIvrNumCall(ctx context.Context, clt *core.SDKClient, req *aliqin.AlibabaAliqinFcIvrNumCallAPIRequest, resp *aliqin.AlibabaAliqinFcIvrNumCallAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
