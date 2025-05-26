package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminDeviceYksSkills 根据设备id获取技能列表
// yunos.tvpubadmin.device.yks.skills
//
// 根据设备id获取技能列表
func YunosTvpubadminDeviceYksSkills(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminDeviceYksSkillsAPIRequest, resp *tvupadmin.YunosTvpubadminDeviceYksSkillsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
