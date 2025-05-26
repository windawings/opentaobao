package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductQuantityUpdate 渠道无仓库存更新接口
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
func TmallSupplychainChannelProductQuantityUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityUpdateAPIRequest, resp *fenxiao.TmallSupplychainChannelProductQuantityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
