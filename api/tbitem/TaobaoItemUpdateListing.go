package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// TaobaoItemUpdateListing 一口价商品上架
// taobao.item.update.listing
//
// * 单个商品上架&lt;br/&gt;* 输入的num_iid必须属于当前会话用户
func TaobaoItemUpdateListing(ctx context.Context, clt *core.SDKClient, req *tbitem.TaobaoItemUpdateListingAPIRequest, resp *tbitem.TaobaoItemUpdateListingAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
