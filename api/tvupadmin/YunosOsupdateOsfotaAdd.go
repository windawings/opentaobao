package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosOsupdateOsfotaAdd 添加系统升级任务
// yunos.osupdate.osfota.add
//
// 添加osupdate系统升级任务
func YunosOsupdateOsfotaAdd(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateOsfotaAddAPIRequest, resp *tvupadmin.YunosOsupdateOsfotaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
