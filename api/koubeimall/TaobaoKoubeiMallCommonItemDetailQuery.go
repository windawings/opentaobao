package koubeimall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonItemDetailQuery 查询商品详情信息
// taobao.koubei.mall.common.item.detail.query
//
// 查询口碑综合体内商品详情信息
func TaobaoKoubeiMallCommonItemDetailQuery(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonItemDetailQueryAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonItemDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
