package smartstore

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/smartstore"
)

// TmallPopupstoreItemDiscountPrice 商品优惠价格查询
// tmall.popupstore.item.discount.price
//
// 商品优惠价格查询
func TmallPopupstoreItemDiscountPrice(ctx context.Context, clt *core.SDKClient, req *smartstore.TmallPopupstoreItemDiscountPriceAPIRequest, resp *smartstore.TmallPopupstoreItemDiscountPriceAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
