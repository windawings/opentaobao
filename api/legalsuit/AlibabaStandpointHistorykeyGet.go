package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaStandpointHistorykeyGet 查询历史数据
// alibaba.standpoint.historykey.get
//
// 查询历史数据
func AlibabaStandpointHistorykeyGet(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaStandpointHistorykeyGetAPIRequest, resp *legalsuit.AlibabaStandpointHistorykeyGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
