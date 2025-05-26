package traderate

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/traderate"
)

// TaobaoTraderatesGet 搜索评价信息
// taobao.traderates.get
//
// 搜索评价信息，只能获取距今180天内的评价记录(只支持查询卖家给出或得到的评价)。
func TaobaoTraderatesGet(ctx context.Context, clt *core.SDKClient, req *traderate.TaobaoTraderatesGetAPIRequest, resp *traderate.TaobaoTraderatesGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
