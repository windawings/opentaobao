package ship

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/ship"
)

// AlitripShipProductSyncnunber 船票班次变更回调
// alitrip.ship.product.syncnunber
//
// 船票班次变更回调
func AlitripShipProductSyncnunber(ctx context.Context, clt *core.SDKClient, req *ship.AlitripShipProductSyncnunberAPIRequest, resp *ship.AlitripShipProductSyncnunberAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
