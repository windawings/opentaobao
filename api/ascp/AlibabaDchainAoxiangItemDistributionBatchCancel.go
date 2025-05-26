package ascp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascp"
)

// AlibabaDchainAoxiangItemDistributionBatchCancel 取消商品分销
// alibaba.dchain.aoxiang.item.distribution.batch.cancel
//
// 取消商品分销
func AlibabaDchainAoxiangItemDistributionBatchCancel(ctx context.Context, clt *core.SDKClient, req *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIRequest, resp *ascp.AlibabaDchainAoxiangItemDistributionBatchCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
