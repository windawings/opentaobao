package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductcatUpdate 修改产品线
// taobao.fenxiao.productcat.update
//
// 修改产品线
func TaobaoFenxiaoProductcatUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductcatUpdateAPIRequest, resp *fenxiao.TaobaoFenxiaoProductcatUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
