package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaWdkUmsHandlingGet 加工单-回流单（新接口）
// alibaba.wdk.ums.handling.get
//
// 加工单-回流单（新接口）
func AlibabaWdkUmsHandlingGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkUmsHandlingGetAPIRequest, resp *wdk.AlibabaWdkUmsHandlingGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
