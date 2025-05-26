package servicecenter

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/servicecenter"
)

// TmallCarLeaseConsume 汽车租赁核销
// tmall.car.lease.consume
//
// 租赁公司回传信息，核销
func TmallCarLeaseConsume(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarLeaseConsumeAPIRequest, resp *servicecenter.TmallCarLeaseConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
