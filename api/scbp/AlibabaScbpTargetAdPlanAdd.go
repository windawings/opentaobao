package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanAdd 定向推广-新建计划
// alibaba.scbp.target.ad.plan.add
//
// 定向推广-新建单条计划
func AlibabaScbpTargetAdPlanAdd(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanAddAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
