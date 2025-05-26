package tbrefund

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbrefund"
)

// TaobaoRefundRefusereasonGet 获取拒绝原因列表
// taobao.refund.refusereason.get
//
// 获取商家拒绝原因列表
func TaobaoRefundRefusereasonGet(ctx context.Context, clt *core.SDKClient, req *tbrefund.TaobaoRefundRefusereasonGetAPIRequest, resp *tbrefund.TaobaoRefundRefusereasonGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
