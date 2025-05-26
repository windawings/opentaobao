package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaLegalSuitCourtOpenPush 开庭信息推送接口
// alibaba.legal.suit.court.open.push
//
// 供ISV推送开庭信息给集团诉讼
func AlibabaLegalSuitCourtOpenPush(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalSuitCourtOpenPushAPIRequest, resp *legalsuit.AlibabaLegalSuitCourtOpenPushAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
