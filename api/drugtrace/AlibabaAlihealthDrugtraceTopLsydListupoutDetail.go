package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydListupoutDetail 上游出库单单据明细查询
// alibaba.alihealth.drugtrace.top.lsyd.listupout.detail
//
// 查询上游出库单明细(带追溯码信息)
func AlibabaAlihealthDrugtraceTopLsydListupoutDetail(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydListupoutDetailAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
