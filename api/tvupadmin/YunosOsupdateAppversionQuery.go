package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosOsupdateAppversionQuery 分页获取桌面升级任务
// yunos.osupdate.appversion.query
//
// 分页获取桌面升级任务
func YunosOsupdateAppversionQuery(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosOsupdateAppversionQueryAPIRequest, resp *tvupadmin.YunosOsupdateAppversionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
