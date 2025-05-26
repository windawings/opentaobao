package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoScitemAdd 发布后端商品
// taobao.scitem.add
//
// 发布后端商品
func TaobaoScitemAdd(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoScitemAddAPIRequest, resp *fenxiao.TaobaoScitemAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
