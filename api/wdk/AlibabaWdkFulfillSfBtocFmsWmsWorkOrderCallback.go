package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallback 顺丰仓作业回传
// alibaba.wdk.fulfill.sf.btoc.fms.wms.work.order.callback
//
// 顺丰仓作业单回传接口
func AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallback(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIRequest, resp *wdk.AlibabaWdkFulfillSfBtocFmsWmsWorkOrderCallbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
