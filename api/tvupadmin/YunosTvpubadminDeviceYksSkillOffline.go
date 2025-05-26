package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkillOffline 技能下架
// yunos.tvpubadmin.device.yks.skill.offline
//
// 迎客松平台技能下架
func YunosTvpubadminDeviceYksSkillOffline(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillOfflineAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
