package usergrowth2

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/usergrowth2"
)

// TaobaoUsergrowthAdMaterialUpdate 素材更新
// taobao.usergrowth.ad.material.update
//
// 素材更新
func TaobaoUsergrowthAdMaterialUpdate(ctx context.Context, clt *core.SDKClient, req *usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIRequest, resp *usergrowth2.TaobaoUsergrowthAdMaterialUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
