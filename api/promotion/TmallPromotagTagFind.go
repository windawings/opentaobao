package promotion

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/promotion"
)

// TmallPromotagTagFind 查询标签接口
// tmall.promotag.tag.find
//
// 查询用户创建的所有标签
func TmallPromotagTagFind(ctx context.Context, clt *core.SDKClient, req *promotion.TmallPromotagTagFindAPIRequest, resp *promotion.TmallPromotagTagFindAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
