package tbk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbk"
)

// TaobaoTbkDgCpaActivityReport 淘宝客-推广者-任务奖励效果报表
// taobao.tbk.dg.cpa.activity.report
//
// 提供给媒体使用的cpa活动报表查询api，当前仅试运行媒体可使用
func TaobaoTbkDgCpaActivityReport(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkDgCpaActivityReportAPIRequest, resp *tbk.TaobaoTbkDgCpaActivityReportAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
