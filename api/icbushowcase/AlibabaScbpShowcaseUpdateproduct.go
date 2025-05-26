package icbushowcase

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbushowcase"
)

// AlibabaScbpShowcaseUpdateproduct 替换橱窗商品
// alibaba.scbp.showcase.updateproduct
//
// 替换橱窗商品
func AlibabaScbpShowcaseUpdateproduct(ctx context.Context, clt *core.SDKClient, req *icbushowcase.AlibabaScbpShowcaseUpdateproductAPIRequest, resp *icbushowcase.AlibabaScbpShowcaseUpdateproductAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
