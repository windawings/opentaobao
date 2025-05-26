package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallCarLeaseOrderidGet 天猫开新车查询订单id
// tmall.car.lease.orderid.get
//
// 天猫开新车查询订单id
func TmallCarLeaseOrderidGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeaseOrderidGetAPIRequest, resp *tmallcar.TmallCarLeaseOrderidGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
