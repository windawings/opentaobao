package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripCorpopApplyApprove 【商旅】更新审批单状态
// alitrip.btrip.corpop.apply.approve
//
// 【商旅】更新审批单状态
func AlitripBtripCorpopApplyApprove(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripCorpopApplyApproveAPIRequest, resp *btrip.AlitripBtripCorpopApplyApproveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
