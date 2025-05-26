package cloudgame

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGameGamepadGet 获取虚拟手柄配置
// alibaba.cloudgame.interactive.game.gamepad.get
//
// 获取虚拟手柄配置
func AlibabaCloudgameInteractiveGameGamepadGet(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGameGamepadGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
