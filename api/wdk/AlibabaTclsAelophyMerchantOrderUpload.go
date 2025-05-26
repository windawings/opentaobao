package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaTclsAelophyMerchantOrderUpload 商家订单数据上传
// alibaba.tcls.aelophy.merchant.order.upload
//
// 商家订单数据上传
func AlibabaTclsAelophyMerchantOrderUpload(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaTclsAelophyMerchantOrderUploadAPIRequest, resp *wdk.AlibabaTclsAelophyMerchantOrderUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
