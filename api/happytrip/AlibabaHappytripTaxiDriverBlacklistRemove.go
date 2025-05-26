package happytrip

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/happytrip"
)

// AlibabaHappytripTaxiDriverBlacklistRemove 移除司机黑名单
// alibaba.happytrip.taxi.driver.blacklist.remove
//
// 移除司机黑名单
func AlibabaHappytripTaxiDriverBlacklistRemove(ctx context.Context, clt *core.SDKClient, req *happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIRequest, resp *happytrip.AlibabaHappytripTaxiDriverBlacklistRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
