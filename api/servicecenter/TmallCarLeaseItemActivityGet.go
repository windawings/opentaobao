package servicecenter

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/servicecenter"
)

// TmallCarLeaseItemActivityGet 查询汽车租赁活动信息
// tmall.car.lease.item.activity.get
//
// 查询汽车租赁活动信息
func TmallCarLeaseItemActivityGet(ctx context.Context, clt *core.SDKClient, req *servicecenter.TmallCarLeaseItemActivityGetAPIRequest, resp *servicecenter.TmallCarLeaseItemActivityGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
