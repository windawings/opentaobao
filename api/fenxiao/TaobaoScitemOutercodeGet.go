package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoScitemOutercodeGet 根据outerCode查询商品
// taobao.scitem.outercode.get
//
// 根据outerCode查询商品
func TaobaoScitemOutercodeGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoScitemOutercodeGetAPIRequest, resp *fenxiao.TaobaoScitemOutercodeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
