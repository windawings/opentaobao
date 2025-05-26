package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaAelophyOrderDesensitizephoneGet 获取订单隐私号
// alibaba.aelophy.order.desensitizephone.get
//
// 获取订单隐私号
func AlibabaAelophyOrderDesensitizephoneGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyOrderDesensitizephoneGetAPIRequest, resp *wdk.AlibabaAelophyOrderDesensitizephoneGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
