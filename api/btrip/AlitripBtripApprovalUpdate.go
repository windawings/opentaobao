package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripApprovalUpdate 更新审批单
// alitrip.btrip.approval.update
//
// 更新审批单
func AlitripBtripApprovalUpdate(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripApprovalUpdateAPIRequest, resp *btrip.AlitripBtripApprovalUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
