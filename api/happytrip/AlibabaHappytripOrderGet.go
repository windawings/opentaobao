package happytrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/happytrip"
)

// AlibabaHappytripOrderGet 获取欢行统一订单模型
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
func AlibabaHappytripOrderGet(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripOrderGetAPIRequest, resp *happytrip.AlibabaHappytripOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
