package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmCustomerUpdate 更新顾客信息
// alibaba.alsc.crm.customer.update
//
// 更新顾客信息
func AlibabaAlscCrmCustomerUpdate(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmCustomerUpdateAPIRequest, resp *alsc.AlibabaAlscCrmCustomerUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
