package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// AlibabaItemOperateUpshelf 商品上架
// alibaba.item.operate.upshelf
//
// 商品上架
func AlibabaItemOperateUpshelf(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemOperateUpshelfAPIRequest, resp *tbitem.AlibabaItemOperateUpshelfAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
