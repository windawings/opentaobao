package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpAdGroupFindForbiddenProduct 查询屏蔽品
// alibaba.scbp.ad.group.find.forbidden.product
//
// 查询屏蔽品
func AlibabaScbpAdGroupFindForbiddenProduct(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpAdGroupFindForbiddenProductAPIRequest, resp *scbp.AlibabaScbpAdGroupFindForbiddenProductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
