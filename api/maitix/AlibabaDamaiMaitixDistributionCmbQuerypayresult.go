package maitix

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixDistributionCmbQuerypayresult 查询招行支付状态api
// alibaba.damai.maitix.distribution.cmb.querypayresult
//
// queryPayResult
func AlibabaDamaiMaitixDistributionCmbQuerypayresult(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIRequest, resp *maitix.AlibabaDamaiMaitixDistributionCmbQuerypayresultAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
