package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytDrugtable 查询药品目录信息
// alibaba.alihealth.drug.kyt.drugtable
//
// 查询药品目录信息
func AlibabaAlihealthDrugKytDrugtable(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytDrugtableAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytDrugtableAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
