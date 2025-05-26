package mos

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/mos"
)

// AlibabaMosStoreGetdefautitems 获取默认状态下商品列表
// alibaba.mos.store.getdefautitems
//
// 获取默认状态下商品列表
func AlibabaMosStoreGetdefautitems(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosStoreGetdefautitemsAPIRequest, resp *mos.AlibabaMosStoreGetdefautitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
