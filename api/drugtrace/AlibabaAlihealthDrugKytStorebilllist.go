package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytStorebilllist 零售端平台单据查询
// alibaba.alihealth.drug.kyt.storebilllist
//
// 零售端平台单据查询
func AlibabaAlihealthDrugKytStorebilllist(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytStorebilllistAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytStorebilllistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
