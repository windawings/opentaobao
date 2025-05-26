package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterModify 修改成本中心
// alitrip.btrip.open.cost.center.modify
//
// 修改成本中心
func AlitripBtripOpenCostCenterModify(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterModifyAPIRequest, resp *btrip.AlitripBtripOpenCostCenterModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
