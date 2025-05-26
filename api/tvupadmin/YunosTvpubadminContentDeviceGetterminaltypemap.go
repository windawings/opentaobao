package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentDeviceGetterminaltypemap 获取终端类型表
// yunos.tvpubadmin.content.device.getterminaltypemap
//
// 获取终端类型表
func YunosTvpubadminContentDeviceGetterminaltypemap(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentDeviceGetterminaltypemapAPIRequest, resp *tvupadmin.YunosTvpubadminContentDeviceGetterminaltypemapAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
