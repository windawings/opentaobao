package alsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alsc"
)

// AlibabaAlscCrmRechargeQryrule 储值规则下行
// alibaba.alsc.crm.recharge.qryrule
//
// 储值规则下行
func AlibabaAlscCrmRechargeQryrule(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRechargeQryruleAPIRequest, resp *alsc.AlibabaAlscCrmRechargeQryruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
