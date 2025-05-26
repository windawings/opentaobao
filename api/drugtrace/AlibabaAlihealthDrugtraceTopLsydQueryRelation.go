package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopLsydQueryRelation 单码关联关系查询
// alibaba.alihealth.drugtrace.top.lsyd.query.relation
//
// 单码关联关系查询
func AlibabaAlihealthDrugtraceTopLsydQueryRelation(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopLsydQueryRelationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
