package ascpffo

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ascpffo"
)

// AliexpressAscpWarehouseInventoryQuery AliExpress在仓库存查询API
// aliexpress.ascp.warehouse.inventory.query
//
// AliExpress在仓库存查询查询API
func AliexpressAscpWarehouseInventoryQuery(ctx context.Context, clt *core.SDKClient, req *ascpffo.AliexpressAscpWarehouseInventoryQueryAPIRequest, resp *ascpffo.AliexpressAscpWarehouseInventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
