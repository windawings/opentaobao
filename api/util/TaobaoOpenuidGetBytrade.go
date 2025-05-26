package util

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/util"
)

// TaobaoOpenuidGetBytrade 通过订单获取对应买家的openUID
// taobao.openuid.get.bytrade
//
// 通过订单获取对应买家的openUID,需要卖家授权
func TaobaoOpenuidGetBytrade(ctx context.Context, clt *core.SDKClient, req *util.TaobaoOpenuidGetBytradeAPIRequest, resp *util.TaobaoOpenuidGetBytradeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
