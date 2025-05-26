package cainiaocntec

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cainiaocntec"
)

// CainiaoCntecLocallifeCommunitylifeSyncorderstatus 订单状态推送
// cainiao.cntec.locallife.communitylife.syncorderstatus
//
// 订单状态推送
func CainiaoCntecLocallifeCommunitylifeSyncorderstatus(ctx context.Context, clt *core.SDKClient, req *cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIRequest, resp *cainiaocntec.CainiaoCntecLocallifeCommunitylifeSyncorderstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
