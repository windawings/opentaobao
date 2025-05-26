package cloudgame

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cloudgame"
)

// AlibabaCloudgameInteractiveGamePlayerStop 用户停止游戏
// alibaba.cloudgame.interactive.game.player.stop
//
// 用户停止游戏
func AlibabaCloudgameInteractiveGamePlayerStop(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIRequest, resp *cloudgame.AlibabaCloudgameInteractiveGamePlayerStopAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
