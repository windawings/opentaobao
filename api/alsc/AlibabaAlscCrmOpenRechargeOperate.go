package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenRechargeOperate 储值操作接口
// alibaba.alsc.crm.open.recharge.operate
//
// 储值操作接口
func AlibabaAlscCrmOpenRechargeOperate(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenRechargeOperateAPIRequest, resp *alsc.AlibabaAlscCrmOpenRechargeOperateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
