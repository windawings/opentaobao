package qianniu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/qianniu"
)

// TaobaoQianniuTaskCancel 取消轻任务
// taobao.qianniu.task.cancel
//
// 由任务发起者调用
func TaobaoQianniuTaskCancel(ctx context.Context, clt *core.SDKClient, req *qianniu.TaobaoQianniuTaskCancelAPIRequest, resp *qianniu.TaobaoQianniuTaskCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
