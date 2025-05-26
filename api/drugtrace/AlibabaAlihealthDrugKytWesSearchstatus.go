package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytWesSearchstatus 单据处理状态查询
// alibaba.alihealth.drug.kyt.wes.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytWesSearchstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytWesSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytWesSearchstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
