package icbu

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/icbu"
)

// AlibabaIcbuProductGroupGet 分组信息获取
// alibaba.icbu.product.group.get
//
// 分组信息获取
func AlibabaIcbuProductGroupGet(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuProductGroupGetAPIRequest, resp *icbu.AlibabaIcbuProductGroupGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
