package aedropshiper

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/aedropshiper"
)

// AliexpressTradeBuyPlaceorder AE下单API
// aliexpress.trade.buy.placeorder
//
// 150欧欧盟税改
func AliexpressTradeBuyPlaceorder(ctx context.Context, clt *core.SDKClient, req *aedropshiper.AliexpressTradeBuyPlaceorderAPIRequest, resp *aedropshiper.AliexpressTradeBuyPlaceorderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
