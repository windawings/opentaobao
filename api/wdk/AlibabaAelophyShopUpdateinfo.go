package wdk

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/wdk"
)

// AlibabaAelophyShopUpdateinfo 更新渠道店基础信息
// alibaba.aelophy.shop.updateinfo
//
// 更新渠道店基础信息
func AlibabaAelophyShopUpdateinfo(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaAelophyShopUpdateinfoAPIRequest, resp *wdk.AlibabaAelophyShopUpdateinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
