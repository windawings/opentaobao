package feedflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemAdzoneRptdailylist 资源包分日数据查询
// taobao.feedflow.item.adzone.rptdailylist
//
// 资源包分日数据查询
func TaobaoFeedflowItemAdzoneRptdailylist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemAdzoneRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemAdzoneRptdailylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
