package tvupadmin

import (
	"context"

	"github.com/windawings/opentaobao/core"
	"github.com/windawings/opentaobao/model/tvupadmin"
)

// YunosTvpubadminContentAppQueryapp 查询应用信息
// yunos.tvpubadmin.content.app.queryapp
//
// 查询应用信息
func YunosTvpubadminContentAppQueryapp(ctx context.Context, clt *core.SDKClient, req *tvupadmin.YunosTvpubadminContentAppQueryappAPIRequest, resp *tvupadmin.YunosTvpubadminContentAppQueryappAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
