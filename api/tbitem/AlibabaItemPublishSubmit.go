package tbitem

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tbitem"
)

// AlibabaItemPublishSubmit 商品发布
// alibaba.item.publish.submit
//
// 新商品发布，提交商品发布信息
func AlibabaItemPublishSubmit(ctx context.Context, clt *core.SDKClient, req *tbitem.AlibabaItemPublishSubmitAPIRequest, resp *tbitem.AlibabaItemPublishSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
