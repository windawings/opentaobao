package drugtrace

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/drugtrace"
)

// AlibabaAlihealthDrugKytScqyPutpackagebind 码拼箱建立父子关系接口
// alibaba.alihealth.drug.kyt.scqy.putpackagebind
//
// 码拼箱建立父子关系接口
func AlibabaAlihealthDrugKytScqyPutpackagebind(ctx context.Context, clt *core.SDKClient, req *drugtrace.AlibabaAlihealthDrugKytScqyPutpackagebindAPIRequest, resp *drugtrace.AlibabaAlihealthDrugKytScqyPutpackagebindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
