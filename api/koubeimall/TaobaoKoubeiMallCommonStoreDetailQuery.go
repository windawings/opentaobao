package koubeimall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/koubeimall"
)

// TaobaoKoubeiMallCommonStoreDetailQuery 查询综合体内的门店详细信息
// taobao.koubei.mall.common.store.detail.query
//
// 查询口碑综合体内的门店详情信息
func TaobaoKoubeiMallCommonStoreDetailQuery(ctx context.Context, clt *core.SDKClient, req *koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIRequest, resp *koubeimall.TaobaoKoubeiMallCommonStoreDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
