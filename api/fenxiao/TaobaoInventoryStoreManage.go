package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoInventoryStoreManage 创建或更新仓库
// taobao.inventory.store.manage
//
// 创建商家仓或者更新商家仓信息
func TaobaoInventoryStoreManage(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryStoreManageAPIRequest, resp *fenxiao.TaobaoInventoryStoreManageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
