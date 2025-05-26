package usergrowth

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/usergrowth"
)

// TaobaoGrowthReachingBuzzwordQuery 淘宝热词榜单数据查询接口
// taobao.growth.reaching.buzzword.query
//
// 查询淘宝热词榜单数据
func TaobaoGrowthReachingBuzzwordQuery(ctx context.Context, clt *core.SDKClient, req *usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIRequest, resp *usergrowth.TaobaoGrowthReachingBuzzwordQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
