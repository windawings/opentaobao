package uscesl

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/uscesl"
)

// TaobaoUsceslIteminfoBatchPut 批量写入商品信息接口
// taobao.uscesl.iteminfo.batch.put
//
// 电子架签批量写入商品数据，用于电子价签展示
func TaobaoUsceslIteminfoBatchPut(ctx context.Context, clt *core.SDKClient, req *uscesl.TaobaoUsceslIteminfoBatchPutAPIRequest, resp *uscesl.TaobaoUsceslIteminfoBatchPutAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
