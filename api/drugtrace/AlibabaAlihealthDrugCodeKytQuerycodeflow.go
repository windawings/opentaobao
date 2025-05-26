package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugCodeKytQuerycodeflow 码流向查询
// alibaba.alihealth.drug.code.kyt.querycodeflow
//
// 追溯码流向查询
func AlibabaAlihealthDrugCodeKytQuerycodeflow(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeflowAPIRequest, resp *drugtrace.AlibabaAlihealthDrugCodeKytQuerycodeflowAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
