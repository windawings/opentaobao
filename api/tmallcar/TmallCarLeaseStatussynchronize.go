package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallCarLeaseStatussynchronize 天猫开新车租后状态同步
// tmall.car.lease.statussynchronize
//
// 天猫开新车租后状态同步
func TmallCarLeaseStatussynchronize(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarLeaseStatussynchronizeAPIRequest, resp *tmallcar.TmallCarLeaseStatussynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
