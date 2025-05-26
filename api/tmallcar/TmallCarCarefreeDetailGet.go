package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallCarCarefreeDetailGet 查询业务单信息
// tmall.car.carefree.detail.get
//
// 查询业务单信息
func TmallCarCarefreeDetailGet(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarCarefreeDetailGetAPIRequest, resp *tmallcar.TmallCarCarefreeDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
