package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytCodereplace 单码替换
// alibaba.alihealth.drug.kyt.codereplace
//
// 单码替换
func AlibabaAlihealthDrugKytCodereplace(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytCodereplaceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
