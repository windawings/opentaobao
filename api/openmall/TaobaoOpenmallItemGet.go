package openmall

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/openmall"
)

// TaobaoOpenmallItemGet 获取商品详情物料
// taobao.openmall.item.get
//
// 获取联盟开放的openmall商品
func TaobaoOpenmallItemGet(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallItemGetAPIRequest, resp *openmall.TaobaoOpenmallItemGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
