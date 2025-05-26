package maitix

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionDeliveryQuery 查询分销物流单
// alibaba.damai.maitix.distribution.delivery.query
//
// 渠道查询物流订单
func AlibabaDamaiMaitixDistributionDeliveryQuery(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionDeliveryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
