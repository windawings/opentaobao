package logistic

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/logistic"
)

// AlibabaEleFengniaoCancelMerchant 商户取消
// alibaba.ele.fengniao.cancel.merchant
//
// 商户取消配送
func AlibabaEleFengniaoCancelMerchant(ctx context.Context, clt *core.SDKClient, req *logistic.AlibabaEleFengniaoCancelMerchantAPIRequest, resp *logistic.AlibabaEleFengniaoCancelMerchantAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
