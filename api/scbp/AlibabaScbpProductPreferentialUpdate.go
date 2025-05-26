package scbp

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/scbp"
)

// AlibabaScbpProductPreferentialUpdate 设置P4P产品优先推广状态
// alibaba.scbp.product.preferential.update
//
// 设置P4P产品优先推广状态
func AlibabaScbpProductPreferentialUpdate(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpProductPreferentialUpdateAPIRequest, resp *scbp.AlibabaScbpProductPreferentialUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
