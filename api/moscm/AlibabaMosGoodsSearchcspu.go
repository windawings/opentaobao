package moscm

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/moscm"
)

// AlibabaMosGoodsSearchcspu cspu查询
// alibaba.mos.goods.searchcspu
//
// 商品信息查询（仅用于商品上传数据验证，不能用于商品下载，有限流）
func AlibabaMosGoodsSearchcspu(ctx context.Context, clt *core.SDKClient, req *moscm.AlibabaMosGoodsSearchcspuAPIRequest, resp *moscm.AlibabaMosGoodsSearchcspuAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
