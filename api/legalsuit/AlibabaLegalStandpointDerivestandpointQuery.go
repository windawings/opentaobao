package legalsuit

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/legalsuit"
)

// AlibabaLegalStandpointDerivestandpointQuery 查询衍生口径
// alibaba.legal.standpoint.derivestandpoint.query
//
// 查询衍生口径
func AlibabaLegalStandpointDerivestandpointQuery(ctx context.Context, clt *core.SDKClient, req *legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIRequest, resp *legalsuit.AlibabaLegalStandpointDerivestandpointQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
