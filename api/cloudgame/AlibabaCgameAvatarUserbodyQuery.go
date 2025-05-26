package cloudgame

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/cloudgame"
)

// AlibabaCgameAvatarUserbodyQuery 用户Avatar body查询
// alibaba.cgame.avatar.userbody.query
//
// Avatar用户body数据查询
func AlibabaCgameAvatarUserbodyQuery(ctx context.Context, clt *core.SDKClient, req *cloudgame.AlibabaCgameAvatarUserbodyQueryAPIRequest, resp *cloudgame.AlibabaCgameAvatarUserbodyQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
