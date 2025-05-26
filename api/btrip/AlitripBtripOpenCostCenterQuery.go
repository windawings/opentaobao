package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterQuery 查询成本中心
// alitrip.btrip.open.cost.center.query
//
// 查询成本中心
func AlitripBtripOpenCostCenterQuery(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterQueryAPIRequest, resp *btrip.AlitripBtripOpenCostCenterQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
