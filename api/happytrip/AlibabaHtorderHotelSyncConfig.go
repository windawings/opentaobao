package happytrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/happytrip"
)

// AlibabaHtorderHotelSyncConfig 同步配置信息
// alibaba.htorder.hotel.sync.config
//
// 同步配置信息
func AlibabaHtorderHotelSyncConfig(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHtorderHotelSyncConfigAPIRequest, resp *happytrip.AlibabaHtorderHotelSyncConfigAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
