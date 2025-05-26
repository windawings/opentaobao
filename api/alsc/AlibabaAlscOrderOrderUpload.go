package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscOrderOrderUpload 订单回流
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
func AlibabaAlscOrderOrderUpload(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscOrderOrderUploadAPIRequest, resp *alsc.AlibabaAlscOrderOrderUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
