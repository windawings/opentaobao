package qianniu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qianniu"
)

// TaobaoQianniuBuyerTagGet 判断买家是否有某些标
// taobao.qianniu.buyer.tag.get
//
// 判断某个买家是否有某些标
func TaobaoQianniuBuyerTagGet(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuBuyerTagGetAPIRequest, resp *qianniu.TaobaoQianniuBuyerTagGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
