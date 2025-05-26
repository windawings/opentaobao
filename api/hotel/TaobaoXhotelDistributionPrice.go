package hotel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/hotel"
)

// TaobaoXhotelDistributionPrice 飞猪分销通用酒店报价接口
// taobao.xhotel.distribution.price
//
// 飞猪分销通用酒店报价接口
func TaobaoXhotelDistributionPrice(ctx context.Context, clt *core.SDKClient, req *hotel.TaobaoXhotelDistributionPriceAPIRequest, resp *hotel.TaobaoXhotelDistributionPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
