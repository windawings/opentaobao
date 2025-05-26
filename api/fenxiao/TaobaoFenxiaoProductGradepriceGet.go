package fenxiao

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/fenxiao"
)

// TaobaoFenxiaoProductGradepriceGet 等级折扣查询
// taobao.fenxiao.product.gradeprice.get
//
// 等级折扣查询
func TaobaoFenxiaoProductGradepriceGet(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoFenxiaoProductGradepriceGetAPIRequest, resp *fenxiao.TaobaoFenxiaoProductGradepriceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
