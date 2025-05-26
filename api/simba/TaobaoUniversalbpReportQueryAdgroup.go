package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryAdgroup 单元报表查询
// taobao.universalbp.report.query.adgroup
//
// 单元报表查询
func TaobaoUniversalbpReportQueryAdgroup(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAdgroupAPIRequest, resp *simba.TaobaoUniversalbpReportQueryAdgroupAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
