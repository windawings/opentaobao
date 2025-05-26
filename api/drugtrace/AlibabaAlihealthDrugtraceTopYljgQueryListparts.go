package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgQueryListparts 往来单位查询
// alibaba.alihealth.drugtrace.top.yljg.query.listparts
//
// 查询往来单位列表
func AlibabaAlihealthDrugtraceTopYljgQueryListparts(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryListpartsAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgQueryListpartsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
