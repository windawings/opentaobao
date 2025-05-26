package tmallsc

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallsc"
)

// AlibabaSscBusinessServicepriceQuery 苹果查询服务供给报价
// alibaba.ssc.business.serviceprice.query
//
// 苹果查询服务供给报价
func AlibabaSscBusinessServicepriceQuery(ctx context.Context, clt *core.SDKClient, req *tmallsc.AlibabaSscBusinessServicepriceQueryAPIRequest, resp *tmallsc.AlibabaSscBusinessServicepriceQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
