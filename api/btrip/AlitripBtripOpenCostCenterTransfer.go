package btrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/btrip"
)

// AlitripBtripOpenCostCenterTransfer 商旅成本中心转换为外部成本中心
// alitrip.btrip.open.cost.center.transfer
//
// 商旅成本中心转换为外部成本中心
func AlitripBtripOpenCostCenterTransfer(ctx context.Context, clt *core.SDKClient, req *btrip.AlitripBtripOpenCostCenterTransferAPIRequest, resp *btrip.AlitripBtripOpenCostCenterTransferAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
