package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrSearchstatus 疫苗企业上传单据后处理状态查询
// alibaba.alihealth.drug.kyt.dr.searchstatus
//
// 单据处理状态查询
func AlibabaAlihealthDrugKytDrSearchstatus(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrSearchstatusAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrSearchstatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
