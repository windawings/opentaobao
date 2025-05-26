package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugtraceTopYljgDrugtable 查询药品目录信息
// alibaba.alihealth.drugtrace.top.yljg.drugtable
//
// 查询药品目录信息
func AlibabaAlihealthDrugtraceTopYljgDrugtable(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugtraceTopYljgDrugtableAPIRequest, resp *drugtrace.AlibabaAlihealthDrugtraceTopYljgDrugtableAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
