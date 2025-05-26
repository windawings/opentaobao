package tmallcar

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tmallcar"
)

// TmallCarFinanceStatusSync 汽车金融状态同步
// tmall.car.finance.status.sync
//
// 汽车金融状态同步
func TmallCarFinanceStatusSync(ctx context.Context, clt *core.SDKClient, req *tmallcar.TmallCarFinanceStatusSyncAPIRequest, resp *tmallcar.TmallCarFinanceStatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
