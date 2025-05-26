package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryCreative 创意报表查询
// taobao.universalbp.report.query.creative
//
// 创意报表查询
func TaobaoUniversalbpReportQueryCreative(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryCreativeAPIRequest, resp *simba.TaobaoUniversalbpReportQueryCreativeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
