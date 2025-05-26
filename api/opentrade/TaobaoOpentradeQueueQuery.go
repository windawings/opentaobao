package opentrade

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/opentrade"
)

// TaobaoOpentradeQueueQuery 尖货交易排队信息查询
// taobao.opentrade.queue.query
//
// 尖货交易排队信息查询
func TaobaoOpentradeQueueQuery(ctx context.Context, clt *core.SDKClient, req *opentrade.TaobaoOpentradeQueueQueryAPIRequest, resp *opentrade.TaobaoOpentradeQueueQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
