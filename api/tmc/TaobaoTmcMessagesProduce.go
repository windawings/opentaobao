package tmc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmc"
)

// TaobaoTmcMessagesProduce 批量发送消息
// taobao.tmc.messages.produce
//
// 批量发送消息
func TaobaoTmcMessagesProduce(ctx context.Context, clt *core.SDKClient, req *tmc.TaobaoTmcMessagesProduceAPIRequest, resp *tmc.TaobaoTmcMessagesProduceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
