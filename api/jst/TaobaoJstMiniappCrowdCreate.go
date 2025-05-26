package jst

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/jst"
)

// TaobaoJstMiniappCrowdCreate 小程序活动创建
// taobao.jst.miniapp.crowd.create
//
// 小程序活动创建
func TaobaoJstMiniappCrowdCreate(ctx context.Context, clt *core.SDKClient, req *jst.TaobaoJstMiniappCrowdCreateAPIRequest, resp *jst.TaobaoJstMiniappCrowdCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
