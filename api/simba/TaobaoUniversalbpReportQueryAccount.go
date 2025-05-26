package simba

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/simba"
)

// TaobaoUniversalbpReportQueryAccount 账户报表查询
// taobao.universalbp.report.query.account
//
// 账户报表查询
func TaobaoUniversalbpReportQueryAccount(ctx context.Context, clt *core.SDKClient, req *simba.TaobaoUniversalbpReportQueryAccountAPIRequest, resp *simba.TaobaoUniversalbpReportQueryAccountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
