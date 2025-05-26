package alihealthoutflow

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/alihealthoutflow"
)

// AlibabaAlihealthRecommendMixcardinfoGet 快应用混合卡片信息
// alibaba.alihealth.recommend.mixcardinfo.get
//
// 快应用混合卡片信息
func AlibabaAlihealthRecommendMixcardinfoGet(ctx context.Context, clt *core.SDKClient, req *alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIRequest, resp *alihealthoutflow.AlibabaAlihealthRecommendMixcardinfoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
