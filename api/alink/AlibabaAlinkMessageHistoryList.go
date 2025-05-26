package alink

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alink"
)

// AlibabaAlinkMessageHistoryList 查询消息列表
// alibaba.alink.message.history.list
//
// 查询消息列表
func AlibabaAlinkMessageHistoryList(ctx context.Context, clt *core.SDKClient, req *alink.AlibabaAlinkMessageHistoryListAPIRequest, resp *alink.AlibabaAlinkMessageHistoryListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
