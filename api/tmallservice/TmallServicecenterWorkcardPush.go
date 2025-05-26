package tmallservice

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardPush 推送服务工单信息
// tmall.servicecenter.workcard.push
//
// 服务商家推送工单信息到天猫。
func TmallServicecenterWorkcardPush(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardPushAPIRequest, resp *tmallservice.TmallServicecenterWorkcardPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
