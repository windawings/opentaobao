package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytGetdruglicense 获取药品资质信息
// alibaba.alihealth.drug.kyt.getdruglicense
//
// 获取药品的资质信息。
func AlibabaAlihealthDrugKytGetdruglicense(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytGetdruglicenseAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytGetdruglicenseAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
