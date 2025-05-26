package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugcodeDrugfactoryGetencrptypk 获取加密公钥
// alibaba.alihealth.drugcode.drugfactory.getencrptypk
//
// 获取服务端给药厂用来加密的公钥
func AlibabaAlihealthDrugcodeDrugfactoryGetencrptypk(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIRequest, resp *drugtrace.AlibabaAlihealthDrugcodeDrugfactoryGetencrptypkAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
