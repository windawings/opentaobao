package icbu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupAdd 增加商品分组
// alibaba.icbu.product.group.add
//
// 增加商品分组
func AlibabaIcbuProductGroupAdd(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupAddAPIRequest, resp *icbu.AlibabaIcbuProductGroupAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
