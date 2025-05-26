package hotel

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/hotel"
)

// TaobaoXhotelDistributionInfo 飞猪分销通用酒店标准信息接口
// taobao.xhotel.distribution.info
//
// 飞猪分销通用酒店标准信息接口
func TaobaoXhotelDistributionInfo(ctx context.Context, clt *core.SDKClient, req *hotel.TaobaoXhotelDistributionInfoAPIRequest, resp *hotel.TaobaoXhotelDistributionInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
