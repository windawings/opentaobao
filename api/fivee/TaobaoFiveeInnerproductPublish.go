package fivee

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fivee"
)

// TaobaoFiveeInnerproductPublish 国产商品发布
// taobao.fivee.innerproduct.publish
//
// 资质共享平台国产商品发布
func TaobaoFiveeInnerproductPublish(ctx context.Context, clt *core.SDKClient, req *fivee.TaobaoFiveeInnerproductPublishAPIRequest, resp *fivee.TaobaoFiveeInnerproductPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
