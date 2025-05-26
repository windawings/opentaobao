package qianniu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskRemove 轻任务删除接口
// taobao.qianniu.task.remove
//
// 轻任务删除接口。
func TaobaoQianniuTaskRemove(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskRemoveAPIRequest, resp *qianniu.TaobaoQianniuTaskRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
